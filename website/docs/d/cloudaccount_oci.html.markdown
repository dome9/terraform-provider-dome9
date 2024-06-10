---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_oci"
sidebar_current: "docs-datasource-cloudaccount-oci"
description: |-
  Get information about oci cloud account onboarded to Dome9.
---

# Data Source: dome9_cloudaccount_oci

Use this data source to get information about an oci cloud account onboarded to Dome9.

## Example Usage

```hcl
data "dome9_cloudaccount_oci" "test" {
  id = "CLOUD_ACCOUNT_ID"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The Dome9 id for the OCI account.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - OCI account id (in Dome9).
* `name` - The name of the OCI account in Dome9
* `creation_date` - Date the account was onboarded to Dome9
* `tenancy_id` - The tenancy id.
* `home_region` - The home region.
* `credentials` - Has the following arguments:
  * `user` - The user ocid.
  * `fingerprint` - Hash code of the public key.
  * `public_key` - The public key.
* `organizational_unit_id` - Organizational unit id.
* `organizational_unit_path` - Organizational unit path
* `organizational_unit_name` - Organizational unit name
* `vendor` - The cloud provider ("oci")
