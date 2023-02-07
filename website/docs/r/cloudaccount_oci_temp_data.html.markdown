---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_oci_temp_data"
sidebar_current: "docs-resource-dome9-cloudaccount-oci-temp-data
description: |- Onboard OCI cloud account
---

# dome9_cloudaccount_oci_temp_data

This resource is used (alongside to 'dome9_cloudaccount_oci' resource) to onboard oci cloud accounts to Dome9. This is the first and pre-requisite step in order to
apply Dome9 features, such as compliance testing, on the account.

**The resource can be used only with 'dome9_cloudaccount_oci' resource.**

## Example Usage

Basic usage:

```hcl
resource "dome9_cloudaccount_oci_temp_data" "test" {
  name = "NAME"
  tenancy_id = "TENANCY_ID"  
  home_region = "HOME_REGION"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the OCI account in Dome9
* `tenancy_id` - (Required) The root tenancy id (root compartment from OCI).
* `home_region` - (Required) The home region (from OCI).

## Attributes Reference

* `id` - OCI account id.
* `name` - The name of the oci account in Dome9
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