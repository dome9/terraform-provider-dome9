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

* `email` - user email. 
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

### SRL 
* `type` - (Optional) Accepted values: AWS, Azure, GCP, OrganizationalUnit.
* `main_id` - (Optional) Cloud Account or Organizational Unit ID.
* `region` - (Optional) Accepted values: "us_east_1", "us_west_1", "eu_west_1", "ap_southeast_1", "ap_northeast_1", "us_west_2", "sa_east_1", "ap_southeast_2", "eu_central_1", "ap_northeast_2", "ap_south_1", "us_east_2", "ca_central_1", "eu_west_2", "eu_west_3", "eu_north_1", "ap_east_1", "me_south_1", "af_south_1", "eu_south_1", "ap_northeast_3", "me_central_1", "ap-south-2", "ap-southeast-3", "ap-southeast-4", "eu-central-2", "eu-south-2".
* `security_group_id` - (Optional) AWS Security Group ID.
* `traffic` - (Optional) Accepted values: "All Traffic", "All Services".
