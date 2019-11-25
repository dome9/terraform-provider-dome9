---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_continuous_compliance_policy"
sidebar_current: "docs-datasource-dome9-continuous-compliance-policy"
description: |-
  Get information about a Dome9 continuous compliance policy.
---

# Data Source: dome9_continuous_compliance_policy

Use this data source to get information about a Dome9 continuous compliance policy.

## Example Usage

```hcl
data "dome9_continuous_compliance_policy" "test" {
  id = "d9-continuous-compliance-policy-id"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The id for the cloud account in Dome9. 

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_account_id` - GCP account name in Dome9.
* `external_account_id` - The account number.
* `cloud_account_type` - creation date for project in Google.
* `bundle_id` - Organizational unit id.
* `notification_ids` - Organizational unit path.
