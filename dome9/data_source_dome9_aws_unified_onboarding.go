package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/unifiedonboarding/aws_unified_onboarding"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
)

func dataSourceAwsUnifiedOnboarding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsUnifiedOnboardingReadInfo,
		Schema: map[string]*schema.Schema{
			providerconst.Id: {
				Type:     schema.TypeString,
				Required: true,
			},
			providerconst.OnboardingId: {
				Type:     schema.TypeString,
				Computed: true,
			},
			providerconst.InitiatedUserName: {
				Type:     schema.TypeString,
				Computed: true,
			},
			providerconst.InitiatedUserId: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			providerconst.EnvironmentId: {
				Type:     schema.TypeString,
				Computed: true,
			},
			providerconst.EnvironmentExternalId: {
				Type:     schema.TypeString,
				Computed: true,
			},
			providerconst.RootStackId: {
				Type:     schema.TypeString,
				Computed: true,
			},
			providerconst.CftVersion: {
				Type:     schema.TypeString,
				Computed: true,
			},
			providerconst.EnvironmentName: {
				Type:     schema.TypeString,
				Computed: true,
			},
			providerconst.UnifiedOnboardingRequest: {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					providerconst.OnboardType: {
						Type:     schema.TypeString,
						Optional: true,
					},
					providerconst.FullProtection: {
						Type:     schema.TypeBool,
						Optional: true,
					},
					providerconst.CloudVendor: {
						Type:     schema.TypeString,
						Optional: true,
					},
					providerconst.EnableStackModify: {
						Type:     schema.TypeBool,
						Optional: true,
					},
					providerconst.PostureManagementConfiguration: {
						Type:     schema.TypeMap,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								providerconst.Rulesets: {
									Type:     schema.TypeList,
									Required: true,
									Elem:     &schema.Schema{Type: schema.TypeInt},
								},
							},
						},
					},
					providerconst.ServerlessConfiguration: {
						Type:     schema.TypeMap,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								providerconst.Enabled: {
									Type:     schema.TypeBool,
									Required: true,
								},
							},
						},
					},
					providerconst.IntelligenceConfigurations: {
						Type:     schema.TypeMap,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								providerconst.Rulesets: {
									Type:     schema.TypeList,
									Required: false,
									Elem:     &schema.Schema{Type: schema.TypeInt},
								},
								providerconst.Enabled: {
									Type:     schema.TypeBool,
									Required: false,
								},
							},
						},
					},
				}}},
			providerconst.Statuses: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					providerconst.Module: {
						Type:     schema.TypeString,
						Computed: true,
					},
					providerconst.Feature: {
						Type:     schema.TypeString,
						Computed: true,
					},
					providerconst.Status: {
						Type:     schema.TypeString,
						Computed: true,
					},
					providerconst.StatusMessage: {
						Type:     schema.TypeString,
						Computed: true,
					},
					providerconst.StackStatus: {
						Type:     schema.TypeString,
						Computed: true,
					},
					providerconst.StackMessage: {
						Type:     schema.TypeString,
						Computed: true,
					},
					providerconst.RemediationRecommendation: {
						Type:     schema.TypeString,
						Computed: true,
					},
				}},
			},
		},
	}
}

func dataSourceAwsUnifiedOnboardingReadInfo(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.awsUnifiedOnboarding.Get(d.Get(providerconst.Id).(string))
	if err != nil {
		return err
	}

	d.SetId(resp.OnboardingId)
	_ = d.Set(providerconst.OnboardingId, resp.OnboardingId)
	_ = d.Set(providerconst.InitiatedUserName, resp.InitiatedUserName)
	_ = d.Set(providerconst.EnvironmentName, resp.EnvironmentName)
	_ = d.Set(providerconst.EnvironmentExternalId, resp.EnvironmentExternalId)
	_ = d.Set(providerconst.RootStackId, resp.RootStackId)
	_ = d.Set(providerconst.CftVersion, resp.CftVersion)
	_ = d.Set(providerconst.UnifiedOnboardingRequest, resp.UnifiedOnboardingRequest)
	_ = d.Set(providerconst.Statuses, expendStatuses(resp.Statuses))
	_ = d.Set(providerconst.EnvironmentId, resp.EnvironmentId)
	_ = d.Set(providerconst.InitiatedUserId, resp.InitiatedUserId)

	return nil
}

func expendStatuses(statuses aws_unified_onboarding.Statuses) interface{} {
	statusesList := make([]interface{}, len(statuses))

	for i, statusItem := range statuses {
		statusesList[i] = map[string]interface{}{
			providerconst.Module:                    statusItem.Module,
			providerconst.Feature:                   statusItem.Feature,
			providerconst.Status:                    statusItem.Status,
			providerconst.StatusMessage:             statusItem.StackMessage,
			providerconst.StackStatus:               statusItem.StackStatus,
			providerconst.StackMessage:              statusItem.StackMessage,
			providerconst.RemediationRecommendation: statusItem.RemediationRecommendation,
		}
	}
	return statusesList
}
