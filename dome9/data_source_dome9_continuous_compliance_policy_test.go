package dome9

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccDataSourceContinuousCompliancePolicyBasic(t *testing.T) {
	policyTypeAndName, policyDataSourceTypeAndName, policyGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousCompliancePolicy)
	azureTypeAndName, _, azureGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAzure)
	notificationTypeAndName, _, notificationGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousComplianceNotification)

	azureHCL := getCloudAccountAzureResourceHCL(azureGeneratedName, variable.CloudAccountAzureCreationResourceName)
	notificationHCL := getContinuousComplianceNotificationResourceHCL(notificationGeneratedName, continuousComplianceNotificationConfig())
	policyHCL := getContinuousCompliancePolicyResourceHCL(azureHCL, azureTypeAndName, notificationHCL, notificationTypeAndName, policyGeneratedName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAzureEnvVarsPreCheck(t) // Azure account used as an input for policy creation thus this check is required
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContinuousCompliancePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckContinuousCompliancePolicyBasic(policyHCL, policyGeneratedName, policyTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(policyDataSourceTypeAndName, "id", policyTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(policyDataSourceTypeAndName, "cloud_account_id", policyTypeAndName, "cloud_account_id"),
					resource.TestCheckResourceAttr(policyDataSourceTypeAndName, "cloud_account_type", "Azure"),
					resource.TestCheckResourceAttr(policyDataSourceTypeAndName, "notification_ids.#", "1"),
				),
			},
		},
	})
}
