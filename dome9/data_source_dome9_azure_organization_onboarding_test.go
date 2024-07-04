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
				Config: testAccCheckAwsOrganizationOnboardingConfigure(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "organization_name", resourceTypeAndName, "organization_name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "app_registration_name", resourceTypeAndName, "app_registration_name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "client_id", resourceTypeAndName, "client_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "client_secret", resourceTypeAndName, "client_secret"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "active_blades", resourceTypeAndName, "active_blades"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "vendor", resourceTypeAndName, "vendor"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "use_cloud_guard_managed_app", resourceTypeAndName, "use_cloud_guard_managed_app"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "use_cloud_guard_managed_app", resourceTypeAndName, "use_cloud_guard_managed_app"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "is_auto_onboarding", resourceTypeAndName, "is_auto_onboarding"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "account_id", resourceTypeAndName, "account_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "external_organization_id", resourceTypeAndName, "external_organization_id"),
				),
				ExpectError: regexp.MustCompile(`.+Failed to assume management account role.+`),
			},
		},
	})
}
