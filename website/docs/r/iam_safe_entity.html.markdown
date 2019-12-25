---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_iam_safe_entity"
sidebar_current: "docs-resource-dome9-iam-safe-entity"
description: |-
    Protect cloud accounts that are managed by Dome9. Control access to them with targeted short-term authorizations (involving the Dome9 mobile app).
---

# dome9_iam_safe_entity

Protect cloud accounts that are managed by Dome9. Control access to them with targeted short-term authorizations (involving the Dome9 mobile app).

## Example Usage

Basic usage:

```hcl
resource "dome9_iam_safe_entity" "dome9_iam_safe_entity_re" {
  protection_mode           = "ProtectWithElevation"
  entity_type               = "User"
  entity_name               = "ENTITY_NAME"
  aws_cloud_account_id      = "00000000-0000-0000-0000-000000000000"
  dome9_users_id_to_protect = ["000000", "111111"]
}

```

## Argument Reference

The following arguments are supported:
 
* `protection_mode` - (Required) Protection mode; can be  "Protect", "ProtectWithElevation".
* `entity_type` - (Required) Entity type to protect; can be  "User", "Role".
* `aws_cloud_account_id` - (Required) AWS cloud account id to protect.
* `entity_name` - (Required) AWS IAM user or role name to protect.
* `dome9_users_id_to_protect` - (Optional) When ProtectWithElevation mode selected, dome9 users ids must be provided.

* Note: To following filed can be updated:
    * `protection_mode`: Switch between `Protect` to `ProtectWithElecation` mode.
    * `dome9_users_id_to_protect`: Update the dome9 users list that can evaluate the aws users or roles. Empty list with switch it from `Protect` to `ProtectWithElevation` mode. 

## Attributes Reference

* `state` - Can be one of the following: `Unattached`, `Attached` or `Restricted`.
* `attached_dome9_users` - List of users in protect with elevation mode.
* `exists_in_aws` - Is exist in aws.
* `arn` - Role or User arn.
