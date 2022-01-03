---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_continuous_compliance_notification"
sidebar_current: "docs-datasource-dome9-continuous-compliance-notification"
description: |-
  Get information about a Dome9 continuous compliance notification policy.
---

# Data Source: dome9_continuous_compliance_notification

Use this data source to get information about a Dome9 continuous compliance notification policy.

## Example Usage

```hcl
data "dome9_continuous_compliance_notification" "test" {
  id = "d9-continuous-compliance-notification-id"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The id for the continuous compliance notification policy in Dome9. 

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Notification policy name.

* `description` -  Description of the notification.

at least one of  `alerts_console`, `scheduled_report`, or `change_detection` must be included

* `alerts_console` -  send findings (also) to the Dome9 alerts console (Boolean); default is False.

* `scheduled_report` - Scheduled email report notification block:
  * `email_sending_state` - send schedule report of findings by email; can be  "Enabled" or "Disabled".

    if `email_sending_state` is Enabled, the following must be included:

    * `schedule_data` - Schedule details:
      * `cron_expression` -  the schedule to issue the email report (in cron expression format)
      * `type` -  type of report; can be  "Detailed", "Summary", "FullCsv" or "FullCsvZip"
      * `recipients` -  comma-separated list of email recipients


* `change_detection` -  Send changes in findings options:
  * `email_sending_stat` - Whether to send email report of changes in findings; can be "Enabled" or "Disabled".

    if `email_sending_stat`  is Enabled, the following must be included:

    * `email_data` - Email notification details:
      * `recipients` -  comma-separated list of email recipients

        <br/>

  * `email_per_finding_sending_state` - send separate email notification for each finding; can be "Enabled" or "
    Disabled"

    if `email_per_finding_sending_state`  is Enabled, the following must be included:

    * `email_per_finding_data` - Email per finding notification details:
      * `recipients` -  comma-separated list of email recipients
      * `notification_output_format` -  format of JSON block for finding; can be  "JsonWithFullEntity"
        , "JsonWithBasicEntity", or "PlainText".

        <br/>

  * `sns_sending_state` - send by AWS SNS for each new finding; can be  "Enabled" or "Disabled".

    if `sns_sending_state`  is Enabled, the following must be included:

    * `sns_data` - SNS notification details:
      * `sns_topic_arn` -  SNS topic ARN
      * `sns_output_format` -  SNS output format; can be  "JsonWithFullEntity", "JsonWithBasicEntity", or "
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
      * `pass` -  password (ServiceNow only)
      * `project_key` - project key (Jira) or API Key (PagerDuty)
      * `issue_type` - issue type (Jira)

        <br/>

  * `webhook_integration_state` - send findings to an HTTP endpoint (webhook); can be  "Enabled" or "Disabled".

    if `webhook_integration_state`  is Enabled, the following must be included:

    * `webhook_data` - Webhook data block supports:
      * `url` -  HTTP endpoint URL
      * `http_method` - HTTP method, "Post" by default.
      * `auth_method` - authentication method; "NoAuth" by default
      * `username` - username in endpoint system
      * `password` - password in endpoint system
      * `format_type` - format for JSON block for finding, can be one of:
        * `JsonWithFullEntity` - JSON - Full entity (default)
        * `SplunkBasic` - Splunk format
        * `ServiceNow` - ServiceNow format
        * `QRadar` - QRadar format
        * `JsonFirstLevelEntity` - Sumo Logic format
        * `Jira` - Jira format
      * `payload_format` - Json Payload
      * `ignore_certificate` - Check this to use self-signed certificates, and ignore validation of them
      * `advanced_url` - Tick this box if you are using a version of Jira that only supports REST API 2

      <br/>

  * `slack_integration_state` - Send report summary to Slack channel (Compliance only); can be  "Enabled" or "
    Disabled".

    if `slack_integration_state`  is Enabled, the following must be included:

    * `slack_data` - Slack data block supports:
      * `url` -  Slack's webhook URL

      <br/>

  * `teams_integration_state` - Send report summary to Teams channel (Compliance only); can be  "Enabled" or "
    Disabled".

    if `teams_integration_state`  is Enabled, the following must be included:

    * `teams_data` - Teams data block supports:
      * `url` -  Teams webhook URL

      <br/>

  * `aws_security_hub_integration_state` - send findings to AWS Secure Hub; can be "Enabled" or "Disabled".

    if `aws_security_hub_integration_state`  is Enabled, the following must be included:

    * `aws_security_hub_integration` - AWS security hub integration block supports:
      * `external_account_id` -  external account id
      * `region` -  AWS region

      <br/>

`gcp_security_command_center_integration` is a change_detection option

* `gcp_security_command_center_integration` - GCP security command center details
  * `state` - send findings to the GCP Security Command Center; can be "Enabled" or "Disabled"

    if `state` is Enabled, the following must be included:

    * `project_id` - GCP Project id
    * `source_id` - GCP Source id

