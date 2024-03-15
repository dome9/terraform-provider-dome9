package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/awp_aws_onboarding"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
)

func TestAccResourceAWPAWSOnboardingBasic(t *testing.T) {
	var awpCloudAccountInfo awp_aws_onboarding.GetAWPOnboardingResponse
	// Generate All Required Random Names for Testing
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwpAwsOnboarding)
	CrossAccountRoleExternalId := os.Getenv(environmentvariable.AwpAwsCrossAccountRoleExternalIdEnvVar)

	// Generate the Awp AWS onboarding HCL Resources
	awpAwsOnboardingHcl := getAwpAwsOnboardingResourceHCL(generatedName, CrossAccountRoleExternalId, false)
	awpAwsOnboardingUpdateHcl := getAwpAwsOnboardingResourceHCL(generatedName, CrossAccountRoleExternalId, true)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAwpAwsEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWPAWSOnboardingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAWPAWSOnboardingBasic(awpAwsOnboardingHcl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwpAccountExists(resourceTypeAndName, &awpCloudAccountInfo),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cloudguard_account_id", variable.OnboardedAwsCloudGuardAccountID),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cross_account_role_name", variable.AwpAwsCrossAccountRoleName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cross_account_role_external_id", CrossAccountRoleExternalId),
					resource.TestCheckResourceAttr(resourceTypeAndName, "scan_mode", variable.ScanMode),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.disabled_regions", variable.DisabledRegions),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.scan_machine_interval_in_hours", variable.ScanMachineIntervalInHours),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.max_concurrence_scans_per_region", variable.MaxConcurrenceScansPerRegion),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.skip_function_apps_scan", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.custom_tags", variable.CustomTags),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "id"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cloud_provider", "aws"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "force_delete", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "should_create_policy", "true"),
				),
			},
			{
				Config: testAccCheckAWPAWSOnboardingBasic(awpAwsOnboardingUpdateHcl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwpAccountExists(resourceTypeAndName, &awpCloudAccountInfo),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cloudguard_account_id", variable.OnboardedAwsCloudGuardAccountID),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cross_account_role_name", variable.AwpAwsCrossAccountRoleName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cross_account_role_external_id", CrossAccountRoleExternalId),
					resource.TestCheckResourceAttr(resourceTypeAndName, "scan_mode", variable.ScanMode),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.disabled_regions", variable.DisabledRegionsUpdate),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.scan_machine_interval_in_hours", variable.ScanMachineIntervalInHoursUpdate),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.max_concurrence_scans_per_region", variable.MaxConcurrenceScansPerRegionUpdate),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.skip_function_apps_scan", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "agentless_account_settings.0.custom_tags", variable.CustomTagsUpdate),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "id"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cloud_provider", "aws"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "force_delete", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "should_create_policy", "true"),
				),
			},
		},
	})
}

func testAccCheckAWPAWSOnboardingDestroy(state *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)
	for _, rs := range state.RootModule().Resources {
		if rs.Type != resourcetype.AwpAwsOnboarding {
			continue
		}
		getOnboardingResponse, _, err := apiClient.awpAwsOnboarding.GetAWPOnboarding("aws", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("error Awp Aws Onboarding still exists, ID: %s", rs.Primary.ID)
		}
		// verify the getOnboardingResponse also is not exists
		if getOnboardingResponse != nil {
			return fmt.Errorf("error Awp Aws Onboarding still exists and wasn't destroyed, ID: %s", rs.Primary.ID)
		}

	}
	return nil
}

func testAccCheckAWPAWSOnboardingBasic(awpAwsOnboardingHcl string) string {
	return fmt.Sprintf(`
// awp aws onboarding resource
%s
`,
		awpAwsOnboardingHcl,
	)
}

func testAccCheckAwpAccountExists(resource string, awpAccount *awp_aws_onboarding.GetAWPOnboardingResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedCloudAccountResponse, _, err := apiClient.awpAwsOnboarding.GetAWPOnboarding("aws", rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*awpAccount = *receivedCloudAccountResponse
		return nil
	}
}

func getAwpAwsOnboardingResourceHCL(generatedResourceName string, externalId string, updateAction bool) string {
	return fmt.Sprintf(`
// awp aws onboarding resource
resource "%s" "%s" {
	cloudguard_account_id = "%s"
	cross_account_role_name = "%s"
	cross_account_role_external_id = "%s"
	scan_mode = "%s"
	agentless_account_settings {
		disabled_regions = "%s"
		scan_machine_interval_in_hours = "%s"
		max_concurrence_scans_per_region = "%s"
		skip_function_apps_scan = "true"
		custom_tags = "%s"
	}
}
`,
		resourcetype.AwpAwsOnboarding,
		generatedResourceName,
		variable.OnboardedAwsCloudGuardAccountID,
		variable.AwpAwsCrossAccountRoleName,
		externalId,
		variable.ScanMode,
		IfThenElse(updateAction, variable.DisabledRegionsUpdate, variable.DisabledRegions),
		IfThenElse(updateAction, variable.ScanMachineIntervalInHoursUpdate, variable.ScanMachineIntervalInHours),
		IfThenElse(updateAction, variable.MaxConcurrenceScansPerRegionUpdate, variable.MaxConcurrenceScansPerRegion),
		IfThenElse(updateAction, variable.CustomTagsUpdate, variable.CustomTags),
	)
}

func testAwpAwsEnvVarsPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.AwpAwsCrossAccountRoleExternalIdEnvVar); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.AwpAwsCrossAccountRoleExternalIdEnvVar)
	}
}
