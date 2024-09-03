---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_awp_azure_onboarding"
sidebar_current: "docs-resource-dome9-awp-azure-onboarding"
description: |-
  Creates an AWP Azure Onboarding in Dome9
---

# dome9_awp_azure_onboarding

This resource is used to create and modify AWP Azure Onboarding in CloudGuard Dome9.

## Example Usage

Basic usage:

```hcl
provider "dome9" {
  dome9_access_id     = "DOME9_ACCESS_ID"
  dome9_secret_key    = "DOME9_SECRET_KEY"
  base_url            = "https://api.dome9.com/v2/"
}


# There is a need to use this terraform module [terraform-dome9-awp-azure] to create all the prerequisites for the onboarding process (All the needed Azure Resources)
# for further information please refer to the module documentation [terraform-dome9-awp-azure](https://registry.terraform.io/modules/dome9/awp-azure/dome9/latest)
# for more examples (simple and complete), you can visit this github examples [terraform-dome9-awp-azure](https://github.com/dome9/terraform-dome9-awp-azure/blob/master/examples)
# Example for the module use:
module "terraform-dome9-awp-azure" {
	source = "github.com/dome9/terraform-dome9-awp-azure"
	awp_cloud_account_id = "<CLOUDGUARD_ACCOUNT_ID or <AZURE_SUBSCRIPTION_ID>"
	awp_scan_mode = "<SCAN_MODE>" # Possible values: "inAccount", "saas", "inAccountHub", "inAccountSub"

	# Optional customizations:
	# e.g:
        # management_group_id = "management group id" # relevat only for inAccountHub mode
        # 
        # the following parameter is required for "InAccountSub" scan mode
        # awp_centralized_account_id = "<CENTRALIZED_CLOUD_ACCOUNT_ID> or <CENTRALIZED_AZURE_SUBSCRIPTION_ID>" # centralized account-id where AWP scanner runs

	# Optional account Settings
	# e.g:
	#   awp_account_settings_azure = {
	#     scan_machine_interval_in_hours = 24
	#     skip_function_apps_scan = false
	#     disabled_regions = ["eastus", "westus", ...] # List of regions to disable
	#     max_concurrent_scans_per_region = 20
	#     in_account_scanner_vpc = "ManagedByAWP"
	#     custom_tags = {
	#       tag1 = "value1"
	#       tag2 = "value2"
	#       tag3 = "value3"
	#       ...
	#     }
	# }
}

# The dome9_awp_azure_onboarding resource defines a Dome9 AWP Azure Onboarding.
# The Dome9 AWP Azure Onboarding resource allows you to onboard an Azure account to Dome9 AWP.
# The cloudguard_account_id attribute is used to specify the CloudGuard account id of the Azure account.
# The management_group_id attribute is used to specify the management group id in that AWP used in the Azure account.
# The scan_mode attribute is used to specify the scan mode of the Dome9 AWP Azure Onboarding. The valid values are "inAccount", "saas", "inAccountHub" and "inAccountSub".
# The agentless_account_settings attribute is used to specify the agentless account settings of the Dome9 AWP Azure Onboarding.
# The disabled_regions attribute is used to specify the disabled regions of the agentless account settings of the Dome9 AWP Azure Onboarding.
# The skip_function_apps_scan attribute is used to specify if skip Azure Function Apps scan in the agentless account settings of the Dome9 AWP Azure Onboarding.
# The scan_machine_interval_in_hours attribute is used to specify the scan machine interval in hours of the agentless account settings of the Dome9 AWP Azure Onboarding.
# The max_concurrent_scans_per_region attribute is used to specify the max concurrent scans per region of the agentless account settings of the Dome9 AWP Azure Onboarding.
# The in_account_scanner_vpc attribute is used to specify the scanner VPC mode of the agentless account settings of the Dome9 AWP AWS Onboarding.
# The custom_tags attribute is used to specify the custom tags of the agentless account settings of the Dome9 AWP Azure Onboarding.
resource "dome9_awp_azure_onboarding" "awp_azure_onboarding_test" {
  cloudguard_account_id = "dome9_cloudaccount_azure.azure_onboarding_account_test.id | <CLOUDGUARD_ACCOUNT_ID> | <AZURE_SUBSCRIPTION_ID>"
  scan_mode = "<SCAN_MODE>" # Possible values: "inAccount", "saas", "inAccountHub", "inAccountSub"
  awp_centralized_account_id = "<CENTRALIZED_CLOUD_ACCOUNT_ID> or <CENTRALIZED_AZURE_SUBSCRIPTION_ID>" # required for "InAccountSub" scan mode, it is the centralized account-id where AWP scanner runs

  # Optional account Settings (supported for 'inAccount', 'inAccountSub' and 'saas' scan modes)
  # e.g:
  agentless_account_settings {
    disabled_regions = ["eastus", "westus"]
    skip_function_apps_scan = false
    scan_machine_interval_in_hours = 24
    max_concurrent_scans_per_region = 20
    in_account_scanner_vpc = "ManagedByAWP"
    custom_tags = {
      tag1 = "value1"
      tag2 = "value2"
      tag3 = "value3"
    }
  }
}

# The dome9_awp_azure_onboarding data source allows you to get the onboarding data of an Azure account (Optional).
data "dome9_awp_azure_onboarding" "awp_azure_onboarding_test" {
  id = dome9_awp_azure_onboarding.awp_azure_onboarding_test.cloudguard_account_id
}
```

## Argument Reference

The following arguments are supported:

* `cloudguard_account_id` - (Required) The CloudGuard account id.
* `scan_mode` - (Required) The scan mode. Valid values are "inAccount", "saas", "inAccountHub", "inAccountSub".
* `awp_centralized_account_id` - (Optional) The centralized cloud account id, required (and only relevant) for "inAccountSub" scan mode
* `management_group_id` -  the management group id, relevat only for inAccountHub mode.
* `agentless_account_settings` - (Optional) The agentless account settings.
  * `disabled_regions` - (Optional) The disabled regions. valid values are "centralus", "eastus", "eastus2", "usgovlowa", "usgovvirginia", "northcentralus", "southcentralus", "westus", "westus2", "westcentralus", "northeurope", "westeurope", "eastasia", "southeastasia", "japaneast", "japanwest", "brazilsouth", "australiaeast", "australiasoutheast", "centralindia", "southindia", "westindia", "chinaeast", "chinanorth", "canadacentral", "canadaeast", "germanycentral", "germanynortheast", "koreacentral", "uksouth", "ukwest", "koreasouth"
  * `scan_machine_interval_in_hours` - (Optional) The scan machine interval in hours
  * `skip_function_apps_scan` - (Optional) Skip Azure Function Apps scan (supported for inAccount and inAccountSub scan modes)
  * `max_concurrent_scans_per_region` - (Optional) The max concurrent scans per region
  * `in_account_scanner_vpc` = optional(string) # The VPC Mode. Valid values: "ManagedByAWP", "ManagedByCustomer" (supported for inAccount and inAccountHub scan modes)
  * `custom_tags` - (Optional) The custom tags.
* `should_create_policy` - (Optional) Whether to create a policy. Default is true.
    
## Attributes Reference

* `missing_awp_private_network_regions` - The missing AWP private network regions.
* `cloud_account_id` - The cloud guard account id.
* `agentless_protection_enabled` - Whether agentless protection is enabled.
* `cloud_provider` - The cloud provider.
* `should_update` - Whether to update.
* `is_org_onboarding` - Whether is org onboarding.

## Import

The AWP Azure Onboarding can be imported; use <ONBOARDING ID> as the import ID.

For example:

```shell
terraform import dome9_awp_azure_onboarding.test_awp_azure_onboarding 00000000-0000-0000-0000-000000000000
```
