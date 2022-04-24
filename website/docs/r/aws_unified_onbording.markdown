---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_unified_onboarding"
sidebar_current: "docs-resource-dome9-aws-unified-onboarding"
description: AWS unified onboarding in Dome9
---

# dome9_aws_unified_onboarding

Begins an aws environment onboarding process. gets onboarding parameters and returns parameters 
that should be set for an aws_cloudformation_stack resource in order to complete the onboarding

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

* `cloud_vendor` - (Optional) the type of the environment that will be onboarded. can be one of the following: `aws`, `awsgov`, `awschina` the defult is `aws`
* `onboard_type` - (Optional) the onboarding type, `simple` for the default onboarding and `Advanced` for customized configuration; default is `Simple`
* `full_protection` - (Optional) The AWS environment operation mode. `true` for "Full-Protection(Read and write)", `false` for "Monitor(ReadOnly)", default is `false`
* `enable_stack_modify` - (Optional)  whether to allow CloudGuard to update and delete this stack upon update or delete of the environment or not, default is `false`
* `posture_management_configuration`:
  * `rulesets` - list of Posture rulesets Ids, that will be associated by policy with the environment
* `serverless_configuration`:
  * `enabled` - whether to enable Serverless protection or not, default is `true`
* `intelligence_configurations`:
  * `enabled` - whether to enable Intelligence (Account Activity) or not, default is `true`
  * `rulesets` - list of Intelligence rulesets Ids that will be associated by policy with the environment
  
## Attributes Reference

* `stack_name` - aws cloudformation stack name
* `parameters` - dictionary with the onboarding template parameters
* `iam_capabilities` - the IAM capabilities
* `template_url` - the Template Url 
 