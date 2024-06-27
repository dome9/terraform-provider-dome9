package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAwsOrganizationOnboardingMemberAccountConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsOrganizationOnboardingMemberAccountConfigurationRead,

		Schema: map[string]*schema.Schema{
			// OrganizationManagementViewModel object fields
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"onboarding_cft_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAwsOrganizationOnboardingMemberAccountConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.awsOrganizationOnboarding.GetMemberAccountConfiguration()
	if err != nil {
		return err
	}
	_ = d.Set("external_id", resp.ExternalId)
	_ = d.Set("content", resp.Content)
	_ = d.Set("onboarding_cft_url", resp.OnboardingCftUrl)

	return nil
}
