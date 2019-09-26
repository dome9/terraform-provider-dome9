---
layout: "dome9"
page_title: "Check Point Cloud Guard Dome9: dome9_cloudaccount_gcp"
sidebar_current: "docs-datasource-dome9-cloudaccount-gcp"
description: |-
  Get information on GCP cloud account.
---

# Data Source: dome9_cloudaccount_gcp

Use this data source to get information about GCP cloud account.

## Example Usage

```hcl
data "dome9_cloudaccount_gcp" "test" {
  account_id         = "my-dome9-id"
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) Account id in Dome9.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Google account name in Dome9.
* `project_id` - the Google project id (that will be onboarded).
* `creation_date` - creation date for project in Google.
* `organizational_unit_id` - Organizational unit id.
* `organizational_unit_path` - Organizational unit path.
* `organizational_unit_name` - Organizational unit name.
* `gsuite_user` - Gsuite user.
* `domain_name` - Domain name.
* `domain_name` - Azure tenant id.
* `vendor` - The cloud provider (gcp).
