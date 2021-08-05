package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/k8s"
	"github.com/dome9/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/variable"

	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccResourceCloudAccountKubernetesBasic(t *testing.T) {

	var cloudAccountResponse k8s.CloudAccountResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountKubernetes)
	anotherResourceTypeAndName, _, anotherGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountKubernetes)
	defaultOrganizationalUnitName := os.Getenv(environmentvariable.OrganizationalUnitName)
	organizationUnitTypeAndName, _, organizationUnitGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.OrganizationalUnit)
	organizationUnitHCL := getOrganizationalUnitResourceHCL(organizationUnitGeneratedName, variable.OrganizationalUnitName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountKubernetesEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAccountKubernetesDestroy,
		Steps: []resource.TestStep{
			{
				//Create Default
				Config: testAccCheckCloudAccountKubernetesBasic(resourceTypeAndName, generatedName, variable.CloudAccountKubernetesOriginalAccountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountKubernetesExists(resourceTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountKubernetesOriginalAccountName),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "creation_date"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "vendor", variable.CloudAccountKubernetesVendor),
					resource.TestCheckResourceAttr(resourceTypeAndName, "organizational_unit_name", defaultOrganizationalUnitName),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "runtime_protection.0.enabled"),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "admission_control.0.enabled"),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "image_assurance.0.enabled"),
				),
			},
			{
				//Update name
				Config: testAccCheckCloudAccountKubernetesBasic(resourceTypeAndName, generatedName, variable.CloudAccountKubernetesUpdatedAccountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountKubernetesExists(resourceTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountKubernetesUpdatedAccountName),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "creation_date"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "vendor", variable.CloudAccountKubernetesVendor),
					resource.TestCheckResourceAttr(resourceTypeAndName, "organizational_unit_name", defaultOrganizationalUnitName),
				),
			},
			{
				//Update OU
				Config: testAccCheckCloudAccountKubernetesBasicWithUpdatedOU(resourceTypeAndName, generatedName, variable.CloudAccountKubernetesUpdatedAccountName, organizationUnitHCL, organizationUnitTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountKubernetesExists(resourceTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.CloudAccountKubernetesUpdatedAccountName),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "creation_date"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "vendor", variable.CloudAccountKubernetesVendor),
					resource.TestCheckResourceAttrPair(resourceTypeAndName, "organizational_unit_id", organizationUnitTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(resourceTypeAndName, "organizational_unit_name", organizationUnitTypeAndName, "name"),
				),
			},
			{
				// Create with features
				Config: testAccCheckCloudAccountKubernetesCreateWithFeatures(anotherResourceTypeAndName, anotherGeneratedName, variable.CloudAccountKubernetesOriginalAccountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountKubernetesExists(anotherResourceTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(anotherResourceTypeAndName, "name", variable.CloudAccountKubernetesOriginalAccountName),
					resource.TestCheckResourceAttrSet(anotherResourceTypeAndName, "creation_date"),
					resource.TestCheckResourceAttr(anotherResourceTypeAndName, "vendor", variable.CloudAccountKubernetesVendor),
					resource.TestCheckResourceAttr(anotherResourceTypeAndName, "organizational_unit_name", defaultOrganizationalUnitName),
					resource.TestCheckResourceAttr(anotherResourceTypeAndName, "runtime_protection.0.enabled", strconv.FormatBool(variable.CloudAccountKubernetesRuntimeProtectionEnabled)),
					resource.TestCheckResourceAttr(anotherResourceTypeAndName, "admission_control.0.enabled", strconv.FormatBool(variable.CloudAccountKubernetesAdmissionControlEnabled)),
					resource.TestCheckResourceAttr(anotherResourceTypeAndName, "image_assurance.0.enabled", strconv.FormatBool(variable.CloudAccountKubernetesImageAssuranceEnabled)),
				),
			},
			{
				// Update features
				Config: testAccCheckCloudAccountKubernetesWithUpdateFeatures(anotherResourceTypeAndName, anotherGeneratedName, variable.CloudAccountKubernetesOriginalAccountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAccountKubernetesExists(anotherResourceTypeAndName, &cloudAccountResponse),
					resource.TestCheckResourceAttr(anotherResourceTypeAndName, "name", variable.CloudAccountKubernetesOriginalAccountName),
					resource.TestCheckResourceAttrSet(anotherResourceTypeAndName, "creation_date"),
					resource.TestCheckResourceAttr(anotherResourceTypeAndName, "vendor", variable.CloudAccountKubernetesVendor),
					resource.TestCheckResourceAttr(anotherResourceTypeAndName, "organizational_unit_name", defaultOrganizationalUnitName),
					resource.TestCheckResourceAttr(anotherResourceTypeAndName, "runtime_protection.0.enabled", strconv.FormatBool(variable.CloudAccountKubernetesRuntimeProtectionUpdateEnabled)),
					resource.TestCheckResourceAttr(anotherResourceTypeAndName, "admission_control.0.enabled", strconv.FormatBool(variable.CloudAccountKubernetesAdmissionControlUpdateEnabled)),
					resource.TestCheckResourceAttr(anotherResourceTypeAndName, "image_assurance.0.enabled", strconv.FormatBool(variable.CloudAccountKubernetesImageAssuranceUpdateEnabled)),
				),
			},
		},
	})
}

func testAccCheckCloudAccountKubernetesExists(resource string, cloudAccount *k8s.CloudAccountResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedCloudAccountResponse, _, err := apiClient.cloudaccountKubernetes.Get(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*cloudAccount = *receivedCloudAccountResponse

		return nil
	}
}

func testAccCheckCloudAccountKubernetesDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.CloudAccountKubernetes {
			continue
		}

		receivedCloudAccountResponse, _, err := apiClient.cloudaccountKubernetes.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if receivedCloudAccountResponse != nil {
			return fmt.Errorf("cloudaccount with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCloudAccountKubernetesEnvVarsPreCheck(t *testing.T) {
	if v := os.Getenv(environmentvariable.OrganizationalUnitName); v == "" {
		t.Fatalf("%s must be set for acceptance tests", environmentvariable.OrganizationalUnitName)
	}
}

func testAccCheckCloudAccountKubernetesBasic(resourceTypeAndName, generatedName, resourceName string) string {
	return fmt.Sprintf(`
// Kubernetes cloud account creation
%s

data "%s" "%s" {
 id = "${%s.id}"
}
`,
		// Kubernetes cloud account
		getBasicCloudAccountKubernetesResourceHCL(generatedName, resourceName),

		// data source variables
		resourcetype.CloudAccountKubernetes,
		generatedName,
		resourceTypeAndName,
	)
}

func testAccCheckCloudAccountKubernetesBasicWithUpdatedOU(resourceTypeAndName, generatedName, resourceName, organizationUnitHCL string, organizationUnitTypeAndName string) string {
	return fmt.Sprintf(`
// OU creation
%s
// Kubernetes cloud account creation
%s

data "%s" "%s" {
 id = "${%s.id}"
}
`,
		// ou arguments
		organizationUnitHCL,

		// Kubernetes cloud account arguments
		getCloudAccountKubernetesResourceHCLWithOU(generatedName, resourceName, organizationUnitTypeAndName),

		// data source variables
		resourcetype.CloudAccountKubernetes,
		generatedName,
		resourceTypeAndName,
	)
}

func testAccCheckCloudAccountKubernetesCreateOrUpdateWithFeatures(resourceTypeAndName, generatedName, resourceName string, isUpdate bool) string {
	var ac, ia, rp bool

	if isUpdate {
		ac = variable.CloudAccountKubernetesAdmissionControlUpdateEnabled
		ia = variable.CloudAccountKubernetesImageAssuranceUpdateEnabled
		rp = variable.CloudAccountKubernetesRuntimeProtectionUpdateEnabled
	} else {
		ac = variable.CloudAccountKubernetesAdmissionControlEnabled
		ia = variable.CloudAccountKubernetesImageAssuranceEnabled
		rp = variable.CloudAccountKubernetesRuntimeProtectionEnabled
	}

	return fmt.Sprintf(`
// Kubernetes cloud account with features
%s

data "%s" "%s" {
 id = "${%s.id}"
}
`,
		// Kubernetes cloud account
		getCloudAccountKubernetesResourceHCLWithfeatures(generatedName, resourceName, rp, ac, ia),

		// data source variables
		resourcetype.CloudAccountKubernetes,
		generatedName,
		resourceTypeAndName,
	)
}

func testAccCheckCloudAccountKubernetesCreateWithFeatures(resourceTypeAndName, generatedName, resourceName string) string {
	return testAccCheckCloudAccountKubernetesCreateOrUpdateWithFeatures(resourceTypeAndName, generatedName, resourceName, false)
}

func testAccCheckCloudAccountKubernetesWithUpdateFeatures(resourceTypeAndName, generatedName, resourceName string) string {
	return testAccCheckCloudAccountKubernetesCreateOrUpdateWithFeatures(resourceTypeAndName, generatedName, resourceName, true)
}

func getBasicCloudAccountKubernetesResourceHCL(generatedName string, resourceName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
 name                   = "%s"
}
`,
		// Kubernetes cloud account variables
		resourcetype.CloudAccountKubernetes,
		generatedName,
		resourceName,
	)
}

func getCloudAccountKubernetesResourceHCLWithOU(generatedName string, resourceName string, organizationUnitTypeAndName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
 name                   = "%s"
 organizational_unit_id = "${%s.id}"
}
`,
		// Kubernetes cloud account variables
		resourcetype.CloudAccountKubernetes,
		generatedName,
		resourceName,
		organizationUnitTypeAndName,
	)
}

func getCloudAccountKubernetesResourceHCLWithfeatures(generatedName string, resourceName string, runtimeProtection, admissionControl, imageAssurance bool) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
 name                   = "%s"
   runtime_protection {
	enabled = %v
  }
  admission_control {
	enabled = %v
  }
  image_assurance {
	enabled = %v
  }
}
`,
		// Kubernetes cloud account variables
		resourcetype.CloudAccountKubernetes,
		generatedName,
		resourceName,
		runtimeProtection,
		admissionControl,
		imageAssurance,
	)
}
