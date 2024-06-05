package dome9

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"testing"
)

func TestAccDataSourceCloudAccountOciBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountOCI)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountOciEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAccountOciDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAccountOciConfigure(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "name", resourceTypeAndName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "tenancy_id", resourceTypeAndName, "tenancy_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "home_region", resourceTypeAndName, "home_region"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "organizational_unit_id", resourceTypeAndName, "organizational_unit_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "credentials", resourceTypeAndName, "credentials"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "organizational_unit_path", resourceTypeAndName, "organizational_unit_path"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "organizational_unit_name", resourceTypeAndName, "organizational_unit_name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "vendor", resourceTypeAndName, "vendor"),
				),
				ExpectError: regexp.MustCompile(`.+Please retry the whole onboarding process \(including downloading a new Terraform file\).+`),
			},
		},
	})
}
