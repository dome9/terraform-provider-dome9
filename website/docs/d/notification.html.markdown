---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_notification"
sidebar_current: "docs-data-source-dome9-notification"
description: |- Fetches details of a specific Dome9 notification.
---

# dome9\_notification

This data source is used to fetch details of a specific Dome9 notification. You can retrieve various details about the notification such as its name, description, alert console settings, and integration settings.

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

- `name` - The name of the Dome9 notification.
- `description` - The description of the Dome9 notification.
- `alerts_console` - Indicates if alerts will be sent to the Dome9 console.
- `integration_settings` - A list of integration settings for the notification. This includes:
  - `reports_integration_settings` - Settings for report integrations.
  - `single_notification_integration_settings` - Settings for single notification integrations.
  - `scheduled_integration_settings` - Settings for scheduled integrations.
    - Each `reports_integration_settings`, `single_notification_integration_settings`, and `scheduled_integration_settings` block contains:
      - `integration_id` - The ID of the integration.
      - `output_type` - The output type of the integration.
      - For `single_notification_integration_settings` only: `payload` - The payload of the Jira notification.
      - For `scheduled_integration_settings` only: `cron_expression` - The cron expression for the scheduled notification.