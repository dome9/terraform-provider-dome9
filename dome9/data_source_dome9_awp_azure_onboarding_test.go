package dome9

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceAwpAzureOnboardingBasic(t *testing.T) {
	// Get dome9_awp_azure_onboarding resource names and values
	awpAzureOnboardingResourceTypeAndName, awpAzureOnboardingDataSourceTypeAndName, resourceGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwpAzureOnboarding)

	// Generate the Awp Azure onboarding HCL Resources
	awpAzureOnboardingHcl := getAwpAzureOnboardingResourceHCL(resourceGeneratedName, false)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwpAzureOnboardingConfig(awpAzureOnboardingHcl, resourceGeneratedName, awpAzureOnboardingResourceTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(awpAzureOnboardingDataSourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(awpAzureOnboardingDataSourceTypeAndName, "scan_mode", awpAzureOnboardingResourceTypeAndName, "scan_mode"),
					resource.TestCheckResourceAttrPair(awpAzureOnboardingDataSourceTypeAndName, "cloud_account_id", awpAzureOnboardingResourceTypeAndName, "cloudguard_account_id"),
					resource.TestCheckResourceAttrPair(awpAzureOnboardingDataSourceTypeAndName, "missing_awp_private_network_regions", awpAzureOnboardingResourceTypeAndName, "missing_awp_private_network_regions"),
					resource.TestCheckResourceAttrPair(awpAzureOnboardingDataSourceTypeAndName, "agentless_protection_enabled", awpAzureOnboardingResourceTypeAndName, "agentless_protection_enabled"),
				),
			},
		},
	})
}

func testAccDataSourceAwpAzureOnboardingConfig(awpAzureOnboardingHcl, dataSourceGeneratedName, awpAzureOnboardingResourceTypeAndName string) string {
	return fmt.Sprintf(`
// awp azure onboarding resource
%s

// awp azure onboarding data source
data "%s" "%s" {
	id = %s.cloudguard_account_id
}
`,
		awpAzureOnboardingHcl,
		resourcetype.AwpAzureOnboarding,
		dataSourceGeneratedName,
		awpAzureOnboardingResourceTypeAndName,
	)
}
