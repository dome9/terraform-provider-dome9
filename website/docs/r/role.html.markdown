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
  name = "ROLE_NAME"
  description = "ROLE_DESC"

  permissions {
    access = ["string"]
    manage = ["string"]
    rulesets = ["string"]
    notifications = ["string"]
    policies = ["string"]
    alert_actions = ["string"]
    create = ["string"]
    view = ["string"]
    on_boarding = ["string"]
    cross_account_access = ["string"]
  }
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Dome9 role name.
* `description` - (Required) Dome9 role description. 
* `permissions` - (Optional) Permissions of the role.

### permissions 

The `permissions` supports the following arguments:
    
* `access` - (Optional) Access permission list (list of SRL).
* `manage` - (Optional) Manage permission list (list of SRL).
* `rulesets` - (Optional) Compliance permission list (list of SRL).
* `notifications` - (Optional) Compliance permission list (list of SRL).
* `policies` - (Optional) Compliance permission list (list of SRL).
* `alert_actions` - (Optional) Compliance permission list (list of SRL).
* `create` - (Optional) Create permission list (list of SRL).
* `view` - (Optional) View permission list (list of SRL).
* `on_boarding` - (Optional) View permission SRL.
* `cross_account_access` - (Optional) Cross account access.

* Note: to create a role, create it with no permisssion, then updated it with the desired permissions.
    
    To understand the roles/permissions refer to to following: `https://gist.github.com/froyke/70ea7d91d01c3ba765a604edf910ebd5`

## Import

IP role can be imported; use `<ROLE ID>` as the import ID. 

For example:

```shell
terraform import dome9_role.role_rs 00000
```
