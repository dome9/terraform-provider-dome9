---
layout: 'dome9'
page_title: 'Check Point CloudGuard Dome9: dome9_image_assurance_policy'
sidebar_current: 'docs-datasource-dome9-image-assurance-policy'
description: |-
    Get information about a Dome9 image assurance policy.
---

# Data Source: dome9_image_assurance_policy

Use this data source to get information about a CloudGuard image assurance policy.

## Example Usage

```hcl
data "dome9_image_assurance_policy" "test-policy" {
  id = "d9-image-assurance-policy-id"
}
```

## Argument Reference

The following arguments are supported:

-   `id` - (Required) The id for the image assurance policy in Dome9.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

-   `target_id` - Environment ID / OU ID.
-   `target_type` - Policy Type (`Environment`, `OrganizationalUnit`).
-   `ruleset_id` - Ruleset ID.
-   `notification_ids` - Notification IDs [list].
-   `admission_control_action` - Image Admission Policy action type (`Detection`, `Prevention`).
-   `admission_control_unscanned_action` - Image Admission Policy action type (`Detection`, `Prevention`).
