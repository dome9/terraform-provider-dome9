package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/unifiedonboarding/aws_unified_onboarding"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"testing"
)

func TestAccResourceAwsUnifiedOnboardingBasic(t *testing.T) {
	var awsUnifiedOnboarding aws_unified_onboarding.UnifiedOnboardingResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwsUnifiedOnboarding)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsUnifiedOnboardingBasic(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsUnifiedOnboardingExists(resourceTypeAndName, &awsUnifiedOnboarding),
					resource.TestCheckResourceAttr(resourceTypeAndName, "iam_capabilities.0", variable.AwsUnifiedOnbordingIamCapabilities0),
					resource.TestCheckResourceAttr(resourceTypeAndName, "iam_capabilities.1", variable.AwsUnifiedOnbordingIamCapabilities1),
					resource.TestCheckResourceAttr(resourceTypeAndName, "iam_capabilities.2", variable.AwsUnifiedOnbordingIamCapabilities2),
					resource.TestCheckResourceAttr(resourceTypeAndName, "onboard_type", variable.AwsUnifiedOnbordingOnboardType),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cloud_vendor", variable.AwsUnifiedOnbordingCloudVendor),
				),
			},
		},
	})
}

func testAccCheckAwsUnifiedOnboardingExists(resource string, awsUnifiedOnboarding *aws_unified_onboarding.UnifiedOnboardingResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			res := fmt.Errorf("didn't find resource: %s", resource)
			return res
		}
		if rs.Primary.ID == "" {
			res := fmt.Errorf("no record ID is set")
			return res

		} else {
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedAwsUnifiedOnboardingResponse, _, err := apiClient.awsUnifiedOnboarding.Get(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		awsUnifiedOnboarding = receivedAwsUnifiedOnboardingResponse

		return nil
	}
}

func testAccCheckAwsUnifiedOnboardingBasic(resourceTypeAndName string, generatedName string) string {
	res := fmt.Sprintf(`
// AwsUnifiedOnbording resource

%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// continuous compliance notification resource
		getAwsUnifiedOnboardingHCL(generatedName),

		// data source variables
		resourcetype.AwsUnifiedOnboarding,
		generatedName+variable.DataSourceSuffix,
		resourceTypeAndName,
	)

	return res
}

func getAwsUnifiedOnboardingHCL(generatedName string) interface{} {
	return fmt.Sprintf(`
// AwsUnifiedOnbording creation
resource "%s" "%s"{ 
onboard_type	 					= "%s"
full_protection					= %s
cloud_vendor						= "%s"
enable_stack_modify				= %s
posture_management_configuration	= %s
serverless_configuration			= %s
intelligence_configurations		= %s

lifecycle {
    prevent_destroy = true
}
	}
`,
		resourcetype.AwsUnifiedOnboarding,
		generatedName,
		variable.AwsUnifiedOnbordingOnboardType,
		variable.AwsUnifiedOnbordingFullProtection,
		variable.AwsUnifiedOnbordingCloudVendor,
		variable.AwsUnifiedOnbordingEnableStackModify,
		variable.AwsUnifiedOnbordingServerlessConfiguration,
		variable.AwsUnifiedOnbordingPostureManagementConfiguration,
		variable.AwsUnifiedOnbordingIntelligenceConfigurations)
}
