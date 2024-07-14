package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/integrations"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"strconv"
	"testing"
)

func TestAccResourceIntegrationBasic(t *testing.T) {
	var integrationResponse integrations.IntegrationViewModel
	IntegrationTypeAndName, _, integrationGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.Integration)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIntegrationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIntegrationBasic(IntegrationTypeAndName, integrationGeneratedName, integrationConfig(integrationGeneratedName)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIntegrationExists(IntegrationTypeAndName, &integrationResponse),
					resource.TestCheckResourceAttr(IntegrationTypeAndName, "name", variable.IntegrationName+"_"+integrationGeneratedName),
					resource.TestCheckResourceAttr(IntegrationTypeAndName, "type", variable.IntegrationType),
				),
			},
			{
				// update name test
				Config: testAccCheckIntegrationBasic(IntegrationTypeAndName, integrationGeneratedName, integrationUpdateConfig(integrationGeneratedName)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIntegrationExists(IntegrationTypeAndName, &integrationResponse),
					resource.TestCheckResourceAttr(IntegrationTypeAndName, "name", variable.IntegrationUpdateName+"_"+integrationGeneratedName),
					resource.TestCheckResourceAttr(IntegrationTypeAndName, "type", variable.IntegrationType),
				),
			},
		},
	})
}

func testAccCheckIntegrationExists(resource string, integration *integrations.IntegrationViewModel) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedIntegrationResponse, _, err := apiClient.integration.GetById(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*integration = *receivedIntegrationResponse

		return nil
	}
}

func testAccCheckIntegrationDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.Integration {
			continue
		}

		receivedIntegrationResponse, _, err := apiClient.integration.GetById(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if receivedIntegrationResponse != nil {
			return fmt.Errorf("integration with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIntegrationBasic(resourceTypeAndName, generatedName, additionalBlock string) string {
	return fmt.Sprintf(`
// integration resource
%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// integration resource
		getIntegrationResourceHCL(generatedName, additionalBlock),

		// data source variables
		resourcetype.Integration,
		generatedName,
		resourceTypeAndName,
	)
}

func getIntegrationResourceHCL(generatedName, additionalBlock string) string {
	return fmt.Sprintf(`
// integration creation
resource "%s" "%s" {
  %s

}
`,
		// resource variables
		resourcetype.Integration,
		generatedName,
		additionalBlock,
	)
}

func integrationConfig(integrationNameSuffix string) string {
	return fmt.Sprintf(`
name            = "%s"
type            = "%s"
configuration = jsonencode({
    Url                = "%s"
    MethodType         = "%s"
    AuthType           = "%s"
    Username           = "%s"
    Password           = "%s"
    IgnoreCertificate  = %s
  })
`,
		variable.IntegrationName+"_"+integrationNameSuffix,
		variable.IntegrationType,

		variable.IntegrationUrl,
		variable.IntegrationMethodType,
		variable.IntegrationAuthType,
		variable.IntegrationUsername,
		variable.IntegrationPassword,
		strconv.FormatBool(variable.IntegrationIgnoreCertificate),
	)
}

func integrationUpdateConfig(integrationSuffixName string) string {
	return fmt.Sprintf(`
name            = "%s"
type            = "%s"
configuration = jsonencode({
    Url                = "%s"
    MethodType         = "%s"
    AuthType           = "%s"
    Username           = "%s"
    Password           = "%s"
    IgnoreCertificate  = %s
  })
`,
		variable.IntegrationUpdateName+"_"+integrationSuffixName,
		variable.IntegrationType,

		variable.IntegrationUpdateUrl,
		variable.IntegrationUpdateMethodType,
		variable.IntegrationUpdateAuthType,
		variable.IntegrationUpdateUsername,
		variable.IntegrationUpdatePassword,
		strconv.FormatBool(variable.IntegrationUpdateIgnoreCertificate),
	)
}
