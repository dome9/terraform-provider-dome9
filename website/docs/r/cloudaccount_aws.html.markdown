---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_aws"
sidebar_current: "docs-resource-dome9-cloudaccount-aws"
description: |-
  Onboard AWS cloud account
---

# dome9_cloudaccount_aws

This resource is used to onboard AWS cloud accounts to Dome9. This is the first and pre-requisite step in order to apply  Dome9 features, such as compliance testing, on the account.

## Example Usage

Basic usage:

```hcl
resource "dome9_cloudaccount_aws" "test" {
  name  = "ACCOUNT NAME"
 
  credentials  {
    arn    = "ARN"
    secret = "SECRET"
    type   = "RoleBased"
  }
  net_sec {
    regions = [
	 	region = "us-east-1"
		new_group_behavior = "ReadOnly"
	 ]	
	 
	}
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of AWS account in Dome9
* `credentials` - (Required) The information needed for Dome9 System in order to connect to the AWS cloud account


### Credentials

`credentials` has the following arguments:

* `arn` - (Required) AWS Role ARN (to be assumed by Dome9)
* `secret` - (Required) The AWS role External ID (Dome9  will have to use this secret in order to assume the role)
* `type` - (Required) The cloud account onboarding method. Set to "RoleBased".

### Network security configuration

`net_sec` has the these arguments:

* `Regions` - (Required) list of the supported regions, and their configuration:
    * `new_group_behavior` - (Required) The network security configuration. Select "ReadOnly", "FullManage", or "Reset".
    * `region` - (Required) AWS region, in AWS format (e.g., "us-east-1")

## Attributes Reference

* `vendor` - The cloud provider ("AWS").
* `external_account_number` - The AWS account number.
* `is_fetching_suspended` - Fetching suspending status.
* `creation_date` - Date the account was onboarded to Dome9.
* `full_protection` - The protection mode for existing security groups in the account.
* `allow_read_only` - The AWS cloud account operation mode. true for "Full-Manage", false for "Readonly".
* `net_sec` - The network security configuration for the AWS cloud account. If not given, sets to default value.

## Import

AWS cloud account can be imported; use `<AWS CLOUD ACCOUNT ID>` as the import ID. 

For example:

```shell
terraform import dome9_cloudaccount_aws.test 00000000-0000-0000-0000-000000000000
```
