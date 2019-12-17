---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_attach_iam_safe"
sidebar_current: "docs-datasource-dome9-attach-iam-safe"
description: |-
    Attach IAM safe to AWS cloud account.
---

# Data Source: dome9_attach_iam_safe

Use this data source to get information about a IAM safe for AWS cloud account that protected by Dome9.

## Example Usage

```hcl
data "dome9_attach_iam_safe" "test" {
  id = "D9-AWS-CLOUD-ACCOUNT"
}

```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The ID of the AWS cloud account that protected by Dome9.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `aws_group_arn` - AWS group arn.
* `aws_policy_arn` - AWS policy arn.
* `mode` - Mode.
