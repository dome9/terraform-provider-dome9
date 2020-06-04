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

func TestAccDataSourceCloudSecurityGroupAWSBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, resourceName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAWSSecurityGroup)
	awsTypeAndName, _, awsGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAWS)

	awsCloudAccountHCL := getCloudAccountAWSResourceHCL(awsGeneratedName, variable.CloudAccountAWSOriginalAccountName, os.Getenv(environmentvariable.CloudAccountAWSEnvVarArn), "")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAWSEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSCloudSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudSecurityGroupAWSBasic(awsCloudAccountHCL, awsTypeAndName, resourceName, resourceTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "dome9_security_group_name", resourceTypeAndName, "dome9_security_group_name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "description", resourceTypeAndName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "aws_region_id", resourceTypeAndName, "aws_region_id"),
				),
			},
		},
	})
}
