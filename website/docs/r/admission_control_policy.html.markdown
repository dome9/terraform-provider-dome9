---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_admission_control_policy"
sidebar_current: "docs-resource-dome9-admission-control-policy"
description: |-
  Creates admission control policies in Dome9
---

# dome9_admission_control_policy

This resource is used to create and modify admission control policy in CloudGuard for Kubernetes. An admission control policy is the combination of a Ruleset applied to a specific Kubernetes environment with specific action.

## Example Usage

Basic usage:

```hcl
resource "dome9_admission_control_policy" "test_ac_policy" {
  target_id    = "Environment ID"
  ruleset_id   = 00000
  target_type  = "Environment or OrganizationalUnit"
  notification_ids    = ["NOTIFICATION IDS"]
  action       = "Detection"
}

```

## Argument Reference

The following arguments are supported:

* `target_id` - (Required) The kubernetes environment id / organizational unit id.
* `ruleset_id` - (Required) The bundle id for the bundle that will be used in the policy.
* `target_type` - (Required) The admission control policy type ("Environment", "OrganizationalUnit").
* `notification_ids` - (Required) The notification policy id's for the policy [list].
* `action` - (Required) The admission control policy action ("Prevention", "Detection").
    
## Attributes Reference

* `id` - Id of the admission control policy.

## Import

The policy can be imported; use `<POLICY ID>` as the import ID. 

For example:

```shell
terraform import dome9_admission_control_policy.test 00000000-0000-0000-0000-000000000000
```
