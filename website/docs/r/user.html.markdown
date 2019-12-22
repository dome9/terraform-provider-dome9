---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_user"
sidebar_current: "docs-resource-dome9-user"
description: |-
  Create users in Dome9
---

# dome9_user

The User resource has methods to create and manage Dome9 users.

## Example Usage

Basic usage:

```hcl
resource "dome9_user" "user_sg" {
  email = "EMAIL"
  first_name = "FIRST_NAME"
  last_name = "LAST_NAME"
  is_sso_enabled = false
}

```

## Argument Reference

The following arguments are supported:

* `email` - (Required) user email address. 
* `first_name` - (Required) user first name. 
* `last_name` - (Required) user last name. 
* `is_sso_enabled` - (Required) indicates user has enabled SSO sign-on. 

Note: The fields  `is_owner` and `role_ids` should be set only after the user is created. So, perform two operations:
* Create user.
* Update the fields.

## Attributes Reference

* `id` - user id.
* `is_suspended` - user is suspended.
* `is_owner` - user is account owner.
* `is_super_user` - user is Super User.
* `is_auditor` - user is auditor.
* `has_api_key` - user has generated an API Key - V1 or V2.
* `has_api_key_v1` - user has generated an API Key - V1.
* `has_api_key_v2` - user has generated an API Key - V2.
* `is_mfa_enabled` - user has enabled MFA authentication.
* `role_ids` - (list) list of roles for the user.
* `iam_safe` - IAM safety details for the user support:
    * `cloud_accounts` - (list) Cloud accounts IAM supports:
        * `cloud_account_id` - cloud account id 
        * `name` - name 
        * `external_account_number` - external account number 
        * `last_lease_time` - last lease time 
        * `state` - state 
        * `iam_entities` - iam entities 
        * `iam_entities_last_lease_time` - (list) iam entities last lease time supports:
            * `iam_entity` - iam entity 
            * `last_lease_time` - last lease time 
        * `cloud_account_state` - cloud account state 
        * `iam_entity` - iam entity 

* `can_switch_role` - user can switch roles.
* `is_locked` - is locked.
* `last_login` - last login.
* `is_mobile_device_paired` - user has paired mobile device.


## Import 

The user can be imported; use `<USER ID>` as the import ID. 

For example:
```shell
terraform import dome9_user.test 000000
```