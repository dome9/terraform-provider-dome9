package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/awp_aws_onboarding"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
)

func TestAccResourceAWPAWSOnboardingBasic(t *testing.T) {
	var awpCloudAccountInfo awp_aws_onboarding.GetAWPOnboardingResponse
	// Generate All Required Random Names for Testing
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwpAwsOnboarding)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAWPAWSOnboardingBasic(generatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwpAccountExists(resourceTypeAndName, &awpCloudAccountInfo),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cloudguard_account_id", "8f9cfb94-4365-4a29-a7b9-cabbb7fe9430"),
					// Add more TestCheckResourceAttr functions for each attribute to check
				),
			},
			{
				Config: testAccCheckAWPAWSOnboardingUpdate(generatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "force_delete", "false"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "should_create_policy", "false"),
				),
			},
		},
	})
}

func testAccCheckAWPAWSOnboardingBasic(generatedName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
	cloudguard_account_id = "%s"
	cross_account_role_name = "%s"
	cross_account_role_external_id = "%s"
	scan_mode = "%s"
	force_delete = true
	should_create_policy = true
	agentless_account_settings {
		disabled_regions = ["us-east-1", "us-west-1"]  # Example disabled regions
		scan_machine_interval_in_hours = 6
		max_concurrence_scans_per_region = 2
		skip_function_apps_scan = true
		custom_tags = {
			tag1 = "value1"
			tag2 = "value2"
		}
	}
}
`,
		resourcetype.AwpAwsOnboarding,
		generatedName,
		"8f9cfb94-4365-4a29-a7b9-cabbb7fe9430",
		"CloudGuardAWPCrossAccountRole",
		"NDYwNjc4MTkzOTI2LThmOWNmYjk0LTQzNjUtNGEyOS1hN2I5LWNhYmJiN2ZlOTQzMA==",
		"inAccount",
	)
}

func testAccCheckAWPAWSOnboardingUpdate(generatedName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
	cloudguard_account_id = "%s"
	cross_account_role_name = "%s"
	cross_account_role_external_id = "%s"
	scan_mode = "%s"
	force_delete = false
	should_create_policy = false
	agentless_account_settings {
		disabled_regions = ["us-east-1"]
		scan_machine_interval_in_hours = 8
		max_concurrence_scans_per_region = 4
		skip_function_apps_scan = true
		custom_tags = {
			tag1 = "value1"
		}
	}
}
`,
		resourcetype.AwpAwsOnboarding,
		generatedName,
		"8f9cfb94-4365-4a29-a7b9-cabbb7fe9430",
		"CloudGuardAWPCrossAccountRole",
		"NDYwNjc4MTkzOTI2LThmOWNmYjk0LTQzNjUtNGEyOS1hN2I5LWNhYmJiN2ZlOTQzMA==",
		"inAccount",
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
