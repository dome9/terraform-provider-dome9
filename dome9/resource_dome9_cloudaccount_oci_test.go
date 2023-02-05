package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/oci"
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
	var cloudAccountOci oci.CloudAccountResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountOCI)

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
				Config: testAccCheckCloudAccountOciConfigure(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountOciExists(resourceTypeAndName, &cloudAccountOci),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountOciCreationResourceName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "vendor", variable.CloudAccountOciVendor),
				),
			},
		},
	})
}

func testAccCheckCloudAccountOciDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.CloudAccountOCI {
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
	if v := os.Getenv(environmentvariable.CloudAccountOciEnvVarTenancyId); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountOciEnvVarTenancyId)
	}
	if v := os.Getenv(environmentvariable.CloudAccountOciEnvVarUserOcid); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountOciEnvVarUserOcid)
	}
}

func testAccCheckCloudAccountOciExists(resource string, resp *oci.CloudAccountResponse) resource.TestCheckFunc {
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

func testAccCheckCloudAccountOciConfigure(resourceTypeAndName, generatedName string) string {
	return fmt.Sprintf(`
// Oci cloud account creation
%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// Oci cloud account
		getCloudAccountOciResourceHCL(generatedName),

		// data source variables
		resourcetype.CloudAccountOCI,
		generatedName,
		resourceTypeAndName,
	)
}

func getCloudAccountOciResourceHCL(cloudAccountName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
	tenancy_id = "%s"
	user_ocid  = "%s"
}
`,
		// Oci cloud account variables
		resourcetype.CloudAccountOCI,
		cloudAccountName,
		os.Getenv(environmentvariable.CloudAccountOciEnvVarTenancyId),
		os.Getenv(environmentvariable.CloudAccountOciEnvVarUserOcid),
	)
}
