---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_alibaba"
sidebar_current: "docs-datasource-cloudaccount-alibaba"
description: |-
  Get information about Alibaba cloud account onboarded to Dome9.
---

# Data Source: dome9_cloudaccount_alibaba

Use this data source to get information about an Alibaba cloud account onboarded to Dome9.

## Example Usage

```hcl
data "dome9_cloudaccount_alibaba" "test" {
  id = "d9-alibaba-cloud-account-id"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The Dome9 id for the Alibaba account.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Account name (in Dome9).
* `alibaba_account_id` - Alibaba account id.
* `creation_date` - Date Alibaba account was onboarded to a Dome9 account.
* `credentials` - Has the following arguments:
  * `access_key` - The access key for the Alibaba account.
* `organizational_unit_id` - Organizational unit id.
* `organizational_unit_path` - Organizational unit path.
* `organizational_unit_name` - Organizational unit name.
* `vendor` - The cloud provider (Alibaba).
