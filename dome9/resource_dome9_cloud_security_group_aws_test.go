package dome9

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/dome9/dome9-sdk-go/services/cloudsecuritygroup/securitygroupaws"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceCloudSecurityGroupAWSBasic(t *testing.T) {
	var cloudSecurityGroupAWSResponse securitygroupaws.CloudSecurityGroupResponse
	securityGroupTypeAndName, _, securityGroupGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAWSSecurityGroup)
	awsTypeAndName, _, awsGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAWS)

	awsCloudAccountHCL := getCloudAccountAWSResourceHCL(awsGeneratedName, variable.CloudAccountAWSOriginalAccountName, os.Getenv(environmentvariable.CloudAccountAWSEnvVarArn), "")

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAWSEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudSecurityGroupAWSDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudSecurityGroupAWSBasic(awsCloudAccountHCL, awsTypeAndName, securityGroupGeneratedName, securityGroupTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudSecurityGroupAWSExists(securityGroupTypeAndName, &cloudSecurityGroupAWSResponse),
					resource.TestCheckResourceAttr(securityGroupTypeAndName, "dome9_security_group_name", variable.AWSSecurityGroupName),
					resource.TestCheckResourceAttr(securityGroupTypeAndName, "description", variable.AWSSecurityGroupDescription),
					resource.TestCheckResourceAttr(securityGroupTypeAndName, "aws_region_id", variable.AWSSecurityGroupRegionID),
				),
			},
		},
	})
}

func testAccCheckCloudSecurityGroupAWSExists(resource string, securityGroup *securitygroupaws.CloudSecurityGroupResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedSecurityGroupResponse, _, err := apiClient.awsSecurityGroup.Get(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*securityGroup = *receivedSecurityGroupResponse

		return nil
	}
}

func testAccCheckCloudSecurityGroupAWSDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.CloudAccountAWSSecurityGroup {
			continue
		}

		receivedAWSSecurityGroupResponse, _, err := apiClient.awsSecurityGroup.Get(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if receivedAWSSecurityGroupResponse != nil {
			return fmt.Errorf("security group with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckCloudSecurityGroupAWSBasic(awsCloudAccountHCL, awsCloudAccountTypeAndName, securityGroupResourceName, securityGroupTypeAndName string) string {
	return fmt.Sprintf(`
// aws cloud account resource
%s

// aws security group creation
resource "%s" "%s" {
  dome9_security_group_name = "%s"
  description               = "%s"
  aws_region_id             = "%s"
  dome9_cloud_account_id    = "${%s.id}"
}

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		awsCloudAccountHCL,

		// resource variables
		resourcetype.CloudAccountAWSSecurityGroup,
		securityGroupResourceName,
		variable.AWSSecurityGroupName,
		variable.AWSSecurityGroupDescription,
		variable.AWSSecurityGroupRegionID,
		awsCloudAccountTypeAndName,

		// data source variables
		resourcetype.CloudAccountAWSSecurityGroup,
		securityGroupResourceName,
		securityGroupTypeAndName,
	)
}
