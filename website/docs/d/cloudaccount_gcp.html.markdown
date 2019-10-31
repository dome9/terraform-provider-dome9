---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_gcp"
sidebar_current: "docs-datasource-dome9-cloudaccount-gcp"
description: |-
  Get information about a GCP cloud account onboarded to Dome9.
---

# Data Source: dome9_cloudaccount_gcp

Use this data source to get information about a GCP cloud account onboarded to Dome9.

## Example Usage

```hcl
data "dome9_cloudaccount_gcp" "test" {
  account_id         = "my-dome9-id"
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The Dome9  id for the GCP account.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - GCP account name in Dome9.
* `project_id` - the Google project id (that will be onboarded).
* `creation_date` - creation date for project in Google.
* `organizational_unit_id` - Organizational unit id.
* `organizational_unit_path` - Organizational unit path.
* `organizational_unit_name` - Organizational unit name.
* `gsuite_user` - Gsuite user.
* `domain_name` - Domain name.
* `domain_name` - Azure tenant id.
* `vendor` - The cloud provider (gcp).
