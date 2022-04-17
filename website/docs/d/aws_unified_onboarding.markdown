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
    cloud_account_id = "ID" CloudAccountId or onbordingId as string
}
```

## Argument Reference

The following arguments are supported:

* `cloud_account_id` - (Required) cloud account id / the onboarding id that create with creation of
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
* `statuses` - the current status per blade of the onboarding process
