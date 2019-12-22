---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_user"
sidebar_current: "docs-datasource-dome9_user"
description: |-
  Get information about a user in Dome9.
---

# Data Source: dome9_user

Use this data source to get information about a user in Dome9.

## Example Usage

```hcl
data "dome9_user" "user_ds" {
  id = "ID"
}

```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The id of the user in Dome9.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `email` - user email address. 
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
        * `iam_entities` - IAM entities 
        * `iam_entities_last_lease_time` - (list) IAM entities last lease time supports:
            * `iam_entity` - IAM entity 
            * `last_lease_time` - last lease time 
        * `cloud_account_state` - cloud account state 
        * `iam_entity` - IAM entity 
* `can_switch_role` - user can switch roles.
* `is_locked` - is locked.
* `last_login` - last login.
* `is_mobile_device_paired` - user has paired mobile device.
