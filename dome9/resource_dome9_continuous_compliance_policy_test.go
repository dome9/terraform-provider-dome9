package dome9

import (
	"fmt"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_policy"

	"github.com/dome9/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceContinuousCompliancePolicyBasic(t *testing.T) {
	var continuousCompliancePolicyResponse continuous_compliance_policy.ContinuousCompliancePolicyResponse
	policyTypeAndName, _, policyGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousCompliancePolicy)
	awsTypeAndName, _, awsGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAWS)
	notificationTypeAndName, _, notificationGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousComplianceNotification)
	notificationUpdateTypeAndName, _, notificationUpdateGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousComplianceNotification)

	awsHCL := getCloudAccountAWSResourceHCL(awsGeneratedName, variable.CloudAccountAWSOriginalAccountName, os.Getenv(environmentvariable.CloudAccountAWSEnvVarArn), "")
	notificationHCL := getContinuousComplianceNotificationResourceHCL(notificationGeneratedName, continuousComplianceNotificationConfig())
	// notificationHCL := getUpdateContinuousCompliancePolicyResourceHCL1(awsHCL, awsTypeAndName, notificationHCL, notificationTypeAndName, notificationUpdateTypeAndName, notificationUpdateHCL, policyGeneratedName)
	notificationUpdateHCL := getContinuousComplianceNotificationResourceHCL(notificationUpdateGeneratedName, continuousComplianceNotificationUpdateConfig())
	policyHCL := getContinuousCompliancePolicyResourceHCL(awsHCL, awsTypeAndName, notificationHCL, notificationTypeAndName, policyGeneratedName)
	updatePolicyHCL := getUpdateContinuousCompliancePolicyResourceHCL(awsHCL, awsTypeAndName, notificationHCL, notificationUpdateTypeAndName, notificationUpdateHCL, policyGeneratedName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAWSEnvVarsPreCheck(t) // Aws account used as an input for policy creation thus this check is required
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContinuousCompliancePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckContinuousCompliancePolicyBasic(policyHCL, policyGeneratedName, policyTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContinuousCompliancePolicyExists(policyTypeAndName, &continuousCompliancePolicyResponse),
					resource.TestCheckResourceAttr(policyTypeAndName, "target_type", strings.Title(variable.CloudAccountAWSVendor)),
					resource.TestCheckResourceAttr(policyTypeAndName, "notification_ids.#", "1"),
				),
			},

			{
				Config: testAccCheckContinuousCompliancePolicyBasic(updatePolicyHCL, policyGeneratedName, policyTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContinuousCompliancePolicyExists(policyTypeAndName, &continuousCompliancePolicyResponse),
					resource.TestCheckResourceAttr(policyTypeAndName, "notification_ids.#", "1"),
				),
			},
		},
	})
}

func testAccCheckContinuousCompliancePolicyExists(resource string, cloudAccount *continuous_compliance_policy.ContinuousCompliancePolicyResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedContinuousCompliancePolicyResponse, _, err := apiClient.continuousCompliancePolicy.Get(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*cloudAccount = *receivedContinuousCompliancePolicyResponse

		return nil
	}
}

func testAccCheckContinuousCompliancePolicyDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.ContinuousCompliancePolicy && rs.Type != resourcetype.ContinuousComplianceNotification && rs.Type != resourcetype.CloudAccountAWS {
			continue
		}

		receivedCloudAccountResponse, _, err := apiClient.continuousCompliancePolicy.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if receivedCloudAccountResponse != nil {
			return fmt.Errorf("continuous compliance policy with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckContinuousCompliancePolicyBasic(policyHCL, policyName, policyTypeAndName string) string {
	return fmt.Sprintf(`
// continuous compliance policy resource
%s

// continuous compliance policy data source
data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// continuous compliance policy resource
		policyHCL,
		// Continuous compliance policy data source variables
		resourcetype.ContinuousCompliancePolicy,
		policyName,
		policyTypeAndName,
	)
}

func getContinuousCompliancePolicyResourceHCL(cloudAccountHCL, cloudAccountTypeAndName, notificationHCL, notificationTypeAndName, policyName string) string {
	return fmt.Sprintf(`
// aws cloud account resource
%s

// continuous compliance notification resource
%s

// continuous compliance policy creation
resource "%s" "%s" {
  target_id    = "${%s.id}"
  ruleset_id   = "%d"
  target_type  = "%s"
  notification_ids    = ["${%s.id}"]
}
`,
		// aws cloud account resource
		cloudAccountHCL,

		// continuous compliance notification resource
		notificationHCL,

		// Continuous compliance policy resource variables
		resourcetype.ContinuousCompliancePolicy,
		policyName,
		cloudAccountTypeAndName,
		variable.ContinuousCompliancePolicyRulesetId,
		strings.Title(variable.CloudAccountAWSVendor),
		notificationTypeAndName,
	)
}

func getUpdateContinuousCompliancePolicyResourceHCL(cloudAccountHCL, cloudAccountTypeAndName, notificationHCL, updateNotificationTypeAndName, updateNotificationHCL, policyName string) string {
	return fmt.Sprintf(`
// aws cloud account resource
%s

// continuous compliance notification resource
%s

%s

// continuous compliance policy creation
resource "%s" "%s" {
  target_id    = "${%s.id}"
  ruleset_id   = "%d"
  target_type  = "%s"
  notification_ids    = ["${%s.id}"]
}
`,
		// aws cloud account resource
		cloudAccountHCL,

		// continuous compliance notification resource
		notificationHCL,
		updateNotificationHCL,

		// Continuous compliance policy resource variables
		resourcetype.ContinuousCompliancePolicy,
		policyName,
		cloudAccountTypeAndName,
		variable.ContinuousCompliancePolicyRulesetId,
		strings.Title(variable.CloudAccountAWSVendor),
		updateNotificationTypeAndName,
	)
}
