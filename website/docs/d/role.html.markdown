---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_role"
sidebar_current: "docs-datasource-dome9-role"
description: |-
  Get information about a role in Dome9.
---

# Data Source: dome9_role

Use this data source to get information about a role in Dome9.

## Example Usage

```hcl
data "dome9_role" "test" {
  id = "d9-role-id"
}

```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The id of the role list in Dome9.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

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
