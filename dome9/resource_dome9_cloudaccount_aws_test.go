package dome9

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"

	"github.com/dome9/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceCloudAccountAWSBasic(t *testing.T) {
	var cloudAccountResponse aws.CloudAccountResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAWS)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAWSPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAccountAWSBasic(resourceTypeAndName, generatedName, variable.CloudAccountAWSOriginalAccountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAWSExists(resourceTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "vendor", variable.CloudAccountAWSVendor),
					resource.TestCheckResourceAttr(resourceTypeAndName, "allow_read_only", "false"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountAWSOriginalAccountName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "net_sec.#", "1"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "net_sec.0.regions.0.region", variable.CloudAccountAWSFetchedRegion),
				),
			},
			{
				Config: testAccCheckCloudAccountAWSBasic(resourceTypeAndName, generatedName, variable.CloudAccountAWSUpdatedAccountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAWSExists(resourceTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountAWSUpdatedAccountName),
				),
			},
		},
	})
}

func testAccCheckCloudAccountAWSExists(resource string, cloudAccount *aws.CloudAccountResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedCloudAccountResponse, _, err := apiClient.cloudaccountAWS.Get(cloudaccounts.QueryParameters{ID: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*cloudAccount = *receivedCloudAccountResponse

		return nil
	}
}

func testAccCloudAccountAWSPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.CloudAccountAWSEnvVarArn); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountAWSEnvVarArn)
	}
	if v := os.Getenv(environmentvariable.CloudAccountAWSEnvVarSecret); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountAWSEnvVarSecret)
	}
}

func testAccCheckCloudAccountDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.CloudAccountAWS {
			continue
		}

		receivedCloudAccountResponse, _, err := apiClient.cloudaccountAWS.Get(cloudaccounts.QueryParameters{ID: rs.Primary.ID})

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if receivedCloudAccountResponse != nil {
			return fmt.Errorf("iplist with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckCloudAccountAWSBasic(resourceTypeAndName, generatedName, resourceName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name        = "%s"
  credentials = {
    arn      = "%s"
    secret   = "%s"
    type     = "RoleBased"
  }
}

data "%s" "%s" {
  id = "${%s.id}"
}

`,
		// resource variable
		resourcetype.CloudAccountAWS,
		generatedName,
		resourceName,
		os.Getenv(environmentvariable.CloudAccountAWSEnvVarArn),
		os.Getenv(environmentvariable.CloudAccountAWSEnvVarSecret),

		// data source variable
		resourcetype.CloudAccountAWS,
		generatedName,
		resourceTypeAndName,
	)
}
