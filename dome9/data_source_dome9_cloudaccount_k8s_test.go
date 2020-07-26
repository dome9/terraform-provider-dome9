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

func TestAccDataSourceCloudAccountK8SBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountK8S)
	defaultOrganizationalUnitName := os.Getenv(environmentvariable.OrganizationalUnitName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountK8SEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAccountK8SDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAccountK8SBasic(resourceTypeAndName, generatedName, variable.CloudAccountK8SOriginalAccountName, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttr(dataSourceTypeAndName, "name", variable.CloudAccountK8SOriginalAccountName), // TODO: why other implementation used pair?
					resource.TestCheckResourceAttr(dataSourceTypeAndName, "vendor", variable.CloudAccountK8SVendor),
					resource.TestCheckResourceAttr(resourceTypeAndName, "organizational_unit_name", defaultOrganizationalUnitName),
				),
			},
		},
	})
}
