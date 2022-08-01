package dome9

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/alibaba"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceAssessmentBasic(t *testing.T) {
	var cloudAccountAlibaba alibaba.CloudAccountResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.Assessment)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAlibabaEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAssessmentDestroy,
		Steps: []resource.TestStep{
			{
				// creation test
				Config: testAccCheckAssessmentConfigure(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountAlibabaExists(resourceTypeAndName, &cloudAccountAlibaba),
					resource.TestCheckResourceAttr(resourceTypeAndName, "bundle_id", strconv.Itoa(variable.BundleID)),
					resource.TestCheckResourceAttr(resourceTypeAndName, "dome9_cloud_account_id", variable.Dome9CloudAccountID),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cloud_account_id", variable.CloudAccountID),
					resource.TestCheckResourceAttr(resourceTypeAndName, "cloud_account_type", variable.CloudAccountType),
					resource.TestCheckResourceAttr(resourceTypeAndName, "request_id", variable.RequestID),
				),
			},
		},
	})
}

func testAccCheckAssessmentDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.CloudAccountAlibaba {
			continue
		}

		resp, _, err := apiClient.cloudaccountAlibaba.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if resp != nil {
			return fmt.Errorf("cloudaccounts with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckAssessmentConfigure(resourceTypeAndName, generatedName string) string {
	return fmt.Sprintf(`
// Assessment creation
%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// Alibaba cloud account
		getAssessmentResourceHCL(generatedName),

		// data source variables
		resourcetype.CloudAccountAlibaba,
		generatedName,
		resourceTypeAndName,
	)
}

func getAssessmentResourceHCL(assessmentName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
      bundle_id = "%v"
	  dome9_cloud_account_id = "%s"
	  cloud_account_id = "%s"
	  cloud_account_type = "%s"
	  should_minimize_result = "%v"
	  request_id = "%s"
}
`,
		// Assessment variables
		resourcetype.Assessment,
		assessmentName,
		variable.BundleID,
		variable.Dome9CloudAccountID,
		variable.CloudAccountID,
		variable.CloudAccountType,
		variable.ShouldMinimizeResult,
		variable.RequestID,
	)
}
