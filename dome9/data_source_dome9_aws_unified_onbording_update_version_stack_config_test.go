package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"testing"
)

func TestAccDataSourceAWSUnifiedOnboardingUpdateVersionStackConfogurationBasic(t *testing.T) {
	resourceTypeAndName, _, resourceName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwsUnifiedOnboarding)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		//CheckDestroy: testAccCheckAWSUnifiedOnboardingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsUnifiedOnbordingBasic(resourceTypeAndName, resourceName),
				Check: resource.ComposeTestCheckFunc(
					//resource.TestCheckResourceAttrPair(resourceTypeAndName+"Data", "cloud_vendor", resourceTypeAndName, "cloud_vendor"),
					//resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "cloud_vendor", resourceTypeAndName, "cloud_vendor"),
					//resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "enable_stack_modify", resourceTypeAndName, "enable_stack_modify"),
					//resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "full_protection", resourceTypeAndName, "full_protection"),
				),
			},
		},
	})
}
