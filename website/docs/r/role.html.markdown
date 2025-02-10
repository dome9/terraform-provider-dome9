---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_role"
sidebar_current: "docs-resource-dome9-role"
description: |-
  Create role in Dome9
---

# dome9_role

The Role resource is used to create and manage CloudGuard roles. Roles are used to manage access permissions for CloudGuard users.

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

Granting "view" permissions for All System Resources:

```hcl
resource "dome9_role" "role_rs" {
  name        = "ROLE_NAME"
  description = "ROLE_DESC"
  
  view {}  // Grants "view" permissions on All System Resources
}
```

Granting "manage" permissions for All System Resources:

```hcl
resource "dome9_role" "role_rs" {
  name        = "ROLE_NAME"
  description = "ROLE_DESC"
  
  manage {}  // Grants "manage" permissions on All System Resources
}
```


## Argument Reference

The following arguments are supported:

* `name` - (Required) CloudGuard role name.
* `description` - (Required) CloudGuard role description. 
* `permit_rulesets` - Is permitted permit rulesets (Optional) .
* `permit_notifications` - Is permitted permit notifications (Optional) .
* `permit_policies` - Is permitted permit policies (Optional) .
* `permit_alert_actions` - Is permitted permit alert actions (Optional) .
* `permit_on_boarding` - Is permitted permit onboarding (Optional)  .
* `cross_account_access` - (Optional) Cross account access.
* `create` - (Optional) Create permission list.
* `access` - (Optional) Access permission list ([SRL](#SRL) Type).
* `view` - (Optional) View permission list ([SRL](#SRL) Type).
* `manage` - (Optional) Manage permission list ([SRL](#SRL) Type).

### SRL 
* `type` - (Optional) Accepted values: AWS, Azure, GCP, OrganizationalUnit, CloudGuardResources, CodeSecurityResources.
* `main_id` - (Optional) Cloud Account, Organizational Unit ID or CodeSecurity Access Level (Admin, Member).
* `region` - (Optional) Accepted values: "us_east_1", "us_west_1", "eu_west_1", "ap_southeast_1", "ap_northeast_1", "us_west_2", "sa_east_1", "ap_southeast_2", "eu_central_1", "ap_northeast_2", "ap_south_1", "us_east_2", "ca_central_1", "eu_west_2", "eu_west_3", "eu_north_1", "il_central_1", "ca_west_1", "mx_central_1", "ap_sotheast_5", "ap_sotheast_7".
* `security_group_id` - (Optional) AWS Security Group ID.
* `traffic` - (Optional) Accepted values: "All Traffic", "All Services".


### Note
* To create a role, create it with no permissions, then updated it with the desired permissions.
* In order to grant "All System Resources" permissions you must specify an empty block following the permission
access level, manage or view, in your Terraform configuration. This instructs the provider to apply the permission access level to ‘All System Resources’.


For more about [Roles and Permissions](https://sc1.checkpoint.com/documents/CloudGuard_Dome9/Documentation/Settings/Users-Roles.htm?tocpath=Settings%20%7C_____4)


## Import

IP role can be imported; use `<ROLE ID>` as the import ID. 

For example:

```shell
terraform import dome9_role.role_rs 00000
```
