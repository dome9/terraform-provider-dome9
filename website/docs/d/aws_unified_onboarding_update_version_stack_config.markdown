---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_unified_onboarding_information"
sidebar_current: "docs-datasource-dome9-aws-unified-onboarding"
description: |-
  GET configuration should be set to the AWS cloud formation SDK for an update
---

# Data Source: dome9_aws_unified_onboarding_update_version_stack_config

Use this data source to get the configuration that should be set to the AWS cloud formation resource for an update

## Example Usage

```hcl
data "dome9_aws_unified_onboarding_update_version_stack_config" "test" {
  onboarding_id = "onboarding_id"
}
```

## Argument Reference

The following arguments are supported:

* `onboarding_id` - (Required) The onboarding id (can be taken for the dome9_aws_unified_onboarding resource 
  (please link it to the resource) or dome9_aws_unified_onboarding data source

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `stack_name` - the aws cloudformation stack name
* `parameters` - dictionary with the onboarding template parameters
* `iam_capabilities` - the iam capabilities
* `template_url` - the template url
