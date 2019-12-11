package dome9

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/dome9/dome9-sdk-go/services/cloudsecuritygroup/securitygroupazure"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceAzureSecurityGroupBasic(t *testing.T) {
	var azureSecurityGroupResponse securitygroupazure.AzureSecurityGroupResponse
	// time.Sleep(20 * time.Second)
	securityGroupTypeAndName, _, securityGroupGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAzureSecurityGroup)
	azureTypeAndName, _, azureGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAzure)
	azureCloudAccountHCL := getCloudAccountAzureResourceHCL(azureGeneratedName, variable.CloudAccountAzureCreationResourceName, variable.CloudAccountAzureUpdateOperationMode)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAzureEnvVarsPreCheck(t)
			testAccAzureSecurityGroupEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAzureSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAzureSecurityGroupBasic(azureCloudAccountHCL, azureTypeAndName, securityGroupGeneratedName, securityGroupTypeAndName, azureSecurityGroupAdditionalBlock()),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAzureSecurityGroupExists(securityGroupTypeAndName, &azureSecurityGroupResponse),
					resource.TestCheckResourceAttr(securityGroupTypeAndName, "dome9_security_group_name", securityGroupGeneratedName),
					resource.TestCheckResourceAttr(securityGroupTypeAndName, "description", variable.AzureSecurityGroupDescription),
					resource.TestCheckResourceAttr(securityGroupTypeAndName, "is_tamper_protected", strconv.FormatBool(variable.AzureSecurityGroupIsTamperProtected)),
					resource.TestCheckResourceAttr(securityGroupTypeAndName, "tags.0.value", variable.AzureSecurityGroupTagValue),
				),
			},
			// Update test
			{
				Config: testAccCheckAzureSecurityGroupBasic(azureCloudAccountHCL, azureTypeAndName, securityGroupGeneratedName, securityGroupTypeAndName, azureSecurityGroupUpdateAdditionalBlock()),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAzureSecurityGroupExists(securityGroupTypeAndName, &azureSecurityGroupResponse),
					resource.TestCheckResourceAttr(securityGroupTypeAndName, "description", variable.AzureSecurityGroupUpdateDescription),
					resource.TestCheckResourceAttr(securityGroupTypeAndName, "is_tamper_protected", strconv.FormatBool(variable.AzureSecurityGroupUpdateIsTamperProtected)),
					resource.TestCheckResourceAttr(securityGroupTypeAndName, "tags.0.value", variable.AzureSecurityGroupUpdateTagValue),
				),
			},
		},
	})
}

func testAccCheckAzureSecurityGroupExists(resource string, securityGroup *securitygroupazure.AzureSecurityGroupResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedSecurityGroupResponse, _, err := apiClient.azureSecurityGroup.Get(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*securityGroup = *receivedSecurityGroupResponse

		return nil
	}
}

func testAccAzureSecurityGroupEnvVarsPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.AzureSecurityGroupResourceGroup); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.AzureSecurityGroupResourceGroup)
	}
}

func testAccCheckAzureSecurityGroupDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.CloudAccountAzureSecurityGroup {
			continue
		}

		receivedAzureSecurityGroupResponse, _, err := apiClient.azureSecurityGroup.Get(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if receivedAzureSecurityGroupResponse != nil {
			return fmt.Errorf("security group with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckAzureSecurityGroupBasic(azureCloudAccountHCL, azureCloudAccountTypeAndName, securityGroupResourceName, securityGroupTypeAndName, additionalBlock string) string {
	return fmt.Sprintf(`
// azure cloud account resource
%s

// azure security group resource
resource "%s" "%s" {
 dome9_security_group_name = "%s"
 region                    = "%s"
 resource_group            = "%s"
 dome9_cloud_account_id    = "${%s.id}"
 %s
}

// azure sequrity group data source
data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// azure cloud account resource variables
		azureCloudAccountHCL,

		// // azure security group resource variables
		resourcetype.CloudAccountAzureSecurityGroup,
		securityGroupResourceName,
		securityGroupResourceName,
		variable.AzureSecurityGroupRegion,
		os.Getenv(environmentvariable.AzureSecurityGroupResourceGroup),
		azureCloudAccountTypeAndName,
		additionalBlock,

		// data source variables
		resourcetype.CloudAccountAzureSecurityGroup,
		securityGroupResourceName,
		securityGroupTypeAndName,
	)
}

func azureSecurityGroupAdditionalBlock() string {
	return fmt.Sprintf(`
description    = "%s"
is_tamper_protected = "%s"
tags {
	key = "tag_key"
	value = "%s"
}

`,
		variable.AzureSecurityGroupDescription,
		strconv.FormatBool(variable.AzureSecurityGroupIsTamperProtected),
		variable.AzureSecurityGroupTagValue,
	)
}

func azureSecurityGroupUpdateAdditionalBlock() string {
	return fmt.Sprintf(`
description    = "%s"
is_tamper_protected = "%s"
tags {
	key = "tag_key"
	value = "%s"
}

`,
		variable.AzureSecurityGroupUpdateDescription,
		strconv.FormatBool(variable.AzureSecurityGroupUpdateIsTamperProtected),
		variable.AzureSecurityGroupUpdateTagValue,
	)
}
