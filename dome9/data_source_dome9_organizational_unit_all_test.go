package dome9

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceOrganizationalUnitAllBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOrganizationalUnitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOrganizationalUnitAllBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.dome9_all_organizational_units.all_units", "items.#"),
				),
			},
		},
	})
}

func testAccCheckOrganizationalUnitAllBasic() string {
	return `
data "dome9_all_organizational_units" "all_units" {}

output "all_organizational_units" {
  value = data.dome9_all_organizational_units.all_units
}
`
}
