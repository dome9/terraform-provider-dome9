package dome9

//import (
//	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
//	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
//	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
//	"testing"
//)
//
//func TestAccDataSourceNotificationBasic(t *testing.T) {
//	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.Notification)
//
//	resource.Test(t, resource.TestCase{
//		PreCheck: func() {
//			testAccPreCheck(t)
//		},
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckNotificationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccCheckNotificationConfigure(resourceTypeAndName, generatedName),
//				Check: resource.ComposeTestCheckFunc(
//					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
//					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "name", resourceTypeAndName, "name"),
//					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "description", resourceTypeAndName, "description"),
//					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "alerts_console", resourceTypeAndName, "alerts_console"),
//					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "send_on_each_occurrence", resourceTypeAndName, "send_on_each_occurrence"),
//					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "origin", resourceTypeAndName, "origin"),
//				),
//			},
//		},
//	})
//}
