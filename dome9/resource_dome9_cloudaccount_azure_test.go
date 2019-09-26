package dome9

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/azure"

	"github.com/dome9/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceCloudAccountAzureBasic(t *testing.T) {
	var cloudAccountAzure azure.CloudAccountResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAzure)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAzureEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAccountAzureDestroy,
		Steps: []resource.TestStep{
			{
				// creation test
				Config: testAccCheckCloudAccountAzureConfigure(resourceTypeAndName, generatedName, variable.CloudAccountAzureCreationResourceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAzureExists(resourceTypeAndName, &cloudAccountAzure),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountAzureCreationResourceName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "operation_mode", variable.CloudAccountAzureOperationMode),
					resource.TestCheckResourceAttr(resourceTypeAndName, "vendor", variable.CloudAccountAzureVendor),
					resource.TestCheckResourceAttr(resourceTypeAndName, "organizational_unit_name", os.Getenv(environmentvariable.OrganizationalUnitName)),
				),
			},
			{
				// update name test
				Config: testAccCheckCloudAccountAzureConfigure(resourceTypeAndName, generatedName, variable.CloudAccountAzureUpdatedAccountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAzureExists(resourceTypeAndName, &cloudAccountAzure),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountAzureUpdatedAccountName),
				),
			},
		},
	})
}

func testAccCheckCloudAccountAzureDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.CloudAccountAzure {
			continue
		}

		getCloudAccountQueryParams := cloudaccounts.QueryParameters{ID: rs.Primary.ID}
		resp, _, err := apiClient.cloudaccountAzure.Get(&getCloudAccountQueryParams)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if resp != nil {
			return fmt.Errorf("cloudaccounts with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCloudAccountAzureEnvVarsPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.CloudAccountAzureEnvVarClientId); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountAzureEnvVarClientId)
	}
	if v := os.Getenv(environmentvariable.CloudAccountAzureEnvVarClientPassword); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountAzureEnvVarClientPassword)
	}
	if v := os.Getenv(environmentvariable.CloudAccountAzureEnvVarSubscriptionId); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountAzureEnvVarSubscriptionId)
	}
	if v := os.Getenv(environmentvariable.CloudAccountAzureEnvVarTenantId); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountAzureEnvVarTenantId)
	}
	if v := os.Getenv(environmentvariable.OrganizationalUnitName); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.OrganizationalUnitName)
	}
}

func testAccCheckCloudAccountAzureExists(resource string, resp *azure.CloudAccountResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		getCloudAccountQueryParams := cloudaccounts.QueryParameters{ID: rs.Primary.ID}
		receivedCloudAccount, _, err := apiClient.cloudaccountAzure.Get(&getCloudAccountQueryParams)
		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}

		*resp = *receivedCloudAccount

		return nil
	}
}

func testAccCheckCloudAccountAzureConfigure(resourceTypeAndName, generatedName, resourceName string) string {
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
data "%s" "%s" {
  account_id = "${%s.id}"
}
`,
		// resource variable
		resourcetype.CloudAccountAzure,
		generatedName,
		os.Getenv(environmentvariable.CloudAccountAzureEnvVarClientId),
		os.Getenv(environmentvariable.CloudAccountAzureEnvVarClientPassword),
		resourceName,
		variable.CloudAccountAzureOperationMode,
		os.Getenv(environmentvariable.CloudAccountAzureEnvVarSubscriptionId),
		os.Getenv(environmentvariable.CloudAccountAzureEnvVarTenantId),

		// data source variable
		resourcetype.CloudAccountAzure,
		generatedName,
		resourceTypeAndName,
	)
}
