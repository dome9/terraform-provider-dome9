---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_organization_onboarding_member_account_configuration"
sidebar_current: "docs-data-source-dome9-aws-organization-onboarding-member-account-configuration"
description: Get information about the AWS organization onboarding member account configuration in CloudGuard.
---

# Data Source: dome9_aws_organization_onboarding_management_stack

Use this data source to retrieve information about the AWS organization onboarding management stack in CloudGuard.

## Example Usage

Basic usage:

```hcl
data "dome9_aws_organization_onboarding_member_account_configuration" "example" {}
```

## Attributes Reference

The following attributes are returned:

* `external_id` - Used in the CloudGuard role.
* `content` - The content of the management stack.
* `onboarding_cft_url` - The s3 URL of the CloudFormation template for management onboarding.