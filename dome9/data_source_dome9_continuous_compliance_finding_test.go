package dome9

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
)

func TestAccDataSourceFindingSearchBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.Assessment)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAssessmentConfigure(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "page_size", resourceTypeAndName, "page_size"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "sorting", resourceTypeAndName, "sorting"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "multi_sorting", resourceTypeAndName, "multi_sorting"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "filter", resourceTypeAndName, "filter"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "search_after", resourceTypeAndName, "search_after"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "data_source", resourceTypeAndName, "data_source"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "search_request", resourceTypeAndName, "search_request"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "findings", resourceTypeAndName, "findings"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "total_findings_count", resourceTypeAndName, "total_findings_count"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "aggregations", resourceTypeAndName, "aggregations"),
				),
			},
		},
	})
}
