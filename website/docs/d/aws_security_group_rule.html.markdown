---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_security_group_rule"
sidebar_current: "docs-datasource-dome9-aws-security-group-rule"
description: |-
  Bound input and output services to AWS Security Group in Dome9
---

# Data Source: dome9_aws_security_group_rule

Use this data source to get inbounds and outbounds services for AWS Security Groups in a cloud account that is managed by Dome9.

## Example Usage

Basic usage:

```hcl
data "dome9_aws_security_group" "aws_sg_ds" {
  id = "SECURITY_GROUP_ID"
}

```

## Argument Reference
In addition to all arguments above, the following attributes are exported:

* `services` - Security Group services.
