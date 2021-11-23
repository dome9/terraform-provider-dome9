package dome9

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/alibaba"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceCloudAccountAlibabaBasic(t *testing.T) {
	var cloudAccountAlibaba alibaba.CloudAccountResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAlibaba)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAlibabaEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAccountAlibabaDestroy,
		Steps: []resource.TestStep{
			{
				// creation test
				Config: testAccCheckCloudAccountAlibabaConfigure(resourceTypeAndName, generatedName, variable.CloudAccountAlibabaCreationResourceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAlibabaExists(resourceTypeAndName, &cloudAccountAlibaba),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountAlibabaCreationResourceName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "vendor", variable.CloudAccountAlibabaVendor),
					resource.TestCheckResourceAttr(resourceTypeAndName, "organizational_unit_name", os.Getenv(environmentvariable.OrganizationalUnitName)),
				),
			},
			{
				// update name test
				Config: testAccCheckCloudAccountAlibabaConfigure(resourceTypeAndName, generatedName, variable.CloudAccountAlibabaUpdatedAccountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAlibabaExists(resourceTypeAndName, &cloudAccountAlibaba),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountAlibabaUpdatedAccountName),
				),
			},
		},
	})
}

func testAccCheckCloudAccountAlibabaDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.CloudAccountAlibaba {
			continue
		}

		resp, _, err := apiClient.cloudaccountAlibaba.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if resp != nil {
			return fmt.Errorf("cloudaccounts with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCloudAccountAlibabaEnvVarsPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.CloudAccountAlibabaEnvVarAccessKey); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountAlibabaEnvVarAccessKey)
	}
	if v := os.Getenv(environmentvariable.CloudAccountAlibabaEnvVarAccessSecret); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountAlibabaEnvVarAccessSecret)
	}
	if v := os.Getenv(environmentvariable.OrganizationalUnitName); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.OrganizationalUnitName)
	}
}

func testAccCheckCloudAccountAlibabaExists(resource string, resp *alibaba.CloudAccountResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedCloudAccount, _, err := apiClient.cloudaccountAlibaba.Get(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}

		*resp = *receivedCloudAccount

		return nil
	}
}

func testAccCheckCloudAccountAlibabaConfigure(resourceTypeAndName, generatedName, resourceName string) string {
	return fmt.Sprintf(`
// Alibaba cloud account creation
%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// Alibaba cloud account
		getCloudAccountAlibabaResourceHCL(generatedName, resourceName),

		// data source variables
		resourcetype.CloudAccountAlibaba,
		generatedName,
		resourceTypeAndName,
	)
}

func getCloudAccountAlibabaResourceHCL(cloudAccountName, generatedAName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
	credentials {
	access_key    = "%s"
	access_secret = "%s"
}
	name          = "%s"
}
`,
		// Alibaba cloud account variables
		resourcetype.CloudAccountAlibaba,
		cloudAccountName,
		os.Getenv(environmentvariable.CloudAccountAlibabaEnvVarAccessKey),
		os.Getenv(environmentvariable.CloudAccountAlibabaEnvVarAccessSecret),
		generatedAName,
	)
}
