package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"testing"
)

func TestAccDataSourceServiceAccountBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ServiceAccount)

	roleTypeAndName, _, roleGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.Role)
	roleHCL := testAccCheckRoleConfigure(roleTypeAndName, roleGeneratedName, variable.RoleDescription, variable.RoleToPermittedAlertActions)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServiceAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckServiceAccountBasic(resourceTypeAndName, generatedName, "test", roleHCL, roleTypeAndName),
				Check: resource.ComposeTestCheckFunc( 
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "name", resourceTypeAndName, "name"),
					resource.TestCheckResourceAttr(dataSourceTypeAndName, "name", "test"),
				),
			},
		},
	})
}
