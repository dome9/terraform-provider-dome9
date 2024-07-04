package dome9

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"testing"
)

func TestAccDataSourceAwsOrganizationOnboardingManagementStack(t *testing.T) {
	_, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AWSOrganizationOnboardingManagementStack)

	hclCode := fmt.Sprintf(`data "%s" "%s" {
	aws_account_id = "%s"
}`, resourcetype.AWSOrganizationOnboardingManagementStack, generatedName, variable.AwsFakeAccountId)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hclCode,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceTypeAndName, "external_id"),
					resource.TestCheckResourceAttrSet(dataSourceTypeAndName, "content"),
					resource.TestCheckResourceAttrSet(dataSourceTypeAndName, "management_cft_url"),
					resource.TestCheckResourceAttr(dataSourceTypeAndName, "is_management_onboarded", "false"),
				),
			},
		},
	})
}
