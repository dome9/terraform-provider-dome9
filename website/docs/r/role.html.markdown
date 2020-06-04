---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_role"
sidebar_current: "docs-resource-dome9-role"
description: |-
  Create role in Dome9
---

# dome9_role

The Role resource is used to create and manage Dome9 roles. Roles are used to manage access permissions for Dome9 users.

## Example Usage

Basic usage:

```hcl
resource "dome9_role" "role_rs" {
  name        = "ROLE_NAME"
  description = "ROLE_DESC"
  access {
    type              = "AWS"
    main_id           = "MAIN_ID"
    region            = "us_east_1"
    security_group_id = "SECURITY_GROUP_ID"
    traffic           = "All Traffic"
  }
  access {
    type    = "OrganizationalUnit"
    main_id = "00000000-0000-0000-0000-000000000000"
  }

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

* `name` - (Required) Dome9 role name.
* `description` - (Required) Dome9 role description. 
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
* `region` - (Optional) Accepted values: "us_east_1", "us_west_1", "eu_west_1", "ap_southeast_1", "ap_northeast_1", "us_west_2", "sa_east_1", "ap_southeast_2", "eu_central_1", "ap_northeast_2", "ap_south_1", "us_east_2", "ca_central_1", "eu_west_2", "eu_west_3", "eu_north_1".
* `security_group_id` - (Optional) AWS Security Group ID.
* `traffic` - (Optional) Accepted values: "All Traffic", "All Services".

* Note: to create a role, create it with no permissions, then updated it with the desired permissions.
    
    To understand the roles/permissions [CLICK HERE](https://gist.github.com/froyke/70ea7d91d01c3ba765a604edf910ebd5).

## Import

IP role can be imported; use `<ROLE ID>` as the import ID. 

For example:

```shell
terraform import dome9_role.role_rs 00000
```
