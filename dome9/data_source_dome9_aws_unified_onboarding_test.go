package dome9

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"testing"
)

func TestAccDataSourceAWSUnifiedOnboardingBasic(t *testing.T) {
	resourceTypeAndName, dataName, resourceName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwsUnifiedOnboarding)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSUnifiedOnboardingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsUnifiedOnboardingBasic(resourceTypeAndName, resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceTypeAndName, "ID", dataName+variable.DataSourceSuffix, "ID"),
					resource.TestCheckResourceAttrPair(resourceTypeAndName, "provider", dataName+variable.DataSourceSuffix, "provider"),
					resource.TestCheckResourceAttrPair(resourceTypeAndName, "parameters.Version", dataName+variable.DataSourceSuffix, "cft_version"),
					resource.TestCheckResourceAttrPair(resourceTypeAndName, "parameters.OnboardingId", dataName+variable.DataSourceSuffix, "onboarding_id"),
				),
				ExpectError: regexp.MustCompile(`.+00000000-0000-0000-0000-000000000000\/DeleteForce, 404.+`),
			},
		},
	})
}

func testAccCheckAWSUnifiedOnboardingDestroy(state *terraform.State) error {
	return nil
}
