package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/serviceaccounts"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"testing"
)

func TestAccResourceServiceAccountBasic(t *testing.T) {
	var serviceAccount serviceaccounts.GetServiceAccountResponse
	serviceAccountTypeAndName, _, serviceAccountGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ServiceAccount)

	roleTypeAndName, _, roleGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.Role)
	roleHCL := testAccCheckRoleConfigure(roleTypeAndName, roleGeneratedName, variable.RoleDescription, variable.RoleToPermittedAlertActions)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServiceAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckServiceAccountBasic(serviceAccountTypeAndName, serviceAccountGeneratedName, variable.ServiceAccountName, roleHCL, roleTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceAccountExists(serviceAccountTypeAndName, &serviceAccount),
					resource.TestCheckResourceAttr(serviceAccountTypeAndName, "name", variable.ServiceAccountName),
				),
			},
			{
				// update name test
				Config: testAccCheckServiceAccountBasic(serviceAccountTypeAndName, serviceAccountGeneratedName, variable.ServiceAccountNameUpdate, roleHCL, roleTypeAndName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceAccountExists(serviceAccountTypeAndName, &serviceAccount),
					resource.TestCheckResourceAttr(serviceAccountTypeAndName, "name", variable.ServiceAccountNameUpdate),
				),
			},
		},
	})
}

func testAccCheckServiceAccountExists(resource string, serviceAccount *serviceaccounts.GetServiceAccountResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedServiceAccountResponse, _, err := apiClient.serviceAccounts.Get(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*serviceAccount = *receivedServiceAccountResponse

		return nil
	}
}

func testAccCheckServiceAccountDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.ServiceAccount {
			continue
		}

		receivedServiceAccountResponse, _, err := apiClient.serviceAccounts.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if receivedServiceAccountResponse != nil {
			return fmt.Errorf("service account with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckServiceAccountBasic(resourceTypeAndName, generatedName, name, roleHCL, roleTypeAndName string) string {
	return fmt.Sprintf(`
// role resource
%s

// service account resource
%s

data "%s" "%s" {
  id = "${%s.id}"
}
`,
		roleHCL,
		// service account resource
		getServiceAccountResourceHCL(generatedName, name, roleTypeAndName),

		// data source variables
		resourcetype.ServiceAccount,
		generatedName,
		resourceTypeAndName,
	)
}

func getServiceAccountResourceHCL(generatedName, name, roleTypeAndName string) string {
	return fmt.Sprintf(`
// service account creation
resource "%s" "%s" {
  name			= "%s"
  role_ids 	   	= ["${%s.id}"]
}
`,
		// resource variables
		resourcetype.ServiceAccount,
		generatedName,
		name,
		roleTypeAndName,
	)
}
