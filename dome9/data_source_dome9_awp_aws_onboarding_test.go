package dome9

import (
	"fmt"
	"os"
	"testing"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceAwpAwsOnboardingBasic(t *testing.T) {
	// Get dome9_awp_aws_onboarding resource names and values
	awpAwsOnboardingResourceTypeAndName, awpAwsOnboardingDataSourceTypeAndName, resourceGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwpAwsOnboarding)
	CrossAccountRoleExternalId := os.Getenv(environmentvariable.AwpAwsCrossAccountRoleExternalIdEnvVar)

	// Generate the Awp AWS onboarding HCL Resources
	awpAwsOnboardingHcl := getAwpAwsOnboardingResourceHCL(resourceGeneratedName, CrossAccountRoleExternalId, false)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAwpAwsEnvVarsPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwpAwsOnboardingConfig(awpAwsOnboardingHcl, resourceGeneratedName, awpAwsOnboardingResourceTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(awpAwsOnboardingDataSourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(awpAwsOnboardingDataSourceTypeAndName, "scan_mode", awpAwsOnboardingResourceTypeAndName, "scan_mode"),
					resource.TestCheckResourceAttrPair(awpAwsOnboardingDataSourceTypeAndName, "cloud_account_id", awpAwsOnboardingResourceTypeAndName, "cloudguard_account_id"),
					resource.TestCheckResourceAttrPair(awpAwsOnboardingDataSourceTypeAndName, "agentless_account_settings.0.disabled_regions.0", awpAwsOnboardingResourceTypeAndName, "agentless_account_settings.0.disabled_regions.0"),
					resource.TestCheckResourceAttrPair(awpAwsOnboardingDataSourceTypeAndName, "agentless_account_settings.0.disabled_regions.1", awpAwsOnboardingResourceTypeAndName, "agentless_account_settings.0.disabled_regions.1"),
					resource.TestCheckResourceAttrPair(awpAwsOnboardingDataSourceTypeAndName, "agentless_account_settings.0.scan_machine_interval_in_hours", awpAwsOnboardingResourceTypeAndName, "agentless_account_settings.0.scan_machine_interval_in_hours"),
					resource.TestCheckResourceAttrPair(awpAwsOnboardingDataSourceTypeAndName, "agentless_account_settings.0.max_concurrent_scans_per_region", awpAwsOnboardingResourceTypeAndName, "agentless_account_settings.0.max_concurrent_scans_per_region"),
					resource.TestCheckResourceAttrPair(awpAwsOnboardingDataSourceTypeAndName, "agentless_account_settings.0.custom_tags.%", awpAwsOnboardingResourceTypeAndName, "agentless_account_settings.0.custom_tags.%"),
					resource.TestCheckResourceAttrPair(awpAwsOnboardingDataSourceTypeAndName, "missing_awp_private_network_regions", awpAwsOnboardingResourceTypeAndName, "missing_awp_private_network_regions"),
					resource.TestCheckResourceAttrPair(awpAwsOnboardingDataSourceTypeAndName, "agentless_protection_enabled", awpAwsOnboardingResourceTypeAndName, "agentless_protection_enabled"),
					resource.TestCheckResourceAttrPair(awpAwsOnboardingDataSourceTypeAndName, "centralized_cloud_account_id", awpAwsOnboardingResourceTypeAndName, "awp_hub_external_account_id"),
				),
			},
		},
	})
}

func testAccDataSourceAwpAwsOnboardingConfig(awpAwsOnboardingHcl, dataSourceGeneratedName, awpAwsOnboardingResourceTypeAndName string) string {
	return fmt.Sprintf(`
// awp aws onboarding resource
%s

// awp aws onboarding data source
data "%s" "%s" {
	id = %s.cloudguard_account_id
}
`,
		awpAwsOnboardingHcl,
		resourcetype.AwpAwsOnboarding,
		dataSourceGeneratedName,
		awpAwsOnboardingResourceTypeAndName,
	)
}
