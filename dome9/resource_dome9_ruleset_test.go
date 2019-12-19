package dome9

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/dome9/dome9-sdk-go/services/rulebundles"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceRuleSetBasic(t *testing.T) {
	var ruleSet rulebundles.RuleBundleResponse
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.RuleSet)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRuleSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckRuleSetConfigure(resourceTypeAndName, generatedName, variable.RuleSetDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRuleSetExists(resourceTypeAndName, &ruleSet),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", generatedName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "description", variable.RuleSetDescription),
					resource.TestCheckResourceAttr(resourceTypeAndName, "rules.#", "1"),
				),
			},

			// Update test
			{
				Config: testAccCheckRuleSetConfigure(resourceTypeAndName, generatedName, variable.RuleSetDescriptionUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRuleSetExists(resourceTypeAndName, &ruleSet),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", generatedName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "description", variable.RuleSetDescriptionUpdate),
					resource.TestCheckResourceAttr(resourceTypeAndName, "rules.#", "1"),
				),
			},
		},
	})
}

func testAccCheckRuleSetDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.RuleSet {
			continue
		}

		ipList, _, err := apiClient.ruleSet.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if ipList != nil {
			return fmt.Errorf("rule set with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckRuleSetExists(resource string, ruleSet *rulebundles.RuleBundleResponse) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedRuleSet, _, err := apiClient.ruleSet.Get(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*ruleSet = *receivedRuleSet

		return nil
	}
}

func testAccCheckRuleSetConfigure(resourceTypeAndName, generatedName, description string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
	name               = "%s"
	description        = "%s"
	cloud_vendor       = "aws"
	language           = "en"
	hide_in_compliance = false
	is_template        = false
	rules {
		name           = "some_rule2"
		logic          = "EC2 should x"
		severity       = "High"
		description    = "rule description here"
		compliance_tag = "ct"
		domain         = "test"
		priority       = "high"
		control_title  = "ct"
		rule_id        = ""
		is_default     = false
	}
}

data "%s" "%s" {
	id = "${%s.id}"
}
`,
		// resource variables
		resourcetype.RuleSet,
		generatedName,
		generatedName,
		description,

		// data source variables
		resourcetype.RuleSet,
		generatedName,
		resourceTypeAndName,
	)
}
