---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_awp_aws_onboarding"
sidebar_current: "docs-datasource-dome9-awp-aws-onboarding"
description: |-
  Get information about AWS AWP onboarding in Dome9.
---

# Data Source: dome9_awp_aws_onboarding

Use this data source to get information about AWS AWP onboarding in Dome9.
it gives details information about the awp aws account scanner configurations.

## Example Usage

```hcl
data "dome9_awp_aws_onboarding" "test" {
  id = "d9-aws-cloudguard-account-id or aws-account-id"
}

```

## Argument Reference

The following arguments supported:

* `id` - (Required) The Dome9 id for the onboarded AWS account.

## Attributes Reference

In addition to all arguments above, the following attributes exported:

* `scan_mode` - The scan mode of the onboarding process
* `agentless_account_settings` - The settings for the agentless account that the awp scanner will be configured with.
* `missing_awp_private_network_regions` - The regions missing AWP private network.
* `account_issues` - The issues related to the awp account.
* `cloud_account_id` - The CloudGuard account ID.
* `agentless_protection_enabled` - Whether agentless protection is enabled or not.
* `cloud_provider` - The cloud provider for the onboarding process.
* `should_update` - Whether the onboarding process should be updated.
* `is_org_onboarding` - Whether the onboarding process is for an organization.