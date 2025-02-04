---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_continuous_compliance_policy"
sidebar_current: "docs-resource-dome9-continuous-compliance-policy"
description: |-
  Creates continuous compliance policies in Dome9
---

# dome9_continuous_compliance_policy

This  resource is used to  create and modify compliance policies in Dome9 for continuous compliance assessments. A continuous compliance policy is the combination of a Rule Bundle applied to a specific cloud account.

## Example Usage

Basic usage:

```hcl
resource "dome9_continuous_compliance_policy" "test_policy" {
  target_id    = "CLOUD ACCOUNT ID"
  ruleset_id           = 00000
  target_type  = "CLOUD ACCOUNT TYPE"
  notification_ids    = ["NOTIFICATION IDS"]
}

```

## Argument Reference

The following arguments are supported:

* `target_id` - (Required) The cloud account id.
* `ruleset_id` - (Required) The bundle id for the bundle that will be used in the policy.
* `target_type` - (Required) The cloud account provider ("Aws", "Azure", "Gcp", "Kubernetes", "OrganizationalUnit").
* `notification_ids` - (Required) The notification policy id's for the policy [list].
    
## Attributes Reference

* `id` - Id of the compliance policy.

## Limitation
The order in which notifications are added to policies may appear differently in the Terraform configuration compared to the actual resource state. Terraform may continuously rearrange the order based on the state it retrieves, but this does not affect execution.

## Import

The policy can be imported; use `<POLICY ID>` as the import ID. 

For example:

```shell
terraform import dome9_continuous_compliance_policy.test 00000000-0000-0000-0000-000000000000
```
