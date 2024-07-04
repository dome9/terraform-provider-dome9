---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_organization_onboarding_management_stack"
sidebar_current: "docs-data-source-dome9-aws-organization-onboarding-management-stack"
description: Get information about the AWS organization onboarding management stack in CloudGuard.
---

# Data Source: dome9_aws_organization_onboarding_management_stack

Use this data source to retrieve information about the AWS organization onboarding management stack in CloudGuard.

## Example Usage

Basic usage:

```hcl
data "dome9_aws_organization_onboarding_management_stack" "example" {
  aws_account_id = "AWS_MANAGEMENT_ACCOUNT_ID"
}
```

## Argument Reference

The following arguments are supported:

* `aws_account_id` - (Required) The AWS management account ID.

## Attributes Reference

The following attributes are returned:

* `external_id` - Used in the CloudGuard role (also called secret).
* `content` - The content of the management stack.
* `management_cft_url` - The s3 URL of the CloudFormation template for management onboarding.
* `is_management_onboarded` - The status of management onboarding (true if onboarded, false otherwise).