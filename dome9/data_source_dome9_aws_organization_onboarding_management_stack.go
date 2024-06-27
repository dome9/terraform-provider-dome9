package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAwsOrganizationOnboardingManagementStack() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsOrganizationOnboardingManagementStackRead,

		Schema: map[string]*schema.Schema{
			// OrganizationManagementViewModel object fields
			"aws_account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_cft_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_management_onboarded": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceAwsOrganizationOnboardingManagementStackRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	orgId := d.Get("aws_account_id").(string)
	resp, _, _ := d9Client.awsOrganizationOnboarding.GetOnboardingConfiguration(orgId)

	//if err != nil {
	//	if err.(*client.ErrorResponse).IsObjectNotFound() {
	//		//log.Printf("[WARN] Removing Aws organization %s from state because it no longer exists in CloudGuard", d.Id())
	//		d.SetId("")
	//		return nil
	//	}
	//
	//	return err
	//}

	_ = d.Set("external_id", resp.ExternalId)
	_ = d.Set("content", resp.Content)
	_ = d.Set("management_cft_url", resp.ManagementCftUrl)
	_ = d.Set("is_management_onboarded", resp.IsManagementOnboarded)

	return nil
}
