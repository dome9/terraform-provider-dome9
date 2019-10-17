---
layout: "dome9"
page_title: "Check Point Cloud Guard Dome9: dome9_cloudaccount_azure"
sidebar_current: "docs-resource-dome9-cloudaccount-azure"
description: |-
  Manages Azure cloud account.
---

# dome9_cloudaccount_azure

The AzureCloudAccounts resource has methods to onboard Azure cloud accounts to Dome9 and to manage some of their settings.

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

* `name` - (Required) The Azure account number.
* `subscription_id` - (Required) Azure subscription id for account.
* `tenant_id` - (Required) Azure tenant id.
* `operation_mode` - (Required) Dome9 operation mode for the Azure account (Read-Only or Managed).
* `credentials` - (Required) Azure account credentials.
* `organizational_unit_id` - Organizational unit id, will apply on update state.

### Credentials

The `credentials` block supports: 

* `client_id` - (Required) Azure account id.
* `client_password` - (Required) Password for account.

## Attributes Reference

* `id` - The ID of the Azure cloud account.
* `vendor` - The cloud provider (Azure).
* `creation_date` - Date Azure account was onboarded to a Dome9 account.
* `organizational_unit_path` - Organizational unit path.
* `organizational_unit_name` - Organizational unit name.

## Import

Azure cloud account can be imported; use `<Azure CLOUD ACCOUNT ID>` as the import ID. For example:

```shell
terraform import dome9_cloudaccount_Azure.test 00000000-0000-0000-0000-000000000000
```
