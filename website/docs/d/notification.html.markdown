---
layout: "dome9"
page_title: "Check Point CloudGuard: dome9_notification"
sidebar_current: "docs-data-source-dome9-notification"
description: |- 
  Fetches details of a specific Dome9 notification.
---

# dome9\_notification

This data source is used to fetch details of a specific CloudGuard notification. You can retrieve various details about the notification such as its name, description, alert console settings, and integration settings.

## Example Usage

```hcl
data "dome9_notification" "example" {
  id = "your-notification-id"
}

output "notification_name" {
  value = data.dome9_notification.example.name
}

output "notification_description" {
  value = data.dome9_notification.example.description
}

output "alerts_console" {
  value = data.dome9_notification.example.alerts_console
}

output "integration_settings" {
  value = data.dome9_notification.example.integration_settings
}
```

## Argument Reference

The following arguments are supported:

- `id` - (Required) The ID of the Dome9 notification to retrieve information for.

## Attribute Reference

The following attributes are exported:

- `name` - The name of the CloudGuard notification.
- `description` - The description of the CloudGuard notification.
- `alerts_console` - Indicates if alerts will be sent to the CloudGuard console.
- `integration_settings` - A list of integration settings for the notification. This includes:
  - `reports_integration_settings` - (Optional) A list of report integration settings blocks. Each block includes:
    - `integration_id` - (Required) The ID of the integration.
    - `output_type` - (Optional) The output type of the integration.
  - `single_notification_integration_settings` - (Optional) A list of single notification integration settings blocks. Each block includes:
    - `integration_id` - (Required) The ID of the integration.
    - `output_type` - (Optional) The output type of the integration.
    - `payload` - (Optional) The payload of the notification (only for Jira).
  - `scheduled_integration_settings` - (Optional) A list of scheduled notification integration settings blocks. Each block includes:
    - `integration_id` - (Required) The ID of the integration.
    - `output_type` - (Optional) The output type of the integration.
    - `cron_expression` - (Required) The cron expression for the scheduled notification.