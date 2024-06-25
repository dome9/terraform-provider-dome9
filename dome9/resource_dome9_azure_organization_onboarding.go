package dome9

import (
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws_org"
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
			// Computed fields - OrganizationManagementViewModel object fields
			"account_id": {
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
	_ = d.Set("update_time", resp.UpdateTime)
	_ = d.Set("creation_time", resp.CreationTime)
	_ = d.Set("is_auto_onboarding", resp.IsAutoOnboarding)

	if err := d.Set("onboarding_configuration", flattenAzureOrganizationOnboardingConfiguration(resp.OnboardingConfiguration)); err != nil {
		return err
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

func resourceAzureOrganizationOnboardingUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Println("An update occurred")

	organizationName := "organization_name"
	if d.HasChange(organizationName) {
		log.Println("The Management Group ID has been changed")

		if resp, err := d9Client.azureOrganizationOnboarding.UpdateOrganizationManagementAsync(d.Id(), azure_org.OnboardingUpdateRequest{
			OrganizationName: d.Get(organizationName).(string),
		}); err != nil {
			return err
		} else {
			log.Printf("resourceAzureOrganizationOnboardingUpdate organization_name response is: %+v\n", resp)
		}
	}

	return nil
}

func expandAzureOrganizationOnboardingRequest(d *schema.ResourceData) azure_org.OnboardingRequest {
	req := azure_org.OnboardingRequest{
		TenantId:            d.Get("tenant_id").(string),
		ManagementGroupId:   d.Get("management_group_id").(string),
		OrganizationName:    d.Get("organization_name").(string),
		AppRegistrationName: d.Get("app_registration_name").(string),
		ClientId:            d.Get("client_id").(string),
		ClientSecret:        d.Get("client_secret").(string),
		ActiveBlades: azure_org.Blades{
			Awp: azure_org.AwpConfiguration{
				BladeConfiguration: azure_org.BladeConfiguration{
					IsEnabled: d.Get("active_blades.serverless.is_enabled").(bool),
				},
				OnboardingMode:            azure_org.AwpOnboardingMode(d.Get("active_blades.awp.onboarding_mode").(string)),
				CentralizedSubscriptionId: d.Get("active_blades.awp.centralized_subscription_id").(string),
				WithFunctionAppsScan:      d.Get("active_blades.awp.with_function_apps_scan").(bool),
			},
			Serverless: azure_org.ServerlessConfiguration{
				BladeConfiguration: azure_org.BladeConfiguration{
					IsEnabled: d.Get("active_blades.serverless.is_enabled").(bool),
				},
			},
			Cdr: azure_org.CdrConfiguration{
				BladeConfiguration: azure_org.BladeConfiguration{
					IsEnabled: d.Get("active_blades.serverless.is_enabled").(bool),
				},
				Accounts: expandCdrAccounts(d.Get("active_blades.cdr.accounts").([]interface{})),
			},
			PostureManagement: azure_org.PostureManagement{
				OnboardingMode: aws_org.OnboardingMode(d.Get("active_blades.posture_management.onboarding_mode").(string)),
			},
		},
		Vendor:                  azure_org.CloudVendor(d.Get("vendor").(string)),
		UseCloudGuardManagedApp: d.Get("use_cloud_guard_managed_app").(bool),
		IsAutoOnboarding:        d.Get("is_auto_onboarding").(bool),
	}
	return req
}

func expandCdrAccounts(accounts []interface{}) []azure_org.StorageAccount {
	result := make([]azure_org.StorageAccount, len(accounts))
	for i, account := range accounts {
		data := account.(map[string]interface{})
		result[i] = azure_org.StorageAccount{
			StorageId: data["storage_id"].(string),
			LogTypes:  expandStringList(data["log_types"].([]interface{})),
		}
	}
	return result
}

func expandStringList(input []interface{}) []string {
	result := make([]string, len(input))
	for i, v := range input {
		result[i] = v.(string)
	}
	return result
}

func flattenAzureOrganizationOnboardingConfiguration(config azure_org.AzureOrganizationOnboardingConfiguration) map[string]interface{} {
	return map[string]interface{}{
		"organization_root_ou_id": config.OrganizationRootOuId,
		"mapping_strategy":        config.MappingStrategy,
		"posture_management":      flattenPostureManagementConfiguration(config.PostureManagement),
	}
}
