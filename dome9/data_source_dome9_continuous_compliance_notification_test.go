package dome9

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
)

func TestAccDataSourceContinuousComplianceNotificationBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.ContinuousComplianceNotification)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContinuousComplianceNotificationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckContinuousComplianceNotificationBasic(resourceTypeAndName, generatedName, continuousComplianceNotificationConfig(generatedName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "name", resourceTypeAndName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "description", resourceTypeAndName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "alerts_console", resourceTypeAndName, "alerts_console"),
					resource.TestCheckResourceAttr(dataSourceTypeAndName, "scheduled_report.#", "1"),
					resource.TestCheckResourceAttr(dataSourceTypeAndName, "change_detection.#", "1"),
				),
			},
		},
	})
}
