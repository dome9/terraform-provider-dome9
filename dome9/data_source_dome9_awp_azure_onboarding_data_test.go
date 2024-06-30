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

func TestAccDataSourceAwpAzureOnboardingDataBasic(t *testing.T) {
	// Get dome9_awp_azure_onboarding_data resource names
	_, awpAzureOnboardingDataSourceTypeAndName, awpAzureOnboardingDataGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwpAzureOnboardingData)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwpAzureOnboardingDataBasic(awpAzureOnboardingDataGeneratedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(awpAzureOnboardingDataSourceTypeAndName, "cloud_account_id", variable.OnboardedAzureCloudGuardAccountID),
					resource.TestCheckResourceAttrSet(awpAzureOnboardingDataSourceTypeAndName, "region"),
					resource.TestCheckResourceAttrSet(awpAzureOnboardingDataSourceTypeAndName, "app_client_id"),
					resource.TestCheckResourceAttrSet(awpAzureOnboardingDataSourceTypeAndName, "awp_cloud_account_id"),
				),
			},
		},
	})
}

func testAccCheckAwpAzureOnboardingDataBasic(awpAzureOnboardingDataGeneratedName string) string {
	res := fmt.Sprintf(`
data "%s" "%s" {
	cloud_account_id = "%s"
}
	`,
		resourcetype.AwpAzureOnboardingData,
		awpAzureOnboardingDataGeneratedName,
		variable.OnboardedAzureCloudGuardAccountID,
	)
	log.Printf("[INFO] testAccCheckAwpAzureOnboardingDataBasic:%+v\n", res)
	return res
}
