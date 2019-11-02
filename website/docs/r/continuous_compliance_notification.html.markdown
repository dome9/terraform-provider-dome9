---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_continuous_compliance_notification"
sidebar_current: "docs-resource-dome9-continuous-compliance-notification"
description: |-
  Creates continuous compliance notification in Dome9
---

# dome9_continuous_compliance_notification

The ContinuousComplianceNotification resource has methods to create and modify notification policies for continuous compliance assessments. With continuous assessments, cloud environments are assessed continuously and autonomously with bundles, and the results are issued to designated recipients as emails or SNS notifications, according to notification policies.

## Example Usage

Basic usage:

```hcl
resource "dome9_continuous_compliance_notification" "test_notification" {
  name           = "NAME"
  description    = "DESCRIPTION"
  alerts_console = "ALERTS_CONSOLE"

  change_detection {
    email_sending_state                = "EMAIL_SENDING_STATE"
    email_per_finding_sending_state    = "EMAIL_PER_FINDING_SENDING_STATE"
    sns_sending_state                  = "SNS_SENDING_STATE"
    external_ticket_creating_state     = "EXTERNAL_TICKET_CREATING_STATE"
    aws_security_hub_integration_state = "AWS_SECURITY_HUB_INTEGRATION_STATE"
    webhook_integration_state          = "WEBHOOK_INTEGRATION_STATE"

    email_data {
      recipients = ["RECIPIENTS"]
    }

    email_per_finding_data {
      recipients                 = ["RECIPIENTS"]
      notification_output_format = "NOTIFICATION_OUTPUT_FORMAT"
    }
  }
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The cloud account id.
* `change_detection` - (Required) The account number.
    
## Attributes Reference

* `description` - Description of the notification.
* `alerts_console` - Bool field include in the alerts console.
* `scheduled_report` - Scheduled report block supports:
    * `email_sending_state` - Email Address, can be one of the following: "Enabled" or "Disabled.
    * `schedule_data` - Set or Get Notification schedule block supports:
        * `cron_expression` - (Required) 
        * `type` - (Required) Can be set to one of the following: "Detailed", "Summary", "FullCsv" or "FullCsvZip"
        * `recipients` - (Required) 
    
* `change_detection` - (Required) Change detection block supports:
    * `email_sending_stat` - Can be one of the following: "Enabled" or "Disabled.
    * `email_per_finding_sending_state` - Can be one of the following: "Enabled" or "Disabled.
    * `sns_sending_state` - Can be one of the following: "Enabled" or "Disabled.
    * `external_ticket_creating_state` - Can be one of the following: "Enabled" or "Disabled.
    * `aws_security_hub_integration_state` - Can be one of the following: "Enabled" or "Disabled.
    * `webhook_integration_state` - Can be one of the following: "Enabled" or "Disabled.
    * `email_data` - Change detection block supports:
        * `recipients` - (Required)
    * `email_per_finding_data` - Email per finding data block supports:
        * `recipients` - (Required)
        * `notification_output_format` - (Required) Can be one of the following: "JsonWithFullEntity", "JsonWithBasicEntity", "PlainText".
    * `sns_data` - SNS data block supports:
        * `sns_topic_arn` - SNS topic arn
        * `sns_output_format` - SNS output format, can be one of the following: "JsonWithFullEntity", "JsonWithBasicEntity", "PlainText".
    * `ticketing_system_data` - Ticketing system data block supports:
        * `system_type` - System type 
        * `should_close_tickets` - Should close tickets 
        * `domain` - Domain 
        * `user` - User 
        * `pass` - Pass 
        * `project_key` - Project key 
        * `issue_type` - Issue type 
    * `aws_security_hub_integration` - AWS security hub integration block supports:
        * `external_account_id` - (Required) external account id
        * `region` - (Required) region
    * `webhook_data` - Webhook data block supports:
        * `url` - url 
        * `http_method` - HTTP method, "Post" by default.
        * `auth_method` - Auth method 
        * `username` - username 
        * `password` - password 
        * `format_type` - Format type 
* `gcp_security_command_center_integration` - GCP security command center integration
    * `state` - State 
    * `project_id` - Project id 
    * `source_id` - Source id 

## Import

The notification can be imported; use `<NOTIFICATION ID>` as the import ID. 

For example:

```shell
terraform import dome9_continuouscompliance_notification.test 00000000-0000-0000-0000-000000000000
```
