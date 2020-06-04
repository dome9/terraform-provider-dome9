package dome9

import (
	"fmt"
	"testing"

	"github.com/dome9/dome9-sdk-go/services/organizationalunits"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceOrganizationalUnitBasic(t *testing.T) {
	var ouResponse organizationalunits.OUResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.OrganizationalUnit)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOrganizationalUnitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOrganizationalUnitConfigure(resourceTypeAndName, generatedName, variable.OrganizationalUnitName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOrganizationalUnitExists(resourceTypeAndName, &ouResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.OrganizationalUnitName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "parent_id", variable.ParentID),
				),
			},

			// Update test
			{
				Config: testAccCheckOrganizationalUnitConfigure(resourceTypeAndName, generatedName, variable.OrganizationalUnitNameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOrganizationalUnitExists(resourceTypeAndName, &ouResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.OrganizationalUnitNameUpdate),
				),
			},
		},
	})
}

func testAccCheckOrganizationalUnitDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.OrganizationalUnit {
			continue
		}

		organizationalUnit, _, err := apiClient.organizationalUnit.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if organizationalUnit != nil {
			return fmt.Errorf("organizational unit with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckOrganizationalUnitExists(resource string, ou *organizationalunits.OUResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		ouResp, _, err := apiClient.organizationalUnit.Get(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*ou = *ouResp

		return nil
	}
}

func testAccCheckOrganizationalUnitConfigure(resourceTypeAndName, resourceName, ouName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
	name               = "%s"
	parent_id          = "%s"
}

data "%s" "%s" {
	id = "${%s.id}"
}
`,
		// resource variables
		resourcetype.OrganizationalUnit,
		resourceName,
		ouName,
		variable.ParentID,

		// data source variables
		resourcetype.OrganizationalUnit,
		resourceName,
		resourceTypeAndName,
	)
}
