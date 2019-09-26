---
layout: "dome9"
page_title: "Check Point Cloud Guard Dome9: dome9_cloudaccount_aws"
sidebar_current: "docs-datasource-dome9-cloudaccount-aws"
description: |-
  Get information on AWS cloud account.
---

# Data Source: dome9_cloudaccount_aws

Use this data source to get information about AWS cloud account.

## Example Usage

```hcl
data "dome9_cloudaccount_aws" "test" {
  id = "my-dome9-id"
}

```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The Dome9 id

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `vendor` - The cloud provider (AWS).
* `name` - The cloud account name.
* `external_account_number` - The AWS account number.
* `error` - Credentials error status.
* `is_fetching_suspended` - Fetching suspending status.
* `creation_date` - Account creation date.
* `full_protection` - The tamper Protection mode for current security groups.
* `allow_read_only` - The AWS cloud account operation mode. true for "Manage", false for "Readonly".
* `net_sec` - The network security configuration for the AWS cloud account.
