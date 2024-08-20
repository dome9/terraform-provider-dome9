package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"regexp"
	"testing"
)

func TestAccDataSourceAzureOrganizationOnboardingBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AzureOrganizationOnboarding)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccAzureOrganizationOnboardingEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAzureOrganizationOnboardingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAzureOrganizationOnboardingConfigure(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "account_id", resourceTypeAndName, "account_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "user_id", resourceTypeAndName, "user_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "organization_name", resourceTypeAndName, "organization_name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "tenant_id", resourceTypeAndName, "tenant_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "management_group_id", resourceTypeAndName, "management_group_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "app_registration_name", resourceTypeAndName, "app_registration_name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "onboarding_configuration", resourceTypeAndName, "onboarding_configuration"),
				),
				ExpectError: regexp.MustCompile(`.+Please ensure that the shell script has completed successfully.+`),
			},
		},
	})
}
