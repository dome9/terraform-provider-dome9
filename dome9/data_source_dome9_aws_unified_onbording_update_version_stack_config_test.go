package dome9

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"log"
	"testing"
)

func TestAccDataSourceAWSUnifiedOnboardingUpdateVersionStackConfogurationBasic(t *testing.T) {
	resourceTypeAndName, _, resourceName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwsUnifiedOnboarding)
	dataTypeAndName := fmt.Sprintf("data.%s.%s", resourcetype.AwsUnifiedOnboardingUpdateVersionStackConfig, resourceName)
	log.Println("TestAccDataSourceAWSUnifiedOnboardingUpdateVersionStackConfogurationBasic ", resourceTypeAndName, dataTypeAndName, resourceName)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSUnifiedOnboardingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsUnifiedOnbordingUpdateVersionStackConfogurationBasic(resourceTypeAndName, resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataTypeAndName+variable.DataSourceSuffix, "ID", resourceTypeAndName, "ID"),
					resource.TestCheckResourceAttrPair(dataTypeAndName+variable.DataSourceSuffix, "onboarding_id", resourceTypeAndName, "parameters.OnboardingId"),
					resource.TestCheckResourceAttrPair(dataTypeAndName+variable.DataSourceSuffix, "stack_name", resourceTypeAndName, "stack_name"),
					resource.TestCheckResourceAttrPair(dataTypeAndName+variable.DataSourceSuffix, "template_url", resourceTypeAndName, "template_url"),
					resource.TestCheckResourceAttrPair(dataTypeAndName+variable.DataSourceSuffix, "provider", resourceTypeAndName, "provider"),
					resource.TestCheckResourceAttrPair(dataTypeAndName+variable.DataSourceSuffix, "iam_capabilities.0", resourceTypeAndName, "iam_capabilities.0"),
					resource.TestCheckResourceAttrPair(dataTypeAndName+variable.DataSourceSuffix, "iam_capabilities.1", resourceTypeAndName, "iam_capabilities.1"),
					resource.TestCheckResourceAttrPair(dataTypeAndName+variable.DataSourceSuffix, "iam_capabilities.2", resourceTypeAndName, "iam_capabilities.2"),
				),
			},
		},
	})
}

func testAccCheckAwsUnifiedOnbordingUpdateVersionStackConfogurationBasic(resourceTypeAndName string, generatedName string) string {
	res := fmt.Sprintf(`
// AwsUnifiedOnbording resource

%s

data "%s" "%s" {
  onboarding_id = "${%s.id}"
}
`,
		// continuous compliance notification resource
		getAwsUnifiedOnboardingHCL(generatedName),

		// data source variables
		resourcetype.AwsUnifiedOnboardingUpdateVersionStackConfig,
		generatedName+variable.DataSourceSuffix,
		resourceTypeAndName,
	)
	log.Printf("[INFO] testAccCheckAwsUnifiedOnboardingBasic:%+v\n", res)

	return res
}
