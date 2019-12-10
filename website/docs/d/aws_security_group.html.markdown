---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_security_group"
sidebar_current: "docs-datasource-dome9-aws-security-group"
description: |-
  Get information about AWS Security Groups in Dome9
---

# Data Source: dome9_aws_security_group

Use this data source to get information about an AWS Security Group onboarded to Dome9.

## Example Usage

Basic usage:

```hcl
data "dome9_aws_security_group" "aws_sg_ds" {
  id = "SECURITY_GROUP_ID"
}

```

## Argument Reference
In addition to all arguments above, the following attributes are exported:

* `dome9_security_group_name` - Name of the Security Group.
* `dome9_cloud_account_id` - Cloud account id in Dome9.
* `description` - Security Group description.
* `aws_region_id` - AWS region; in AWS format (e.g., "us-east-1").
* `is_protected` - indicates whether the Security Group is protected.
* `vpc_id` - Security Group id.
* `vpc_name` - name of VPC containing the Security Group .
* `tags` - Security Group tags.
* `services` - Security Group services.
* `cloud_account_name` - AWS cloud account name.
* `external_id` - Security Group external id.
