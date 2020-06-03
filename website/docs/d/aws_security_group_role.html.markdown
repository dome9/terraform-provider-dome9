---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_security_group_role"
sidebar_current: "docs-datasource-dome9-aws-security-group-role"
description: |-
  Bound input and output services to AWS Security Group in Dome9
---

# Data Source: dome9_aws_security_group_role

This resource has methods to add and manage input and output services to Security Groups in a cloud account that is managed by Dome9.

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
