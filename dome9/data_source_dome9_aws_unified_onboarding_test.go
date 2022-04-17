package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
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
				Config: testAccCheckAwsUnifiedOnbordingBasic(resourceTypeAndName, resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceTypeAndName, "ID", dataName+"Data", "ID"),
				),
			},
		},
	})
}

func testAccCheckAWSUnifiedOnboardingDestroy(state *terraform.State) error {
	return nil
}