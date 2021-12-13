---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_continuous_compliance_notification"
sidebar_current: "docs-resource-dome9-continuous-compliance-notification"
description: |- Creates continuous compliance notification in Dome9
---

# dome9_continuous_compliance_notification

This resource is used to create and modify Dome9 notification policies for Continuous Compliance assessments of cloud
accounts. Continuous assessments apply bundles of compliance rules to your cloud account continuously, and send
notifications of issues according to the Notification Policy.

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
    slack_integration_state            = "SLACK_INTEGRATION_STATE"
    teams_integration_state            = "TEAMS_INTEGRATION_STATE"

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

* `name` - (Required) The cloud account id in Dome9.
* `description` - (Optional) Description of the notification.

at least one of  `alerts_console`, `scheduled_report`, or `change_detection` must be included

* `alerts_console` - (Optional) send findings (also) to the Dome9 web app alerts console (Boolean); default is False.

* `scheduled_report` - Scheduled email report notification block:
    * `email_sending_state` - send schedule report of findings by email; can be  "Enabled" or "Disabled".

      if `email_sending_state` is Enabled, the following must be included:

        * `schedule_data` - Schedule details:
            * `cron_expression` - (Required) the schedule to issue the email report (in cron expression format)
            * `type` - (Required) type of report; can be  "Detailed", "Summary", "FullCsv" or "FullCsvZip"
            * `recipients` - (Required) comma-separated list of email recipients


* `change_detection` - (Required) Send changes in findings options:
    * `email_sending_stat` - send email report of changes in findings; can be "Enabled" or "Disabled".

      if `email_sending_stat`  is Enabled, the following must be included:

        * `email_data` - Email notification details:
            * `recipients` - (Required) comma-separated list of email recipients

              <br/>

    * `email_per_finding_sending_state` - send separate email notification for each finding; can be "Enabled" or "
      Disabled"

      if `email_per_finding_sending_state`  is Enabled, the following must be included:

        * `email_per_finding_data` - Email per finding notification details:
            * `recipients` - (Required) comma-separated list of email recipients
            * `notification_output_format` - (Required) format of JSON block for finding; can be  "JsonWithFullEntity"
              , "JsonWithBasicEntity", or "PlainText".

            <br/>

    * `sns_sending_state` - send by AWS SNS for each new finding; can be  "Enabled" or "Disabled".

      if `sns_sending_state`  is Enabled, the following must be included:

        * `sns_data` - SNS notification details:
            * `sns_topic_arn` - (Required) SNS topic ARN
            * `sns_output_format` - (Required) SNS output format; can be  "JsonWithFullEntity", "JsonWithBasicEntity", or "
              PlainText".

            <br/>
    * `external_ticket_creating_state` - send each finding to an external ticketing system; can be  "Enabled" or "
      Disabled".

      if `external_ticket_creating_state`  is Enabled, the following must be included:

        * `ticketing_system_data` - Ticketing system details:
            * `system_type` - system type; can be "ServiceOne", "Jira", or "PagerDuty"
            * `should_close_tickets` - ticketing system should close tickets when resolved (bool)
            * `domain` - ServiceNow domain name (ServiceNow only)
            * `user` - user name (ServiceNow only)
            * `pass` - (Required) password (ServiceNow only)
            * `project_key` - project key (Jira) or API Key (PagerDuty)
            * `issue_type` - issue type (Jira)

            <br/>

    * `webhook_integration_state` - send findings to an HTTP endpoint (webhook); can be  "Enabled" or "Disabled".

      if `webhook_integration_state`  is Enabled, the following must be included:

        * `webhook_data` - Webhook data block supports:
            * `url` - (Required) HTTP endpoint URL
            * `http_method` - HTTP method, "Post" by default.
            * `auth_method` - authentication method; "NoAuth" by default
            * `username` - username in endpoint system
            * `password` - password in endpoint system
            * `format_type` - format for JSON block for finding; can be "Basic" or "ServiceNow"
            * `payload_format` - Json Payload
            * `ignore_certificate` - Check this to use self-signed certificates, and ignore validation of them
            * `advanced_url` - Tick this box if you are using a version of Jira that only supports REST API 2

          <br/>

    * `slack_integration_state` - Send report summary to Slack channel (Compliance only); can be  "Enabled" or "
      Disabled".

      if `slack_integration_state`  is Enabled, the following must be included:

        * `slack_data` - Slack data block supports:
            * `url` - (Required) Slack's webhook URL

          <br/>

    * `teams_integration_state` - Send report summary to Teams channel (Compliance only); can be  "Enabled" or "
      Disabled".

      if `teams_integration_state`  is Enabled, the following must be included:

        * `teams_data` - Teams data block supports:
            * `url` - (Required) Teams webhook URL

          <br/>

    * `aws_security_hub_integration_state` - send findings to AWS Secure Hub; can be "Enabled" or "Disabled".

      if `aws_security_hub_integration_state`  is Enabled, the following must be included:

        * `aws_security_hub_integration` - AWS security hub integration block supports:
            * `external_account_id` - (Required) external account id
            * `region` - (Required) AWS region

          <br/>

`gcp_security_command_center_integration` is a change_detection option

* `gcp_security_command_center_integration` - GCP security command center details
    * `state` - send findings to the GCP Security Command Center; can be "Enabled" or "Disabled"

      if `state` is Enabled, the following must be included:

        * `project_id` - GCP Project id
        * `source_id` - GCP Source id

## Import

The notification can be imported; use `<NOTIFICATION ID>` as the import ID.

For example:

```shell
terraform import dome9_continuouscompliance_notification.test 00000000-0000-0000-0000-000000000000
```
