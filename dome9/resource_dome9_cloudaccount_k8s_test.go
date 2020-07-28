package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/k8s"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceCloudAccountK8SBasic(t *testing.T) {
	var cloudAccountResponse k8s.CloudAccountResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountK8S)
	defaultOrganizationalUnitName := os.Getenv(environmentvariable.OrganizationalUnitName)
	organizationUnitTypeAndName, _, organizationUnitGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.OrganizationalUnit)
	organizationUnitHCL := getOrganizationalUnitResourceHCL(organizationUnitGeneratedName, variable.OrganizationalUnitName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountK8SEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAccountK8SDestroy,
		Steps: []resource.TestStep{
			{
				//Create Default
				Config: testAccCheckCloudAccountK8SBasic(resourceTypeAndName, generatedName, variable.CloudAccountK8SOriginalAccountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountK8SExists(resourceTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountK8SOriginalAccountName),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "creation_date"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "vendor", variable.CloudAccountK8SVendor),
					resource.TestCheckResourceAttr(resourceTypeAndName, "organizational_unit_name", defaultOrganizationalUnitName),
				),
			},
			{
				//Update name
				Config: testAccCheckCloudAccountK8SBasic(resourceTypeAndName, generatedName, variable.CloudAccountK8SUpdatedAccountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountK8SExists(resourceTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountK8SUpdatedAccountName),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "creation_date"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "vendor", variable.CloudAccountK8SVendor),
					resource.TestCheckResourceAttr(resourceTypeAndName, "organizational_unit_name", defaultOrganizationalUnitName),
				),
			},
			{
				//Update OU
				Config: testAccCheckCloudAccountK8SBasicWithUpdatedOU(resourceTypeAndName, generatedName, variable.CloudAccountK8SUpdatedAccountName, organizationUnitHCL, organizationUnitTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountK8SExists(resourceTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountK8SUpdatedAccountName),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "creation_date"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "vendor", variable.CloudAccountK8SVendor),
					resource.TestCheckResourceAttrPair(resourceTypeAndName, "organizational_unit_id", organizationUnitTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(resourceTypeAndName, "organizational_unit_name", organizationUnitTypeAndName, "name"),
				),
			},
		},
	})
}

func testAccCheckCloudAccountK8SExists(resource string, cloudAccount *k8s.CloudAccountResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedCloudAccountResponse, _, err := apiClient.cloudaccountK8S.Get(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*cloudAccount = *receivedCloudAccountResponse

		return nil
	}
}

func testAccCheckCloudAccountK8SDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.CloudAccountK8S {
			continue
		}

		receivedCloudAccountResponse, _, err := apiClient.cloudaccountK8S.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if receivedCloudAccountResponse != nil {
			return fmt.Errorf("cloudaccount with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCloudAccountK8SEnvVarsPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.OrganizationalUnitName); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.OrganizationalUnitName)
	}
}

func testAccCheckCloudAccountK8SBasic(resourceTypeAndName, generatedName, resourceName string) string {
	return fmt.Sprintf(`
// k8s cloud account creation
%s

data "%s" "%s" {
 id = "${%s.id}"
}
`,
		// k8s cloud account
		getBasicCloudAccountK8SResourceHCL(generatedName, resourceName),

		// data source variables
		resourcetype.CloudAccountK8S,
		generatedName,
		resourceTypeAndName,
	)
}

func testAccCheckCloudAccountK8SBasicWithUpdatedOU(resourceTypeAndName , generatedName, resourceName, organizationUnitHCL string, organizationUnitTypeAndName string) string {
	return fmt.Sprintf(`
// OU creation
%s
// k8s cloud account creation
%s

data "%s" "%s" {
 id = "${%s.id}"
}
`,
		// ou arguments
		organizationUnitHCL,

		// k8s cloud account arguments
		getCloudAccountK8SResourceHCLWithOU(generatedName, resourceName, organizationUnitTypeAndName),

		// data source variables
		resourcetype.CloudAccountK8S,
		generatedName,
		resourceTypeAndName,
	)
}

func getBasicCloudAccountK8SResourceHCL(generatedName string, resourceName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
 name                   = "%s"
}
`,
		// k8s cloud account variables
		resourcetype.CloudAccountK8S,
		generatedName,
		resourceName,
	)
}

func getCloudAccountK8SResourceHCLWithOU(generatedName string, resourceName string, organizationUnitTypeAndName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
 name                   = "%s"
 organizational_unit_id = "${%s.id}"
}
`,
		// k8s cloud account variables
		resourcetype.CloudAccountK8S,
		generatedName,
		resourceName,
		organizationUnitTypeAndName,
	)
}
