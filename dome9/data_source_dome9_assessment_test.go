package dome9

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
)

func TestAccDataSourceAssessmentBasic(t *testing.T) {
	resourceTypeAndName, dataSourceTypeAndName, generatedName := method.GenerateRandomSourcesTypeAndName(resourcetype.Assessment)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAssessmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAssessmentConfigure(resourceTypeAndName, generatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "id", resourceTypeAndName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "request", resourceTypeAndName, "request"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "tests", resourceTypeAndName, "tests"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "test_entities", resourceTypeAndName, "test_entities"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "exclusions", resourceTypeAndName, "exclusions"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "remediations", resourceTypeAndName, "remediations"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "data_sync_status", resourceTypeAndName, "data_sync_status"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "created_time", resourceTypeAndName, "created_time"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "assessment_id", resourceTypeAndName, "assessment_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "triggered_by", resourceTypeAndName, "triggered_by"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "assessment_passed", resourceTypeAndName, "assessment_passed"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "has_errors", resourceTypeAndName, "has_errors"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "stats", resourceTypeAndName, "stats"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "has_data_sync_status_issues", resourceTypeAndName, "has_data_sync_status_issues"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "comparison_custom_id", resourceTypeAndName, "comparison_custom_id"),
					resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "additional_fields", resourceTypeAndName, "additional_fields"),
				),
			},
		},
	})
}
