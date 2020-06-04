---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_user"
sidebar_current: "docs-resource-dome9-user"
description: |-
  Create user in Dome9
---

# dome9_user

The User resource has methods to create and manage Dome9 users.

## Example Usage

Basic usage:

```hcl
resource "dome9_users" "users_sg" {
  email                = "EMAIL"
  first_name           = "FIRST_NAME"
  last_name            = "LAST_NAME"
  is_sso_enabled       = false
  permit_notifications = false
  permit_rulesets      = false
  permit_policies      = false
  permit_alert_actions = false
  permit_on_boarding   = false
  create               = []
  cross_account_access = []
}

```

## Argument Reference

The following arguments are supported:

* `email` - (Required) user email. 
* `first_name` - (Required) userfirst name. 
* `last_name` - (Required) user last name. 
* `is_sso_enabled` - (Required) user has enabled SSO sign-on. 
* permission fields:
    * `permit_rulesets` - Is permitted permit rulesets (Optional) .
    * `permit_notifications` - Is permitted permit notifications (Optional) .
    * `permit_policies` - Is permitted permit policies (Optional) .
    * `permit_alert_actions` - Is permitted permit alert actions (Optional) .
    * `permit_on_boarding` - Is permitted permit on boarding (Optional)  .
    * `cross_account_access` - (Optional) Cross account access.
    * `create` - (Optional) Create permission list.
    * `access` - (Optional) Access permission list ([SRL](#SRL) Type).
    * `view` - (Optional) View permission list ([SRL](#SRL) Type).
    * `manage` - (Optional) Manage permission list ([SRL](#SRL) Type).

Note: The filed that can be updated are `is_owner`, `role_ids` and "permission" fields. The update occur in two steps:
* Create user.
* Then update the desired field.

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
* `is_mobile_device_paired` - user has paired mobile device.


## Import 

The user can be imported; use `<USER ID>` as the import ID. 

For example:
```shell
terraform import dome9_user.test 000000
```