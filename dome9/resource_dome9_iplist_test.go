package dome9

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/dome9/dome9-sdk-go/services/iplist"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func TestAccResourceIPListBasic(t *testing.T) {
	var ipList iplist.IpList
	resourceTypeAndName, _, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.IPList)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIPListConfigure(resourceTypeAndName, generatedName, variable.IPListDescriptionResource),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPListExists(resourceTypeAndName, &ipList),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.IPListCreationResourceName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "description", variable.IPListDescriptionResource),
					resource.TestCheckResourceAttr(resourceTypeAndName, "items.#", "1"),
				),
			},

			// Update test
			{
				Config: testAccCheckIPListConfigure(resourceTypeAndName, generatedName, variable.IPListUpdateDescriptionResource),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPListExists(resourceTypeAndName, &ipList),
					resource.TestCheckResourceAttr(resourceTypeAndName, "name", variable.IPListCreationResourceName),
					resource.TestCheckResourceAttr(resourceTypeAndName, "description", variable.IPListUpdateDescriptionResource),
					resource.TestCheckResourceAttr(resourceTypeAndName, "items.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIPListDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourcetype.IPList {
			continue
		}

		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			log.Println("Failed in conversion with error:", err)
			return err
		}
		ipList, _, err := apiClient.iplist.Get(id)

		if err == nil {
			return fmt.Errorf("id %d already exists", id)
		}

		if ipList != nil {
			return fmt.Errorf("iplist with id %d exists and wasn't destroyed", id)
		}
	}

	return nil
}

func testAccCheckIPListExists(resource string, ipList *iplist.IpList) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("didn't find resource: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			log.Println("Failed in conversion with error:", err)
			return err
		}

		apiClient := testAccProvider.Meta().(*Client)
		receivedIpList, _, err := apiClient.iplist.Get(id)

		if err != nil {
			return fmt.Errorf("failed fetching resource %s. Recevied error: %s", resource, err)
		}
		*ipList = *receivedIpList

		return nil
	}
}

func testAccCheckIPListConfigure(resourceTypeAndName, generatedName, description string) string {
	return fmt.Sprintf(`
resource "%s" "%s" {
  name        = "%s"
  description = "%s"

  items {
      ip      = "1.1.4.4/32"
      comment = "Unused ip, just for testing"
    }
}
data "%s" "%s" {
  id = "${%s.id}"
}
`,
		// resource variables
		resourcetype.IPList,
		generatedName,
		variable.IPListCreationResourceName,
		description,

		// data source variables
		resourcetype.IPList,
		generatedName,
		resourceTypeAndName,
	)
}
