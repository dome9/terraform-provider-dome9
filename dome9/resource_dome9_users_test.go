package dome9

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/dome9/dome9-sdk-go/services/users"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceUsersBasic(t *testing.T) {
	var usersResponse users.UserResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.Users)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckUsersConfigure(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists(resourceTypeAndName, &usersResponse),
					resource.TestCheckResourceAttr(resourceTypeAndName, "email", variable.UserEmail),
					resource.TestCheckResourceAttr(resourceTypeAndName, "first_name", variable.UserFirstName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "last_name", variable.UserLastName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "is_sso_enabled", strconv.FormatBool(variable.UserIsSsoEnabled)),
				),
			},
		},
	})
}

func testAccCheckUserDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.Users {
			continue
		}

		ipList, _, err := apiClient.users.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if ipList != nil {
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
  email = "%s"
  first_name = "%s"
  last_name = "%s"
  is_sso_enabled = "%s"
}

data "%s" "%s" {
	id = "${%s.id}"
}
`,
		// resource variables
		resourcetype.Users,
		generatedName,
		variable.UserEmail,
		variable.UserFirstName,
		variable.UserLastName,
		strconv.FormatBool(variable.UserIsSsoEnabled),

		// data source variables
		resourcetype.Users,
		generatedName,
		resourceTypeAndName,
	)
}
