package dome9

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"testing"
)

func TestAccResourceFindingSearchBasic(t *testing.T) {
	_, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousComplianceFinding)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFindingSearchDestroy,
		Steps: []resource.TestStep{
			{
				// creation test
				Config: testAccCheckFindingSearchConfigure(generatedName),
				Check: resource.ComposeTestCheckFunc(
				),
			},
		},
	})
}

func testAccCheckFindingSearchDestroy(s *terraform.State) error {
	return nil
}

func testAccCheckFindingSearchConfigure(generatedName string) string {
	return fmt.Sprintf(`
data "%s" "%s" {

}
`,

		// data source variables
		resourcetype.ContinuousComplianceFinding,
		generatedName,
	)
}

