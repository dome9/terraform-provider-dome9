---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_attach_iam_safe"
sidebar_current: "docs-resource-dome9-attach-iam-safe"
description: |-
  Attach Dome9 IAM Safety to an AWS cloud account.
---

# dome9_attach_iam_safe

Attach IAM safe to AWS cloud account.

## Example Usage

Basic usage:

```hcl
resource "dome9_attach_iam_safe" "test" {
  aws_cloud_account_id = "00000000-0000-0000-0000-000000000000"
  aws_group_arn        = "AWS_GROUP_ARN"
  aws_policy_arn       = "AWS_POLICY_ARN"
}

```

## Argument Reference

The following arguments are supported:

* `aws_cloud_account_id` - (Required) AWS cloud account to which IAM safety will be attached. 
* `aws_group_arn` - (Required) AWS group ARN.
* `aws_policy_arn` - (Required) AWS policy ARN.

## Attributes Reference

* `mode` - Mode.

## Import

Cloud account IAM safety can be imported; use `<AWS CLOUD ACCOUNT ID>` as the import ID. 

For example:

```shell
terraform import dome9_attach_iam_safe_re.test 00000000-0000-0000-0000-000000000000
```
