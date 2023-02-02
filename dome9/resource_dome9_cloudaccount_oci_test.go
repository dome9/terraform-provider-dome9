package dome9

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceCloudAccountOciBasic(t *testing.T) {
	var cloudAccountOci Oci.CloudAccountResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountOci)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountOciEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAccountOciDestroy,
		Steps: []resource.TestStep{
			{
				// creation test
				Config: testAccCheckCloudAccountOciConfigure(resourceTypeAndName, generatedName, variable.CloudAccountOciCreationResourceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountOciExists(resourceTypeAndName, &cloudAccountOci),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountOciCreationResourceName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "vendor", variable.CloudAccountOciVendor),
					resource.TestCheckResourceAttr(resourceTypeAndName, "organizational_unit_name", os.Getenv(environmentvariable.OrganizationalUnitName)),
				),
			},
			{
				// update name test
				Config: testAccCheckCloudAccountOciConfigure(resourceTypeAndName, generatedName, variable.CloudAccountOciUpdatedAccountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountOciExists(resourceTypeAndName, &cloudAccountOci),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountOciUpdatedAccountName),
				),
			},
		},
	})
}

func testAccCheckCloudAccountOciDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.CloudAccountOci {
			continue
		}

		resp, _, err := apiClient.cloudaccountOci.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if resp != nil {
			return fmt.Errorf("cloudaccounts with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCloudAccountOciEnvVarsPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.CloudAccountOciEnvVarAccessKey); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountOciEnvVarAccessKey)
	}
	if v := os.Getenv(environmentvariable.CloudAccountOciEnvVarAccessSecret); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountOciEnvVarAccessSecret)
	}
	if v := os.Getenv(environmentvariable.OrganizationalUnitName); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.OrganizationalUnitName)
	}
}

func testAccCheckCloudAccountOciExists(resource string, resp *Oci.CloudAccountResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedCloudAccount, _, err := apiClient.cloudaccountOci.Get(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}

		*resp = *receivedCloudAccount

		return nil
	}
}

func testAccCheckCloudAccountOciConfigure(resourceTypeAndName, generatedName, resourceName string) string {
	return fmt.Sprintf(`
// oci cloud account creation
%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// oci cloud account
		getCloudAccountOciResourceHCL(generatedName, resourceName),

		// data source variables
		resourcetype.CloudAccountOci,
		generatedName,
		resourceTypeAndName,
	)
}

func getCloudAccountOciResourceHCL(cloudAccountName, generatedAName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
	credentials = {
		access_key    = "%s"
		access_secret = "%s"
}
	name          = "%s"
}
`,
		// oci cloud account variables
		resourcetype.CloudAccountOci,
		cloudAccountName,
		os.Getenv(environmentvariable.CloudAccountOciEnvVarAccessKey),
		os.Getenv(environmentvariable.CloudAccountOciEnvVarAccessSecret),
		generatedAName,
	)
}
