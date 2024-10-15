package dome9

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/awp"
	"github.com/dome9/dome9-sdk-go/services/awp/azure_onboarding"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
)

func resourceAwpAzureOnboarding() *schema.Resource {
	return &schema.Resource{
		Create: resourceAWPAzureOnboardingCreate,
		Read:   resourceAWPAzureOnboardingRead,
		Update: resourceAWPAzureOnboardingUpdate,
		Delete: resourceAWPAzureOnboardingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cloudguard_account_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scan_mode": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"inAccount",
					"saas",
					"inAccountHub",
					"inAccountSub",
				}, false),
			},
			"centralized_cloud_account_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
			"management_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
			"agentless_account_settings": {
				Type:     schema.TypeList,
				MaxItems: 1,
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
						"skip_function_apps_scan": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"scan_machine_interval_in_hours": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  24,
						},
						"max_concurrent_scans_per_region": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  20,
						},
						"in_account_scanner_vpc": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "ManagedByAWP",
						},
						"sse_cmk_encrypted_disks_scan": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
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
			"awp_version": {
				Type:     schema.TypeString,
			},
			"missing_awp_private_network_regions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Default:  true,
			},
		},
	}
}

func resourceAWPAzureOnboardingCreate(d *schema.ResourceData, meta interface{}) error {
	d9client := meta.(*Client)
	cloudguardAccountId := d.Get("cloudguard_account_id").(string)
	req, err := expandAWPOnboardingRequestAzure(d, meta)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Creating AWP Azure Onboarding request %+v\n", req)
	options := awp_onboarding.CreateOptions{
		ShouldCreatePolicy: strconv.FormatBool(d.Get("should_create_policy").(bool)),
	}
	_, err = d9client.awpAzureOnboarding.CreateAWPOnboarding(cloudguardAccountId, req, options)
	if err != nil {
		return err
	}
	d.SetId(cloudguardAccountId) // set the resource ID to the CloudGuard Account ID
	log.Printf("[INFO] Created AWP Azure Onboarding with CloudGuard Account ID: %v\n", cloudguardAccountId)

	return resourceAWPAzureOnboardingRead(d, meta)
}

func expandAWPOnboardingRequestAzure(d *schema.ResourceData, meta interface{}) (awp_azure_onboarding.CreateAWPOnboardingRequestAzure, error) {
	cloudGuardHubAccountID, err := checkCentralizedAzure(d, meta)
	agentlessAccountSettings, err := expandAgentlessAccountSettingsAzure(d)
	if err != nil {
		return awp_azure_onboarding.CreateAWPOnboardingRequestAzure{}, err
	}
	return awp_azure_onboarding.CreateAWPOnboardingRequestAzure{
		ScanMode:                  d.Get("scan_mode").(string),
		IsTerraform:               true,
		ManagementGroupId:         d.Get("management_group_id").(string),
		AgentlessAccountSettings:  agentlessAccountSettings,
		CentralizedCloudAccountId: cloudGuardHubAccountID,
	}, nil
}

func checkCentralizedAzure(d *schema.ResourceData, meta interface{}) (string, error) {
	scanMode := d.Get("scan_mode").(string)
	if scanMode == "inAccountSub" {
		d9client := meta.(*Client)
		hubExternalAccountId, exist := d.Get("centralized_cloud_account_id").(string)
		if !exist || hubExternalAccountId == "" {
			errorMsg := fmt.Sprintf("centralized_cloud_account_id is required when scan_mode is inAccountSub, got '%s'", hubExternalAccountId)
			return "", errors.New(errorMsg)
		}

		getCloudAccountQueryParams := cloudaccounts.QueryParameters{ID: hubExternalAccountId}
		cloudAccountresp, _, err := d9client.cloudaccountAzure.Get(&getCloudAccountQueryParams)
		if err != nil {
			return "", err
		}
		return cloudAccountresp.ID, nil
	}
	return "", nil
}

func resourceAWPAzureOnboardingRead(d *schema.ResourceData, meta interface{}) error {
	d9client := meta.(*Client)
	resp, _, err := d9client.awpAzureOnboarding.GetAWPOnboarding(d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing Azure cloud account %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	log.Printf("[INFO] Reading AWP Azure Onbaording account data: %+v\n", resp)

	// set the schema fields from the response
	_ = d.Set("missing_awp_private_network_regions", resp.MissingAwpPrivateNetworkRegions)
	_ = d.Set("cloud_account_id", resp.CloudAccountId)
	_ = d.Set("agentless_protection_enabled", resp.AgentlessProtectionEnabled)
	_ = d.Set("scan_mode", resp.ScanMode)
	_ = d.Set("cloud_provider", resp.Provider)
	_ = d.Set("should_update", resp.ShouldUpdate)
	_ = d.Set("is_org_onboarding", resp.IsOrgOnboarding)

	err = setAgentlessAccountSettingsAzure(resp, d)
	if err != nil {
		return err
	}

	return nil
}

func resourceAWPAzureOnboardingDelete(d *schema.ResourceData, meta interface{}) error {
	d9client := meta.(*Client)
	log.Printf("[INFO] Offboarding AWP Account with cloud guard id : %v\n", d.Id())

	_, err := d9client.awpAzureOnboarding.DeleteAWPOnboarding(d.Id())
	if err != nil {
		return err
	}
	return nil
}

func expandAgentlessAccountSettingsAzure(d *schema.ResourceData) (*awp_onboarding.AgentlessAccountSettings, error) {
	if _, ok := d.GetOk("agentless_account_settings"); !ok {
		// If "agentless_account_settings" key doesn't exist, return nil (since these settings are optional)
		return nil, nil
	}
	agentlessAccountSettingsList := d.Get("agentless_account_settings").([]interface{})
	agentlessAccountSettingsMap := agentlessAccountSettingsList[0].(map[string]interface{})
	scanMode := d.Get("scan_mode").(string)

	var scanMachineIntervalInHours int
	if scanMode == "saas" {
		scanMachineIntervalInHours = providerconst.DefaultScanMachineIntervalInHoursSaas
	} else {
		scanMachineIntervalInHours = providerconst.DefaultScanMachineIntervalInHoursInAccount
	}

	// Initialize the AgentlessAccountSettings struct with default values
	agentlessAccountSettings := &awp_onboarding.AgentlessAccountSettings{
		DisabledRegions:              make([]string, 0),
		SkipFunctionAppsScan:         false,
		CustomTags:                   make(map[string]string),
		ScanMachineIntervalInHours:   scanMachineIntervalInHours,
		InAccountScannerVPC:          providerconst.DefaultInAccountScannerVPCMode,
		SseCmkEncryptedDisksScan:     false,
		MaxConcurrenceScansPerRegion: providerconst.DefaultMaxConcurrentScansPerRegion,
	}

	// Check if the key exists and is not nil
	if disabledRegionsInterface, ok := agentlessAccountSettingsMap["disabled_regions"].([]interface{}); ok {
		disabledRegions := make([]string, len(disabledRegionsInterface))
		for i, disabledRegion := range disabledRegionsInterface {
			disabledRegions[i] = disabledRegion.(string)
		}
		_, err := validateDisabledRegionsAzure(disabledRegions)
		if err != nil {
			return agentlessAccountSettings, err
		}
		agentlessAccountSettings.DisabledRegions = disabledRegions
	}

	if scanMachineInterval, ok := agentlessAccountSettingsMap["scan_machine_interval_in_hours"].(int); ok {
		if scanMode == "saas" && (scanMachineInterval < providerconst.DefaultScanMachineIntervalInHoursSaas || scanMachineInterval > providerconst.MaxScanMachineIntervalInHours) {
			return nil, fmt.Errorf("scan_machine_interval_in_hours must be between %d and %d for saas mode", providerconst.DefaultScanMachineIntervalInHoursSaas, providerconst.MaxScanMachineIntervalInHours)
		} else if scanMode == "inAccount" && (scanMachineInterval < providerconst.DefaultScanMachineIntervalInHoursInAccount || scanMachineInterval > providerconst.MaxScanMachineIntervalInHours) {
			return nil, fmt.Errorf("scan_machine_interval_in_hours must be between %d and %d for inAccount mode", providerconst.DefaultScanMachineIntervalInHoursInAccount, providerconst.MaxScanMachineIntervalInHours)
		}
		agentlessAccountSettings.ScanMachineIntervalInHours = scanMachineInterval
	}

	if maxConcurrentScans, ok := agentlessAccountSettingsMap["max_concurrent_scans_per_region"].(int); ok {
		if maxConcurrentScans < providerconst.MinMaxConcurrentScansPerRegion || maxConcurrentScans > providerconst.DefaultMaxConcurrentScansPerRegion {
			return nil, fmt.Errorf("max_concurrent_scans_per_region must be between 1 and 20")
		}
		agentlessAccountSettings.MaxConcurrenceScansPerRegion = maxConcurrentScans
	}

	if inAccountScannerVPC, ok := agentlessAccountSettingsMap["in_account_scanner_vpc"].(string); ok {
		agentlessAccountSettings.InAccountScannerVPC = inAccountScannerVPC
	}

	if sseCmkEncryptedDisksScan, ok := agentlessAccountSettingsMap["sse_cmk_encrypted_disks_scan"].(bool); ok {
		agentlessAccountSettings.SseCmkEncryptedDisksScan = sseCmkEncryptedDisksScan
	}

	if customTagsInterface, ok := agentlessAccountSettingsMap["custom_tags"].(map[string]interface{}); ok {
		customTags := make(map[string]string)
		for k, v := range customTagsInterface {
			customTags[k] = v.(string)
		}
		agentlessAccountSettings.CustomTags = customTags
	}

	if skipFunctionAppsScan, ok := agentlessAccountSettingsMap["skip_function_apps_scan"].(bool); ok {
		agentlessAccountSettings.SkipFunctionAppsScan = skipFunctionAppsScan
	}

	return agentlessAccountSettings, nil
}

func setAgentlessAccountSettingsAzure(resp *awp_onboarding.GetAWPOnboardingResponse, d *schema.ResourceData) error {
	if resp.AgentlessAccountSettings != nil {
		// Check if all fields of AgentlessAccountSettings are nil
		if resp.AgentlessAccountSettings.DisabledRegions != nil ||
			resp.AgentlessAccountSettings.ScanMachineIntervalInHours != 0 ||
			resp.AgentlessAccountSettings.MaxConcurrenceScansPerRegion != 0 ||
			resp.AgentlessAccountSettings.CustomTags != nil {
			if err := d.Set("agentless_account_settings", flattenAgentlessAccountSettingsAzure(resp.AgentlessAccountSettings)); err != nil {
				return err
			}
		}
	}
	return nil
}

func flattenAgentlessAccountSettingsAzure(settings *awp_onboarding.AgentlessAccountSettings) []interface{} {

	m := map[string]interface{}{
		"disabled_regions":                settings.DisabledRegions,
		"skip_function_apps_scan":         settings.SkipFunctionAppsScan,
		"scan_machine_interval_in_hours":  settings.ScanMachineIntervalInHours,
		"max_concurrent_scans_per_region": settings.MaxConcurrenceScansPerRegion,
		"in_account_scanner_vpc":          settings.InAccountScannerVPC,
		"sse_cmk_encrypted_disks_scan":    settings.SseCmkEncryptedDisksScan,
		"custom_tags":                     settings.CustomTags,
	}
	return []interface{}{m}
}

func resourceAWPAzureOnboardingUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Println("An update occurred")

	if d.HasChange("should_create_policy") {
		log.Println("should_create_policy has been changed")
		if err := d.Set("should_create_policy", d.Get("should_create_policy").(bool)); err != nil {
			return err
		}
	}
	// Check if there are changes in the AgentlessAccountSettings fields
	if d.HasChange("agentless_account_settings") || d.HasChange("awp_version") {
		log.Println("agentless_account_settings has been changed")
		// Build the update request
		newAgentlessAccountSettings, err := expandAgentlessAccountSettingsAzure(d)
		if err != nil {
			return err
		}
		// Send the update request
		scanMode := d.Get("scan_mode").(string)

		_, err = d9Client.awpAzureOnboarding.UpdateAWPSettings(d.Id(), scanMode, *newAgentlessAccountSettings)
		if err != nil {
			return err
		}
		log.Printf("[INFO] Updated agentless account settings for cloud account %s\n", d.Id())
	}

	return nil
}

func validateDisabledRegionsAzure(regions []string) (bool, error) {
	hyphenatedAzureRegions := convertRegionsFormatAzure(providerconst.AzureSecurityGroupRegions)
	validate, invalidRegions := checkDisabledRegionsAzure(regions, hyphenatedAzureRegions)
	if !validate {
		errorMsg := fmt.Sprintf("Expected disabled-regions to be one of %v, got %v", hyphenatedAzureRegions, invalidRegions)
		return false, errors.New(errorMsg)
	}
	return true, nil
}

func convertRegionsFormatAzure(regions []string) []string {
	hyphenatedRegions := make([]string, len(regions))
	for i, region := range regions {
		hyphenatedRegions[i] = strings.ReplaceAll(region, "_", "-")
	}
	return hyphenatedRegions
}

func checkDisabledRegionsAzure(regions []string, regionsToCompare []string) (bool, []string) {
	invalidRegions := make([]string, 0)
	for _, val := range regions {
		flag := false
		for _, region := range regionsToCompare {
			if val == region {
				flag = true
				break
			}
		}
		if !flag {
			invalidRegions = append(invalidRegions, val)
		}
	}
	return len(invalidRegions) == 0, invalidRegions
}
