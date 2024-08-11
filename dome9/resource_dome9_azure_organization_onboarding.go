package dome9

import (
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/azure_org"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
)

func resourceAzureOrganizationOnboarding() *schema.Resource {
	return &schema.Resource{
		Create: resourceAzureOrganizationOnboardingCreate,
		Read:   resourceAzureOrganizationOnboardingRead,
		Update: resourceAzureOrganizationOnboardingUpdate,
		Delete: resourceAzureOrganizationOnboardingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// OnboardingRequest object fields
			"workflow_id": {
				Type:     schema.TypeString,
				Optional: true,
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
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"awp": {
							Type:     schema.TypeMap,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_enabled": {
										Type:    schema.TypeBool,
										Default: false, //is it right for default value?
									},
									"onboarding_mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"centralized_subscription_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"with_function_apps_scan": {
										Type:    schema.TypeBool,
										Default: false, //is it right for default value?
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
										Type:    schema.TypeBool,
										Default: false, //is it right for default value?
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
										Type:    schema.TypeBool,
										Default: false, //is it right for default value?
									},
									"accounts": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"storage_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"log_types": {
													Type:     schema.TypeList,
													Computed: true,
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
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"azure",
					"azurechina",
					"azuregov",
				}, false),
			},
			"use_cloud_guard_managed_app": {
				Type:    schema.TypeBool,
				Default: false, //is it right for default value?
			},
			"is_auto_onboarding": {
				Type:    schema.TypeBool,
				Default: false, //is it right for default value?
			},
			// OrganizationManagementViewModel object fields
			//is it needed here? required?
			"id": {
				Type:     schema.TypeString,
				Required: true, //is it right to set it to required here?
			},
			"account_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeInt,
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
						"awp_configuration": {
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
										Optional: true,
									},
									"with_function_apps_scan": {
										Type:    schema.TypeBool,
										Default: false, //is it right for default value?
									},
								},
							},
						},
						"serverless_configuration": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"cdr_configuration": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"accounts": {
										Type:     schema.TypeList,
										Computed: true,
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
					},
				},
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

func resourceAzureOrganizationOnboardingCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandAzureOrganizationOnboardingRequest(d)
	log.Printf("[INFO] Creating Azure organization with request %+v\n", req)

	resp, _, err := d9Client.azureOrganizationOnboarding.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created Azure organization. ID: %v\n", resp.Id)
	d.SetId(resp.Id)

	return resourceAzureOrganizationOnboardingRead(d, meta)
}

func resourceAzureOrganizationOnboardingRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.azureOrganizationOnboarding.Get(d.Id())

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing Azure organization %s from state because it no longer exists in CloudGuard", d.Id())
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
	_ = d.Set("onboarding_configuration", flattenAzureOrganizationOnboardingConfiguration(resp.OnboardingConfiguration))
	_ = d.Set("is_auto_onboarding", resp.IsAutoOnboarding)
	_ = d.Set("update_time", resp.UpdateTime)
	_ = d.Set("creation_time", resp.CreationTime)

	//other error checks?

	return nil
}

func resourceAzureOrganizationOnboardingUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Println("An update occurred")

	if d.HasChange("organization_name") {
		log.Println("The configuration has been changed")

		updateConfigReq := azure_org.OnboardingUpdateRequest{
			OrganizationName: d.Get("organization_name").(string),
		}

		if resp, err := d9Client.azureOrganizationOnboarding.UpdateOrganizationManagementAsync(d.Id(), updateConfigReq); err != nil {
			return err
		} else {
			log.Printf("resourceAzureOrganizationOnboardingUpdate Configuration response is: %+v\n", resp)
		}
	}

	return nil
}

func resourceAzureOrganizationOnboardingDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting Azure organization ID: %v\n", d.Id())
	if _, err := d9Client.azureOrganizationOnboarding.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func expandAzureOrganizationOnboardingRequest(d *schema.ResourceData) azure_org.OnboardingRequest {
	req := azure_org.OnboardingRequest{
		WorkflowId:              d.Get("workflow_id").(string),
		TenantId:                d.Get("tenant_id").(string),
		ManagementGroupId:       d.Get("management_group_id").(string),
		OrganizationName:        d.Get("organization_name").(string),
		AppRegistrationName:     d.Get("app_registration_name").(string),
		ClientId:                d.Get("client_id").(string),
		ClientSecret:            d.Get("client_secret").(string),
		UseCloudGuardManagedApp: d.Get("use_cloud_guard_managed_app").(bool),
		IsAutoOnboarding:        d.Get("is_auto_onboarding").(bool),
		ActiveBlades: azure_org.Blades{
			Awp: azure_org.AwpConfiguration{
				BladeConfiguration: azure_org.BladeConfiguration{
					IsEnabled: d.Get("active_blades.0.awp.is_enabled").(bool),
				},
				OnboardingMode:            azure_org.AwpOnboardingMode(d.Get("active_blades.0.awp.onboarding_mode").(string)),
				CentralizedSubscriptionId: d.Get("active_blades.0.awp.centralized_subscription_id").(string),
				WithFunctionAppsScan:      d.Get("active_blades.0.awp.with_function_apps_scan").(bool),
			},
			Serverless: azure_org.ServerlessConfiguration{
				BladeConfiguration: azure_org.BladeConfiguration{
					IsEnabled: d.Get("active_blades.0.serverless.is_enabled").(bool),
				},
			},
			Cdr: azure_org.CdrConfiguration{
				BladeConfiguration: azure_org.BladeConfiguration{
					IsEnabled: d.Get("active_blades.0.cdr.is_enabled").(bool),
				},
				Accounts: []azure_org.StorageAccount{
					{
						StorageId: d.Get("active_blades.0.cdr.accounts.0.storage_id").(string),
					},
				},
			},
			PostureManagement: azure_org.PostureManagement{
				OnboardingMode: azure_org.OnboardingMode(d.Get("active_blades.0.posture_management.onboarding_mode").(string)),
			},
		},
		Vendor: azure_org.CloudVendor(d.Get("vendor").(string)),
	}

	return req
}
