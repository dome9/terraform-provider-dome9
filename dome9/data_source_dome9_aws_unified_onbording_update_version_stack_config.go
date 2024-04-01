package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
)

func dataSourceAwsUnifiedOnboardingUpdateVersionStackConfig() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsUnifiedOnboardingReadConfig,
		Schema: map[string]*schema.Schema{
			providerconst.OnboardingId: {
				Type:     schema.TypeString,
				Required: true,
			},
			providerconst.StackName: {
				Type:     schema.TypeString,
				Computed: true,
			},
			providerconst.Parameters: {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			providerconst.IamCapabilities: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			providerconst.TemplateUrl: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAwsUnifiedOnboardingReadConfig(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.awsUnifiedOnboarding.GetUpdateStackConfig(d.Get(providerconst.OnboardingId).(string))
	if err != nil {
		return err
	}

	d.SetId(d.Get(providerconst.OnboardingId).(string))
	_ = d.Set(providerconst.StackName, resp.StackName)
	_ = d.Set(providerconst.Parameters, resp.Parameters)
	_ = d.Set(providerconst.IamCapabilities, resp.IamCapabilities)
	_ = d.Set(providerconst.TemplateUrl, resp.TemplateUrl)

	return nil
}
