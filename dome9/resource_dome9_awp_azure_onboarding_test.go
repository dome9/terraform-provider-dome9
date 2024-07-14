package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/awp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"testing"
	"time"
)

func TestAccResourceAWPAzureOnboardingBasic(t *testing.T) {
	var awpCloudAccountInfo awp_onboarding.GetAWPOnboardingResponse
	// Generate All Required Random Names for Testing
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwpAzureOnboarding)
	disabledRegion1, _ := getRegionByIndex(variable.AzureDisabledRegions, 0)
	disabledRegion2, _ := getRegionByIndex(variable.AzureDisabledRegions, 1)
	disabledRegionUpdate3, _ := getRegionByIndex(variable.AzureDisabledRegionsUpdate, 2)
	disabledRegionUpdate4, _ := getRegionByIndex(variable.AzureDisabledRegionsUpdate, 3)

	// Generate the Awp Azure onboarding HCL Resources
	awpAzureOnboardingHcl := getAwpAzureOnboardingResourceHCL(generatedName, false)
	awpAzureOnboardingUpdateHcl := getAwpAzureOnboardingResourceHCL(generatedName, true)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWPAzureOnboardingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAWPAzureOnboardingBasic(awpAzureOnboardingHcl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwpAzureAccountExists(resourceTypeAndName, &awpCloudAccountInfo),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cloudguard_account_id", variable.OnboardedAzureCloudGuardAccountID),
					resource.TestCheckResourceAttr(resourceTypeAndName, "scan_mode", variable.ScanMode),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.disabled_regions.0", disabledRegion1),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.disabled_regions.1", disabledRegion2),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.scan_machine_interval_in_hours", variable.ScanMachineIntervalInHours),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.max_concurrent_scans_per_region", variable.MaxConcurrentScansPerRegion),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.custom_tags.%", "2"),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "id"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cloud_provider", "azure"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "should_create_policy", "true"),
				),
			},
			{
				Config: testAccCheckAWPAzureOnboardingBasic(awpAzureOnboardingUpdateHcl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwpAzureAccountExists(resourceTypeAndName, &awpCloudAccountInfo),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cloudguard_account_id", variable.OnboardedAzureCloudGuardAccountID),
					resource.TestCheckResourceAttr(resourceTypeAndName, "scan_mode", variable.ScanMode),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.disabled_regions.0", disabledRegion1),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.disabled_regions.1", disabledRegion2),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.disabled_regions.2", disabledRegionUpdate3),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.disabled_regions.3", disabledRegionUpdate4),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.scan_machine_interval_in_hours", variable.ScanMachineIntervalInHoursUpdate),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.max_concurrent_scans_per_region", variable.MaxConcurrentScansPerRegionUpdate),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.custom_tags.%", "3"),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "id"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cloud_provider", "azure"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "should_create_policy", "true"),
				),
			},
		},
	})
}

func testAccCheckAWPAzureOnboardingDestroy(state *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)
	for _, rs := range state.RootModule().Resources {
		if rs.Type != resourcetype.AwpAzureOnboarding {
			continue
		}
		maxRetries := 3
		retryInterval := time.Second * 5
		var getOnboardingResponse *awp_onboarding.GetAWPOnboardingResponse
		var err error
		for i := 0; i < maxRetries; i++ {
			getOnboardingResponse, _, err = apiClient.awpAzureOnboarding.GetAWPOnboarding(rs.Primary.ID)
			if err == nil || getOnboardingResponse != nil {
				// If the request was successful or the resource still exists, wait for the retry interval before trying again
				time.Sleep(retryInterval)
			} else {
				// If the request failed with a 404 status code, break the loop
				break
			}
		}
		if err == nil {
			return fmt.Errorf("error Awp Azure Onboarding still exists, ID: %s", rs.Primary.ID)
		}
		// verify the getOnboardingResponse also is not exists
		if getOnboardingResponse != nil {
			return fmt.Errorf("error Awp Azure Onboarding still exists and wasn't destroyed, ID: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckAWPAzureOnboardingBasic(awpAzureOnboardingHcl string) string {
	return fmt.Sprintf(`
// awp azure onboarding resource
%s
`,
		awpAzureOnboardingHcl,
	)
}

func getAwpAzureOnboardingResourceHCL(generatedResourceName string, updateAction bool) string {
	return fmt.Sprintf(`
// awp azure onboarding resource
resource "%s" "%s" {
	cloudguard_account_id = "%s"
	scan_mode = "%s"
	agentless_account_settings {
		disabled_regions = %s
		scan_machine_interval_in_hours = "%s"
		max_concurrent_scans_per_region = "%s"
		custom_tags = %s
	}
}
`,
		resourcetype.AwpAzureOnboarding,
		generatedResourceName,
		variable.OnboardedAzureCloudGuardAccountID,
		variable.ScanMode,
		IfThenElse(updateAction, variable.AzureDisabledRegionsUpdate, variable.AzureDisabledRegions),
		IfThenElse(updateAction, variable.ScanMachineIntervalInHoursUpdate, variable.ScanMachineIntervalInHours),
		IfThenElse(updateAction, variable.MaxConcurrentScansPerRegionUpdate, variable.MaxConcurrentScansPerRegion),
		IfThenElse(updateAction, variable.CustomTagsUpdate, variable.CustomTags),
	)
}

func testAccCheckAwpAzureAccountExists(resource string, awpAccount *awp_onboarding.GetAWPOnboardingResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedCloudAccountResponse, _, err := apiClient.awpAzureOnboarding.GetAWPOnboarding(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*awpAccount = *receivedCloudAccountResponse
		return nil
	}
}
