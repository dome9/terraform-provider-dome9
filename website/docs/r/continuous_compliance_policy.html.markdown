---
layout: "dome9"
page_title: "Check Point Cloud Guard Dome9: dome9_continuous_compliance_policy"
sidebar_current: "docs-resource-dome9-continuous-compliance-policy"
description: |-
  Manages continuous compliance policy.
---

# dome9_continuous_compliance_policy

The ContinuousCompliancePolicy resource has methods to create and modify compliance policies for continuous compliance assessments. A continuous compliance policy is the combination of a Rule Bundle applied to a specific cloud account.
With continuous compliance, compliance policies are assessed continuously and autonomously, and the results are issued to designated recipients as emails or SNS notifications, according to notification policies.

## Example Usage

Basic usage:

```hcl
resource "dome9_continuouscompliance_policy" "test_policy" {
  cloud_account_id    = "CLOUD ACCOUNT ID"
  external_account_id = "EXTERNAL ACCOUNT ID"
  bundle_id           = 00000
  cloud_account_type  = "CLOUD ACCOUNT TYPE"
  notification_ids    = ["NOTIFICATION IDS"]
}

```

## Argument Reference

The following arguments are supported:

* `cloud_account_id` - (Required) The cloud account id.
* `external_account_id` - (Required) The account number.
* `bundle_id` - (Required) The bundle id for the policy.
* `cloud_account_type` - (Required) The cloud account provider (AWS/Azure/Google).
* `notification_ids` - (Required) The notifications id's for the policy.
    
## Attributes Reference

* `id` - ID of the policy.

## Import

The policy can be imported; use `<POLICY ID>` as the import ID. For example:

```shell
terraform import dome9_continuouscompliance_policy.test 00000000-0000-0000-0000-000000000000
```
