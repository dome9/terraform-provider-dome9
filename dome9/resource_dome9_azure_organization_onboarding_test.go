package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/azure_org"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"os"
	"regexp"
	"testing"
)

func TestAccResourceAzureOrganizationOnboardingBasic(t *testing.T) {
	var response azure_org.OrganizationManagementViewModel
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AzureOrganizationOnboarding)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccAzureOrganizationOnboardingEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAzureOrganizationOnboardingDestroy,
		Steps: []resource.TestStep{
			{
				// creation test
				Config: testAccCheckAzureOrganizationOnboardingConfigure(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAzureOrganizationOnboardingExists(resourceTypeAndName, &response),
					resource.TestCheckResourceAttr(resourceTypeAndName, "tenant_id", os.Getenv(environmentvariable.CloudAccountAzureEnvVarTenantId)),
				),
				ExpectError: regexp.MustCompile(`.+Please ensure that the shell script has completed successfully.+`),
			},
		},
	})
}

func testAccCheckAzureOrganizationOnboardingExists(resource string, resp *azure_org.OrganizationManagementViewModel) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedCloudAccount, _, err := apiClient.azureOrganizationOnboarding.Get(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}

		*resp = *receivedCloudAccount

		return nil
	}
}

func testAccCheckAzureOrganizationOnboardingDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.AzureOrganizationOnboarding {
			continue
		}

		resp, _, err := apiClient.azureOrganizationOnboarding.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if resp != nil {
			return fmt.Errorf("azure org entity with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccAzureOrganizationOnboardingEnvVarsPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.CloudAccountAzureEnvVarTenantId); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountAzureEnvVarTenantId)
	}
}

func testAccCheckAzureOrganizationOnboardingConfigure(resourceTypeAndName, generatedName string) string {
	return fmt.Sprintf(`
// Resource creation
%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// Resource code
		getAzureOrganizationOnboardingResourceHCL(generatedName),

		// Data source variables
		resourcetype.AzureOrganizationOnboarding,
		generatedName,
		resourceTypeAndName,
	)
}

func getAzureOrganizationOnboardingResourceHCL(resourceName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  tenant_id = "%s"
  management_group_id = "%s"
  active_blades {
    awp {
      is_enabled                  = false
    }
    cdr {
      is_enabled = false
    }
    serverless {
      is_enabled = false
    }
    posture_management {
      onboarding_mode = "read"
    }
  }
  use_cloud_guard_managed_app = true
  is_auto_onboarding = true
  vendor = "%s"
}
`,
		//Resource variables
		resourcetype.AzureOrganizationOnboarding,
		resourceName,
		os.Getenv(environmentvariable.CloudAccountOrgAzureEnvVarTenantId),
		os.Getenv(environmentvariable.CloudAccountOrgAzureEnvVarManagementGroupId),
		variable.AzureOrganizationOnboardingVendorName,
	)
}
