---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_unified_onboarding"
sidebar_current: "docs-datasource-dome9-aws-unified-onboarding"
description: |-
  GET configuration should be set to the AWS cloud formation SDK for an update
---

# Data Source: dome9_aws_unified_onboarding

Use this data source to get the configuration that should be set to the AWS cloud formation SDK for an update.

## Example Usage

```hcl
data "dome9_aws_unified_onboarding" "test" {
  onboarding_id = "onboarding_id"
}
```

## Argument Reference

The following arguments are supported:

* `onboarding_id` - (Required) The onboarding that create with creation of dome9_aws_unified_onboarding resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `stack_name` - the Google project id (that will be onboarded).
* `parameters` - map of parameters. ???? need to explain more 
* `iam_capabilities` - Organizational unit id.
* `template_url` - Organizational unit path.
