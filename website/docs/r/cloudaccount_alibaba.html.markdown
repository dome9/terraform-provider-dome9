---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_alibaba"
sidebar_current: "docs-resource-dome9-cloudaccount-alibaba"
description: |- Onboard Alibaba cloud account
---

# dome9_cloudaccount_alibaba

This resource is used to onboard Alibaba cloud accounts to Dome9. This is the first and pre-requisite step in order to
apply Dome9 features, such as compliance testing, on the account.

## Example Usage

Basic usage:

```hcl
resource "dome9_cloudaccount_alibaba" "test" {
  name        = "NAME"
  credentials = {
    access_key    = "ACCESS_KEY"
    access_secret = "ACCESS_SECRET"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Alibaba account in Dome9
* `credentials` - (Required) Has the following arguments:
    * `access_key` - (Required) The access key for the Alibaba account.
    * `access_secret` - (Required) The access secret for the Alibaba account.
* `organizational_unit_id` - (optional) Organizational unit id.

## Attributes Reference

* `name` - The name of the Alibaba account in Dome9
* `vendor` - The cloud provider ("Alibaba")
* `alibaba_account_id` - Alibaba account id.
* `creation_date` - Date the account was onboarded to Dome9
* `organizational_unit_path` - Organizational unit path
* `organizational_unit_name` - Organizational unit name
* `organizational_unit_id` - Organizational unit id.

## Import

Alibaba cloud account can be imported; use `<Alibaba CLOUD ACCOUNT ID>` as the import ID.

For example:

```shell
terraform import dome9_cloudaccount_alibaba.test 00000000-0000-0000-0000-000000000000
```
