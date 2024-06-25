package dome9

import (
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceAzureOrganizationOnboarding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAzureOrganizationOnboardingRead,

		Schema: map[string]*schema.Schema{
			// Computed fields - OrganizationManagementViewModel object fields
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"organization_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_registration_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"active_blades": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"awp": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"onboarding_mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"centralized_subscription_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"with_function_apps_scan": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"serverless": {
							Type:     schema.TypeMap,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"cdr": {
							Type:     schema.TypeMap,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"accounts": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"storage_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"log_types": {
													Type:     schema.TypeList,
													Computed: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
								},
							},
						},
						"posture_management": {
							Type:     schema.TypeMap,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"onboarding_mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"vendor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_cloud_guard_managed_app": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_auto_onboarding": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"account_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_organization_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"onboarding_configuration": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"organization_root_ou_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mapping_strategy": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"posture_management": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rulesets_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeInt},
									},
									"onboarding_mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"user_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAzureOrganizationOnboardingRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	orgId := d.Get("id").(string)
	resp, _, err := d9Client.azureOrganizationOnboarding.Get(orgId)

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing Aws organization %s from state because it no longer exists in CloudGuard", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	d.SetId(resp.Id)
	_ = d.Set("account_id", resp.AccountId)
	_ = d.Set("user_id", resp.UserId)
	_ = d.Set("organization_name", resp.OrganizationName)
	_ = d.Set("tenant_id", resp.TenantId)
	_ = d.Set("management_group_id", resp.ManagementGroupId)
	_ = d.Set("app_registration_name", resp.AppRegistrationName)
	_ = d.Set("update_time", resp.UpdateTime)
	_ = d.Set("creation_time", resp.CreationTime)
	_ = d.Set("is_auto_onboarding", resp.IsAutoOnboarding)

	if err := d.Set("onboarding_configuration", flattenAwsOrganizationOnboardingConfiguration(resp.OnboardingConfiguration)); err != nil {
		return err
	}

	return nil
}
