package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/unifiedOnbording/awsUnifiedOnbording"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
	"log"
)
func resourceAwsUnifiedOnbording() *schema.Resource {
	return &schema.Resource{
		Create: resourceUnifiedOnboardingCreate,
		Read:   unifiedOnbordingResourceRead,
		Update: resourceUnifiedOnbordingUpdat,
		Delete: resourceUnifiedOnbordingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			providerconst.OnboardType: {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			providerconst.FullProtection: {
				Type:     schema.TypeBool,
				Optional: true,
			},
			providerconst.CloudVendor: {
				Type:     schema.TypeString,
				Optional: true,
			},
			providerconst.EnableStackModify: {
				Type:     schema.TypeBool,
				Optional: true,
			},
			providerconst.PostureManagementConfiguration: {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						providerconst.Rulesets: {
							Type:     schema.TypeList,
							Required: true,
						},
					},
				},
			},
			providerconst.ServerlessConfiguration: {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						providerconst.Enabled: {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
			providerconst.IntelligenceConfigurations: {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						providerconst.Rulesets: {
							Type:     schema.TypeList,
							Required: false,
						},
						providerconst.Enabled: {
							Type:     schema.TypeBool,
							Required: false,
						},
					},
				},
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

func resourceUnifiedOnboardingCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandAwsUnifiedOnbordingRequest(d)
	log.Printf("[INFO] Creating Unified Onbording request %+v\n", req)
	resp, _, err := d9Client.AwsUnifiedOnbording.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created UnifiedOnbording resource with stackName: %v\n", resp.StackName)

	addOnboardingIdAsSchemaId(d, resp)

	log.Printf("[INFO] Getting Unified Onbording:\n%+v\n", resp)
	_ = d.Set(providerconst.StackName, resp.StackName)
	_ = d.Set(providerconst.Parameters, resp.Parameters)
	_ = d.Set(providerconst.IamCapabilities, resp.IamCapabilities)
	_ = d.Set(providerconst.TemplateUrl, resp.TemplateUrl)

	return nil
}

func expandAwsUnifiedOnbordingRequest(d *schema.ResourceData) awsUnifiedOnbording.UnifiedOnbordingRequest {

	return awsUnifiedOnbording.UnifiedOnbordingRequest{
		CloudVendor:                    d.Get(providerconst.CloudVendor).(string),
		OnboardType:                    d.Get(providerconst.OnboardType).(string),
		EnableStackModify:              d.Get(providerconst.EnableStackModify).(bool),
		FullProtection:                 d.Get(providerconst.FullProtection).(bool),
		PostureManagementConfiguration: expendPostureManagementConfiguration(d),
		ServerlessConfiguration:        expendServerlessConfiguration(d),
		IntelligenceConfigurations:     expendIntelligenceConfigurations(d),
	}

}

func expendIntelligenceConfigurations(d *schema.ResourceData) awsUnifiedOnbording.IntelligenceConfigurations {
	var intelligenceConfigurations awsUnifiedOnbording.IntelligenceConfigurations
	intelligenceConfigurations.Enabled = true // d.Get(variable.Enabled).(bool)
	intelligenceConfigurations.Rulesets = *getRulesets(d)

	return intelligenceConfigurations
}

func expendServerlessConfiguration(d *schema.ResourceData) awsUnifiedOnbording.ServerlessConfiguration {
	var serverlessConfiguration awsUnifiedOnbording.ServerlessConfiguration
	item := true// d.Get(Enabled).(bool)
	serverlessConfiguration.Enabled = item
	return serverlessConfiguration
}

func expendPostureManagementConfiguration(d *schema.ResourceData) awsUnifiedOnbording.PostureManagementConfiguration {
	var postureManagementConfiguration awsUnifiedOnbording.PostureManagementConfiguration
	postureManagementConfiguration.Rulesets = *getRulesets(d)
	return postureManagementConfiguration
}

func getRulesets(d *schema.ResourceData) *[]int {
	var rulesets []int
	if itemsInterface, ok := d.GetOk(providerconst.Rulesets); ok {
		items := itemsInterface.([]interface{})
		rulesets = make([]int, len(items))
		for i, item := range items {
			rulesets[i] = item.(int)
		}
	}

	if rulesets == nil {
		rulesets = make([]int, 0)
	}

	return &rulesets
}

func addOnboardingIdAsSchemaId(d *schema.ResourceData, resp *awsUnifiedOnbording.UnifiedOnbordingConfigurationResponse) {
	var p = resp.Parameters
	var schemaId string
	for _, value := range p {
		if value.Key == providerconst.OnboardingId {
			schemaId = value.Value
		}
	}

	if len(schemaId) > 0 {
		d.SetId(schemaId)
	}
}

func resourceUnifiedOnbordingDelete(data *schema.ResourceData, i interface{}) error {
	return nil
}

func resourceUnifiedOnbordingUpdat(data *schema.ResourceData, i interface{}) error {
	return nil
}

func unifiedOnbordingResourceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
