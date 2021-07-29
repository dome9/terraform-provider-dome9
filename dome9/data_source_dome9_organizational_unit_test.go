package dome9

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/dome9/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccDataSourceOrganizationalUnitBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.OrganizationalUnit)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOrganizationalUnitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOrganizationalUnitConfigure(resourceTypeAndName, generatedName, variable.OrganizationalUnitName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "name", resourceTypeAndName, "name"),
				),
			},
		},
	})
}
