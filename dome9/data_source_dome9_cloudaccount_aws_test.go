package dome9

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/dome9/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/environmentvariable"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccDataSourceCloudAccountAWSBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.CloudAccountAWS)
	originalArn := os.Getenv(environmentvariable.CloudAccountAWSEnvVarArn)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccCloudAccountAWSPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAccountAWSBasic(resourceTypeAndName, generatedName, variable.CloudAccountAWSCreationResourceName, originalArn, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "name", resourceTypeAndName, "name"),
					resource.TestCheckResourceAttr(dataSourceTypeAndName, "vendor", variable.CloudAccountAWSVendor),
					resource.TestCheckResourceAttr(dataSourceTypeAndName, "net_sec.#", "1"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "net_sec.0.regions.0.region", variable.CloudAccountAWSFetchedRegion),
				),
			},
		},
	})
}
