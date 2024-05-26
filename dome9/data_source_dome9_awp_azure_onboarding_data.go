package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/dome9/dome9-sdk-go/services/awp_azure_onboarding"
)

func dataSourceAwpAzureOnboardingData() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwpAzureOnboardingDataRead,

		Schema: map[string]*schema.Schema{
			"cloud_account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"centralized_cloud_account_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default: nil,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_client_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"awp_cloud_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"awp_centralized_cloud_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAwpAzureOnboardingDataRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	cloudguardAccountId := d.Get("cloud_account_id").(string)
	req, err := expandAWPOnboardingDataRequest(d)

	resp, _, err := d9Client.awpAzureOnboarding.Get(cloudguardAccountId, req)
	if err != nil {
		return err
	}

	d.SetId(resp.AwpCloudAccountId)
	_ = d.Set("region", resp.Region)
	_ = d.Set("app_client_id", resp.AppClientId)
	_ = d.Set("awp_cloud_account_id", resp.AwpCloudAccountId)
	_ = d.Set("awp_centralized_cloud_account_id", resp.AwpCentralizedCloudAccountId)

	if err != nil {
		return err
	}
	return nil
}

func expandAWPOnboardingDataRequest(d *schema.ResourceData) (awp_azure_onboarding.CreateAWPOnboardingDataRequest, error) {

	return awp_azure_onboarding.CreateAWPOnboardingDataRequest{
		CentralizedId:            d.Get("centralized_cloud_account_id").(string),
	}, nil
}