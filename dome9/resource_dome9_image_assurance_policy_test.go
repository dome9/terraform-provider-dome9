package dome9

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/dome9/dome9-sdk-go/services/imageassurance/imageassurance_policy"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceImagePolicyPolicyBasic(t *testing.T) {
	var response imageassurance_policy.ImageAssurancePolicyResponse

	// Generate All Required Random Names for Testing
	policyTypeAndName, _, policyGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ImageAssurancePolicy)
	notificationTypeAndName, _, notificationGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousComplianceNotification)
	kubernetesAccountResourceTypeAndName, _, kubernetesAccountGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountKubernetes)

	// Generate Kubernetes Account HCL Resource
	kubernetesAccountHCL := getCloudAccountKubernetesResourceHCLWithfeatures(kubernetesAccountGeneratedName, variable.ImageAssuranceKubernetesAccountName,
		variable.CloudAccountKubernetesRuntimeProtectionEnabled,
		variable.CloudAccountKubernetesAdmissionControlEnabled,
		variable.CloudAccountKubernetesImageAssuranceEnabled,
		variable.CloudAccountKubernetesThreatIntelligenceEnabled)

	// Generate Compliance Notification HCL Resource
	notificationHCL := getContinuousComplianceNotificationResourceHCL(notificationGeneratedName, continuousComplianceNotificationConfig(notificationGeneratedName))

	// Generate Image Assurance Policy HCL Resource
	imageAssurancePolicyHCL := getImageAssurancePolicyResourceHCL(kubernetesAccountHCL, kubernetesAccountResourceTypeAndName, notificationHCL,
		notificationTypeAndName, policyGeneratedName, false)
	imageAssurancePolicyUpdatedHCL := getImageAssurancePolicyResourceHCL(kubernetesAccountHCL, kubernetesAccountResourceTypeAndName, notificationHCL,
		notificationTypeAndName, policyGeneratedName, true)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckImageAssurancePolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Create Policy Test Step
				Config: testCheckImageAssurancePolicyBasic(imageAssurancePolicyHCL, policyGeneratedName, policyTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testCheckImageAssurancePolicyExists(policyTypeAndName, &response),
					resource.TestCheckResourceAttr(policyTypeAndName, "admission_control_action", variable.ImageAssurancePolicyDetectAction),
					resource.TestCheckResourceAttr(policyTypeAndName, "admission_control_unscanned_action", variable.ImageAssurancePolicyDetectAction),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "id"),
					resource.TestCheckResourceAttr(policyTypeAndName, "notification_ids.#", "2"),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "notification_ids.0"),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "notification_ids.1"),
					resource.TestCheckResourceAttr(policyTypeAndName, "ruleset_id", strconv.Itoa(variable.ImageAssurancePolicyDefaultRulesetId)),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "target_id"),
					resource.TestCheckResourceAttr(policyTypeAndName, "target_type", variable.ImageAssurancePolicyTargetType),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				// Update Policy Test Step from Detection to Prevention
				Config: testCheckImageAssurancePolicyBasic(imageAssurancePolicyUpdatedHCL, policyGeneratedName, policyTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testCheckImageAssurancePolicyExists(policyTypeAndName, &response),
					resource.TestCheckResourceAttr(policyTypeAndName, "admission_control_action", variable.ImageAssurancePolicyPreventAction),
					resource.TestCheckResourceAttr(policyTypeAndName, "admission_control_unscanned_action", variable.ImageAssurancePolicyPreventAction),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "id"),
					resource.TestCheckResourceAttr(policyTypeAndName, "notification_ids.#", "1"),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "notification_ids.0"),
					resource.TestCheckResourceAttr(policyTypeAndName, "ruleset_id", strconv.Itoa(variable.ImageAssurancePolicyDefaultRulesetId)),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "target_id"),
					resource.TestCheckResourceAttr(policyTypeAndName, "target_type", variable.ImageAssurancePolicyTargetType),
				),
			},
		},
	})
}

func testCheckImageAssurancePolicyExists(resource string, acPolicy *imageassurance_policy.ImageAssurancePolicyResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}
		apiClient := testAccProvider.Meta().(*Client)
		receivedImageAssurancePolicyResponse, _, err := apiClient.imageAssurancePolicy.Get(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*acPolicy = *receivedImageAssurancePolicyResponse

		return nil
	}
}

func testAccCheckImageAssurancePolicyDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.ImageAssurancePolicy {
			continue
		}

		ImageAssurancePolicyResponse, _, err := apiClient.imageAssurancePolicy.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if ImageAssurancePolicyResponse != nil {
			return fmt.Errorf("image assurance policy with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testCheckImageAssurancePolicyBasic(policyHCL, policyName, policyTypeAndName string) string {
	return fmt.Sprintf(`
// image assurance policy resource
%s

// image assurance policy data source
data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// image assurance policy resource
		policyHCL,
		// image assurance policy data source variables
		resourcetype.ImageAssurancePolicy,
		policyName,
		policyTypeAndName,
	)
}

func getImageAssurancePolicyResourceHCL(kubernetesAccountHCL, kubernetesCloudAccountTypeAndName, notificationHCL, notificationTypeAndName, policyName string, updateAction bool) string {
	return fmt.Sprintf(`
// kubernetes cloud account resource
%s

// continuous compliance notification resource
%s

// image assurance policy resource creation
resource "%s" "%s" {
  target_id    = "${%s.id}"
  ruleset_id   = "%d"
  target_type  = "%s"
  notification_ids    = ["${%s.id}"]
  admission_control_action       = "%s"
  admission_control_unscanned_action       = "%s"
}
`,
		// kubernetes cloud account resource
		kubernetesAccountHCL,
		// continuous compliance notification resource
		notificationHCL,
		// image assurance Policy resource type
		resourcetype.ImageAssurancePolicy,
		policyName,
		kubernetesCloudAccountTypeAndName,
		variable.ImageAssurancePolicyDefaultRulesetId,
		variable.ImageAssurancePolicyTargetType,
		notificationTypeAndName,
		IfThenElse(updateAction, variable.ImageAssurancePolicyPreventAction, variable.ImageAssurancePolicyDetectAction),
		IfThenElse(updateAction, variable.ImageAssurancePolicyPreventAction, variable.ImageAssurancePolicyDetectAction),
	)
}
