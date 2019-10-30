package dome9

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_policy"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceContinuousCompliancePolicyBasic(t *testing.T) {
	var continuousCompliancePolicyResponse continuous_compliance_policy.ContinuousCompliancePolicyResponse
	policyResourceTypeAndName, _, policyGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousCompliancePolicy)
	cloudAccountResourceTypeAndName, _, cloudAccountGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAzure)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAzureEnvVarsPreCheck(t)                // Azure account used as an input for policy creation thus this check is required
			testAccContinuousComplianceNotificationEnvVarsPreCheck(t) // Continuous compliance notifications IDs used as an input for test policy creation thus this check is required
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContinuousCompliancePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckContinuousCompliancePolicyBasic(policyGeneratedName, cloudAccountGeneratedName, cloudAccountResourceTypeAndName, policyResourceTypeAndName, getNotificationIDsConfig()),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContinuousCompliancePolicyExists(policyResourceTypeAndName, &continuousCompliancePolicyResponse),
					resource.TestCheckResourceAttr(policyResourceTypeAndName, "cloud_account_type", strings.Title(variable.CloudAccountAzureVendor)),
					resource.TestCheckResourceAttr(policyResourceTypeAndName, "notification_ids.#", "1"),
				),
			},
			{
				Config: testAccCheckContinuousCompliancePolicyBasic(policyGeneratedName, cloudAccountGeneratedName, cloudAccountResourceTypeAndName, policyResourceTypeAndName, getUpdateNotificationIDsConfig()),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContinuousCompliancePolicyExists(policyResourceTypeAndName, &continuousCompliancePolicyResponse),
					resource.TestCheckResourceAttr(policyResourceTypeAndName, "notification_ids.#", "2"),
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
		if rs.Type != resourcetype.ContinuousCompliancePolicy {
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

func testAccContinuousComplianceNotificationEnvVarsPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.ContinuousCompliancePolicyEnvVarNotificationId1); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.ContinuousCompliancePolicyEnvVarNotificationId1)
	}
	if v := os.Getenv(environmentvariable.ContinuousCompliancePolicyEnvVarNotificationId2); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.ContinuousCompliancePolicyEnvVarNotificationId2)
	}
}

func testAccCheckContinuousCompliancePolicyBasic(policyGeneratedName, cloudAccountGeneratedName, cloudAccountResourceTypeAndName, policyResourceTypeAndName, NotificationIDsConfig string) string {
	return fmt.Sprintf(`
// azure cloud account creation
%s

// continuous compliance policy creation
resource "%s" "%s" {
  cloud_account_id    = "${%s.id}"
  external_account_id = "${%s.subscription_id}"
  cloud_account_type  = "%s"
  notification_ids    = %s
}

// continuous compliance policy data source
data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// Azure cloud account
		getCloudAccountAzureConfig(cloudAccountGeneratedName, cloudAccountResourceTypeAndName),

		// Continuous compliance policy resource variables
		resourcetype.ContinuousCompliancePolicy,
		policyGeneratedName,
		cloudAccountResourceTypeAndName,
		cloudAccountResourceTypeAndName,
		strings.Title(variable.CloudAccountAzureVendor),
		NotificationIDsConfig,

		// Continuous compliance policy data source variables
		resourcetype.ContinuousCompliancePolicy,
		policyGeneratedName,
		policyResourceTypeAndName,
	)
}

func getNotificationIDsConfig() string {
	return fmt.Sprintf("[\"%s\"]", os.Getenv(environmentvariable.ContinuousCompliancePolicyEnvVarNotificationId1))
}

func getUpdateNotificationIDsConfig() string {
	return fmt.Sprintf("[\"%s\", \"%s\"]", os.Getenv(environmentvariable.ContinuousCompliancePolicyEnvVarNotificationId1), os.Getenv(environmentvariable.ContinuousCompliancePolicyEnvVarNotificationId2))
}
