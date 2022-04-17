---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_unified_onboarding"
sidebar_current: "docs-resource-dome9-aws-unified-onboarding"
description: AWS unified onboarding in Dome9
---

# dome9_aws_unified_onboarding

Create an onboarding configuration including an onboardingId and use the configuration to create
"aws_cloudformation_stack" in order to onboard to "Cloud Guard"

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


resource "aws_cloudformation_stack" "stack"{
  name = dome9_aws_unified_onboarding.test.stack_name
  template_url = dome9_aws_unified_onboarding.test.template_url
  parameters = dome9_aws_unified_onboarding.test.parameters
  capabilities = dome9_aws_unified_onboarding.test.iam_capabilities
}

```

## Argument Reference

The following arguments are supported:

* `cloud_vendor` - (Optional) the type of the cloud account that will be onboarded. can be one of the following: `aws`
  , `awsgov`, `awschina` the defult is `aws`
* `onboard_type` - (Optional) "simple" for pre-configured "one-click" onboarding and "Advanced" for customized
  configuration (String); default is "Simple"
* `full_protection` - (Optional) The AWS cloud account operation mode. `true` for "Full-Protection(Read and write)", `false` for "
  Monitor(ReadOnly)"; default is `false`
* `enable_stack_modify` - (Optional) Allow Cloud Guard tom update and delete this stack upon update or delete of the environment, default is `false`
* `posture_management_configuration` - (Optional) :
    * `rulesets` - List of Posture Management ruleset Ids (String) default is "[]"
* `serverless_configuration` - (Optional):
    * `enabled` - whether to enable Serverless protection or not, default is `true`
* `intelligence_configurations` - (Optional):
    * `enabled` - whether to enable Intelligence (Account Activity) or not, default is `true`
    * `rulesets` - list of Intelligence ruleset Ids that will be associated with a policy, default is []

## Attributes Reference

* `stack_name` - aws cloudformation stack name
* `parameters` - dictionary with the onboarding template parameters
* `iam_capabilities` - the IAM capabilities
* `template_url` - the Template Url 
 