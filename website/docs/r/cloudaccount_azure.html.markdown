---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_azure"
sidebar_current: "docs-resource-dome9-cloudaccount-azure"
description: |-
  Onboard Azure cloud account
---

# dome9_cloudaccount_azure

This resource is used to onboard Azure cloud accounts to Dome9. This is the first and pre-requisite step in order to apply  Dome9 features, such as compliance testing, on the account.

## Example Usage

Basic usage:

```hcl
resource "dome9_cloudaccount_azure" "test" {
  name            = "NAME"
  operation_mode  = "OPERATION MODE"
  subscription_id = "SUBSCRIPTION ID"
  tenant_id       = "TENANT ID"
  credentials = {
    client_id       = "CLIENT ID"
    client_password = "CLIENT PASSWORD"
  }
  organizational_unit_id = "ORGANIZATIONAL UNIT ID"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Azure account in Dome9
* `subscription_id` - (Required) The Azure subscription id for account
* `tenant_id` - (Required) The Azure tenant id
* `operation_mode` - (Required) Dome9 operation mode for the Azure account ("Read-Only" or "Managed")
* `credentials` - (Required) Azure account credentials
* `organizational_unit_id` - (Optional) Organizational Unit that this cloud account will be attached to

### Credentials

The `credentials` block supports: 

* `client_id` - (Required) Azure account id
* `client_password` - (Required) Password for account

## Attributes Reference

* `id` - The ID of the Azure cloud account
* `vendor` - The cloud provider ("Azure")
* `creation_date` - Date the account was onboarded to Dome9
* `organizational_unit_path` - Organizational unit path
* `organizational_unit_name` - Organizational unit name

## Import

Azure cloud account can be imported; use `<Azure CLOUD ACCOUNT ID>` as the import ID. 

For example:

```shell
terraform import dome9_cloudaccount_Azure.test 00000000-0000-0000-0000-000000000000
```
