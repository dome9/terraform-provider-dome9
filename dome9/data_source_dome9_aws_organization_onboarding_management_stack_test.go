package dome9

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"testing"
)

func TestAccDataSourceAwsOrganizationOnboardingManagementStack(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AWSOrganizationOnboardingManagementStack)

	hclCode := fmt.Sprintf(`data "%s" "%s" {
	aws_account_id = "111111111111"
}`, resourcetype.AWSOrganizationOnboardingManagementStack, generatedName)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hclCode,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "aws_account_id", resourceTypeAndName, "aws_account_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "external_id", resourceTypeAndName, "external_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "content", resourceTypeAndName, "content"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "management_cft_url", resourceTypeAndName, "onboarding_cft_url"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "is_management_onboarded", resourceTypeAndName, "onboarding_cft_status"),
				),
			},
		},
	})
}
