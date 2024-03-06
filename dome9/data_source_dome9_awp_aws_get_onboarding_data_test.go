package dome9

import (
	"fmt"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceAwpAwsOnboardingDataBasic(t *testing.T) {

	// Get dome9_aws_unified_onboarding resource to do aws onboarding
	awsUnifiedOnboardingResourceTypeAndName, awsUnifiedOnboardingDataResourceName, awsUnifiedOnboardingResourceName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwsUnifiedOnboarding)
	awsUnifiedOnboardingResourceHCL := getAwsUnifiedOnboardingResourceHCL(awsUnifiedOnboardingResourceName, awsUnifiedOnboardingResourceTypeAndName)

	// Get aws_cloudformation_stack resource
	_, _, cloudFormationStackGeneratedName := method.GenerateRandomSourcesTypeAndName(providerconst.AwsCloudFormationStack)
	awsCloudFormationStackHcl := getAwsCloudFormationStackResourceHCL(cloudFormationStackGeneratedName, awsUnifiedOnboardingResourceTypeAndName)

	// Get dome9_awp_aws_get_onboarding_data resource names
	_, awpAwsOnboardingDataSourceTypeAndName, awpAwsOnboardingDataGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwpAwsOnboardingData)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwpAwsOnboardingDataBasic(awsUnifiedOnboardingResourceHCL, awsCloudFormationStackHcl, awpAwsOnboardingDataGeneratedName, awsUnifiedOnboardingDataResourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(awpAwsOnboardingDataSourceTypeAndName, "external_aws_account_id", variable.AwpAwsAccountID),
					resource.TestCheckResourceAttrSet(awpAwsOnboardingDataSourceTypeAndName, "stage"),
					resource.TestCheckResourceAttrSet(awpAwsOnboardingDataSourceTypeAndName, "region"),
					resource.TestCheckResourceAttrSet(awpAwsOnboardingDataSourceTypeAndName, "cloud_guard_backend_account_id"),
					resource.TestCheckResourceAttrSet(awpAwsOnboardingDataSourceTypeAndName, "agentless_bucket_name"),
					resource.TestCheckResourceAttrSet(awpAwsOnboardingDataSourceTypeAndName, "remote_functions_prefix_key"),
					resource.TestCheckResourceAttrSet(awpAwsOnboardingDataSourceTypeAndName, "remote_snapshots_utils_function_name"),
					resource.TestCheckResourceAttrSet(awpAwsOnboardingDataSourceTypeAndName, "remote_snapshots_utils_function_run_time"),
					resource.TestCheckResourceAttrSet(awpAwsOnboardingDataSourceTypeAndName, "remote_snapshots_utils_function_time_out"),
					resource.TestCheckResourceAttrSet(awpAwsOnboardingDataSourceTypeAndName, "awp_client_side_security_group_name"),
					resource.TestCheckResourceAttrSet(awpAwsOnboardingDataSourceTypeAndName, "cross_account_role_external_id"),
				),
			},
		},
	})
}

func testAccCheckAwpAwsOnboardingDataBasic(awsUnifiedOnboardingHcl string, awsCloudFormationStackHcl string,
	awpAwsGetOnboardingDataGeneratedName string, awsUnifiedOnboardingDataResourceAndType string) string {
	res := fmt.Sprintf(`
// AwsUnifiedOnbording resource
%s
// AwsCloudFormationStack resource
%s
data "%s" "%s" {
	external_aws_account_id = "%s.environment_external_id"
}
	`,
		awsUnifiedOnboardingHcl,
		awsCloudFormationStackHcl,
		resourcetype.AwpAwsOnboardingData,
		awpAwsGetOnboardingDataGeneratedName,
		awsUnifiedOnboardingDataResourceAndType,
	)
	log.Printf("[INFO] testAccCheckAwpAwsOnboardingDataBasic:%+v\n", res)
	return res
}

func getAwsUnifiedOnboardingResourceHCL(awsUnifiedOnboardingResourceName string, awsUnifiedOnboardingResourceTypeAndName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
	cloud_vendor = "aws"
	onboard_type = "Simple"
	full_protection = true
	enable_stack_modify = true
	posture_management_configuration = {
		rulesets = "[0]"
	}
	serverless_configuration  = {
		enabled = false
	}
	intelligence_configurations = {
		rulesets = "[0]"
		enabled = false
	}
}
data "%s" "%s" {
	id = "%s.id"
}
	`, resourcetype.AwsUnifiedOnboarding,
		awsUnifiedOnboardingResourceName,
		resourcetype.AwsUnifiedOnboarding,
		awsUnifiedOnboardingResourceName,
		awsUnifiedOnboardingResourceTypeAndName)
}

func getAwsCloudFormationStackResourceHCL(stackGeneratedName string, awsUnifiedOnboardingResourceTypeAndName string) string {
	return fmt.Sprintf(`
resource "%s" "%s"{
	name = %s.stack_name
	template_url = %s.template_url
	parameters = %s.parameters
	capabilities = %s.iam_capabilities
}
	`, providerconst.AwsCloudFormationStack,
		stackGeneratedName,
		awsUnifiedOnboardingResourceTypeAndName,
		awsUnifiedOnboardingResourceTypeAndName,
		awsUnifiedOnboardingResourceTypeAndName,
		awsUnifiedOnboardingResourceTypeAndName)
}
