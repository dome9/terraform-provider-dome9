---
layout: "dome9"
page_title: "Check Point Cloud Guard Dome9: dome9_cloudaccount_gcp"
sidebar_current: "docs-resource-dome9-cloudaccount-gcp"
description: |-
  Manages GCP cloud account.
---

# dome9_cloudaccount_gcp

The GoogleCloudAccount resource has methods to onboard Google cloud accounts to a Dome9 account, and to get details for a Google accounts Dome9.

## Example Usage

Basic usage:

```hcl
resource "dome9_cloudaccount_gcp" "gcp_ca" {
  name = "sandbox"

  service_account_credentials = {
    auth_provider_x509_cert_url = "https://www.googleapis.com/oauth2/v1/certs"
    auth_uri                    = "https://accounts.google.com/o/oauth2/auth"
    client_email                = "EMAIL@ADDRESS.COM"
    client_id                   = "CID"
    client_x509_cert_url        = "CERT URL"
    private_key                 = "KEY"
    private_key_id              = "PRIVATE"
    project_id                  = "ID"
    token_uri                   = "https://oauth2.googleapis.com/token"
    type                        = "service_account"
  }
  gsuite_user = "GSUITE USER"
  domain_name = "DOMAIN NAME"
  organizational_unit_id = "ORGANIZATIONAL UNIT ID"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Google account name in Dome9.
* `gsuite_user` - (Optional) The gsuite user.
* `service_account_credentials` - (Required) The service account JSON block (from the GCP console).
* `domain_name` - (Optional) The domain name.
* `organizational_unit_id` - (Optional) Organizational unit id, Will apply on update state.

### Service Account Credentials

The `service_account_credentials` block supports: 

* `type` - (Required) type. i.e "service_account"
* `project_id` - (Required) Project ID
* `private_key_id` - (Required) Private key ID
* `private_key` - (Required) Private key
* `client_email` - (Required) GCP client email
* `client_id` - (Required) Client id
* `auth_uri` - (Required) Auth URI. i.e "https://accounts.google.com/o/oauth2/auth"
* `token_uri` - (Required) Token URI. i.e "https://oauth2.googleapis.com/token"
* `auth_provider_x509_cert_url` - (Required) auth_provider_x509_cert_url. i.e "https://www.googleapis.com/oauth2/v1/certs"
* `client_x509_cert_url` - (Required) client_x509_cert_url

## Attributes Reference

* `id` - The ID of the GCP cloud account.
* `creation_date` - creation date for project in Google.
* `vendor` - The cloud provider (gcp).
* `organizational_unit_path` - Organizational unit path.
* `organizational_unit_name` - Organizational unit name.

## Import

GCP cloud account can be imported; use `<GCP CLOUD ACCOUNT ID>` as the import ID. For example:

```shell
terraform import dome9_cloudaccount_gcp.test 00000000-0000-0000-0000-000000000000
```
