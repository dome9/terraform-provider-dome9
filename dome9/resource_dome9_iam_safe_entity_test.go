package dome9

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"
	"github.com/dome9/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceIAMSafeEntityBasic(t *testing.T) {
	var cloudAccountResponse aws.CloudAccountResponse
	iamSafeEntityTypeAndName, _, iamSafeEntityGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.IAMSafeEntity)
	userTypeAndName, _, userGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.User)
	awsTypeAndName, _, awsGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAWS)
	attachIAMSafeTypeAndName, _, attachIAMSafeGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AttachIAMSafeToAwsCloudAccount)
	awsCloudAccountHCL := getCloudAccountAWSResourceHCL(awsGeneratedName, variable.CloudAccountAWSOriginalAccountName, os.Getenv(environmentvariable.CloudAccountAWSEnvVarArn), "")
	awsCloudAccountAndAttachIAMSafeHCL := testAccCheckAttachIAMSafeConfigure(awsCloudAccountHCL, awsTypeAndName, attachIAMSafeGeneratedName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAWSEnvVarsPreCheck(t)
			testAccAttachIAMSafeEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIAMEntityDestroy,
		Steps: []resource.TestStep{
			{
				Config: awsCloudAccountAndAttachIAMSafeHCL,
			},
			{
				PreConfig: func() { time.Sleep(time.Second * variable.WaitUntilAttachIAMSafeDone) },
				Config:    testAccCheckProtectIAMEntityConfigure(awsCloudAccountAndAttachIAMSafeHCL, attachIAMSafeTypeAndName, awsTypeAndName, iamSafeEntityGeneratedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProtectIAMEntityExists(iamSafeEntityTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(iamSafeEntityTypeAndName, "protection_mode", variable.IAMSafeEntityProtect),
					resource.TestCheckResourceAttr(iamSafeEntityTypeAndName, "entity_type", variable.IAMSafeEntityTypeUser),
					resource.TestCheckResourceAttr(iamSafeEntityTypeAndName, "entity_name", variable.IAMSafeEntityName),
				),
			},
			// update
			{
				Config: testAccCheckProtectIAMEntityUpdateConfigure(awsCloudAccountAndAttachIAMSafeHCL, attachIAMSafeTypeAndName, awsTypeAndName, userTypeAndName, iamSafeEntityGeneratedName, userGeneratedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProtectIAMEntityExists(iamSafeEntityTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(iamSafeEntityTypeAndName, "protection_mode", variable.IAMSafeEntityProtectWithElevation),
					resource.TestCheckResourceAttr(iamSafeEntityTypeAndName, "dome9_users_id_to_protect.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIAMEntityDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.IAMSafeEntity {
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

func testAccCheckProtectIAMEntityExists(resource string, awsCloudAccount *aws.CloudAccountResponse) resource.TestCheckFunc {
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
		*awsCloudAccount = *receivedAWSCloudAccount

		return nil
	}
}

func testAccCheckProtectIAMEntityConfigure(awsCloudAccountAndAttachIAMSafeHCL, attachIAMSafeTypeAndName, awsCloudAccountTypeAndName, iamGeneratedName string) string {
	return fmt.Sprintf(`
// aws cloud account and attach iam safe
%s

resource "%s" "%s" {
  depends_on           = [%s, %s]
  protection_mode      = "%s"
  entity_type          = "%s"
  entity_name          = "%s"
  aws_cloud_account_id = "${%s.id}"
}
`,
		awsCloudAccountAndAttachIAMSafeHCL,

		// protect IAM entity resource variables
		resourcetype.IAMSafeEntity,
		iamGeneratedName,
		attachIAMSafeTypeAndName,
		awsCloudAccountTypeAndName,

		variable.IAMSafeEntityProtect,
		variable.IAMSafeEntityTypeUser,
		variable.IAMSafeEntityName,
		awsCloudAccountTypeAndName,
	)
}

func testAccCheckProtectIAMEntityUpdateConfigure(awsCloudAccountAndAttachIAMSafeHCL, attachIAMSafeTypeAndName, awsCloudAccountTypeAndName, userTypeAndName, iamGeneratedName, userGeneratedName string) string {
	return fmt.Sprintf(`
// aws cloud account and attach iam safe
%s

// user resource
resource "%s" "%s" {
  email          = "%s"
  first_name     = "%s"
  last_name      = "%s"
  is_sso_enabled = "%s"
}

resource "%s" "%s" {
  depends_on                = [%s, %s]
  protection_mode           = "%s"
  entity_type               = "%s"
  entity_name               = "%s"
  aws_cloud_account_id      = "${%s.id}"
  dome9_users_id_to_protect = ["${%s.id}"]
}
`,
		awsCloudAccountAndAttachIAMSafeHCL,

		// user resource
		resourcetype.User,
		userGeneratedName,
		composeGenerateEmail(userGeneratedName),
		variable.UserFirstName,
		variable.UserLastName,
		strconv.FormatBool(variable.UserIsSsoEnabled),

		// protect IAM entity resource variables
		resourcetype.IAMSafeEntity,
		iamGeneratedName,
		attachIAMSafeTypeAndName,
		awsCloudAccountTypeAndName,
		variable.IAMSafeEntityProtectWithElevation,
		variable.IAMSafeEntityTypeUser,
		variable.IAMSafeEntityName,
		awsCloudAccountTypeAndName,
		userTypeAndName,
	)
}
