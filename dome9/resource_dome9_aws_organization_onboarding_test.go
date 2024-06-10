package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws_org"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"os"
	"regexp"
	"testing"
)

func TestAccResourceAwsOrganizationOnboardingBasic(t *testing.T) {
	var response aws_org.OrganizationManagementViewModel
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AWSOrganizationOnboarding)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccAWSOrganizationOnboardingEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsOrganizationOnboardingDestroy,
		Steps: []resource.TestStep{
			{
				// creation test
				Config: testAccCheckAwsOrganizationOnboardingConfigure(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsOrganizationOnboardingExists(resourceTypeAndName, &response),
					resource.TestCheckResourceAttr(resourceTypeAndName, "aws_organization_name", variable.AwsOrganizationOnboardingCreationResourceName),
				),
				ExpectError: regexp.MustCompile(`.+Failed to assume management account role.+`),
			},
		},
	})
}

func testAccCheckAwsOrganizationOnboardingDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.AWSOrganizationOnboarding {
			continue
		}

		resp, _, err := apiClient.awsOrganizationOnboarding.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if resp != nil {
			return fmt.Errorf("aws org entity with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccAWSOrganizationOnboardingEnvVarsPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.AwsOrganizationOnboardingEnvVarRoleArn); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.AwsOrganizationOnboardingEnvVarRoleArn)
	}
	if v := os.Getenv(environmentvariable.AwsOrganizationOnboardingEnvVarSecret); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.AwsOrganizationOnboardingEnvVarSecret)
	}
	if v := os.Getenv(environmentvariable.AwsOrganizationOnboardingEnvVarStackSetArn); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.AwsOrganizationOnboardingEnvVarStackSetArn)
	}
}

func testAccCheckAwsOrganizationOnboardingExists(resource string, resp *aws_org.OrganizationManagementViewModel) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedCloudAccount, _, err := apiClient.awsOrganizationOnboarding.Get(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}

		*resp = *receivedCloudAccount

		return nil
	}
}

func testAccCheckAwsOrganizationOnboardingConfigure(resourceTypeAndName, generatedName string) string {
	return fmt.Sprintf(`
// Resource creation
%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// Resource code
		getAwsOrganizationOnboardingResourceHCL(generatedName),

		// Data source variables
		resourcetype.AWSOrganizationOnboarding,
		generatedName,
		resourceTypeAndName,
	)
}

func getAwsOrganizationOnboardingResourceHCL(resourceName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
	role_arn = "%s"
	secret  = "%s"
	stack_set_arn  = "%s"
	aws_organization_name  = "%s"
}
`,
		// Resource variables
		resourcetype.AWSOrganizationOnboarding,
		resourceName,
		os.Getenv(environmentvariable.AwsOrganizationOnboardingEnvVarRoleArn),
		os.Getenv(environmentvariable.AwsOrganizationOnboardingEnvVarSecret),
		os.Getenv(environmentvariable.AwsOrganizationOnboardingEnvVarStackSetArn),
		variable.AwsOrganizationOnboardingCreationResourceName,
	)
}
