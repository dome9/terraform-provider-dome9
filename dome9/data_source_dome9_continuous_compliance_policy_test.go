package dome9

import (
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccDataSourceContinuousCompliancePolicyBasic(t *testing.T) {
	policyTypeAndName, policyDataSourceTypeAndName, policyGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousCompliancePolicy)
	awsTypeAndName, _, awsGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAWS)
	notificationTypeAndName, _, notificationGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousComplianceNotification)

	awsHCL := getCloudAccountAWSResourceHCL(awsGeneratedName, variable.CloudAccountAWSOriginalAccountName, os.Getenv(environmentvariable.CloudAccountAWSEnvVarArn), "")
	notificationHCL := getContinuousComplianceNotificationResourceHCL(notificationGeneratedName, continuousComplianceNotificationConfig())
	policyHCL := getContinuousCompliancePolicyResourceHCL(awsHCL, awsTypeAndName, notificationHCL, notificationTypeAndName, policyGeneratedName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAWSEnvVarsPreCheck(t) // Aws account used as an input for policy creation thus this check is required
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContinuousCompliancePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckContinuousCompliancePolicyBasic(policyHCL, policyGeneratedName, policyTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(policyDataSourceTypeAndName, "id", policyTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(policyDataSourceTypeAndName, "target_id", policyTypeAndName, "target_internal_id"),
					resource.TestCheckResourceAttr(policyDataSourceTypeAndName, "target_type", "Aws"),
					resource.TestCheckResourceAttr(policyDataSourceTypeAndName, "notification_ids.#", "1"),
				),
			},
		},
	})
}
