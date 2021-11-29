---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_azure"
sidebar_current: "docs-datasource-cloudaccount-azure"
description: |-
  Get information about Azure cloud account onboarded to Dome9.
---

# Data Source: dome9_cloudaccount_azure

Use this data source to get information about an Azure cloud account onboarded to Dome9.

## Example Usage

```hcl
data "dome9_cloudaccount_azure" "test" {
  id = "d9-azure-cloud-account-id"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The Dome9 id for the Azure account.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Account name (in Dome9).
* `subscription_id` - Azure subscription id for account.
* `tenant_id` - Azure tenant id.
* `operation_mode` - Dome9 operation mode for the Azure account ("Read" or "Manage")
* `vendor` - The cloud provider (Azure).
* `creation_date` - Date Azure account was onboarded to a Dome9 account.
* `organizational_unit_id` - Organizational unit id.
* `organizational_unit_path` - Organizational unit path.
* `organizational_unit_name` - Organizational unit name.
