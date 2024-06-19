package dome9

import (
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
				Required: true,
			},
			"management_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"organization_name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AzureOrg",
			},
			"app_registration_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_secret": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"active_blades": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"awp": {
							Type:     schema.TypeMap,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"onboarding_mode": {
										Type:         schema.TypeString,
										Optional:     true,
										Default:      "inAccountHub",
										ValidateFunc: validation.StringInSlice([]string{"saas", "inAccount", "inAccountHub"}, false),
									},
									"centralized_subscription_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"with_function_apps_scan": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
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
										Optional: true,
										Default:  false,
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
										Optional: true,
										Default:  true,
									},
									"accounts": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"storage_id": {
													Type:     schema.TypeString,
													Required: true,
												},
												"log_types": {
													Type:     schema.TypeList,
													Required: true,
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
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"Read", "Manage"}, false),
									},
								},
							},
						},
					},
				},
			},
			"vendor": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"azure", "azuregov", "azurechina"}, false),
				Default:      "azure",
			},
			"use_cloud_guard_managed_app": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"is_auto_onboarding": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"account_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_organization_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_management_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_account_stack_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_account_stack_region": {
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
	resp, _, err := d9Client.awsOrganizationOnboarding.Get(orgId)

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing Aws organization %s from state because it no longer exists in CloudGuard", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	d.SetId(resp.Id)
	_ = d.Set("accountId", resp.AccountId)
	_ = d.Set("externalOrganizationId", resp.ExternalOrganizationId)
	_ = d.Set("externalManagementAccountId", resp.ExternalManagementAccountId)
	_ = d.Set("managementAccountStackId", resp.ManagementAccountStackId)
	_ = d.Set("managementAccountStackRegion", resp.ManagementAccountStackRegion)
	_ = d.Set("userId", resp.UserId)
	_ = d.Set("organizationName", resp.OrganizationName)
	_ = d.Set("updateTime", resp.UpdateTime)
	_ = d.Set("creationTime", resp.CreationTime)
	_ = d.Set("stackSetRegions", resp.StackSetRegions)
	_ = d.Set("stackSetOrganizationalUnitIds", resp.StackSetOrganizationalUnitIds)

	if err := d.Set("onboarding_configuration", flattenAwsOrganizationOnboardingConfiguration(resp.OnboardingConfiguration)); err != nil {
		return err
	}

	return nil
}
