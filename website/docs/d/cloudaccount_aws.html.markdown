---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_AWS"
sidebar_current: "docs-datasource-dome9-cloudaccount-AWS"
description: |-
  Get information about  AWS cloud account onboarded to Dome9.
---

# Data Source: dome9_cloudaccount_AWS

Use this data source to get information about an AWS cloud account onboarded to Dome9.

## Example Usage

```hcl
data "dome9_cloudaccount_aws" "test" {
  id = "d9-AWS-cloud-account-id"
}

```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The Dome9  id for the AWS account 

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `vendor` - The cloud provider ("aws/awsgov").
* `name` - The cloud account name in Dome9.
* `external_account_number` - The AWS account number.
* `error` - Credentials error status.
* `is_fetching_suspended` - Fetching suspending status.
* `creation_date` - Date account was onboarded to Dome9.
* `full_protection` - The tamper Protection mode for current security groups.
* `allow_read_only` - The AWS cloud account operation mode. true for "Manage", false for "Readonly".
* `net_sec` - The network security configuration for the AWS cloud account.
* `organizational_unit_id` - Organizational unit id.
* `IAM_safe` - IAM safe entity details
    * `AWS_group_ARN` - AWS group ARN  
    * `AWS_policy_ARN` - AWS policy ARN  
    * `mode` - Mode  
    * `restricted_IAM_entities` - Restricted IAM safe entities which has the following:  
		* `roles_ARNs` - Restricted IAM safe entities roles ARNs
		* `users_ARNs` - Restricted IAM safe entities users ARNs
