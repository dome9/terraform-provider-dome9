package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceAwpAzureOnboarding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwpAzureOnboardingRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scan_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"agentless_account_settings": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disabled_regions": {
							Type:     schema.TypeList,
							Computed: true,
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
							Computed: true,
						},
						"max_concurrent_scans_per_region": {
							Type:     schema.TypeInt,
							Computed: true,
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
							Computed: true,
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
		},
	}
}

func dataSourceAwpAzureOnboardingRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	cloudguardAccountId := d.Get("id").(string)
	log.Printf("Getting data for AWP Azure Onboarding id: %s\n", cloudguardAccountId)

	resp, _, err := d9Client.awpAzureOnboarding.GetAWPOnboarding(cloudguardAccountId)
	if err != nil {
		return err
	}

	d.SetId(resp.CloudAccountId)
	// Set other schema fields here
	_ = d.Set("scan_mode", resp.ScanMode)
	_ = d.Set("missing_awp_private_network_regions", resp.MissingAwpPrivateNetworkRegions)
	_ = d.Set("cloud_account_id", resp.CloudAccountId)
	_ = d.Set("agentless_protection_enabled", resp.AgentlessProtectionEnabled)
	_ = d.Set("cloud_provider", resp.Provider)
	_ = d.Set("should_update", resp.ShouldUpdate)
	_ = d.Set("is_org_onboarding", resp.IsOrgOnboarding)

	if resp.AgentlessAccountSettings != nil {
		if err := d.Set("agentless_account_settings", flattenAgentlessAccountSettingsAzure(resp.AgentlessAccountSettings)); err != nil {
			return err
		}
	}

	return nil
}
