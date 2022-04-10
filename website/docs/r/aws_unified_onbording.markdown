---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_unified_onboarding"
sidebar_current: "docs-resource-dome9-continuous-compliance-notification"
description: AWS unified onboarding in Dome9
---

# dome9_aws_unified_onboarding

Create an onboarding configuration including an onbordingId

## Example Usage

Basic usage:

```hcl
resource resource "dome9_aws_unified_onboarding" "test" {
	cloud_vendor = "aws"
        onboard_type = "Simple"
        full_protection = true
        enable_stack_modify = true

        posture_management_configuration = {
        rulesets = "[0]"
    }
    serverless_configuration  = {
        enabled = true
    }
    intelligence_configurations = {
        rulesets = "[0]"
        enabled = true
    }
}

```

## Argument Reference

The following arguments are supported:

* `cloud_vendor ` - (Optional) Cloud vendor that the ruleset is associated with, can be one of the following: `aws`, `azure`, `google` etc.
* `onboard_type ` - (Optional) "Simple" for oneClick onbording. (String); default is "Simple".
* `full_protection ` - (Optional) The protection mode for security groups in the account. (Boolean); default is False.
* `enable_stack_modify ` - (Optional) Enable stack modify. (Boolean); default is False.
* `posture_management_configuration  ` - (Optional) :
    * `rulesets ` - List of ruleset Ids. default is [0]
* `serverlessConfiguration` - (Optional) Send changes in findings options:
    * `enabled` - true or false to Enables serverless.  (Boolean); default is true.
* `intelligenceConfigurations` - (Optional) Send changes in findings options:
    * `enabled` - true or false to Enables intelligence.  (Boolean); default is true.
    * `rulesets ` - List of ruleset Ids. default is [0]
 
