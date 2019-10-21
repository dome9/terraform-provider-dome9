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
			testAccCloudAccountAzureEnvVarsPreCheck(t) // Azure account used as an input for policy creation thus this check is required
			testAccContinuousCompliancePolicyEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContinuousCompliancePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckContinuousCompliancePolicyBasic(policyGeneratedName, cloudAccountGeneratedName, cloudAccountResourceTypeAndName, policyResourceTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContinuousCompliancePolicyExists(policyResourceTypeAndName, &continuousCompliancePolicyResponse),
					resource.TestCheckResourceAttr(policyResourceTypeAndName, "cloud_account_type", strings.Title(variable.CloudAccountAzureVendor)),
					resource.TestCheckResourceAttr(policyResourceTypeAndName, "notification_ids.#", "1"),
				),
			},
		},
	})
}

func testAccContinuousCompliancePolicyEnvVarsPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.ContinuousCompliancePolicyEnvVarNotificationId); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.ContinuousCompliancePolicyEnvVarNotificationId)
	}
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
		if rs.Type != resourcetype.CloudAccountAWS {
			continue
		}

		receivedCloudAccountResponse, _, err := apiClient.continuousCompliancePolicy.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if receivedCloudAccountResponse != nil {
			return fmt.Errorf("iplist with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckContinuousCompliancePolicyBasic(policyGeneratedName, cloudAccountGeneratedName, cloudAccountResourceTypeAndName, policyResourceTypeAndName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  credentials = {
    client_id = "%s"
    client_password = "%s"
  }

  name = "%s"
  operation_mode = "%s"
  subscription_id = "%s"
  tenant_id = "%s"
}

// continuous compliance policy creation
resource "%s" "%s" {
  cloud_account_id = "${%s.id}"
  external_account_id = "${%s.subscription_id}"
  cloud_account_type = "%s"
  notification_ids = ["%s"]
}

// continuous compliance policy data source
data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// Azure cloud account resource variable
		resourcetype.CloudAccountAzure,
		cloudAccountGeneratedName,
		os.Getenv(environmentvariable.CloudAccountAzureEnvVarClientId),
		os.Getenv(environmentvariable.CloudAccountAzureEnvVarClientPassword),
		variable.CloudAccountAzureCreationResourceName,
		variable.CloudAccountAzureOperationMode,
		os.Getenv(environmentvariable.CloudAccountAzureEnvVarSubscriptionId),
		os.Getenv(environmentvariable.CloudAccountAzureEnvVarTenantId),

		// Continuous Compliance Policy resource
		resourcetype.ContinuousCompliancePolicy,
		policyGeneratedName,
		cloudAccountResourceTypeAndName,
		cloudAccountResourceTypeAndName,
		strings.Title(variable.CloudAccountAzureVendor),
		os.Getenv(environmentvariable.ContinuousCompliancePolicyEnvVarNotificationId),

		// Continuous Compliance Policy data source
		resourcetype.ContinuousCompliancePolicy,
		policyGeneratedName,
		policyResourceTypeAndName,
	)
}
