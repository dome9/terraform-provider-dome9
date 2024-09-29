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
				Computed: true,
			},
			"app_registration_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"awp": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"onboarding_mode": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "inAccount",
										ValidateFunc: validation.StringInSlice([]string{
											"inAccount",
											"inAccountHub",
											"saas",
										}, false),
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
									"with_sse_cmk_encrypted_disks_scan": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
								},
							},
						},
						"serverless": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
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
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"accounts": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"storage_id": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"log_types": {
													Type:     schema.TypeList,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
								},
							},
						},
						"posture_management": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"onboarding_mode": {
										Type:     schema.TypeString,
										Optional: true,
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
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"with_sse_cmk_encrypted_disks_scan": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
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
	_ = d.Set("is_auto_onboarding", resp.IsAutoOnboarding)
	_ = d.Set("update_time", resp.UpdateTime)
	_ = d.Set("creation_time", resp.CreationTime)

	if err := d.Set("onboarding_configuration", flattenAzureOrganizationOnboardingConfiguration(resp.OnboardingConfiguration)); err != nil {
		return err
	}

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

	activeBladesList := d.Get("active_blades").([]interface{})
	activeBlades := activeBladesList[0].(map[string]interface{})

	awp := activeBlades["awp"].([]interface{})[0].(map[string]interface{})
	serverless := activeBlades["serverless"].([]interface{})[0].(map[string]interface{})
	cdr := activeBlades["cdr"].([]interface{})[0].(map[string]interface{})
	postureManagement := activeBlades["posture_management"].([]interface{})[0].(map[string]interface{})

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
					IsEnabled: awp["is_enabled"].(bool),
				},
				OnboardingMode:               azure_org.AwpOnboardingMode(awp["onboarding_mode"].(string)),
				CentralizedSubscriptionId:    awp["centralized_subscription_id"].(string),
				WithFunctionAppsScan:         awp["with_function_apps_scan"].(bool),
				WithSseCmkEncryptedDisksScan: awp["with_sse_cmk_encrypted_disks_scan"].(bool),
			},
			Serverless: azure_org.ServerlessConfiguration{
				BladeConfiguration: azure_org.BladeConfiguration{
					IsEnabled: serverless["is_enabled"].(bool),
				},
			},
			Cdr: azure_org.CdrConfiguration{
				BladeConfiguration: azure_org.BladeConfiguration{
					IsEnabled: cdr["is_enabled"].(bool),
				},
				Accounts: extractCdrAccounts(cdr),
			},
			PostureManagement: azure_org.PostureManagement{
				OnboardingMode: aws_org.OnboardingMode(postureManagement["onboarding_mode"].(string)),
			},
		},
		Vendor: azure_org.CloudVendor(d.Get("vendor").(string)),
	}

	return req
}

func extractCdrAccounts(cdrMap map[string]interface{}) []azure_org.StorageAccount {
	var accounts []azure_org.StorageAccount

	if accountsList, ok := cdrMap["accounts"].([]interface{}); ok {
		for _, account := range accountsList {
			if accountMap, ok := account.(map[string]interface{}); ok {
				storageId := accountMap["storage_id"].(string)
				logTypesInterface := accountMap["log_types"].([]interface{})
				var logTypes []string
				for _, logType := range logTypesInterface {
					logTypes = append(logTypes, logType.(string))
				}

				storageAccount := azure_org.StorageAccount{
					StorageId: storageId,
					LogTypes:  logTypes,
				}
				accounts = append(accounts, storageAccount)
			}
		}
	}

	return accounts
}
