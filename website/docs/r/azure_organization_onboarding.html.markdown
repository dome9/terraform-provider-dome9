---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_azure_organization_onboarding"
sidebar_current: "docs-resource-dome9-azure-organization-onboarding"
description: Onboard Azure organization to CloudGuard
---

# dome9_azure_organization_onboarding

Connect an Azure organization to CloudGuard in one quick process.

### Important
First you need to fill all onboarding configuration in the CloudGuard portal in the [Azure onboarding page](https://secure.dome9.com/v2/azure-onboarding) to get the Cloud Shell script.
The script should be run and complete before applying the `dome9_azure_organization_onboarding` resource.

## Example Usage

Basic usage:

```hcl
  resource "dome9_azure_organization_onboarding" "test" {
  tenant_id           = "TENANT_ID"
  management_group_id = "MANAGEMENT_GROUP_ID"

  active_blades {
    awp {
      is_enabled       = false
    }
    cdr {
      is_enabled = false
    }
    serverless {
      is_enabled = false
    }
    posture_management {
      onboarding_mode = "ONBOARDING_MODE"
    }
  }

  use_cloud_guard_managed_app = false
  is_auto_onboarding          = false
  vendor                      = "VENDOR"
}
```

Advanced usage:

```hcl
resource "dome9_azure_organization_onboarding" "test" {
  tenant_id           = "TENANT_ID"
  management_group_id = "MANAGEMENT_GROUP_ID"

  active_blades {
    awp {
      is_enabled       = true
      onboarding_mode  = "ONBOARDING_MODE"
    }
    cdr {
      is_enabled = true
      accounts {
        storage_id = "STORAGE_ID"
        log_types  = ["LOG_TYPE1", "LOG_TYPE2"]
      }
    }
    serverless {
      is_enabled = false
    }
    posture_management {
      onboarding_mode = "ONBOARDING_MODE"
    }
  }

  use_cloud_guard_managed_app = false
  is_auto_onboarding          = false
  vendor                      = "VENDOR"
}

```

## Argument Reference

The following arguments are supported:

* `workflow_id` - (Optional) The workflow ID.
* `tenant_id` - (Required) The Tenant ID to onboard.
* `management_group_id` - (Optional) The Management Group ID to onboard.
* `organization_name` - (Optional) Organization name in CloudGuard. Default is `AzureOrg`.
* `app_registration_name` - (Optional) The name of the application created in the script. Required only if non UseCloudGuardManagedApp mode is used.
* `client_id` - (Optional) Application (client) ID. Required only if non UseCloudGuardManagedApp mode is used.
* `client_secret` - (Optional) Azure client secret. Required only if non UseCloudGuardManagedApp mode is used.
* `active_blades` - (Required) Indicates which blades to Activate.
    * `awp` - (Required) Azure Workload Protection configuration.
        * `is_enabled` - (Required) Boolean flag to enable Azure Workload Protection.
        * `onboarding_mode` - (Optional) Onboarding mode. Can be: `saas`, `inAccount`, `inAccountHub`.
        * `centralized_subscription_id` - (Optional) Centralized subscription ID.
        * `with_function_apps_scan` - (Optional) Boolean flag to enable function apps scan.
    * `cdr` - (Required) CloudGuard Data Protection configuration.
        * `is_enabled` - (Required) Boolean flag to enable CloudGuard Data Protection.
        * `accounts` - (Optional) List of storage accounts.
            * `storage_id` - (Required) Storage account ID.
            * `log_types` - (Optional) List of log types.
    * `serverless` - (Required) Serverless configuration.
        * `is_enabled` - (Required) Boolean flag to enable serverless protection. - Not supported yet.
        * `accounts` - (Optional) List of serverless accounts.
            * `storage_id` - (Required) Storage account ID.
            * `log_types` - (Optional) List of log types.
    * `posture_management` - (Required) Posture management configuration.
        * `onboarding_mode` - (Required) Onboarding mode. Can be: `Read`, `Manage`.
        * `rulesets_ids` - (Optional) List of ruleset IDs that will run automatically on the organization cloud accounts.
* `vendor` - (Required) Vendor name. Can be: `azure`, `azurechina`, `azuregov`. Default is `azure`.
* `use_cloud_guard_managed_app` - (Optional) Specifies whether to use the Check Point application to connect the subscriptions to CloudGuard. Default is false.
* `is_auto_onboarding` - (Optional) Declares if the onboarding pipeline automatically onboards newly discovered subscriptions after the initial onboarding. Default is true and cannot change to false.


## Attributes Reference

* `account_id` - CloudGuard account ID.
* `user_id` - CloudGuard user ID.
* `organization_name` - Organization name in CloudGuard.
* `tenant_id` - Azure tenant ID.
* `management_group_id` - Azure management group ID.
* `app_registration_name` - Azure app registration name.
* `onboarding_configuration` - Onboarding configuration.
    * `organization_root_ou_id` - Organization root OU ID.
    * `mapping_strategy` - Mapping strategy type.
    * `posture_management` - Posture management configuration.
        * `rulesets_ids` - List of ruleset IDs that will run automatically on the organization cloud accounts.
        * `onboarding_mode` - Onboarding mode. Can be: `Read`, `Manage`.
    * `awp_configuration` - Azure Workload Protection configuration.
        * `is_enabled` - Boolean flag to enable Azure Workload Protection.
        * `onboarding_mode` - Onboarding mode. Can be: `saas`, `inAccount`, `inAccountHub`.
        * `centralized_subscription_id` - Centralized subscription ID.
        * `with_function_apps_scan` - Boolean flag to enable function apps scan.
    * `serverless_configuration` - Serverless configuration.
        * `is_enabled` - Boolean flag to enable serverless protection.
        * `accounts` - List of serverless accounts.
            * `storage_id` - Storage account ID.
            * `log_types` - List of log types.
* `is_auto_onboarding` - Boolean flag to enable auto onboarding.
* `update_time` - last update time of the stackSet.
* `creation_time` - Creation time of the organization.

































 