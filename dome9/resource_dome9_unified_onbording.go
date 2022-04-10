package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/unifiedOnbording/awsUnifiedOnbording"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"log"
)
func resourceUnifiedOnbording() *schema.Resource {
	return &schema.Resource{
		Create: resourceUnifiedOnbordingCreate,
		Read:   unifiedOnbordingResourceRead,
		Update: resourceUnifiedOnbordingUpdat,
		Delete: resourceUnifiedOnbordingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"onboard_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"full_protection": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cloud_vendor": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_stack_modify": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"posture_management_configuration": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Rulesets: {
							Type:     schema.TypeList,
							Required: true,
						},
					},
				},
			},
			"serverless_configuration": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Enabled: {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
			"intelligence_configurations": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Rulesets: {
							Type:     schema.TypeList,
							Required: false,
						},
						Enabled: {
							Type:     schema.TypeBool,
							Required: false,
						},
					},
				},
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

//todo to consider the resource aws_cloudformation_stack
func resourceUnifiedOnbordingCreate(d *schema.ResourceData, meta interface{}) error {
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
	_ = d.Set(StackName, resp.StackName)
	_ = d.Set(Parameters, resp.Parameters)
	_ = d.Set(IamCapabilities, resp.IamCapabilities)
	_ = d.Set(TemplateUrl, resp.TemplateUrl)

	return nil
}

func expandAwsUnifiedOnbordingRequest(d *schema.ResourceData) awsUnifiedOnbording.UnifiedOnbordingRequest {

	return awsUnifiedOnbording.UnifiedOnbordingRequest{
		CloudVendor:                    d.Get("cloud_vendor").(string),
		OnboardType:                    d.Get("onboard_type").(string),
		EnableStackModify:              d.Get("enable_stack_modify").(bool),
		FullProtection:                 d.Get("full_protection").(bool),
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
	if itemsInterface, ok := d.GetOk(Rulesets); ok {
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
		if value.Key == variable.OnboardingId {
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
