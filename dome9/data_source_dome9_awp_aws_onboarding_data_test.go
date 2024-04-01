package dome9

import (
	"fmt"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceAwpAwsOnboardingDataBasic(t *testing.T) {
	// Get dome9_awp_aws_onboarding_data resource names
	_, awpAwsOnboardingDataSourceTypeAndName, awpAwsOnboardingDataGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwpAwsOnboardingData)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwpAwsOnboardingDataBasic(awpAwsOnboardingDataGeneratedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(awpAwsOnboardingDataSourceTypeAndName, "cloud_account_id", variable.OnboardedAwsCloudGuardAccountID),
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
					resource.TestCheckResourceAttrSet(awpAwsOnboardingDataSourceTypeAndName, "remote_snapshots_utils_function_s3_pre_signed_url"),
				),
			},
		},
	})
}

func testAccCheckAwpAwsOnboardingDataBasic(awpAwsOnboardingDataGeneratedName string) string {
	res := fmt.Sprintf(`
data "%s" "%s" {
	cloud_account_id = "%s"
}
	`,
		resourcetype.AwpAwsOnboardingData,
		awpAwsOnboardingDataGeneratedName,
		variable.OnboardedAwsCloudGuardAccountID,
	)
	log.Printf("[INFO] testAccCheckAwpAwsOnboardingDataBasic:%+v\n", res)
	return res
}
