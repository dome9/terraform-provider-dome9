---
layout: "dome9"
page_title: "Check Point CloudGuard: dome9_notification"
sidebar_current: "docs-data-source-dome9-notification"
description: |- 
  Fetches details of a specific CloudGuard integration.
---

# dome9\_notification

This data source is used to fetch details of a specific CloudGuard integration. You can retrieve various details about the integration such as its name, type, and configuration.

## Example Usage

```hcl
data "dome9_integration" "example" {
  id = "your-integration-id"
}

output "integration_name" {
  value = data.dome9_integration.example.name
}

output "integration_type" {
  value = data.dome9_integration.example.type
}

output "integration_configuration" {
  value = data.dome9_integration.example.configuration
}
```
## Argument Reference

The following arguments are supported:

- `id` - (Required) The ID of the CloudGuard integration to retrieve information for.

## Attribute Reference

The following attributes are exported:

- `name` - The name of the CloudGuard integration.
- `type` - The type of the CloudGuard integration.
- `configuration` - The configuration details of the CloudGuard integration.
