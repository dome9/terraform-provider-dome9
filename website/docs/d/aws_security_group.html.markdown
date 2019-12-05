---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_security_group"
sidebar_current: "docs-datasource-dome9-aws-security-group"
description: |-
  Get information about AWS security group in Dome9
---

# Data Source: dome9_aws_security_group

Use this data source to get information about an AWS security group onboarded to Dome9.

## Example Usage

Basic usage:

```hcl
data "dome9_aws_security_group" "aws_sg_ds" {
  id = "SECURITY_GROUP_ID"
}

```

## Argument Reference
In addition to all arguments above, the following attributes are exported:

* `dome9_security_group_name` - Name of the security group.
* `dome9_cloud_account_id` - Cloud account id in Dome9.
* `description` - Security group description.
* `aws_region_id` - AWS region; in AWS format (e.g., "us-east-1").
* `is_protected` - Is security group protected.
* `vpc_id` - Security group id.
* `vpc_name` - Security group vpc name.
* `tags` - Security group tags.
* `services` - Security group services.
* `cloud_account_name` - AWS cloud account name.
* `external_id` - Security group external id.
