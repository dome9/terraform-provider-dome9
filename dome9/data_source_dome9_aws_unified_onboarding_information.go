package dome9

import (
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	CloudAccountId            = "cloud_account_id"
	InitiatedUserName         = "initiated_user_name"
	InitiatedUserId           = "initiated_user_id"
	EnvironmentId             = "environment_id"
	EnvironmentName           = "environment_name"
	EnvironmentExternalId     = "environment_external_id"
	RootStackId               = "root_stack_id"
	CftVersion                = "cft_version"
	UnifiedOnboardingRequest  = "onbording_request"
	Statuses                  = "statuses"
	Module                    = "module"
	Feature                   = "feature"
	Status                    = "status"
	StatusMessage             = "status_message"
	StackStatus               = "stack_status"
	StackMessage              = "stack_message"
	remediationRecommendation = "remediation_recommendation"
)

func dataSourceAwsUnifiedOnboardingInformation() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsUnifiedOnboardingReadInfo,
		Schema: map[string]*schema.Schema{
			CloudAccountId: {
				Type:     schema.TypeString,
				Required: true,
			},
			OnboardingId: {
				Type:     schema.TypeString,
				Computed: true,
			},
			InitiatedUserName: {
				Type:     schema.TypeString,
				Computed: true,
			},
			InitiatedUserId: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			EnvironmentId: {
				Type:     schema.TypeString,
				Computed: true,
			},
			EnvironmentExternalId: {
				Type:     schema.TypeString,
				Computed: true,
			},
			RootStackId: {
				Type:     schema.TypeString,
				Computed: true,
			},
			CftVersion: {
				Type:     schema.TypeString,
				Computed: true,
			},
			EnvironmentName: {
				Type:     schema.TypeString,
				Computed: true,
			},
			UnifiedOnboardingRequest: {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"onboard_type": {
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"full_protection": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"cloud_vendor": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"enable_stack_modify": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"posture_management_configuration": {
						Type:     schema.TypeMap,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								Rulesets: {
									Type:     schema.TypeList,
									Required: true,
								},
							},
						},
					},
					"serverless_configuration": {
						Type:     schema.TypeMap,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								Enabled: {
									Type:     schema.TypeBool,
									Required: true,
								},
							},
						},
					},
					"intelligence_configurations": {
						Type:     schema.TypeMap,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								Rulesets: {
									Type:     schema.TypeList,
									Required: false,
								},
								Enabled: {
									Type:     schema.TypeBool,
									Required: false,
								},
							},
						},
					},
				}}},
			Statuses: {
				Type:     schema.TypeString,
				Computed: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					Module: {
						Type:     schema.TypeString,
						Computed: true,
						ForceNew: true,
					},
					Feature: {
						Type:     schema.TypeBool,
						Computed: true,
					},
					Status: {
						Type:     schema.TypeString,
						Computed: true,
					},
					StatusMessage: {
						Type:     schema.TypeString,
						Computed: true,
					},
					StackStatus: {
						Type:     schema.TypeString,
						Computed: true,
					},
					StackMessage: {
						Type:     schema.TypeString,
						Computed: true,
					},
					remediationRecommendation: {
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
	resp, _, err := d9Client.AwsUnifiedOnbording.Get(d.Get("CloudAccountId").(string))
	if err != nil {
		return err
	}

	log.Printf("[INFO] Get UnifiedOnbording Information with OnbordingId: %s\n", resp.OnboardingId)

	return unifiedOnbordingReadInfo(d, meta)
}

func unifiedOnbordingReadInfo(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.AwsUnifiedOnbording.Get(d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing rule set %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	log.Printf("[INFO] Getting Unified Onbording:\n%+v\n", resp)
	_ = d.Set(OnboardingId, resp.OnboardingId)
	_ = d.Set(InitiatedUserName, resp.InitiatedUserName)
	_ = d.Set(InitiatedUserId, resp.InitiatedUserId)
	_ = d.Set(EnvironmentId, resp.EnvironmentId)
	_ = d.Set(EnvironmentName, resp.EnvironmentName)
	_ = d.Set(EnvironmentExternalId, resp.EnvironmentExternalId)
	_ = d.Set(RootStackId, resp.RootStackId)
	_ = d.Set(CftVersion, resp.CftVersion)
	_ = d.Set(UnifiedOnboardingRequest, resp.UnifiedOnbordingRequest)
	_ = d.Set(Status, resp.Statuses)

	return nil
}
