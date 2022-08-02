package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/assessment"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceAssessmentBasic(t *testing.T) {
	var assessmentData assessment.RunBundleResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.Assessment)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAssessmentDestroy,
		Steps: []resource.TestStep{
			{
				// creation test
				Config: testAccCheckAssessmentConfigure(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAssessmentExists(resourceTypeAndName, &assessmentData),
					resource.TestCheckResourceAttr(resourceTypeAndName, "triggered_by", variable.TriggeredBy),
					//resource.TestCheckResourceAttr(resourceTypeAndName, "assessment_passed", variable.AssessmentPassed),
					//resource.TestCheckResourceAttr(resourceTypeAndName, "has_errors", variable.HasErrors),
					//resource.TestCheckResourceAttr(resourceTypeAndName, "has_data_sync_status_issues", variable.HasDataSyncStatusIssues),
				),
			},
		},
	})
}

func testAccCheckAssessmentDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.Assessment {
			continue
		}

		resp, _, err := apiClient.assessment.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if resp != nil {
			return fmt.Errorf("assessment with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckAssessmentExists(resource string, resp *assessment.RunBundleResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedCloudAccount, _, err := apiClient.assessment.Get(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}

		*resp = *receivedCloudAccount

		return nil
	}
}

func testAccCheckAssessmentConfigure(resourceTypeAndName, generatedName string) string {
	return fmt.Sprintf(`
// Assessment creation
%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// assessment
		getAssessmentResourceHCL(generatedName),

		// data source variables
		resourcetype.Assessment,
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
