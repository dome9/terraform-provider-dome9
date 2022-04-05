package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/unifiedOnbording/awsUnifiedOnbording"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"testing"
)

func TestAccResourceAwsUnifiedOnbordingBasic(t *testing.T) {
	var awsUnifiedOnbording awsUnifiedOnbording.UnifiedOnbordingConfigurationResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwsUnifiedOnbording)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsUnifiedOnbordingBasic(resourceTypeAndName, generatedName, variable.AwsUnifiedOnbordingOnboardType,
					variable.AwsUnifiedOnbordingFullProtection, variable.AwsUnifiedOnbordingCloudVendor,
					variable.AwsUnifiedOnbordingEnableStackModify, variable.AwsUnifiedOnbordingPostureManagementConfiguration,
					variable.AwsUnifiedOnbordingServerlessConfiguration, variable.AwsUnifiedOnbordingIntelligenceConfigurations),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsUnifiedOnbordingExists(resourceTypeAndName, &awsUnifiedOnbording),
					resource.TestCheckResourceAttr(resourceTypeAndName, "template_url", variable.AwsUnifiedOnbordingTemplateUrl),
					resource.TestCheckResourceAttr(resourceTypeAndName, "iam_capabilities", variable.AwsUnifiedOnbordingIamCapabilities),
				),
			},
		},
	})
}

func testAccCheckAwsUnifiedOnbordingExists(resource string, awsUnifiedOnbording *awsUnifiedOnbording.UnifiedOnbordingConfigurationResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedAwsUnifiedOnbordingResponse, _, err := apiClient.AwsUnifiedOnbording.Get(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*awsUnifiedOnbording = *receivedAwsUnifiedOnbordingResponse

		return nil
	}
}

func testAccCheckAwsUnifiedOnbordingBasic(resourceTypeAndName string, generatedName string, onboardType string, fullProtection string, cloudVendor string, enableStackModify string, postureManagementConfiguration string, serverlessConfiguration string, intelligenceConfigurations string) string {
	return fmt.Sprintf(`
// AwsUnifiedOnbording resource
%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// continuous compliance notification resource
		getContinuousComplianceAwsUnifiedOnbordingHCL(generatedName, resourceTypeAndName),

		// data source variables
		resourcetype.AwsUnifiedOnbording,
		generatedName,
		resourceTypeAndName,
	)
}

func getContinuousComplianceAwsUnifiedOnbordingHCL(generatedName string, resourceTypeAndName string) interface{} {
	return fmt.Sprintf(`{
// AwsUnifiedOnbording creation
resource "%s" "%s"{ 
	"onboardType" 						= "%s",
	"fullProtection"					= "%s",
	  "cloudVendor"						= "%s",
	  "enableStackModify"				= "%s",
	  "postureManagementConfiguration"	= "%s",
	  "serverlessConfiguration"			= "%s",
	  "intelligenceConfigurations"		= "%s"
	}
}`,
		generatedName,
		resourceTypeAndName,
		variable.AwsUnifiedOnbordingOnboardType,
		variable.AwsUnifiedOnbordingFullProtection,
		variable.AwsUnifiedOnbordingCloudVendor,
		variable.AwsUnifiedOnbordingEnableStackModify,
		variable.AwsUnifiedOnbordingServerlessConfiguration,
		variable.AwsUnifiedOnbordingPostureManagementConfiguration,
		variable.AwsUnifiedOnbordingIntelligenceConfigurations)
}
