---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_notification"
sidebar_current: "docs-resource-dome9-notification"
description: |- 
  Creates and manages Dome9 notifications.
---

# dome9\_notification

This resource is used to create and manage Dome9 notifications. Notifications in Dome9 allow you to send alerts based on specific events or criteria to various destinations such as email, Slack, or an HTTP endpoint.

## Example Usage

Basic usage:

```hcl
resource "dome9_notification" "example_notification" {
  name                    = "Example Notification"
  description             = "This is an example notification."
  alerts_console          = true
  send_on_each_occurrence = false

  integration_settings {
    single_notification_integration_settings {
      integration_id = "example-integration-id-1"
      output_type    = "Default"
    }

    reports_integration_settings {
      integration_id = "example-integration-id-2"
      output_type    = "Default"
    }

    scheduled_integration_settings {
      integration_id  = "example-integration-id-3"
      output_type     = "Detailed"
      cron_expression = "0 0 22 * * ?"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) The name of the notification.
- `description` - (Optional) A description of the notification.
- `alerts_console` - (Optional) Boolean flag to send alerts to the Dome9 web app alerts console. Defaults to `true`.
- `send_on_each_occurrence` - (Optional) Boolean flag to send notifications on each occurrence. Defaults to `false`.
- `origin` - (Optional) The origin of the notification. Defaults to `"ComplianceEngine"`.
- `integration_settings` - (Required) A block of integration settings for the notification. The block supports:
    - `reports_integration_settings` - (Optional) A list of report integration settings blocks.
        - `integration_id` - (Required) The ID of the integration.
        - `output_type` - (Optional) The output type of the integration.
    - `single_notification_integration_settings` - (Optional) A list of single notification integration settings blocks.
        - `integration_id` - (Required) The ID of the integration.
        - `output_type` - (Optional) The output type of the integration.
        - `payload` - (Optional) The payload of the notification.
    - `scheduled_integration_settings` - (Optional) A list of scheduled notification integration settings blocks.
        - `integration_id` - (Required) The ID of the integration.
        - `output_type` - (Optional) The output type of the integration.
        - `cron_expression` - (Required) The cron expression for the scheduled notification.
      

## Import

This resource can be imported using the continuous compliance notification ID, which can be found in the Dome9 console.

```shell
terraform import dome9_notification.example <NOTIFICATION_ID>