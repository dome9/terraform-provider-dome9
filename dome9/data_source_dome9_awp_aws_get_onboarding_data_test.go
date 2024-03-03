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
	_, awsOnboardingDataSourceTypeAndName, randomDataSourceName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwpAwsOnboardingData)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwpAwsOnboardingDataBasic(randomDataSourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(awsOnboardingDataSourceTypeAndName, "external_aws_account_id", variable.AwpAwsAccountID),
					resource.TestCheckResourceAttrSet(awsOnboardingDataSourceTypeAndName, "stage"),
					resource.TestCheckResourceAttrSet(awsOnboardingDataSourceTypeAndName, "region"),
					resource.TestCheckResourceAttrSet(awsOnboardingDataSourceTypeAndName, "cloud_guard_backend_account_id"),
					resource.TestCheckResourceAttrSet(awsOnboardingDataSourceTypeAndName, "agentless_bucket_name"),
					resource.TestCheckResourceAttrSet(awsOnboardingDataSourceTypeAndName, "remote_functions_prefix_key"),
					resource.TestCheckResourceAttrSet(awsOnboardingDataSourceTypeAndName, "remote_snapshots_utils_function_name"),
					resource.TestCheckResourceAttrSet(awsOnboardingDataSourceTypeAndName, "remote_snapshots_utils_function_run_time"),
					resource.TestCheckResourceAttrSet(awsOnboardingDataSourceTypeAndName, "remote_snapshots_utils_function_time_out"),
					resource.TestCheckResourceAttrSet(awsOnboardingDataSourceTypeAndName, "awp_client_side_security_group_name"),
					resource.TestCheckResourceAttrSet(awsOnboardingDataSourceTypeAndName, "cross_account_role_external_id"),
				),
			},
		},
	})
}

func testAccCheckAwpAwsOnboardingDataBasic(resourceName string) string {
	res := fmt.Sprintf(`
		data "%s" "%s" {
			external_aws_account_id = "%s"
		}
	`,
		// Add the HCL configuration for the resource here
		resourcetype.AwpAwsOnboardingData,
		resourceName,
		variable.AwpAwsAccountID,
	)
	log.Printf("[INFO] testAccCheckAwpAwsOnboardingDataBasic:%+v\n", res)
	return res
}
