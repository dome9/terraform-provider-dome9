---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_awp_azure_onboarding_data"
sidebar_current: "docs-datasource-dome9-awp-azure-get-onboarding-data"
description: |-
  Get information about Azure AWP onboarding data in Dome9.
---

# Data Source: dome9_awp_azure_onboarding_data

Use this data source to get information about Azure AWP onboarding data in Dome9.

## Example Usage

```hcl
data "dome9_awp_azure_onboarding_data" "test" {
  cloud_account_id = "d9-azure-cloud-account-id"
  centralized_cloud_account_id = "d9-azure-centralized-cloud-account-id"
}

```

## Argument Reference

The following arguments supported:

* `cloud_account_id` - (Required) The Dome9 id for the onboarded Azure account, 
  * it can be the dome9 cloudguard account id or the azure subscription id.
* `centralized_cloud_account_id` - (Optional) The Dome9 id for the Azure scanner account, 
  * it can be the dome9 cloudguard account id or the azure subscription id.

## Attributes Reference

In addition to all arguments above, the following attributes exported:

* `region` - The region of the AWP Azure onboarding process.
* `app_client_id` - The Azure App client ID.
* `awp_cloud_account_id` - The Dome9 id for the onboarded Azure account.
* `awp_centralized_cloud_account_id` - The Dome9 id for the Azure scanner account. 