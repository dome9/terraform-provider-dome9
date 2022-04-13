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

* `cloud_vendor` - (Optional) Cloud vendor that the ruleset is associated with, can be one of the following: `aws`, `awsgov`, `awschina` the defult is `aws`
* `onboard_type` - (Optional) "Simple" for pre-configured "one-click" onboarding and "Advanced" for customized configuration (String); default is "Simple"
* `full_protection` - (Optional) The AWS cloud account operation mode. `true` for "Full-Protection(R/W)", `false` for "Monitor(ReadOnly)"
* `enable_stack_modify` - (Optional) Enable stack modify (Boolean); default is `false`
* `posture_management_configuration` - (Optional) :
    * `rulesets` - List of Posture Management ruleset Ids (String) default is "[0]"
* `serverlessConfiguration` - (Optional):
    * `enabled` - `true` or `false` to enable Serverless Protection (Boolean); default is `true`
* `intelligenceConfigurations` - (Optional):
    * `enabled` - `true` or `false` to enable Intelligence (Account Activity) (Boolean); default is `true`
    * `rulesets` - List of Intelligence ruleset Ids (String) default is [0]
 
