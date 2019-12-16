---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_gcp"
sidebar_current: "docs-resource-dome9-cloudaccount-gcp"
description: |-
  Onboard GCP cloud account
---

# dome9_cloudaccount_gcp

This resource is used to onboard GCP cloud accounts to Dome9. This is the first and pre-requisite step in order to apply  Dome9 features, such as compliance testing, on the account.

## Example Usage

Basic usage:

```hcl
resource "dome9_cloudaccount_gcp" "gcp_ca" {
  name                 = "sandbox"
  project_id           = "ID"
  private_key_id       = "PRIVATE"
  private_key          = "KEY"
  client_email         = "EMAIL@ADDRESS.COM"
  client_id            = "CID"
  client_x509_cert_url = "https://www.googleapis.com/oauth2/v1/certs"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Google account name in Dome9
* `project_id` - (Required) Project ID
* `private_key_id` - (Required) Private key ID
* `private_key` - (Required) Private key
* `client_email` - (Required) GCP client email
* `client_id` - (Required) Client id
* `client_x509_cert_url` - (Required) client_x509_cert_url
* `gsuite_user` - (Optional) The gsuite user
* `domain_name` - (Optional) The domain name
* `organizational_unit_id` - (Optional) Organizational Unit that this cloud account will be attached to

## Attributes Reference

* `id` - The ID of the GCP cloud account
* `creation_date` - creation date for project in Google.
* `vendor` - The cloud provider (gcp).
* `organizational_unit_path` - Organizational unit path.
* `organizational_unit_name` - Organizational unit name.

## Import

GCP cloud account can be imported; use `<GCP CLOUD ACCOUNT ID>` as the import ID. 

For example:

```shell
terraform import dome9_cloudaccount_gcp.test 00000000-0000-0000-0000-000000000000
```
