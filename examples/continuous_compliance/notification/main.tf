resource "dome9_continuouscompliance_notification" "test_notification" {
  name           = "NOTIFICATION_NAME"
  description    = "NOTIFICATION_DESCRIPTION"
  alerts_console = true
  scheduled_report {
    email_sending_state = "EMAIL_SENDING_STATE" // options: ["Enabled", "Disabled"]
    schedule_data {
      cron_expression = "CRON_EXPRESSION"
      type            = "TYPE" // options: ["Detailed", "Summary", "FullCsv", "FullCsvZip"]
      recipients = [
        "test@test.com"
      ]
    }
  }

  change_detection {
    //    All the states can be Enabled or Disabled
    email_sending_state                = "EMAIL_SENDING_STATE"
    email_per_finding_sending_state    = "EMAIL_PER_FINDING_SENDING_STATE"
    sns_sending_state                  = "SNS_SENDING_STATE"
    external_ticket_creating_state     = "EXTERNAL_TICKET_CREATING_STATE"
    aws_security_hub_integration_state = "AWS_SECURITY_HUB_INTEGRATION_STATE"
    webhook_integration_state          = "WEBHOOK_INTEGRATION_STATE"

    email_data {
      recipients = [
        "test@test.com"
      ]
    }

    email_per_finding_data {
      notification_output_format = "NOTIFICATION_OUTPUT_FORMAT" // options: ["JsonWithFullEntity", "JsonWithBasicEntity", "PlainText"]
      recipients = [
        "test@test.com"
      ]
    }

    sns_data {
      sns_topic_arn     = "SNS_TOPIC_ARN"
      sns_output_format = "SNS_OUTPUT_FORMAT" // options: ["JsonWithFullEntity", "JsonWithBasicEntity", "PlainText"]
    }

    ticketing_system_data {
      system_type          = "SYSTEM_TYPE"
      should_close_tickets = false
      domain               = "DOMAIN"
      user                 = "USER"
      pass                 = "PASS"
      project_key          = "PROJECT_KE"
      issue_type           = "ISSUE_TYPE"
    }

    aws_security_hub_integration {
      external_account_id = "EXTERNAL_ACCOUNT_ID"
      region              = "AWS_REGION"
    }

    webhook_data {
      url         = "URL"
      http_method = "Post"
      auth_method = "AUTH_METHOD" // options: ["NoAuth", "BasicAuth"]
      username    = "USERNAME"
      password    = "PASSWORD"
      format_type = "FORMAT_TYPE" // options: ["JsonWithFullEntity", "SplunkBasic", "ServiceNow"]
    }
  }
  gcp_security_command_center_integration {
    state      = "STATE" // options: ["Enabled", "Disabled"]
    project_id = "PROJECT_ID"
    source_id  = "SOURCE_ID"
  }
}
