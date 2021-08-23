---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_service_account"
sidebar_current: "docs-resource-dome9-service-account"
description: |-
  Get information about a Service Account in Dome9
---

# Data Source: dome9_service_account

Use this data source to get information about a service account in Dome9.

## Example Usage

Basic usage:

```hcl
data "dome9_service_account" "service_account" {
  id = "ID"
}

```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The id of the service account in Dome9.

## Attributes Reference

* `name` - service account name.
* `id` - service account id.
* `api_key_id` - api key.
* `role_ids` - service account role ids.
* `date_created` - service account creation time.
* `last_used` - service account last used.