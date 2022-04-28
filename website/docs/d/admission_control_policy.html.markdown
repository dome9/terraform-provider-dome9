---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_admission_control_policy"
sidebar_current: "docs-datasource-dome9-admission-control-policy"
description: |-
  Get information about a Dome9 Admission Control policy.
---

# Data Source: dome9_admission_control_policy

Use this data source to get information about a Dome9 admission control policy.

## Example Usage

```hcl
data "dome9_admission_control_policy" "test-policy" {
  id = dome9_admission_control_policy.test-policy.id
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The id for the cloud account in Dome9. 

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `target_id` - Cloud account id in Dome9.
* `target_type` - Kubernetes Account type (`Environment`, `OrganizationalUnit`).
* `ruleset_id` - Rule bundle ID.
* `notification_ids` - Dome9 Notifications IDs [list] .
* `action` - Policy action type (`Detection`, `Prevention`).