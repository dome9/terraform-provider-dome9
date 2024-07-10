package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/notifications"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"strconv"
	"testing"
)

func TestAccResourceNotificationBasic(t *testing.T) {
	var notificationResponse notifications.ResponseNotificationViewModel
	notificationTypeAndName, _, notificationGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.Notification)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNotificationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckNotificationBasic(notificationTypeAndName, notificationGeneratedName, notificationConfig(notificationGeneratedName)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationExists(notificationTypeAndName, &notificationResponse),
					resource.TestCheckResourceAttr(notificationTypeAndName, "name", variable.NotificationName+"_"+notificationGeneratedName),
					resource.TestCheckResourceAttr(notificationTypeAndName, "description", variable.NotificationDescription),
					resource.TestCheckResourceAttr(notificationTypeAndName, "alerts_console", strconv.FormatBool(variable.NotificationAlertsConsole)),
					resource.TestCheckResourceAttr(notificationTypeAndName, "send_on_each_occurrence", strconv.FormatBool(variable.NotificationSendOnEachOccurrence)),
				),
			},
			{
				// update name test
				Config: testAccCheckNotificationBasic(notificationTypeAndName, notificationGeneratedName, notificationUpdateConfig(notificationGeneratedName)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationExists(notificationTypeAndName, &notificationResponse),
					resource.TestCheckResourceAttr(notificationTypeAndName, "name", variable.NotificationUpdateName+"_"+notificationGeneratedName),
					resource.TestCheckResourceAttr(notificationTypeAndName, "description", variable.NotificationUpdateDescription),
					resource.TestCheckResourceAttr(notificationTypeAndName, "alerts_console", strconv.FormatBool(variable.NotificationUpdateAlertsConsole)),
					resource.TestCheckResourceAttr(notificationTypeAndName, "send_on_each_occurrence", strconv.FormatBool(variable.NotificationSendOnEachOccurrence)),
				),
			},
		},
	})
}

func testAccCheckNotificationExists(resource string, notification *notifications.ResponseNotificationViewModel) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedNotificationResponse, _, err := apiClient.notifications.GetById(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*notification = *receivedNotificationResponse

		return nil
	}
}

func testAccCheckNotificationDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.Notification {
			continue
		}

		receivedNotificationResponse, _, err := apiClient.notifications.GetById(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if receivedNotificationResponse != nil {
			return fmt.Errorf("notification with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckNotificationBasic(resourceTypeAndName, generatedName, additionalBlock string) string {
	return fmt.Sprintf(`
// notification resource
%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// notification resource
		getNotificationResourceHCL(generatedName, additionalBlock),

		// data source variables
		resourcetype.Notification,
		generatedName,
		resourceTypeAndName,
	)
}

func getNotificationResourceHCL(generatedName, additionalBlock string) string {
	return fmt.Sprintf(`
// notification creation
resource "%s" "%s" {
  %s

}
`,
		// resource variables
		resourcetype.Notification,
		generatedName,
		additionalBlock,
	)
}

func notificationConfig(notificationNameSuffix string) string {
	return fmt.Sprintf(`
name           = "%s"
description    = "%s"
alerts_console = "%s"
send_on_each_occurrence = "%s"

`,
		variable.NotificationName+"_"+notificationNameSuffix,
		variable.NotificationDescription,
		strconv.FormatBool(variable.NotificationAlertsConsole),
		strconv.FormatBool(variable.NotificationSendOnEachOccurrence),
	)
}

func notificationUpdateConfig(notificationSuffixName string) string {
	return fmt.Sprintf(`
name           = "%s"
description    = "%s"
alerts_console = "%s"
send_on_each_occurrence = "%s"
`,
		variable.NotificationUpdateName+"_"+notificationSuffixName,
		variable.NotificationUpdateDescription,
		strconv.FormatBool(variable.NotificationUpdateAlertsConsole),
		strconv.FormatBool(variable.NotificationSendOnEachOccurrence),
	)
}
