package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/unifiedonboarding/aws_unified_onboarding"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"log"
	"testing"
)

func TestAccResourceAwsUnifiedOnboardingBasic(t *testing.T) {
	var awsUnifiedOnboarding aws_unified_onboarding.UnifiedOnboardingResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwsUnifiedOnboarding)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {testAccPreCheck(t)},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsUnifiedOnbordingBasic(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsUnifiedOnboardingExists(resourceTypeAndName, &awsUnifiedOnboarding),
					resource.TestCheckResourceAttr(resourceTypeAndName, "template_url", variable.AwsUnifiedOnbordingTemplateUrl),
					resource.TestCheckResourceAttr(resourceTypeAndName, "iam_capabilities.0", variable.AwsUnifiedOnbordingIamCapabilities0),
					resource.TestCheckResourceAttr(resourceTypeAndName, "iam_capabilities.1", variable.AwsUnifiedOnbordingIamCapabilities1),
					resource.TestCheckResourceAttr(resourceTypeAndName, "iam_capabilities.2", variable.AwsUnifiedOnbordingIamCapabilities2),
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
			log.Printf("[INFO] testAccCheckAwsUnifiedOnboardingExists:%+v\n", res)
			return res
		}
		if rs.Primary.ID == "" {
			res :=  fmt.Errorf("no record ID is set")
			log.Printf("[INFO] testAccCheckAwsUnifiedOnboardingExists:%+v\n", res)
			return res

		} else {
			log.Printf("[INFO] testAccCheckAwsUnifiedOnboardingExists:%+v\n", rs)
			log.Printf("[INFO] testAccCheckAwsUnifiedOnboardingExists OK:%+v\n", ok)
		}
		log.Printf("[INFO] testAccCheckAwsUnifiedOnboardingExists:%+v\n", rs)

		apiClient := testAccProvider.Meta().(*Client)
		receivedAwsUnifiedOnboardingResponse, _, err := apiClient.awsUnifiedOnboarding.Get(rs.Primary.ID)
		log.Printf("[INFO] testAccCheckAwsUnifiedOnboardingExists apiClient:%+v\n", receivedAwsUnifiedOnboardingResponse)
		log.Printf("[INFO] testAccCheckAwsUnifiedOnboardingExists err:%+v\n", err)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		awsUnifiedOnboarding = receivedAwsUnifiedOnboardingResponse

		return nil
	}
}

func testAccCheckAwsUnifiedOnbordingBasic(resourceTypeAndName string, generatedName string) string {
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
	log.Printf("[INFO] testAccCheckAwsUnifiedOnbordingBasic:%+v\n", res)

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
