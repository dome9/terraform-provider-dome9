package dome9

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"testing"
)

func TestAccDataSourceAWSOrganizationOnboardingMemberAccountConfig(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AWSOrganizationOnboardingMemberAccountConfig)

	hclCode := fmt.Sprintf(`data "dome9_aws_organization_onboarding_member_account_configuration" "%s" {}`, generatedName)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hclCode,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "external_id", resourceTypeAndName, "external_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "content", resourceTypeAndName, "content"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "onboarding_cft_url", resourceTypeAndName, "onboarding_cft_url"),
				),
			},
		},
	})
}
