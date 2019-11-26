package dome9

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccDataSourceRuleSetBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.RuleSet)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRuleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRuleSetConfigure(resourceTypeAndName, generatedName, variable.RuleSetDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "name", resourceTypeAndName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "description", resourceTypeAndName, "description"),
					resource.TestCheckResourceAttr(dataSourceTypeAndName, "rules.#", "1"),
				),
			},
		},
	})
}
