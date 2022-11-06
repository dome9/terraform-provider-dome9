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
	notificationHCL := getContinuousComplianceNotificationResourceHCL(notificationGeneratedName, continuousComplianceNotificationConfig())

	// Generate Admission Control Policy HCL Resource
	admissionPolicyHCL := getImageAssurancePolicyResourceHCL(kubernetesAccountHCL, kubernetesAccountResourceTypeAndName, notificationHCL,
		notificationTypeAndName, policyGeneratedName, false)
	admissionPolicyUpdatedHCL := getImageAssurancePolicyResourceHCL(kubernetesAccountHCL, kubernetesAccountResourceTypeAndName, notificationHCL,
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
				Config: testAccCheckImageAssurancePolicyBasic(admissionPolicyHCL, policyGeneratedName, policyTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageAssurancePolicyExists(policyTypeAndName, &response),
					resource.TestCheckResourceAttr(policyTypeAndName, "action", variable.ImageAssurancePolicyDetectAction),
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
				Config: testAccCheckImageAssurancePolicyBasic(admissionPolicyUpdatedHCL, policyGeneratedName, policyTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageAssurancePolicyExists(policyTypeAndName, &response),
					resource.TestCheckResourceAttr(policyTypeAndName, "action", variable.ImageAssurancePolicyPreventAction),
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

func testAccCheckImageAssurancePolicyExists(resource string, acPolicy *imageAssurancePolicy.ImageAssurancePolicyResponse) resource.TestCheckFunc {
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
			return fmt.Errorf("admission control policy with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckImageAssurancePolicyBasic(policyHCL, policyName, policyTypeAndName string) string {
	return fmt.Sprintf(`
// admission control policy resource
%s

// admission control policy data source
data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// admission control policy resource
		policyHCL,
		// Admission Control policy data source variables
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

// admission control policy resource creation
resource "%s" "%s" {
  target_id    = "${%s.id}"
  ruleset_id   = "%d"
  target_type  = "%s"
  notification_ids    = ["${%s.id}"]
  action       = "%s"
}
`,
		// kubernetes cloud account resource
		kubernetesAccountHCL,
		// continuous compliance notification resource
		notificationHCL,
		// Admission Control Policy resource type
		resourcetype.ImageAssurancePolicy,
		policyName,
		kubernetesCloudAccountTypeAndName,
		variable.ImageAssurancePolicyDefaultRulesetId,
		variable.ImageAssurancePolicyTargetType,
		notificationTypeAndName,
		IfThenElse(updateAction, variable.ImageAssurancePolicyPreventAction, variable.ImageAssurancePolicyDetectAction),
	)
}