package dome9

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_notification"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceContinuousComplianceNotificationBasic(t *testing.T) {
	var continuousComplianceNotificationResponse continuous_compliance_notification.ContinuousComplianceNotificationResponse
	notificationTypeAndName, _, notificationGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousComplianceNotification)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContinuousComplianceNotificationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckContinuousComplianceNotificationBasic(notificationTypeAndName, notificationGeneratedName, continuousComplianceNotificationConfig(notificationGeneratedName)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContinuousComplianceNotificationExists(notificationTypeAndName, &continuousComplianceNotificationResponse),
					resource.TestCheckResourceAttr(notificationTypeAndName, "name", variable.ContinuousComplianceNotificationName+"_"+notificationGeneratedName),
					resource.TestCheckResourceAttr(notificationTypeAndName, "description", variable.ContinuousComplianceNotificationDescription),
					resource.TestCheckResourceAttr(notificationTypeAndName, "alerts_console", strconv.FormatBool(variable.ContinuousComplianceNotificationAlertsConsole)),
					resource.TestCheckResourceAttr(notificationTypeAndName, "scheduled_report.#", "1"),
					resource.TestCheckResourceAttr(notificationTypeAndName, "change_detection.#", "1"),
				),
			},
			{
				// update name test
				Config: testAccCheckContinuousComplianceNotificationBasic(notificationTypeAndName, notificationGeneratedName, continuousComplianceNotificationUpdateConfig(notificationGeneratedName)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContinuousComplianceNotificationExists(notificationTypeAndName, &continuousComplianceNotificationResponse),
					resource.TestCheckResourceAttr(notificationTypeAndName, "name", variable.ContinuousComplianceNotificationUpdateName),
					resource.TestCheckResourceAttr(notificationTypeAndName, "description", variable.ContinuousComplianceNotificationUpdateDescription),
					resource.TestCheckResourceAttr(notificationTypeAndName, "alerts_console", strconv.FormatBool(variable.ContinuousComplianceNotificationUpdateAlertsConsole)),
					resource.TestCheckResourceAttr(notificationTypeAndName, "scheduled_report.#", "0"),
					resource.TestCheckResourceAttr(notificationTypeAndName, "change_detection.#", "1"),
				),
			},
		},
	})
}

func testAccCheckContinuousComplianceNotificationExists(resource string, continuousComplianceNotification *continuous_compliance_notification.ContinuousComplianceNotificationResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedContinuousComplianceNotificationResponse, _, err := apiClient.continuousComplianceNotification.Get(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*continuousComplianceNotification = *receivedContinuousComplianceNotificationResponse

		return nil
	}
}

func testAccCheckContinuousComplianceNotificationDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.ContinuousComplianceNotification {
			continue
		}

		receivedContinuousComplianceNotificationResponse, _, err := apiClient.continuousComplianceNotification.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if receivedContinuousComplianceNotificationResponse != nil {
			return fmt.Errorf("notification with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckContinuousComplianceNotificationBasic(resourceTypeAndName, generatedName, additionalBlock string) string {
	return fmt.Sprintf(`
// continuous compliance notification resource
%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// continuous compliance notification resource
		getContinuousComplianceNotificationResourceHCL(generatedName, additionalBlock),

		// data source variables
		resourcetype.ContinuousComplianceNotification,
		generatedName,
		resourceTypeAndName,
	)
}

func getContinuousComplianceNotificationResourceHCL(generatedName, additionalBlock string) string {
	return fmt.Sprintf(`
// continuous compliance notification creation
resource "%s" "%s" {
  %s
  change_detection {
    email_sending_state                = "%s"
    email_per_finding_sending_state    = "%s"
    sns_sending_state                  = "%s"
    external_ticket_creating_state     = "%s"
    aws_security_hub_integration_state = "%s"
    webhook_integration_state          = "%s"

    email_data {
      recipients = ["%s"]
    }

    email_per_finding_data {
      recipients                 = ["%s"]
      notification_output_format = "%s"
    }
  }
}
`,
		// resource variables
		resourcetype.ContinuousComplianceNotification,
		generatedName,

		additionalBlock,

		variable.ContinuousComplianceNotificationEnabled,
		variable.ContinuousComplianceNotificationEnabled,
		variable.ContinuousComplianceNotificationDisabled,
		variable.ContinuousComplianceNotificationDisabled,
		variable.ContinuousComplianceNotificationDisabled,
		variable.ContinuousComplianceNotificationDisabled,

		variable.ContinuousComplianceNotificationRecipient,

		variable.ContinuousComplianceNotificationRecipient,
		variable.ContinuousComplianceNotificationJsonWithFullEntity,
	)
}

func continuousComplianceNotificationConfig(notificationNameSuffix string) string {
	return fmt.Sprintf(`
name           = "%s"
description    = "%s"
alerts_console = "%s"

scheduled_report {
  email_sending_state = "%s"
  schedule_data {
    cron_expression = "%s"
    type            = "%s"
    recipients      = ["%s"]
  }
}
`,
		variable.ContinuousComplianceNotificationName+"_"+notificationNameSuffix,
		variable.ContinuousComplianceNotificationDescription,
		strconv.FormatBool(variable.ContinuousComplianceNotificationAlertsConsole),

		variable.ContinuousComplianceNotificationEnabled,
		variable.ContinuousComplianceNotificationCronExpression,
		variable.ContinuousComplianceNotificationType,
		variable.ContinuousComplianceNotificationRecipient,
	)
}

func continuousComplianceNotificationUpdateConfig(notificationSuffixName string) string {
	return fmt.Sprintf(`
name           = "%s"
description    = "%s"
alerts_console = "%s"
`,
		variable.ContinuousComplianceNotificationUpdateName+"_"+notificationSuffixName,
		variable.ContinuousComplianceNotificationUpdateDescription,
		strconv.FormatBool(variable.ContinuousComplianceNotificationUpdateAlertsConsole),
	)
}
