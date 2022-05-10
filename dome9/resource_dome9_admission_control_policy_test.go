package dome9

import (
	"encoding/json"
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/admissioncontrol/admission_policy"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceAdmissionControlPolicyBasic(t *testing.T) {
	var admissionControlPolicyResponse admission_policy.AdmissionControlPolicyResponse

	// Generate All Required Random Names for Testing
	policyTypeAndName, _, policyGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AdmissionControlPolicy)
	t.Log("policyTypeAndName: " + policyTypeAndName)
	t.Log("policyGeneratedName: " + policyGeneratedName)

	notificationTypeAndName, _, notificationGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousComplianceNotification)
	t.Log("notificationTypeAndName: " + notificationTypeAndName)
	t.Log("notificationGeneratedName: " + notificationGeneratedName)

	kubernetesAccountResourceTypeAndName, _, kubernetesAccountGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountKubernetes)
	t.Log("kubernetesAccountResourceTypeAndName: " + kubernetesAccountResourceTypeAndName)
	t.Log("kubernetesAccountGeneratedName: " + kubernetesAccountGeneratedName)

	// Create Kubernetes Account HCL Resource
	kubernetesAccountHCL := getCloudAccountKubernetesResourceHCLWithfeatures(kubernetesAccountGeneratedName, variable.AdmissionControlKubernetesAccountName,
		variable.CloudAccountKubernetesRuntimeProtectionEnabled,
		variable.CloudAccountKubernetesAdmissionControlEnabled,
		variable.CloudAccountKubernetesImageAssuranceEnabled)

	t.Log("Kubernetes HCL" + kubernetesAccountHCL)

	// Create Compliance Notification HCL Resource
	notificationHCL := getContinuousComplianceNotificationResourceHCL(notificationGeneratedName, continuousComplianceNotificationConfig())
	t.Log("notificationHCL : " + notificationHCL)

	// Create Admission Control Policy HCL Resource
	admissionPolicyHCL := getAdmissionControlPolicyResourceHCL(kubernetesAccountHCL, kubernetesAccountResourceTypeAndName, notificationHCL,
		notificationTypeAndName, policyGeneratedName, false)
	t.Log("admissionPolicyHCL : " + admissionPolicyHCL)
	admissionPolicyUpdatedHCL := getAdmissionControlPolicyResourceHCL(kubernetesAccountHCL, kubernetesAccountResourceTypeAndName, notificationHCL,
		notificationTypeAndName, policyGeneratedName, true)
	t.Log("admissionPolicyUpdatedHCL : " + admissionPolicyUpdatedHCL)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAdmissionControlPolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Create Policy Test Step
				Config: testAccCheckAdmissionControlPolicyBasic(admissionPolicyHCL, policyGeneratedName, policyTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmissionControlPolicyExists(policyTypeAndName, &admissionControlPolicyResponse),
					resource.TestCheckResourceAttr(policyTypeAndName, "action", variable.AdmissionControlPolicyDetectAction),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "id"),
					resource.TestCheckResourceAttr(policyTypeAndName, "notification_ids.#", "2"),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "notification_ids.0"),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "notification_ids.1"),
					resource.TestCheckResourceAttr(policyTypeAndName, "ruleset_platform", variable.AdmissionControlPolicyRuleSetPlatform),
					resource.TestCheckResourceAttr(policyTypeAndName, "ruleset_id", strconv.Itoa(variable.AdmissionControlPolicyDefaultRulesetId)),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "target_id"),
					resource.TestCheckResourceAttr(policyTypeAndName, "target_type", variable.AdmissionControlPolicyTargetType),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				// Update Policy Test Step from Detection to Prevention
				Config: testAccCheckAdmissionControlPolicyBasic(admissionPolicyUpdatedHCL, policyGeneratedName, policyTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmissionControlPolicyExists(policyTypeAndName, &admissionControlPolicyResponse),
					resource.TestCheckResourceAttr(policyTypeAndName, "action", variable.AdmissionControlPolicyPreventAction),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "id"),
					resource.TestCheckResourceAttr(policyTypeAndName, "notification_ids.#", "1"),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "notification_ids.0"),
					resource.TestCheckResourceAttr(policyTypeAndName, "ruleset_platform", variable.AdmissionControlPolicyRuleSetPlatform),
					resource.TestCheckResourceAttr(policyTypeAndName, "ruleset_id", strconv.Itoa(variable.AdmissionControlPolicyDefaultRulesetId)),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "target_id"),
					resource.TestCheckResourceAttr(policyTypeAndName, "target_type", variable.AdmissionControlPolicyTargetType),
				),
			},
		},
	})
}

func testAccCheckAdmissionControlPolicyExists(resource string, acPolicy *admission_policy.AdmissionControlPolicyResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}
		b, err := json.MarshalIndent(rs, "", "\t")
		fmt.Println("Resource State : " + string(b))
		apiClient := testAccProvider.Meta().(*Client)
		receivedAdmissionControlPolicyResponse, _, err := apiClient.admissionControlPolicy.Get(rs.Primary.ID)

		policy, err := json.MarshalIndent(receivedAdmissionControlPolicyResponse, "", "\t")
		fmt.Println("Admission Control Policy : " + string(policy))

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*acPolicy = *receivedAdmissionControlPolicyResponse

		return nil
	}
}

func testAccCheckAdmissionControlPolicyDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.AdmissionControlPolicy {
			continue
		}

		admissionControlPolicyResponse, _, err := apiClient.admissionControlPolicy.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if admissionControlPolicyResponse != nil {
			return fmt.Errorf("admission control policy with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckAdmissionControlPolicyBasic(policyHCL, policyName, policyTypeAndName string) string {
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
		resourcetype.AdmissionControlPolicy,
		policyName,
		policyTypeAndName,
	)
}

func getAdmissionControlPolicyResourceHCL(kubernetesAccountHCL, kubernetesCloudAccountTypeAndName, notificationHCL, notificationTypeAndName, policyName string, updateAction bool) string {
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
  ruleset_platform = "%s"
}
`,
		// kubernetes cloud account resource
		kubernetesAccountHCL,
		// continuous compliance notification resource
		notificationHCL,
		// Admission Control Policy resource type
		resourcetype.AdmissionControlPolicy,
		policyName,
		kubernetesCloudAccountTypeAndName,
		variable.AdmissionControlPolicyDefaultRulesetId,
		variable.AdmissionControlPolicyTargetType,
		notificationTypeAndName,
		IfThenElse(updateAction, variable.AdmissionControlPolicyPreventAction, variable.AdmissionControlPolicyDetectAction),
		variable.AdmissionControlPolicyRuleSetPlatform,
	)
}

// IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
