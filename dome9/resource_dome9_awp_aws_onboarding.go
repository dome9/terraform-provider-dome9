package dome9

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/awp"
	"github.com/dome9/dome9-sdk-go/services/awp/aws_onboarding"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
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
			"awp_centralized_cloud_account_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"cross_account_role_name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"cross_account_role_external_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
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
			"force_delete": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceAWPAWSOnboardingCreate(d *schema.ResourceData, meta interface{}) error {
	d9client := meta.(*Client)
	cloudguardAccountId := d.Get("cloudguard_account_id").(string)
	cloudGuardHubAccountID, err := checkCentralized(d, meta)
	if err != nil {
		return err
	}
	agentlessAccountSettings, err := expandAgentlessAccountSettings(d)
	if err != nil {
		return err
	}

	req := awp_aws_onboarding.CreateAWPOnboardingRequestAws{
		CrossAccountRoleName:       d.Get("cross_account_role_name").(string),
		CentralizedCloudAccountId:  cloudGuardHubAccountID,
		CrossAccountRoleExternalId: d.Get("cross_account_role_external_id").(string),
		IsTerraform:                true,
		AgentlessAccountSettings:   agentlessAccountSettings,
		ScanMode:                   d.Get("scan_mode").(string),
	}

	log.Printf("[INFO] Creating AWP AWS Onboarding request %+v\n", req)

	options := awp_onboarding.CreateOptions{
		ShouldCreatePolicy: strconv.FormatBool(d.Get("should_create_policy").(bool)),
	}
	_, err = d9client.awpAwsOnboarding.CreateAWPOnboarding(cloudguardAccountId, req, options)
	if err != nil {
		return err
	}

	d.SetId(cloudguardAccountId) // set the resource ID to the CloudGuard Account ID
	log.Printf("[INFO] Created AWP AWS Onboarding with CloudGuard Account ID: %v\n", cloudguardAccountId)

	return resourceAWPAWSOnboardingRead(d, meta)
}

func checkCentralized(d *schema.ResourceData, meta interface{}) (string, error) {
	scanMode := d.Get("scan_mode").(string)
	if scanMode == "inAccountSub" {
		if _, ok := d.GetOk("agentless_account_settings"); ok {
			agentlessAccountSettingsList := d.Get("agentless_account_settings").([]interface{})
			if len(agentlessAccountSettingsList) < 1 {
				errorMsg := fmt.Sprintf("currently account settings not supported for centralized onboarding (%s)", scanMode)
				return "", errors.New(errorMsg)
			}
		}
	}
	if scanMode == "inAccountSub" {
		d9client := meta.(*Client)
		hubExternalAccountId, exist := d.Get("awp_centralized_cloud_account_id").(string)
		if !exist || hubExternalAccountId == "" {
			errorMsg := fmt.Sprintf("awp_centralized_cloud_account_id is required when scan_mode is inAccountSub, got '%s'", hubExternalAccountId)
			return "", errors.New(errorMsg)
		}

		getCloudAccountQueryParams := cloudaccounts.QueryParameters{ID: hubExternalAccountId}
		cloudAccountresp, _, err := d9client.cloudaccountAWS.Get(&getCloudAccountQueryParams)
		if err != nil {
			return "", err
		}
		return cloudAccountresp.ID, nil
	}
	return "", nil
}

func resourceAWPAWSOnboardingRead(d *schema.ResourceData, meta interface{}) error {
	d9client := meta.(*Client)
	resp, _, err := d9client.awpAwsOnboarding.GetAWPOnboarding(d.Id())
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

	err = setAgentlessAccountSettings(resp, d)
	if err != nil {
		return err
	}

	return nil
}

func resourceAWPAWSOnboardingDelete(d *schema.ResourceData, meta interface{}) error {
	d9client := meta.(*Client)
	log.Printf("[INFO] Offboarding AWP Account with cloud guard id : %v\n", d.Id())
	options := awp_onboarding.DeleteOptions{
		ForceDelete: strconv.FormatBool(d.Get("force_delete").(bool)),
	}
	_, err := d9client.awpAwsOnboarding.DeleteAWPOnboarding(d.Id(), options)
	if err != nil {
		return err
	}
	if d.Get("scan_mode").(string) == "inAccountSub" {
		// delay for 30 seconds to allow the account to be removed from the hub
		time.Sleep(30 * time.Second)
	}
	return nil
}

func expandAgentlessAccountSettings(d *schema.ResourceData) (*awp_onboarding.AgentlessAccountSettings, error) {
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
		CustomTags:                   make(map[string]string),
		ScanMachineIntervalInHours:   scanMachineIntervalInHours,
		InAccountScannerVPC:          providerconst.DefaultInAccountScannerVPCMode,
		MaxConcurrenceScansPerRegion: providerconst.DefaultMaxConcurrentScansPerRegion,
	}

	// Check if the key exists and is not nil
	if disabledRegionsInterface, ok := agentlessAccountSettingsMap["disabled_regions"].([]interface{}); ok {
		disabledRegions := make([]string, len(disabledRegionsInterface))
		for i, disabledRegion := range disabledRegionsInterface {
			disabledRegions[i] = disabledRegion.(string)
		}
		_, err := validateDisabledRegions(disabledRegions)
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

	if customTagsInterface, ok := agentlessAccountSettingsMap["custom_tags"].(map[string]interface{}); ok {
		customTags := make(map[string]string)
		for k, v := range customTagsInterface {
			customTags[k] = v.(string)
		}
		agentlessAccountSettings.CustomTags = customTags
	}

	return agentlessAccountSettings, nil
}

func setAgentlessAccountSettings(resp *awp_onboarding.GetAWPOnboardingResponse, d *schema.ResourceData) error {
	if resp.AgentlessAccountSettings != nil {
		// Check if all fields of AgentlessAccountSettings are nil
		if resp.AgentlessAccountSettings.DisabledRegions != nil ||
			resp.AgentlessAccountSettings.ScanMachineIntervalInHours != 0 ||
			resp.AgentlessAccountSettings.MaxConcurrenceScansPerRegion != 0 ||
			resp.AgentlessAccountSettings.CustomTags != nil {
			if err := d.Set("agentless_account_settings", flattenAgentlessAccountSettings(resp.AgentlessAccountSettings)); err != nil {
				return err
			}
		}
	}
	return nil
}

func flattenAgentlessAccountSettings(settings *awp_onboarding.AgentlessAccountSettings) []interface{} {

	m := map[string]interface{}{
		"disabled_regions":                settings.DisabledRegions,
		"scan_machine_interval_in_hours":  settings.ScanMachineIntervalInHours,
		"max_concurrent_scans_per_region": settings.MaxConcurrenceScansPerRegion,
		"in_account_scanner_vpc":          settings.InAccountScannerVPC,
		"custom_tags":                     settings.CustomTags,
	}
	return []interface{}{m}
}

func resourceAWPAWSOnboardingUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Println("An update occurred")

	if d.HasChange("delete_force") {
		log.Println("delete_force has been changed")
		if err := d.Set("delete_force", d.Get("delete_force").(bool)); err != nil {
			return err
		}
	}

	if d.HasChange("should_create_policy") {
		log.Println("should_create_policy has been changed")
		if err := d.Set("should_create_policy", d.Get("should_create_policy").(bool)); err != nil {
			return err
		}
	}

	_, err := checkCentralized(d, meta)
	if err != nil {
		return err
	}

	// Check if there are changes in the AgentlessAccountSettings fields
	if d.HasChange("agentless_account_settings") {
		log.Println("agentless_account_settings has been changed")

		// Build the update request
		newAgentlessAccountSettings, err := expandAgentlessAccountSettings(d)
		if err != nil {
			return err
		}

		scanMode := d.Get("scan_mode").(string)

		// Send the update request
		_, err = d9Client.awpAwsOnboarding.UpdateAWPSettings(d.Id(), scanMode, *newAgentlessAccountSettings)
		if err != nil {
			return err
		}
		log.Printf("[INFO] Updated agentless account settings for cloud account %s\n", d.Id())
	}

	return nil
}

func validateDisabledRegions(regions []string) (bool, error) {
	hyphenatedAWSRegions := convertRegionsFormat(providerconst.AWSRegions)
	validate, invalidRegions := checkDisabledRegions(regions, hyphenatedAWSRegions)
	if !validate {
		errorMsg := fmt.Sprintf("Expected disabled-regions to be one of %v, got %v", hyphenatedAWSRegions, invalidRegions)
		return false, errors.New(errorMsg)
	}
	return true, nil
}

func convertRegionsFormat(regions []string) []string {
	hyphenatedRegions := make([]string, len(regions))
	for i, region := range regions {
		hyphenatedRegions[i] = strings.ReplaceAll(region, "_", "-")
	}
	return hyphenatedRegions
}

func checkDisabledRegions(regions []string, regionsToCompare []string) (bool, []string) {
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
