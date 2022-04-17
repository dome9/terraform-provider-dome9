package dome9

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/method"
	"testing"
)

func TestAccDataSourceAWSUnifiedOnboardingBasic(t *testing.T) {
	resourceTypeAndName, _, resourceName := method.GenerateRandomSourcesTypeAndName(resourcetype.AwsUnifiedOnboarding)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		//CheckDestroy: testAccCheckAWSUnifiedOnboardingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsUnifiedOnbordingBasic(resourceTypeAndName, resourceName),
				Check: resource.ComposeTestCheckFunc(
					//resource.TestCheckResourceAttrPair(resourceTypeAndName+"Data", "cloud_vendor", resourceTypeAndName, "cloud_vendor"),
					//resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "cloud_vendor", resourceTypeAndName, "cloud_vendor"),
					//resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "enable_stack_modify", resourceTypeAndName, "enable_stack_modify"),
					//resource.TestCheckResourceAttrPair(dataSourceTypeAndName, "full_protection", resourceTypeAndName, "full_protection"),
				),
			},
		},
	})
}

func testAccCheckAWSUnifiedOnboardingDestroy(state *terraform.State) error {
	apiClient := testAccProvider.Meta().(*Client)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != resourcetype.AwsUnifiedOnboarding {
			continue
		}

		response, _, err := apiClient.awsUnifiedOnboarding.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("id %s already exists", rs.Primary.ID)
		}

		if response != nil {
			return fmt.Errorf("cloudaccount with id %s exists and wasn't destroyed", rs.Primary.ID)
		}
	}

	return nil
}
//
//func dataSourceAwsUnifiedOnboarding() *schema.Resource {
//	return &schema.Resource{
//		Read: dataSourceAwsUnifiedOnboardingReadInfo,
//		Schema: map[string]*schema.Schema{
//			providerconst.Id: {
//				Type:     schema.TypeString,
//				Required: true,
//			},
//			providerconst.OnboardingId: {
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//			providerconst.InitiatedUserName: {
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//			providerconst.InitiatedUserId: {
//				Type:     schema.TypeInt,
//				Computed: true,
//			},
//			providerconst.EnvironmentId: {
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//			providerconst.EnvironmentExternalId: {
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//			providerconst.RootStackId: {
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//			providerconst.CftVersion: {
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//			providerconst.EnvironmentName: {
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//			providerconst.UnifiedOnboardingRequest: {
//				Type:     schema.TypeMap,
//				Computed: true,
//				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
//					providerconst.OnboardType: {
//						Type:     schema.TypeString,
//						Optional: true,
//					},
//					providerconst.FullProtection: {
//						Type:     schema.TypeBool,
//						Optional: true,
//					},
//					providerconst.CloudVendor: {
//						Type:     schema.TypeString,
//						Optional: true,
//					},
//					providerconst.EnableStackModify: {
//						Type:     schema.TypeBool,
//						Optional: true,
//					},
//					providerconst.PostureManagementConfiguration: {
//						Type:     schema.TypeMap,
//						Optional: true,
//						Elem: &schema.Resource{
//							Schema: map[string]*schema.Schema{
//								providerconst.Rulesets: {
//									Type:     schema.TypeList,
//									Required: true,
//									Elem:     &schema.Schema{Type: schema.TypeInt},
//								},
//							},
//						},
//					},
//					providerconst.ServerlessConfiguration: {
//						Type:     schema.TypeMap,
//						Optional: true,
//						Elem: &schema.Resource{
//							Schema: map[string]*schema.Schema{
//								providerconst.Enabled: {
//									Type:     schema.TypeBool,
//									Required: true,
//								},
//							},
//						},
//					},
//					providerconst.IntelligenceConfigurations: {
//						Type:     schema.TypeMap,
//						Optional: true,
//						Elem: &schema.Resource{
//							Schema: map[string]*schema.Schema{
//								providerconst.Rulesets: {
//									Type:     schema.TypeList,
//									Required: false,
//									Elem:     &schema.Schema{Type: schema.TypeInt},
//								},
//								providerconst.Enabled: {
//									Type:     schema.TypeBool,
//									Required: false,
//								},
//							},
//						},
//					},
//				}}},
//			providerconst.Statuses: {
//				Type:     schema.TypeList,
//				Computed: true,
//				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
//					providerconst.Module: {
//						Type:     schema.TypeString,
//						Computed: true,
//					},
//					providerconst.Feature: {
//						Type:     schema.TypeString,
//						Computed: true,
//					},
//					providerconst.Status: {
//						Type:     schema.TypeString,
//						Computed: true,
//					},
//					providerconst.StatusMessage: {
//						Type:     schema.TypeString,
//						Computed: true,
//					},
//					providerconst.StackStatus: {
//						Type:     schema.TypeString,
//						Computed: true,
//					},
//					providerconst.StackMessage: {
//						Type:     schema.TypeString,
//						Computed: true,
//					},
//					providerconst.RemediationRecommendation: {
//						Type:     schema.TypeString,
//						Computed: true,
//					},
//				}},
//			},
//		},
//	}
//}
