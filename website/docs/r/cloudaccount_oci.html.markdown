---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_oci"
sidebar_current: "docs-resource-dome9-cloudaccount-oci"
description: |- Onboard OCI cloud account
---

# dome9_cloudaccount_oci

This resource is used to onboard oci cloud accounts to Dome9. This is the first and pre-requisite step in order to
apply Dome9 features, such as compliance testing, on the account.

**The resource can be used only after 'dome9_cloudaccount_oci_temp_data' resource is created.**

## Example Usage

Basic usage:

```hcl
resource "dome9_cloudaccount_oci" "test" {
  tenancy_id = "TENANCY_ID"  
  user_ocid  = "USER_OCID"
}
```

Advanced full usage (you must use these OCI resources, so we can get all permissions. Also, the OCI provider must be
used to create these resources):

```hcl
resource "dome9_cloudaccount_oci_temp_data" "test" {
  name = "NAME"
  tenancy_id = "TENANCY_ID"
  home_region = "HOME_REGION"
}

resource "dome9_cloudaccount_oci" "test" {
  tenancy_id = dome9_cloudaccount_oci_temp_data.test.tenancy_id
  user_ocid  = oci_identity_user.user.id
}

resource "oci_identity_user" "user" {
    name           = "CloudGuard-user"
    description    = "CloudGuard Onboarding"
    compartment_id = dome9_cloudaccount_oci_temp_data.test.tenancy_id
}

resource "oci_identity_group" "group" {
    name           = "CloudGuard-group"
    description    = "CloudGuard Onboarding"
    compartment_id = oci_identity_user.user.compartment_id
}

resource "oci_identity_policy" "policy" {
    name           = "CloudGuard-policy"
    description    = "CloudGuard Onboarding"
    compartment_id = oci_identity_user.user.compartment_id
    statements = [
              "allow group ${oci_identity_group.group.name} to inspect all-resources in tenancy",
              "allow group ${oci_identity_group.group.name} to read instances in tenancy",
              "allow group ${oci_identity_group.group.name} to read load-balancers in tenancy",
              "allow group ${oci_identity_group.group.name} to read buckets in tenancy",
              "allow group ${oci_identity_group.group.name} to read nat-gateways in tenancy",
              "allow group ${oci_identity_group.group.name} to read public-ips in tenancy",
              "allow group ${oci_identity_group.group.name} to read file-family in tenancy",
              "allow group ${oci_identity_group.group.name} to read instance-configurations in tenancy",
              "allow group ${oci_identity_group.group.name} to read network-security-groups in tenancy",
              "allow group ${oci_identity_group.group.name} to read resource-availability in tenancy",
              "allow group ${oci_identity_group.group.name} to read audit-events in tenancy",
              "allow group ${oci_identity_group.group.name} to read users in tenancy",
              "allow group ${oci_identity_group.group.name} to read vss-family in tenancy",       
              "allow group ${oci_identity_group.group.name} to read usage-budgets in tenancy" ,
              "allow group ${oci_identity_group.group.name} to read usage-reports in tenancy",
              "allow group ${oci_identity_group.group.name} to read data-safe-family in tenancy"
            ]
}

resource "oci_identity_user_group_membership" "user_group_membership" {
    group_id = oci_identity_group.group.id
    user_id = oci_identity_user.user.id
}

resource "oci_identity_api_key" "api_key" {
    user_id = oci_identity_user.user.id
    key_value = dome9_cloudaccount_oci_temp_data.test.public_key
}
```

## Argument Reference

The following arguments are supported:

* `tenancy_id` - (Required) The tenancy id.
* `user_ocid` - (Required) The user ocid.
* `organizational_unit_id` - (optional) Organizational unit id.

## Attributes Reference

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


## Import

oci cloud account can be imported; use `<OCI CLOUD ACCOUNT ID>` as the import ID.

For example:

```shell
terraform import dome9_cloudaccount_oci.test 00000000-0000-0000-0000-000000000000
```
