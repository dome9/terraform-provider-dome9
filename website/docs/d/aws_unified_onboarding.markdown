---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_unified_onboarding"
sidebar_current: "docs-datasource-dome9-aws-unified-onboarding"
description: |- GET Onboarding information by onboarding_id/cloud_account_id in the "Required" field
---

# Data Source: dome9_aws_unified_onboarding

Use this data source to get information about onboarding of an AWS cloud account.

## Example Usage

```hcl
data "dome9_aws_unified_onboarding" "aws_unidied_onboarding_ds" {
    id = "ID"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) environment_id / onboarding_id 

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `onboarding_id` - the onboarding id
* `InitiatedUserName` - the name of the initiated User
* `initiated_user_id` - the id of the initiated User
* `environment_external_id` - the AWS cloud account id
* `environment_id` - the AWS cloud account internal id
* `environment_name` - the aws environment name
* `root_stack_id` - the arn of the root stack created by the onboarding
* `cft_version` - the current Cloud Formation Template version
* `onbording_request` - [the onboarding request](https://registry.terraform.io/providers/dome9/dome9/latest/docs/resources/aws_unified_onbording)
* `statuses` - list of statuses, a status per blade of the onboarding process
  * `module` - the module name (Permissions / Continuous Posture / Serverless Protection / Account Activity)
  * `feature` - the feature name (Intelligence / Workload / Posture / Onboarding / General)
  * `remediation_recommendation` - remediation recommendation
  * `stack_message` - the stack massage
  * `stack_status` - the stack status
  * `status` - The status
  * `status_message` - the status massage

