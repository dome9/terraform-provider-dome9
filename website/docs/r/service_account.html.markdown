---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_service_account"
sidebar_current: "docs-resource-dome9-service-account"
description: |-
  Create Service Account in Dome9
---

# dome9_service_account

This resource is used to create and manage Dome9 Service Account. Service Account is an account created explicitly to provide credentials and security context for a service.

## Example Usage

Basic usage:

```hcl
resource "dome9_service_account" "service_account" {
  name        = "SERVICE_ACCOUNT_NAME"
  role_ids    = []
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Dome9 service account name.
* `role_ids` - (Required) Dome9 role ids for the service account. 

## Attributes Reference

* `id` - service account id.
* `name` - service account name.
* `api_key_id` - api key.
* `api_key_secret` - secret.
* `role_ids` - service account role ids.

## Import

The service account can be imported; use `<SERVICE ACCOUNT ID>` as the import ID. 

For example:

```shell
terraform import dome9_service_account.service_account 00000
```
