---
layout: 'dome9'
page_title: 'Check Point CloudGuard Dome9: dome9_image_assurance_policy'
sidebar_current: 'docs-resource-dome9-admission-control-policy'
description: |-
    Creates image assurance policies in Dome9
---

# dome9_image_assurance_policy

This resource is used to create and modify image assurance policy in CloudGuard for Kubernetes. An image assurance policy is the combination of a Ruleset applied to a specific Kubernetes environment with specific action.

## Example Usage

Basic usage:

```hcl
resource "dome9_image_assurance_policy" "test_ia_policy" {
  target_id    = "Environment ID"
  ruleset_id   = 00000
  target_type  = "Environment"
  notification_ids    = ["NOTIFICATION IDS"]
  admission_control_action       = "Detection"
  admission_control_unscanned_action       = "Detection"
}

```

## Argument Reference

The following arguments are supported:

-   `target_id` - (Required) The kubernetes environment id / organizational unit id.
-   `ruleset_id` - (Required) The bundle id for the bundle that will be used in the policy.
-   `target_type` - (Required) The imageassurance policy type ("Environment", "OrganizationalUnit").
-   `notification_ids` - (Required) The notification policy id's for the policy [list].
-   `admission_control_action` - (Required) The imageassurance policy action ("Prevention", "Detection").
-   `admission_control_unscanned_action` - (Required) The imageassurance policy action ("Prevention", "Detection").

## Attributes Reference

-   `id` - Id of the imageassurance policy.

## Import

The policy can be imported; use `<POLICY ID>` as the import ID.

For example:

```shell
terraform import dome9_image_assurance_policy.test 00000000-0000-0000-0000-000000000000
```
