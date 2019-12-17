package dome9

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccDataSourceAttachIAMSafeBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AttachIAMSafeToAwsCloudAccount)
	awsTypeAndName, _, awsGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAWS)

	awsCloudAccountHCL := getCloudAccountAWSResourceHCL(awsGeneratedName, variable.CloudAccountAWSOriginalAccountName, os.Getenv(environmentvariable.CloudAccountAWSEnvVarArn), "")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAWSEnvVarsPreCheck(t)
			testAccAttachIAMSafeEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAttachIAMSafeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAttachIAMSafeConfigure(awsCloudAccountHCL, awsTypeAndName, resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "aws_group_arn", resourceTypeAndName, "aws_group_arn"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "aws_policy_arn", resourceTypeAndName, "aws_policy_arn"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "mode", resourceTypeAndName, "mode"),
				),
			},
		},
	})
}
