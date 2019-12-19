package dome9

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/dome9/dome9-sdk-go/services/roles"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceRoleBasic(t *testing.T) {
	var role roles.RoleResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.Role)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRoleConfigure(resourceTypeAndName, generatedName, variable.RoleDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoleExists(resourceTypeAndName, &role),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.RoleName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "description", variable.RoleDescription),
				),
			},

			// Update test
			{
				Config: testAccCheckRoleConfigure(resourceTypeAndName, generatedName, variable.RoleUpdateDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoleExists(resourceTypeAndName, &role),
					resource.TestCheckResourceAttr(resourceTypeAndName, "description", variable.RoleUpdateDescription),
				),
			},
		},
	})
}

func testAccCheckRoleDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.Role {
			continue
		}

		role, _, err := apiClient.role.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if role != nil {
			return fmt.Errorf("role with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckRoleExists(resource string, role *roles.RoleResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedRole, _, err := apiClient.role.Get(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*role = *receivedRole

		return nil
	}
}

func testAccCheckRoleConfigure(resourceTypeAndName, generatedName, description string) string {
	return fmt.Sprintf(`
// role resource
%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// resource variables
		RoleResourceHCL(generatedName, description),

		// data source variables
		resourcetype.Role,
		generatedName,
		resourceTypeAndName,
	)
}

func RoleResourceHCL(generatedName, description string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name        = "%s"
  description = "%s"
  permissions {
    access               = []
    manage               = []
    rulesets             = []
    notifications        = []
    policies             = []
    alert_actions        = []
    create               = []
    view                 = []
    on_boarding          = []
    cross_account_access = []
  }
}
`,
		// resource variables
		resourcetype.Role,
		generatedName,
		variable.RoleName,
		description,
	)
}
