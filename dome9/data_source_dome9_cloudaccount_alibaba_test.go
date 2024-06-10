package dome9

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccDataSourceCloudAccountAlibabaBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAlibaba)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAlibabaEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAccountAlibabaDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAccountAlibabaConfigure(resourceTypeAndName, generatedName, variable.CloudAccountAlibabaCreationResourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "name", resourceTypeAndName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "alibaba_account_id", resourceTypeAndName, "alibaba_account_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "organizational_unit_id", resourceTypeAndName, "organizational_unit_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "credentials", resourceTypeAndName, "credentials"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "organizational_unit_path", resourceTypeAndName, "organizational_unit_path"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "organizational_unit_name", resourceTypeAndName, "organizational_unit_name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "vendor", resourceTypeAndName, "vendor"),
				),
			},
		},
	})
}
