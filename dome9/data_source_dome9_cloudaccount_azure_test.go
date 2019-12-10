package dome9

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccDataSourceCloudAccountAzureBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAzure)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAzureEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAccountAzureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAccountAzureConfigure(resourceTypeAndName, generatedName, variable.CloudAccountAzureCreationResourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "name", resourceTypeAndName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "operation_mode", resourceTypeAndName, "operation_mode"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "subscription_id", resourceTypeAndName, "subscription_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "tenant_id", resourceTypeAndName, "tenant_id"),
				),
			},
		},
	})
}
