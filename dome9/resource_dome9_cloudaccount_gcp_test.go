package dome9

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/gcp"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceCloudAccountGCPBasic(t *testing.T) {
	var cloudAccountGCP gcp.CloudAccountResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountGCP)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountGCPEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAccountGCPDestroy,
		Steps: []resource.TestStep{
			{
				// creation test
				Config: testAccCheckCloudAccountGCPConfigure(resourceTypeAndName, generatedName, variable.CloudAccountGCPCreationResourceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountGCPExists(resourceTypeAndName, &cloudAccountGCP),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountGCPCreationResourceName),
				),
			},
			{
				// update name test
				Config: testAccCheckCloudAccountGCPConfigure(resourceTypeAndName, generatedName, variable.CloudAccountGCPUpdatedAccountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountGCPExists(resourceTypeAndName, &cloudAccountGCP),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountGCPUpdatedAccountName),
				),
			},
		},
	})
}

func testAccCheckCloudAccountGCPDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.CloudAccountGCP {
			continue
		}

		getCloudAccountQueryParams := cloudaccounts.QueryParameters{ID: rs.Primary.ID}
		resp, _, err := apiClient.cloudaccountGCP.Get(&getCloudAccountQueryParams)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if resp != nil {
			return fmt.Errorf("cloudaccounts with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCloudAccountGCPEnvVarsPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.CloudAccountGCPEnvVarClientEmail); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountGCPEnvVarClientEmail)
	}
	if v := os.Getenv(environmentvariable.CloudAccountGCPEnvVarClientId); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountGCPEnvVarClientId)
	}
	if v := os.Getenv(environmentvariable.CloudAccountGCPEnvVarClientX509CertUrl); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountGCPEnvVarClientX509CertUrl)
	}
	if v := os.Getenv(environmentvariable.CloudAccountGCPEnvVarPrivateKey); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountGCPEnvVarPrivateKey)
	}
	if v := os.Getenv(environmentvariable.CloudAccountGCPEnvVarPrivateKeyId); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountGCPEnvVarPrivateKeyId)
	}
	if v := os.Getenv(environmentvariable.CloudAccountGCPEnvVarProjectId); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.CloudAccountGCPEnvVarProjectId)
	}
}

func testAccCheckCloudAccountGCPExists(resource string, resp *gcp.CloudAccountResponse) resource.TestCheckFunc {
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
		receivedCloudAccount, _, err := apiClient.cloudaccountGCP.Get(&getCloudAccountQueryParams)
		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}

		*resp = *receivedCloudAccount

		return nil
	}
}

func testAccCheckCloudAccountGCPConfigure(resourceTypeAndName, generatedName, resourceName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name = "%s"
  service_account_credentials ={
    type = "%s"
    project_id = "%s"
    private_key_id = "%s"
    private_key = "%s"
    client_email = "%s"
    client_id = "%s"
    auth_uri = "%s"
    token_uri = "%s"
    auth_provider_x509_cert_url = "%s"
    client_x509_cert_url = "%s"
  }
}

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// resource variables
		resourcetype.CloudAccountGCP,
		generatedName,
		resourceName,
		variable.CloudAccountGCPType,
		os.Getenv(environmentvariable.CloudAccountGCPEnvVarProjectId),
		os.Getenv(environmentvariable.CloudAccountGCPEnvVarPrivateKeyId),
		os.Getenv(environmentvariable.CloudAccountGCPEnvVarPrivateKey),
		os.Getenv(environmentvariable.CloudAccountGCPEnvVarClientEmail),
		os.Getenv(environmentvariable.CloudAccountGCPEnvVarClientId),
		variable.CloudAccountGCPAuthURL,
		variable.CloudAccountGCPTokenURL,
		variable.CloudAccountGCPAuthProviderX509CertURL,
		os.Getenv(environmentvariable.CloudAccountGCPEnvVarClientX509CertUrl),

		// data source variables
		resourcetype.CloudAccountGCP,
		generatedName,
		resourceTypeAndName,
	)
}
