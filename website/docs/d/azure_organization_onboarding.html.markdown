---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_azure_organization_onboarding"
sidebar_current: "docs-resource-dome9-azure-organization-onboarding"
description: Onboard Azure organization to CloudGuard
---

# Data Source: dome9_azure_organization_onboarding

Use this data source to get information about connected Azire organization to CloudGuard.

## Example Usage

Basic usage:

```hcl
data "dome9_azure_organization_onboarding" "test" {
  id = "ORGANIZATION_ID"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) CloudGuard Azure organization onboarding entity ID.

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
        * `with_sse_cmk_encrypted_disks_scan` - Boolean flag to enable sse cmk scan.
    * `serverless_configuration` - Serverless configuration.
        * `is_enabled` - Boolean flag to enable serverless protection.
        * `accounts` - List of serverless accounts.
            * `storage_id` - Storage account ID.
            * `log_types` - List of log types.
* `is_auto_onboarding` - Declares if the onboarding pipeline automatically onboards newly discovered subscriptions after the initial onboarding.
* `update_time` - last update time of the stackSet.
* `creation_time` - Creation time of the organization.
































 