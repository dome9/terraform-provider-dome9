---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_unified_onboarding"
sidebar_current: "docs-datasource-dome9-aws-unified-onboarding"
description: |- GET Onboarding information by onboarding_id/account_id in the "Required" field
---

# Data Source: dome9_aws_unified_onboarding

Use this data source to get the information about the onboarding to the AWS cloud account.

## Example Usage

```hcl
data "dome9_aws_unified_onboarding" "T" {
    cloud_account_id = "ID" CloudAccountId or onbordingId as string
}
```

## Argument Reference

The following arguments are supported:

* `cloud_account_id` - (Required) cloud_account_id OR The onboardingId that create with creation of
  dome9_aws_unified_onboarding resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `onboarding_id` - The OnboardingId that create with creation of dome9_aws_unified_onboarding resource..
* `InitiatedUserName` - The name of the initiated User.
* `initiated_user_id` - The ID of the initiated User.
* `environment_external_id` - The AWS cloud account id.
* `environment_id` - environment_id ?. "4d76adb3-6c97-4984-948b-119c2cce6252"
* `environment_name` - Organizational unit path."chkp-aws-rnd-menahema-base"
* `root_stack_id` - arn:aws:cloudformation:*
* `cft_version` - ?
* `onbording_request` - The request data for the creation of the current Data Source.
* `statuses` - ?.
