---
layout: "dome9"
page_title: "Check Point Cloud Guard Dome9: dome9_cloudaccount_aws"
sidebar_current: "docs-resource-dome9-cloudaccount-aws"
description: |-
  Manages AWS cloud account.
---

# dome9_cloudaccount_aws

The AWS cloud accounts resource has methods to onboard AWS cloud accounts to Dome9 and to manage some of their settings.

## Example Usage

Basic usage:

```hcl
resource "dome9_cloudaccount_aws" "test" {
  name  = "ACCOUNT NAME"
 
  credentials = {
    arn    = "ARN"
    secret = "SECRET"
    type   = "RoleBased"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The AWS account name.
* `credentials` - (Required) The information needed for Dome9 System in order to connect to the AWS cloud account.

The `credentials` block supports:
    
* `arn` - (Required) AWS Role ARN (to be assumed by Dome9 System)
* `secret` - (Required) The AWS role External ID (Dome9 System will have to use this secret in order to assume the role)
* `type` - (Required) The cloud account onboarding method. Should be set to "RoleBased" as other methods are deprecated.

## Attributes Reference

* `id` - The ID of the AWS cloud account.
* `vendor` - The cloud provider (AWS).
* `external_account_number` - The AWS account number.
* `is_fetching_suspended` - Fetching suspending status.
* `creation_date` - Account creation date.
* `full_protection` - The tamper Protection mode for current security groups.
* `allow_read_only` - The AWS cloud account operation mode. true for "Manage", false for "Readonly".
* `net_sec` - The network security configuration for the AWS cloud account.

## Import

AWS cloud account can be imported; use `<AWS CLOUD ACCOUNT ID>` as the import ID. For example:

```shell
terraform import dome9_cloudaccount_aws.test 00000000-0000-0000-0000-000000000000
```
