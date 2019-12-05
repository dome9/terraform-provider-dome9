---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_users"
sidebar_current: "docs-resource-dome9-users"
description: |-
  Create user in Dome9
---

# dome9_users

The User resource has methods to create and manage Dome9 users.

## Example Usage

Basic usage:

```hcl
resource "dome9_users" "users_sg" {
  email = "EMAIL"
  first_name = "FIRST_NAME"
  last_name = "LAST_NAME"
  is_sso_enabled = false
}

```

## Argument Reference

The following arguments are supported:

* `email` - (Required) user email. 
* `first_name` - (Required) userfirst name. 
* `last_name` - (Required) user last name. 
* `is_sso_enabled` - (Required) user has enabled SSO sign-on. 

## Attributes Reference

* `id` - user id.
* `is_suspended` - user is suspended.
* `is_owner` - user is account owner.
* `is_super_user` - user is Super User.
* `is_auditor` - user is auditor user.
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
* `permissions` - User direct permissions.
    * `access` - (list) access
    * `manage` - (list) manage
    * `rulesets` - (list) rulesets
    * `notifications` - (list) notifications
    * `policies` - (list) policies
    * `alert_actions` - (list) alert_actions
    * `create` - (list) create
    * `view` - (list) view
    * `on_boarding` - (list) on_boarding
    * `cross_account_access` - (list) cross_account_access
* `calculated_permissions` - calculated_permissions User Id supports:
    * `access` - (list) access
    * `manage` - (list) manage
    * `rulesets` - (list) rulesets
    * `notifications` - (list) notifications
    * `policies` - (list) policies## Import
    * `alert_actions` - (list) alert_actions
    * `create` - (list) createUser can be imported; use `<USER ID>` as the import ID. 
    * `view` - (list) view
    * `on_boarding` - (list) on_boardingFor example:
    * `cross_account_access` - (list) cross_account_access
* `is_mobile_device_paired` - user has paired mobile device.


## Import 

The user can be imported; use `<USER ID>` as the import ID. 

For example:
```shell
terraform import dome9_users.test 000000
```