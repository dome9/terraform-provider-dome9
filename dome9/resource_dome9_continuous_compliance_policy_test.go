package dome9

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_policy"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceContinuousCompliancePolicyBasic(t *testing.T) {
	var continuousCompliancePolicyResponse continuous_compliance_policy.ContinuousCompliancePolicyResponse
	policyTypeAndName, _, policyGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousCompliancePolicy)
	azureTypeAndName, _, azureGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAzure)
	notificationTypeAndName, _, notificationGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousComplianceNotification)
	notificationUpdateTypeAndName, _, notificationUpdateGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousComplianceNotification)

	azureHCL := getCloudAccountAzureResourceHCL(azureGeneratedName, variable.CloudAccountAzureCreationResourceName)
	notificationHCL := getContinuousComplianceNotificationResourceHCL(notificationGeneratedName, continuousComplianceNotificationConfig())
	// notificationHCL := getUpdateContinuousCompliancePolicyResourceHCL1(azureHCL, azureTypeAndName, notificationHCL, notificationTypeAndName, notificationUpdateTypeAndName, notificationUpdateHCL, policyGeneratedName)
	notificationUpdateHCL := getContinuousComplianceNotificationResourceHCL(notificationUpdateGeneratedName, continuousComplianceNotificationUpdateConfig())
	policyHCL := getContinuousCompliancePolicyResourceHCL(azureHCL, azureTypeAndName, notificationHCL, notificationTypeAndName, policyGeneratedName)
	updatePolicyHCL := getUpdateContinuousCompliancePolicyResourceHCL(azureHCL, azureTypeAndName, notificationHCL, notificationUpdateTypeAndName, notificationUpdateHCL, policyGeneratedName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAzureEnvVarsPreCheck(t) // Azure account used as an input for policy creation thus this check is required
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContinuousCompliancePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckContinuousCompliancePolicyBasic(policyHCL, policyGeneratedName, policyTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContinuousCompliancePolicyExists(policyTypeAndName, &continuousCompliancePolicyResponse),
					resource.TestCheckResourceAttr(policyTypeAndName, "cloud_account_type", strings.Title(variable.CloudAccountAzureVendor)),
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
		if rs.Type != resourcetype.ContinuousCompliancePolicy && rs.Type != resourcetype.ContinuousComplianceNotification && rs.Type != resourcetype.CloudAccountAzure {
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
// azure cloud account resource
%s

// continuous compliance notification resource
%s

// continuous compliance policy creation
resource "%s" "%s" {
  cloud_account_id    = "${%s.id}"
  external_account_id = "${%s.subscription_id}"
  cloud_account_type  = "%s"
  notification_ids    = ["${%s.id}"]
}
`,
		// azure cloud account resource
		cloudAccountHCL,

		// continuous compliance notification resource
		notificationHCL,

		// Continuous compliance policy resource variables
		resourcetype.ContinuousCompliancePolicy,
		policyName,
		cloudAccountTypeAndName,
		cloudAccountTypeAndName,
		strings.Title(variable.CloudAccountAzureVendor),
		notificationTypeAndName,
	)
}

func getUpdateContinuousCompliancePolicyResourceHCL(cloudAccountHCL, cloudAccountTypeAndName, notificationHCL, updateNotificationTypeAndName, updateNotificationHCL, policyName string) string {
	return fmt.Sprintf(`
// azure cloud account resource
%s

// continuous compliance notification resource
%s

%s

// continuous compliance policy creation
resource "%s" "%s" {
  cloud_account_id    = "${%s.id}"
  external_account_id = "${%s.subscription_id}"
  cloud_account_type  = "%s"
  notification_ids    = ["${%s.id}"]
}
`,
		// azure cloud account resource
		cloudAccountHCL,

		// continuous compliance notification resource
		notificationHCL,
		updateNotificationHCL,

		// Continuous compliance policy resource variables
		resourcetype.ContinuousCompliancePolicy,
		policyName,
		cloudAccountTypeAndName,
		cloudAccountTypeAndName,
		strings.Title(variable.CloudAccountAzureVendor),
		updateNotificationTypeAndName,
	)
}
