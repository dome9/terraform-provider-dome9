package dome9

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
)

func TestAccDataSourceContinuousCompliancePolicyBasic(t *testing.T) {
	policyResourceTypeAndName, policyDataSourceTypeAndName, policyGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousCompliancePolicy)
	cloudAccountResourceTypeAndName, _, cloudAccountGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAzure)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAzureEnvVarsPreCheck(t) // Azure account used as an input for policy creation thus this check is required
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContinuousCompliancePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckContinuousCompliancePolicyBasic(policyGeneratedName, cloudAccountGeneratedName, cloudAccountResourceTypeAndName, policyResourceTypeAndName, getNotificationIDsConfig()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(policyDataSourceTypeAndName, "id", policyResourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(policyDataSourceTypeAndName, "cloud_account_id", policyResourceTypeAndName, "cloud_account_id"),
					resource.TestCheckResourceAttr(policyDataSourceTypeAndName, "cloud_account_type", "Azure"),
					resource.TestCheckResourceAttr(policyDataSourceTypeAndName, "notification_ids.#", "1"),
				),
			},
		},
	})
}
