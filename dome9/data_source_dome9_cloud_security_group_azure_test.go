package dome9

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccDataSourceCloudSecurityGroupAzureBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, resourceName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAzureSecurityGroup)
	azureTypeAndName, _, azureGeneratedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAzure)
	azureCloudAccountHCL := getCloudAccountAzureResourceHCL(azureGeneratedName, variable.CloudAccountAzureCreationResourceName, variable.CloudAccountAzureUpdateOperationMode)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAzureEnvVarsPreCheck(t)
			testAccAzureSecurityGroupEnvVarsPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAzureSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAzureSecurityGroupBasic(azureCloudAccountHCL, azureTypeAndName, resourceName, resourceTypeAndName, azureSecurityGroupAdditionalBlock()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "dome9_security_group_name", resourceTypeAndName, "dome9_security_group_name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "description", resourceTypeAndName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "region", resourceTypeAndName, "region"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "resource_group", resourceTypeAndName, "resource_group"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "tags.0.key", resourceTypeAndName, "tags.0.key"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "tags.0.value", resourceTypeAndName, "tags.0.value"),
				),
			},
		},
	})
}
