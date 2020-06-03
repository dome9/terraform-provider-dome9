package dome9

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceCloudSecurityGroupAWSRoleBasic(t *testing.T) {
	securityGroupTypeAndName, _, securityGroupGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAWSSecurityGroup)
	securityGroupRoleTypeAndName, _, securityGroupRoleGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAWSSecurityGroupRole)
	awsTypeAndName, _, awsGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAWS)

	awsCloudAccountHCL := getCloudAccountAWSResourceHCL(awsGeneratedName, variable.CloudAccountAWSOriginalAccountName, os.Getenv(environmentvariable.CloudAccountAWSEnvVarArn), "")
	awsSecurityGroupHCL := getCloudAccountSecurityGroupAWSResourceHCL(securityGroupGeneratedName, securityGroupGeneratedName, awsTypeAndName, "")

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudSecurityGroupAWSRoleBasic(awsCloudAccountHCL, awsSecurityGroupHCL, securityGroupTypeAndName, securityGroupRoleGeneratedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(securityGroupRoleTypeAndName, "services.#", "1"),
				),
			},
		},
	})
}

func testAccCheckCloudSecurityGroupAWSRoleBasic(awsCloudAccountHCL, awsSecurityGroupHCL, securityGroupTypeAndName, securityGroupRoleGeneratedName string) string {
	return fmt.Sprintf(`
// aws cloud account resource
%s

// aws security group creation
%s

resource "%s" "%s" {
  dome9_security_group_id = "${%s.id}"
  services {
    inbound {
      name          = "inbound-test-aws-sg-role"
      description   = "inbound-test-aws-sg-role"
      protocol_type = "ALL"
      port          = ""
      open_for_all  = true
      scope {
        type = "CIDR"
        data = {
          cidr = "0.0.0.0/0"
          note = "Allow All Traffic"
        }
      }
    }
  }
}
`,
		awsCloudAccountHCL,
		awsSecurityGroupHCL,

		// resource variables
		resourcetype.CloudAccountAWSSecurityGroupRole,
		securityGroupRoleGeneratedName,
		securityGroupTypeAndName,
	)
}
