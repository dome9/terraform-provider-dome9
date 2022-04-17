---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_unified_onboarding"
sidebar_current: "docs-datasource-dome9-aws-unified-onboarding"
description: |- GET Onboarding information by onboarding_id/cloud_account_id in the "Required" field
---

# Data Source: dome9_aws_unified_onboarding

Use this data source to get the information about the onboarding of an AWS cloud account.

## Example Usage

```hcl
data "dome9_aws_unified_onboarding" "aws_unidied_onboarding_ds" {
    id = "ID" CloudAccountId or onbordingId as string
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) cloud account id / the onboarding id that create with creation of
  dome9_aws_unified_onboarding resource

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `onboarding_id` - the onboarding id
* `InitiatedUserName` - the name of the initiated User
* `initiated_user_id` - the id of the initiated User
* `environment_external_id` - the AWS cloud account id
* `environment_id` - the AWS cloud account internal environment id
* `environment_name` - aws environment name
* `root_stack_id` - arn:aws:cloudformation:*
* `cft_version` - the current Cloud Formation Template version
* `onbording_request` - the request data for the creation of the onboarding
  * `cloud_vendor` - Cloud vendor that the ruleset is associated with, can be one of the following: `aws`, `awsgov`, `awschina` the defult is `aws`
  * `onboard_type` - "Simple" for pre-configured "one-click" onboarding and "Advanced" for customized configuration (String); default is "simple"
  * `full_protection` - The AWS cloud account operation mode. `true` for "Full-Protection(Read and write)", `false` for "Monitor(ReadOnly)"; default is `false`
  * `enable_stack_modify` -(Optional) Allow Cloud Guard tom update and delete this stack upon update or delete of the environment, default is `false`
  * `posture_management_configuration`:
    * `rulesets` - List of Posture Management ruleset Ids (String) default is "[]"
  * `serverless_configuration`:
    * `enabled` - whether to enable Serverless protection or not, default is `true`
  * `intelligence_configurations`:
    * `enabled` - whether to enable Intelligence (Account Activity) or not, default is `true`
    * `rulesets` - list of Intelligence ruleset Ids that will be associated with a policy, default is []
* `statuses` - list of statuses, a status per blade of the onboarding process
  * `module` - the module name(Role Permissions/Continuous Posture/Serverless Protection/Account Activity)
  * `feature` - the feature name(Intelligence/Workload/Posture/Onbording/General)
  * `remediation_recommendation` - remediation recommendation
  * `stack_message` - the stack massage
  * `stack_status` - the stack status
  * `status` - The status (Error/Active)
  * `status_message` - the status massage

