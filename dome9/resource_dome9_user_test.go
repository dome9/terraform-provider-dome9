package dome9

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/dome9/dome9-sdk-go/services/users"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceUsersBasic(t *testing.T) {
	var usersResponse users.UserResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.User)

	roleTypeAndName, _, roleName := method.GenerateRandomSourcesTypeAndName(resourcetype.Role)
	roleHCL := RoleResourceHCL(roleName, variable.RoleDescription, variable.RoleToPermittedAlertActions)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckUsersConfigure(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists(resourceTypeAndName, &usersResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "email", composeGenerateEmail(generatedName)),
					resource.TestCheckResourceAttr(resourceTypeAndName, "first_name", variable.UserFirstName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "last_name", variable.UserLastName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "is_sso_enabled", strconv.FormatBool(variable.UserIsSsoEnabled)),
				),
			},
			{
				Config: testAccCheckUsersUpdateConfigure(roleHCL, roleTypeAndName, resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists(resourceTypeAndName, &usersResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "role_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "permit_alert_actions", strconv.FormatBool(variable.RoleUpdateToPermittedAlertActions)),
				),
			},
		},
	})
}

func testAccCheckUserDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.User {
			continue
		}

		user, _, err := apiClient.users.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if user != nil {
			return fmt.Errorf("user with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckUserExists(resource string, user *users.UserResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedUser, _, err := apiClient.users.Get(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*user = *receivedUser

		return nil
	}
}

func testAccCheckUsersConfigure(resourceTypeAndName, generatedName string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  email          = "%s"
  first_name     = "%s"
  last_name      = "%s"
  is_sso_enabled = "%s"
}

data "%s" "%s" {
	id = "${%s.id}"
}
`,

		// user resource variables
		resourcetype.User,
		generatedName,
		composeGenerateEmail(generatedName),
		variable.UserFirstName,
		variable.UserLastName,
		strconv.FormatBool(variable.UserIsSsoEnabled),

		// data source variables
		resourcetype.User,
		generatedName,
		resourceTypeAndName,
	)
}

func testAccCheckUsersUpdateConfigure(roleHCL, roleTypeAndName, resourceTypeAndName, generatedName string) string {
	return fmt.Sprintf(`
// role
%s

resource "%s" "%s" {
  email                = "%s"
  first_name           = "%s"
  last_name            = "%s"
  is_sso_enabled       = "%s"
  role_ids             = ["${%s.id}"]
  permit_alert_actions = "%s"
}

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// role resource HCL
		roleHCL,

		// user resource variables
		resourcetype.User,
		generatedName,
		composeGenerateEmail(generatedName),
		variable.UserFirstName,
		variable.UserLastName,
		strconv.FormatBool(variable.UserIsSsoEnabled),
		roleTypeAndName,
		strconv.FormatBool(variable.RoleUpdateToPermittedAlertActions),

		// data source variables
		resourcetype.User,
		generatedName,
		resourceTypeAndName,
	)
}

func composeGenerateEmail(emailName string) string {
	return fmt.Sprintf("%s@%s.com", emailName, emailName)
}
