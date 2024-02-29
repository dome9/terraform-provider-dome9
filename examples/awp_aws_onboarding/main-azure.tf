# Dome9 Provider Configurations
terraform {
  required_providers {
    dome9 = {
      source = "dome9/dome9"
      version = "1.29.6"
    }
  }
}

/**
provider "dome9" {
  dome9_access_id     = "DOME9_CLOUDGUARD_API_ACCESS_ID"
  dome9_secret_key    = "DOME9_CLOUDGUARD_API_SECRET_KEY"
  base_url            = "DOME9_API_BASE_URL"
}
**/

# Define variables
# Regular Account or Sub Account Subscription ID
variable "subscription_id" {} # Customer should provide it with the onboarding resource
variable "management_group" {} # Customer should provide it with the onboarding resource
variable "tenant_id" {} # should be calculated from azurerm_subscription data source
variable "app_object_id" {} # should be calculated from app_client_id (it should be provided from get onboarding data api)
variable "hub_subscription_id" {} # Customer should provide it with the onboarding resource
variable "scan_mode" {} # Customer should provide it with the onboarding resource
variable "region" {} # should be taken from get onboarding data azure api
variable "awp_version" {} # should be taken from get onboarding data azure api


data "dome9_awp_azure_generate_onboarding_data" "dome9_awp_azure_generating_onboarding_data_source" {
  subscription_id = "d0dd3387-d9c5-487b-8b18-4fa969fd98cd"
  scan_mode = "inAccountHub" # The scan mode (valid options are: $SCAN_MODE_SAAS, $SCAN_MODE_IN_ACCOUNT, $SCAN_MODE_IN_ACCOUNT_HUB, $SCAN_MODE_IN_ACCOUNT_SUB)
  hub_subscription_id = "d0dd3387-d9c5-487b-8b18-4fa969fd98cd" # The hub subscription id, this param is relevant in case scan_mode is $SCAN_MODE_IN_ACCOUNT_HUB
                                                                # or $SCAN_MODE_IN_ACCOUNT_SUB and represents the subscription where the AWP scans will be executed
  skip_function_apps_scan  = false # currently this attribute not supported with the azure resources (the default is false as we understood)
}


# locals
locals {
  SCAN_MODE_SAAS = "saas"
  SCAN_MODE_IN_ACCOUNT = "inAccount"
  SCAN_MODE_IN_ACCOUNT_SUB = "inAccountSub"
  SCAN_MODE_IN_ACCOUNT_HUB = "inAccountHub"

  AWP_VM_SCAN_OPERATOR_ROLE_NAME_PREFIX = "CloudGuard AWP VM Scan Operator"
  AWP_VM_SCAN_OPERATOR_ROLE_DESCRIPTION = "Grants all needed permissions for CloudGuard app registration to scan VMs (version: ${var.awp_version})"
  AWP_VM_SCAN_OPERATOR_ROLE_ACTIONS = [
    "Microsoft.Compute/disks/read",
    "Microsoft.Compute/disks/write",
    "Microsoft.Compute/disks/delete",
    "Microsoft.Compute/disks/beginGetAccess/action",
    "Microsoft.Compute/snapshots/read",
    "Microsoft.Compute/snapshots/write",
    "Microsoft.Compute/snapshots/delete",
    "Microsoft.Compute/snapshots/beginGetAccess/action",
    "Microsoft.Compute/snapshots/endGetAccess/action",
    "Microsoft.Network/networkInterfaces/join/action",
    "Microsoft.Network/networkInterfaces/write",
    "Microsoft.Compute/virtualMachines/write",
    "Microsoft.Compute/virtualMachines/delete",
    "Microsoft.Network/networkSecurityGroups/write",
    "Microsoft.Network/networkSecurityGroups/join/action",
    "Microsoft.Network/virtualNetworks/write",
    "Microsoft.Network/virtualNetworks/subnets/join/action"
  ]

  AWP_VM_DATA_SHARE_ROLE_NAME_PREFIX = "CloudGuard AWP VM Data Share"
  AWP_VM_DATA_SHARE_ROLE_DESCRIPTION = "Grants needed permissions for CloudGuard app registration to read VMs data (version: ${var.awp_version})"
  AWP_VM_DATA_SHARE_ROLE_ACTIONS = [
    "Microsoft.Compute/disks/beginGetAccess/action",
    "Microsoft.Compute/virtualMachines/read"
  ]

  AWP_FA_MANAGED_IDENTITY_NAME = "CloudGuardAWPScannerManagedIdentity"

  AWP_FA_SCANNER_ROLE_NAME_PREFIX = "CloudGuard AWP Function Apps Scanner"
  AWP_FA_SCANNER_ROLE_DESCRIPTION = "Grants needed permissions for CloudGuard AWP function-apps scanner (version: ${var.awp_version})"
  AWP_FA_SCANNER_ROLE_ACTIONS = [
    "Microsoft.Web/sites/publish/Action",
    "Microsoft.Web/sites/config/list/Action",
    "microsoft.web/sites/functions/read"
  ]

  AWP_FA_SCAN_OPERATOR_ROLE_NAME_PREFIX = "CloudGuard AWP FunctionApp Scan Operator"
  AWP_FA_SCAN_OPERATOR_ROLE_DESCRIPTION = "Grants all needed permissions for CloudGuard app registration to scan function-apps (version: ${var.awp_version})"
  AWP_FA_SCAN_OPERATOR_ROLE_ACTIONS = [
    "Microsoft.Compute/virtualMachines/write",
    "Microsoft.Compute/virtualMachines/extensions/write",
    "Microsoft.Network/networkSecurityGroups/write",
    "Microsoft.Network/networkSecurityGroups/join/action",
    "Microsoft.Network/virtualNetworks/write",
    "Microsoft.Network/virtualNetworks/subnets/join/action",
    "Microsoft.ManagedIdentity/userAssignedIdentities/assign/action"
  ]

  AWP_RESOURCE_GROUP_NAME_PREFIX = "cloudguard-AWP"
  AWP_OWNER_TAG = "Owner=CG.AWP"
  AWP_VERSION_TAG = "CloudGuard.AWP.Version=${var.awp_version}"
  LOCATION = var.region
}


# Provider block for the hub account (used only in In-Account-Sub mode)
provider "azurerm" {
  alias   = "hub"
  features {}

  subscription_id = var.hub_subscription_id
  # Add any other necessary authentication details for the hub account
}

# Provider block for the sub account (used only in In-Account-Sub mode)
provider "azurerm" {
  alias   = "sub"
  features {}

  subscription_id = var.subscription_id
  # Add any other necessary authentication details for the sub account
}

# Data source to retrieve information about the current Azure subscription
data "azurerm_subscription" "hub" {
  provider = azurerm.hub
}

data "azurerm_subscription" "sub" {
  provider = azurerm.sub
}

# Define the resource group where CloudGuard resources will be deployed
resource "azurerm_resource_group" "cloudguard" {
  count = var.scan_mode == local.SCAN_MODE_IN_ACCOUNT || var.scan_mode == local.SCAN_MODE_IN_ACCOUNT_HUB ? 1 : 0
  name     = local.AWP_RESOURCE_GROUP_NAME_PREFIX
  location = local.LOCATION
  tags     = {
    Owner   = local.AWP_OWNER_TAG
    Version = local.AWP_VERSION_TAG
  }
}

resource "azurerm_resource_group" "cloudguard_hub" {
  count = var.scan_mode == local.SCAN_MODE_IN_ACCOUNT_SUB ? 1 : 0
  provider = azurerm.hub
  name     = "${local.AWP_RESOURCE_GROUP_NAME_PREFIX}_${var.subscription_id}"
  location = local.LOCATION
  tags     = {
    Owner   = local.AWP_OWNER_TAG
    Version = local.AWP_VERSION_TAG
  }
}

# Define custom roles based on scan mode
resource "azurerm_role_definition" "cloudguard_vm_data_share" {
  count           = var.scan_mode == local.SCAN_MODE_IN_ACCOUNT || var.scan_mode == local.SCAN_MODE_IN_ACCOUNT_HUB || var.scan_mode == local.SCAN_MODE_SAAS ? 1 : 0
  name            = "CloudGuard AWP VM Data Share ${var.subscription_id}" # need to change subscription id to hub subscription id when hub mode
  description = local.AWP_VM_DATA_SHARE_ROLE_DESCRIPTION
  scope           = var.scan_mode == local.SCAN_MODE_IN_ACCOUNT || var.scan_mode == local.SCAN_MODE_SAAS ? "/subscriptions/${var.subscription_id}" : "/providers/Microsoft.Management/managementGroups/${var.management_group}:-${var.tenant_id}"
  permissions {
    actions     = local.AWP_VM_DATA_SHARE_ROLE_ACTIONS
    not_actions = []
  }
}

# Define the managed identity for CloudGuard AWP
resource "azurerm_managed_identity" "cloudguard_identity" {
  count    = var.scan_mode == local.SCAN_MODE_IN_ACCOUNT || var.scan_mode == local.SCAN_MODE_IN_ACCOUNT_HUB ? 1 : 0
  name     = local.AWP_FA_MANAGED_IDENTITY_NAME
  location = azurerm_resource_group.cloudguard.location
  resource_group_name = azurerm_resource_group.cloudguard.name
}

resource "azurerm_role_definition" "cloudguard_vm_scan_operator" {
  count           = var.scan_mode == local.SCAN_MODE_IN_ACCOUNT || var.scan_mode == local.SCAN_MODE_IN_ACCOUNT_HUB ? 1 : 0
  description = local.AWP_VM_SCAN_OPERATOR_ROLE_DESCRIPTION
  name            = "${local.AWP_VM_SCAN_OPERATOR_ROLE_NAME_PREFIX} ${var.subscription_id}"
  scope           = var.scan_mode == local.SCAN_MODE_IN_ACCOUNT ? "/subscriptions/${var.subscription_id}" : "/providers/Microsoft.Management/managementGroups/${var.management_group}:-${var.tenant_id}"
  permissions {
    actions     = local.AWP_VM_SCAN_OPERATOR_ROLE_ACTIONS
    not_actions = []
  }
}

resource "azurerm_role_definition" "cloudguard_function_apps_scanner" {
  count           = var.scan_mode == local.SCAN_MODE_IN_ACCOUNT || var.scan_mode == local.SCAN_MODE_IN_ACCOUNT_HUB ? 1 : 0
  name            = "${local.AWP_FA_SCANNER_ROLE_NAME_PREFIX} ${var.subscription_id}"
  description = local.AWP_FA_SCANNER_ROLE_DESCRIPTION
  scope           = var.scan_mode == local.SCAN_MODE_IN_ACCOUNT ? "/subscriptions/${var.subscription_id}" : "/providers/Microsoft.Management/managementGroups/${var.management_group}:-${var.tenant_id}"
  permissions {
    actions     = local.AWP_FA_SCANNER_ROLE_ACTIONS
    not_actions = []
  }
}

resource "azurerm_role_definition" "cloudguard_function_apps_scan_operator" {
  count           = var.scan_mode == local.SCAN_MODE_IN_ACCOUNT || var.scan_mode == local.SCAN_MODE_IN_ACCOUNT_HUB ? 1 : 0
  name            = "${local.AWP_FA_SCAN_OPERATOR_ROLE_NAME_PREFIX} ${var.subscription_id}"
  description = local.AWP_FA_SCAN_OPERATOR_ROLE_DESCRIPTION
  scope           = var.scan_mode == local.SCAN_MODE_IN_ACCOUNT ? "/subscriptions/${var.subscription_id}" : "/providers/Microsoft.Management/managementGroups/${var.management_group}:-${var.tenant_id}"
  permissions {
    actions     = local.AWP_FA_SCAN_OPERATOR_ROLE_ACTIONS
    not_actions = []
  }
}

# Assign custom roles based on scan mode
resource "azurerm_role_assignment" "cloudguard_vm_data_share_assignment" {
  count           = var.scan_mode == local.SCAN_MODE_SAAS || var.scan_mode == local.SCAN_MODE_IN_ACCOUNT || var.scan_mode == local.SCAN_MODE_IN_ACCOUNT_SUB ? 1 : 0
  provider = azurerm.sub
  name = "${local.AWP_VM_DATA_SHARE_ROLE_NAME_PREFIX} ${var.subscription_id}"
  scope           = "/subscriptions/${var.subscription_id}"
  role_definition_name = azurerm_role_definition.cloudguard_vm_data_share[count.index].name
  principal_id    = var.app_object_id
}

resource "azurerm_role_assignment" "cloudguard_vm_scan_operator_assignment" {
  count           = var.scan_mode == local.SCAN_MODE_IN_ACCOUNT || var.scan_mode == local.SCAN_MODE_IN_ACCOUNT_HUB ? 1 : 0
  scope           = "/subscriptions/${var.subscription_id}"
  role_definition_name = azurerm_role_definition.cloudguard_vm_scan_operator[count.index].name
  principal_id    = var.app_object_id
}

resource "azurerm_role_assignment" "cloudguard_function_apps_scanner_assignment" {
  count           = var.scan_mode == local.SCAN_MODE_IN_ACCOUNT || var.scan_mode == local.SCAN_MODE_IN_ACCOUNT_HUB ? 1 : 0
  provider = azurerm.sub
  scope           = "/subscriptions/${var.subscription_id}"
  role_definition_name = azurerm_role_definition.cloudguard_function_apps_scanner[count.index].name
  principal_id    = azurerm_managed_identity.cloudguard_identity[count.index].id
}

resource "azurerm_role_assignment" "cloudguard_function_apps_scan_operator_assignment" {
  count           = var.scan_mode == local.SCAN_MODE_IN_ACCOUNT || var.scan_mode == local.SCAN_MODE_IN_ACCOUNT_HUB ? 1 : 0
  scope           = "/subscriptions/${var.subscription_id}"
  role_definition_name = azurerm_role_definition.cloudguard_function_apps_scan_operator[count.index].name
  principal_id    = var.app_object_id
}

resource "azurerm_resource_group" "cloudguard_hub" {
  count     = var.scan_mode == "inAccountSub" ? 1 : 0
  name     = "cloudguard-AWP-${var.subscription_id}"
  location = var.region
}

resource "dome9_awp_azure_onboarding" "awp_azure_onboarding_resource" {
  subscription_id = "d0dd3387-d9c5-487b-8b18-4fa969fd98cd"
  scan_mode = "inAccountHub" # The scan mode (valid options are: $SCAN_MODE_SAAS, $SCAN_MODE_IN_ACCOUNT, $SCAN_MODE_IN_ACCOUNT_HUB, $SCAN_MODE_IN_ACCOUNT_SUB)
  hub_subscription_id = "d0dd3387-d9c5-487b-8b18-4fa969fd98cd"

  # azure role name customizations currently unsupported
  onboarding_customizations = {
    virtual_machine_data_share_role_name   = "string"
    virtual_machine_scan_operator_role_name = "string"
    function_app_scan_operator_role_name   = "string"
    function_app_scanner_role_name         = "string"
    resource_group_name                    = "string"
    scanner_managed_identity_name          = "string"
  }
  agentless_account_settings = {
    disabled_regions               = ["string"]
    scan_machine_interval_in_hours = 0
    max_concurrence_scans_per_region = 0
    skip_function_apps_scan       = false #
    custom_tags                    = {}
  }
}
