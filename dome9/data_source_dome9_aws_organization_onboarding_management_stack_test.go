package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"testing"
)

func TestAccDataSourceAwsOrganizationOnboardingManagementStack(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AWSOrganizationOnboardingManagementStack)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccAWSOrganizationOnboardingEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsOrganizationOnboardingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsOrganizationOnboardingConfigure(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "account_id", resourceTypeAndName, "account_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "external_organization_id", resourceTypeAndName, "external_organization_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "external_management_account_id", resourceTypeAndName, "external_management_account_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "management_account_stack_id", resourceTypeAndName, "management_account_stack_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "management_account_stack_region", resourceTypeAndName, "management_account_stack_region"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "onboarding_configuration", resourceTypeAndName, "onboarding_configuration"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "user_id", resourceTypeAndName, "user_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "organization_name", resourceTypeAndName, "organization_name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "update_time", resourceTypeAndName, "update_time"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "creation_time", resourceTypeAndName, "creation_time"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "stack_set_regions", resourceTypeAndName, "stack_set_regions"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "stack_set_organizational_unit_ids", resourceTypeAndName, "stack_set_organizational_unit_ids"),
				),
			},
		},
	})
}
