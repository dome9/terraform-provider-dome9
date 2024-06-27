package dome9

import (
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccDataSourceCloudAccountKubernetesBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountKubernetes)
	defaultOrganizationalUnitName := os.Getenv(environmentvariable.OrganizationalUnitName)
	resourceName := variable.CloudAccountKubernetesOriginalAccountName + "_" + generatedName

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountKubernetesEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAccountKubernetesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAccountKubernetesBasic(resourceTypeAndName, generatedName, resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttr(dataSourceTypeAndName, "name", resourceName),
					resource.TestCheckResourceAttr(dataSourceTypeAndName, "vendor", variable.CloudAccountKubernetesVendor),
					resource.TestCheckResourceAttr(resourceTypeAndName, "organizational_unit_name", defaultOrganizationalUnitName),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "runtime_protection.0.enabled"),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "admission_control.0.enabled"),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "image_assurance.0.enabled"),
					resource.TestCheckResourceAttrSet(resourceTypeAndName, "threat_intelligence.0.enabled"),
				),
			},
		},
	})
}
