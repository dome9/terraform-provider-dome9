package dome9

import (
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	Rulesets        = "rulesets"
	Enabled         = "enabled"
	StackName       = "stack_name"
	Parameters      = "parameters"
	IamCapabilities = "iam_capabilities"
	TemplateUrl     = "template_url"
	OnboardingId    = "onboarding_id"
)

func dataSourceAwsUnifiedOnboarding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsUnifiedOnboardingReadConfig,
		Schema: map[string]*schema.Schema{
			OnboardingId: {
				Type:     schema.TypeString,
				Required: true,
			},
			StackName: {
				Type:     schema.TypeString,
				Computed: true,
			},
			Parameters: {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			IamCapabilities: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			TemplateUrl: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAwsUnifiedOnboardingReadConfig(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.AwsUnifiedOnbording.GetUpdateStackConfig(d.Get("onboarding_id").(string))
	if err != nil {
		return err
	}

	log.Printf("[INFO] Get UnifiedOnbording data resource configuration with stackName: %v\n", resp.StackName)

	return unifiedOnbordingReadConfiguration(d, meta)
}

func unifiedOnbordingReadConfiguration(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.AwsUnifiedOnbording.GetUpdateStackConfig(d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing rule set %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	log.Printf("[INFO] Getting Unified Onbording:\n%+v\n", resp)
	_ = d.Set(StackName, resp.StackName)
	_ = d.Set(Parameters, resp.Parameters)
	_ = d.Set(IamCapabilities, resp.IamCapabilities)
	_ = d.Set(TemplateUrl, resp.TemplateUrl)

	return nil
}
