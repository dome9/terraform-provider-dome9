package dome9

import (
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
)

func TestAccDataSourceAdmissionControlPolicyBasic(t *testing.T) {
	// Generate All Required Random Names for Testing
	policyTypeAndName, policyDataSourceTypeAndName, policyGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.AdmissionControlPolicy)
	notificationTypeAndName, _, notificationGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousComplianceNotification)
	kubernetesAccountResourceTypeAndName, _, kubernetesAccountGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountKubernetes)

	// Generate Kubernetes Account HCL Resource
	kubernetesAccountHCL := getCloudAccountKubernetesResourceHCLWithfeatures(kubernetesAccountGeneratedName, variable.AdmissionControlKubernetesAccountName,
		variable.CloudAccountKubernetesRuntimeProtectionEnabled,
		variable.CloudAccountKubernetesAdmissionControlEnabled,
		variable.CloudAccountKubernetesImageAssuranceEnabled,
		variable.CloudAccountKubernetesThreatIntelligenceEnabled)

	// Generate Notification HCL Configurations
	notificationHCL := getContinuousComplianceNotificationResourceHCL(notificationGeneratedName, continuousComplianceNotificationConfig())

	// Generate Admission Control Policy HCL Resource
	admissionPolicyHCL := getAdmissionControlPolicyResourceHCL(kubernetesAccountHCL, kubernetesAccountResourceTypeAndName, notificationHCL,
		notificationTypeAndName, policyGeneratedName, false)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAdmissionControlPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAdmissionControlPolicyBasic(admissionPolicyHCL, policyGeneratedName, policyTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(policyDataSourceTypeAndName, "action", policyTypeAndName, "action"),
					resource.TestCheckResourceAttrPair(policyDataSourceTypeAndName, "id", policyTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(policyDataSourceTypeAndName, "target_id", policyTypeAndName, "target_id"),
					resource.TestCheckResourceAttr(policyTypeAndName, "notification_ids.#", "2"),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "notification_ids.0"),
					resource.TestCheckResourceAttrSet(policyTypeAndName, "notification_ids.1"),
					resource.TestCheckResourceAttrPair(policyDataSourceTypeAndName, "ruleset_id", policyTypeAndName, "ruleset_id"),
					resource.TestCheckResourceAttrPair(policyDataSourceTypeAndName, "target_type", policyTypeAndName, "target_type"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
