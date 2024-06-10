---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_organization_onboarding"
sidebar_current: "docs-resource-dome9-aws-organization-onboarding"
description: Onboard AWS organization to CloudGuard
---

# dome9_aws_organization_onboarding

Connect an AWS organization to CloudGuard in one quick process.

## Example Usage

Basic usage:

```hcl
  resource "dome9_aws_organization_onboarding" "test" {
  role_arn              = "ROLE_ARN"
  secret                = "SECRET"
  stack_set_arn         = "STACK_SET_ARN"
  aws_organization_name = "AWS_ORG_NAME"
}
```

## Argument Reference

The following arguments are supported:

* `role_arn` - (Required) CloudGuard role ARN from AWS.
* `secret ` - (Required) External ID from the management-stack API.
* `api_key` - (Optional) API key, needed only for 'UserBased' type.
* `stack_set_arn` - (Required) The created StackSet ARN.
* `aws_organization_name` - (Optional) Organization name in CloudGuard.
* `enable_stack_modify` - (Optional) Boolean flag to enable stack modification. Default is false.
* `type` - (Optional) Credential type. Default is RoleBased. Can be: `UserBased`, `RoleBased`.

  
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

































 