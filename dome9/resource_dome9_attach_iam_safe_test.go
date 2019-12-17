package dome9

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceAttachIAMSafeBasic(t *testing.T) {
	var cloudAccountResponse aws.CloudAccountResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AttachIAMSafeToAwsCloudAccount)
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
					testAccCheckAttachIAMSafeExists(resourceTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "aws_group_arn", os.Getenv(environmentvariable.AttachIAMSafeEnvVarGroupArn)),
					resource.TestCheckResourceAttr(resourceTypeAndName, "aws_policy_arn", os.Getenv(environmentvariable.AttachIAMSafeEnvVarPolicyArn)),
				),
			},
		},
	})
}

func testAccCheckAttachIAMSafeDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.AttachIAMSafeToAwsCloudAccount {
			continue
		}

		awsCloudAccountResponse, _, err := apiClient.cloudaccountAWS.Get(cloudaccounts.QueryParameters{ID: rs.Primary.ID})

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if awsCloudAccountResponse != nil {
			return fmt.Errorf("aws cloud account with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckAttachIAMSafeExists(resource string, awsCloudAcccount *aws.CloudAccountResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedAWSCloudAccount, _, err := apiClient.cloudaccountAWS.Get(cloudaccounts.QueryParameters{ID: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*awsCloudAcccount = *receivedAWSCloudAccount

		return nil
	}
}

func testAccAttachIAMSafeEnvVarsPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.AttachIAMSafeEnvVarGroupArn); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.AttachIAMSafeEnvVarGroupArn)
	}
	if v := os.Getenv(environmentvariable.AttachIAMSafeEnvVarPolicyArn); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.AttachIAMSafeEnvVarPolicyArn)
	}
}

func testAccCheckAttachIAMSafeConfigure(awsCloudAccountHCL, awsCloudAccountTypeAndName, resourceTypeAndName, generatedName string) string {
	return fmt.Sprintf(`
// aws cloud account resource
%s

resource "%s" "%s" {
  aws_cloud_account_id = "${%s.id}"
  aws_group_arn        = "%s"
  aws_policy_arn       = "%s"
}

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		awsCloudAccountHCL,

		// resource variables
		resourcetype.AttachIAMSafeToAwsCloudAccount,
		generatedName,
		awsCloudAccountTypeAndName,
		os.Getenv(environmentvariable.AttachIAMSafeEnvVarGroupArn),
		os.Getenv(environmentvariable.AttachIAMSafeEnvVarPolicyArn),

		// data source variables
		resourcetype.AttachIAMSafeToAwsCloudAccount,
		generatedName,
		resourceTypeAndName,
	)
}
