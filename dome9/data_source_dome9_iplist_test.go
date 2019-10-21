package dome9

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccDataSourceIpListBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.IPList)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIPListConfigure(resourceTypeAndName, generatedName, variable.IPListDescriptionResource),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "name", resourceTypeAndName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "description", resourceTypeAndName, "description"),
					resource.TestCheckResourceAttr(dataSourceTypeAndName, "items.#", "1"),
				),
			},
		},
	})
}
