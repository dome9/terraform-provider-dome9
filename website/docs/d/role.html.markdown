---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_role"
sidebar_current: "docs-datasource-dome9-role"
description: |-
  Get information about a role in Dome9.
---

# Data Source: dome9_role

Use this data source to get information about a role in CloudGuard.

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
* `region` - (Optional) Accepted values: "us_east_1", "us_west_1", "eu_west_1", "ap_southeast_1", "ap_northeast_1", "us_west_2", "sa_east_1", "ap_southeast_2", "eu_central_1", "ap_northeast_2", "ap_south_1", "us_east_2", "ca_central_1", "eu_west_2", "eu_west_3", "eu_north_1", "ap_east_1", "me_south_1", "af_south_1", "eu_south_1", "ap_northeast_3", "me_central_1", "ap_south_2", "ap_southeast_3", "ap_southeast_4", "eu_central_2", "eu_south_2", "il_central_1", "ca_west_1".
* `security_group_id` - (Optional) AWS Security Group ID.
* `traffic` - (Optional) Accepted values: "All Traffic", "All Services".
