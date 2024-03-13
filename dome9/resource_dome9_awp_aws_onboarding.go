package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/awp_aws_onboarding"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"log"
	"strconv"
)

func resourceAwpAwsOnboarding() *schema.Resource {
	return &schema.Resource{
		Create: resourceAWPAWSOnboardingCreate,
		Read:   resourceAWPAWSOnboardingRead,
		Update: resourceAWPAWSOnboardingUpdate,
		Delete: resourceAWPAWSOnboardingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cloudguard_account_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"centralized_cloud_account_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cross_account_role_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cross_account_role_external_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scan_mode": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"inAccount",
					"saas",
					"in-account-hub",
					"in-account-sub",
				}, false),
			},
			"agentless_account_settings": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disabled_regions": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"scan_machine_interval_in_hours": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  4,
						},
						"max_concurrence_scans_per_region": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1,
						},
						"skip_function_apps_scan": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"custom_tags": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
					},
				},
			},
			"missing_awp_private_network_regions": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"account_issues": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"regions": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"account": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"issue_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"cloud_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"agentless_protection_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"cloud_provider": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"should_update": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_org_onboarding": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"should_create_policy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"force_delete": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
		CustomizeDiff: func(diff *schema.ResourceDiff, v interface{}) error {
			scanMode, scanModeOk := diff.GetOk("scan_mode")
			centralizedCloudAccountId, centralizedCloudAccountIdOk := diff.GetOk("centralized_cloud_account_id")

			if scanModeOk && scanMode == "in-account-sub" {
				if !centralizedCloudAccountIdOk || centralizedCloudAccountId == "" {
					return fmt.Errorf("'centralized_cloud_account_id' must be set and not empty when 'scan_mode' is 'in-account-sub'")
				}
			}

			return nil
		},
	}
}

func resourceAWPAWSOnboardingCreate(d *schema.ResourceData, meta interface{}) error {
	d9client := meta.(*Client)
	cloudguardAccountId := d.Get("cloudguard_account_id").(string)
	req := expandAWPOnboardingRequest(d)
	log.Printf("[INFO] Creating AWP AWS Onboarding request %+v\n", req)
	options := awp_aws_onboarding.CreateOptions{
		ShouldCreatePolicy: strconv.FormatBool(d.Get("should_create_policy").(bool)),
	}
	_, err := d9client.awpAwsOnboarding.CreateAWPOnboarding(cloudguardAccountId, req, options)
	if err != nil {
		return err
	}
	d.SetId(cloudguardAccountId)
	log.Printf("[INFO] Created AWP AWS Onboarding with CloudGuard Account ID: %v\n", cloudguardAccountId)
	d.SetId(cloudguardAccountId) // set the resource ID to the CloudGuard Account ID

	return resourceAWPAWSOnboardingRead(d, meta)
}

func expandAWPOnboardingRequest(d *schema.ResourceData) awp_aws_onboarding.CreateAWPOnboardingRequest {

	return awp_aws_onboarding.CreateAWPOnboardingRequest{
		CrossAccountRoleName:       d.Get("cross_account_role_name").(string),
		CrossAccountRoleExternalId: d.Get("cross_account_role_external_id").(string),
		ScanMode:                   d.Get("scan_mode").(string),
		IsTerraform:                true,
		AgentlessAccountSettings:   expandAgentlessAccountSettings(d),
	}
}

func resourceAWPAWSOnboardingRead(d *schema.ResourceData, meta interface{}) error {
	d9client := meta.(*Client)
	resp, _, err := d9client.awpAwsOnboarding.GetAWPOnboarding("aws", d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing AWS cloud account %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	log.Printf("[INFO] Reading AWP AWS Onbaording account data: %+v\n", resp)

	// set the schema fields from the response
	_ = d.Set("missing_awp_private_network_regions", resp.MissingAwpPrivateNetworkRegions)
	_ = d.Set("cloud_account_id", resp.CloudAccountId)
	_ = d.Set("agentless_protection_enabled", resp.AgentlessProtectionEnabled)
	_ = d.Set("scan_mode", resp.ScanMode)
	_ = d.Set("cloud_provider", resp.Provider)
	_ = d.Set("should_update", resp.ShouldUpdate)
	_ = d.Set("is_org_onboarding", resp.IsOrgOnboarding)
	_ = d.Set("centralized_cloud_account_id", resp.CentralizedCloudAccountId)

	if resp.AgentlessAccountSettings != nil {
		if err := d.Set("agentless_account_settings", flattenAgentlessAccountSettings(resp.AgentlessAccountSettings)); err != nil {
			return err
		}
	}

	if resp.AccountIssues != nil {
		if err := d.Set("account_issues", flattenAccountIssues(resp.AccountIssues)); err != nil {
			return err
		}
	}

	return nil
}

func resourceAWPAWSOnboardingDelete(d *schema.ResourceData, meta interface{}) error {
	d9client := meta.(*Client)
	log.Printf("[INFO] Offboarding AWP Account with cloud guard id : %v\n", d.Id())
	options := awp_aws_onboarding.DeleteOptions{
		ForceDelete: strconv.FormatBool(d.Get("force_delete").(bool)),
	}
	_, err := d9client.awpAwsOnboarding.DeleteAWPOnboarding(d.Id(), options)
	if err != nil {
		return err
	}
	return nil
}

func expandAgentlessAccountSettings(d *schema.ResourceData) *awp_aws_onboarding.AgentlessAccountSettings {
	if _, ok := d.GetOk("agentless_account_settings"); !ok {
		// If "agentless_account_settings" key doesn't exist, return nil (since these settings are optional)
		return nil
	}
	agentlessAccountSettingsItem := d.Get("agentless_account_settings").(*schema.Set).List()[0]
	agentlessAccountSettingsMap := agentlessAccountSettingsItem.(map[string]interface{})

	// Initialize the AgentlessAccountSettings struct with default values
	agentlessAccountSettings := &awp_aws_onboarding.AgentlessAccountSettings{
		DisabledRegions:              make([]string, 0),
		CustomTags:                   make(map[string]string),
		ScanMachineIntervalInHours:   4,
		MaxConcurrenceScansPerRegion: 1,
		SkipFunctionAppsScan:         true,
	}

	// Check if the key exists and is not nil
	if disabledRegionsInterface, ok := agentlessAccountSettingsMap["disabled_regions"].([]interface{}); ok {
		disabledRegions := make([]string, len(disabledRegionsInterface))
		for i, disabledRegion := range disabledRegionsInterface {
			disabledRegions[i] = disabledRegion.(string)
		}
		agentlessAccountSettings.DisabledRegions = disabledRegions
	}

	if scanMachineInterval, ok := agentlessAccountSettingsMap["scan_machine_interval_in_hours"].(int); ok {
		agentlessAccountSettings.ScanMachineIntervalInHours = scanMachineInterval
	}

	if maxConcurrenceScans, ok := agentlessAccountSettingsMap["max_concurrence_scans_per_region"].(int); ok {
		agentlessAccountSettings.MaxConcurrenceScansPerRegion = maxConcurrenceScans
	}

	if skipFunctionAppsScan, ok := agentlessAccountSettingsMap["skip_function_apps_scan"].(bool); ok {
		agentlessAccountSettings.SkipFunctionAppsScan = skipFunctionAppsScan
	}

	if customTagsInterface, ok := agentlessAccountSettingsMap["custom_tags"].(map[string]interface{}); ok {
		customTags := make(map[string]string)
		for k, v := range customTagsInterface {
			customTags[k] = v.(string)
		}
		agentlessAccountSettings.CustomTags = customTags
	}

	return agentlessAccountSettings
}

func flattenAgentlessAccountSettings(settings *awp_aws_onboarding.AgentlessAccountSettings) []interface{} {

	m := map[string]interface{}{
		"disabled_regions":                 settings.DisabledRegions,
		"scan_machine_interval_in_hours":   settings.ScanMachineIntervalInHours,
		"max_concurrence_scans_per_region": settings.MaxConcurrenceScansPerRegion,
		"skip_function_apps_scan":          settings.SkipFunctionAppsScan,
		"custom_tags":                      settings.CustomTags,
	}
	return []interface{}{m}
}

func flattenAccountIssues(accountIssues *awp_aws_onboarding.AccountIssues) []interface{} {
	m := map[string]interface{}{
		"regions": accountIssues.Regions,
		"account": accountIssues.Account,
	}

	return []interface{}{m}
}

func resourceAWPAWSOnboardingUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
