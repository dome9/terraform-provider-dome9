---
layout: "dome9"
page_title: "Check Point Cloud Guard Dome9: dome9_cloudaccount_azure"
sidebar_current: "docs-datasource-cloudaccount-azure"
description: |-
  Get information on Azure cloud account.
---

# Data Source: dome9_cloudaccount_azure

Use this data source to get information about Azure cloud account.

## Example Usage

```hcl
data "dome9_cloudaccount_azure" "test" {
  account_id         = "my-dome9-id"
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) Account id in Dome9.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Account name (in Dome9).
* `subscription_id` - Azure subscription id for account.
* `tenant_id` - Azure tenant id.
* `operation_mode` - Dome9 operation mode for the Azure account (Read-Only or Managed).
* `creation_date` - Date Azure account was onboarded to a Dome9 account.
* `organizational_unit_id` - Organizational unit id.
* `organizational_unit_path` - Organizational unit path.
* `organizational_unit_name` - Organizational unit name.
* `vendor` - The cloud provider (Azure).
