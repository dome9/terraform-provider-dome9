---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_organization_onboarding"
sidebar_current: "docs-resource-dome9-aws-organization-onboarding"
description: Get management stack configuration for given AWS account.
---

# Data Source: dome9_aws_organization_onboarding

Get management stack configuration for given AWS account.
For empty awsAccountId, return the default configuration.

## Example Usage

Basic usage:

```hcl
data "dome9_aws_organization_onboarding" "test" {
  id = "ORGANIZATION_ID"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) CloudGuard AWS organization onboarding entity ID.

## Attributes Reference

* `account_id` - CloudGuard account ID.
* `external_organization_id` - External management account ID (Account ID in AWS).
* `management_account_stack_id` - Management account stack ID.
* `management_account_stack_region` - Management account stack region.
* `onboarding_configuration` - Onboarding configuration.
    * `organization_root_ou_id` - Organization root OU ID.
    * `mapping_strategy` - Mapping strategy type.
    * `posture_management` - Posture management configuration.
        * `rulesets_ids` - List of ruleset IDs that will run automatically on the organization cloud accounts.
        * `onboarding_mode` - Onboarding mode. Can be: `Read`, `Manage`.
* `user_id` - CloudGuard user ID.
* `organization_name` - Organization name in CloudGuard.
* `update_time` - last update time of the stackSet.
* `creation_time` - Creation time of the organization.
* `stack_set_regions` - List of AWS regions the StackSet has stack instances deployed in.
* `stack_set_organizational_unit_ids` - List of organization root ID or organizational unit (OU) IDs.

































 