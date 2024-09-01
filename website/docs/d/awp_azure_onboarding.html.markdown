---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_awp_azure_onboarding"
sidebar_current: "docs-datasource-dome9-awp-azure-onboarding"
description: |-
  Get information about Azure AWP onboarding in Dome9.
---

# Data Source: dome9_awp_azure_onboarding

Use this data source to get information about Azure AWP onboarding in Dome9.
it gives details information about the awp azure account scanner configurations.

## Example Usage

```hcl
data "dome9_awp_azure_onboarding" "test" {
  id = "d9-azure-cloudguard-account-id or azure-subscription-id"
}

```

## Argument Reference

The following arguments supported:

* `id` - (Required) The Dome9 id for the onboarded Azure account.

## Attributes Reference

In addition to all arguments above, the following attributes exported:

* `scan_mode` - The scan mode of the onboarding process
* `agentless_account_settings` - The settings for the agentless account that the awp scanner will be configured with.
* `missing_awp_private_network_regions` - The regions missing AWP private network.
* `cloud_account_id` - The CloudGuard account ID.
* `agentless_protection_enabled` - Whether agentless protection is enabled or not.
* `cloud_provider` - The cloud provider for the onboarding process.
* `should_update` - Whether the onboarding process should be updated.
* `is_org_onboarding` - Whether the onboarding process is for an organization.