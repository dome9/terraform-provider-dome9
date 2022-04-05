package dome9

import (
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/unifiedOnbording/awsUnifiedOnbording"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

const (
	Rulesets     = "rulesets"
	Enabled      = "enabled"
	OnboardingId = "OnboardingId"
)

func resourceUnifiedOnbording() *schema.Resource {
	return &schema.Resource{
		Create: resourceUnifiedOnbordingCreate,
		Read:   resourceUnifiedOnbordingRead,
		Update: resourceUnifiedOnbordingUpdat,
		Delete: resourceUnifiedOnbordingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"onboard_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"full_protection": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"cloud_vendor": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enable_stack_modify": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"posture_management_configuration": {
				Type:     schema.TypeMap,
				Required: true,
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
				Required: true,
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
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Rulesets: {
							Type:     schema.TypeList,
							Required: true,
						},
						Enabled: {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceUnifiedOnbordingDelete(data *schema.ResourceData, i interface{}) error {
	return nil
}

func resourceUnifiedOnbordingUpdat(data *schema.ResourceData, i interface{}) error {
	return nil
}

func resourceUnifiedOnbordingRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.AwsUnifiedOnbording.Get(d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing rule set %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	log.Printf("[INFO] Getting Unified Onbording:\n%+v\n", resp)
	_ = d.Set("stack_name", resp.StackName)
	_ = d.Set("parameters", resp.Parameters)
	_ = d.Set("iam_capabilities", resp.IamCapabilities)
	_ = d.Set("template_url", resp.TemplateUrl)

	return nil
}

func resourceUnifiedOnbordingCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandAwsUnifiedOnbordingRequest(d)
	log.Printf("[INFO] Creating Unified Onbording request %+v\n", req)
	resp, _, err := d9Client.AwsUnifiedOnbording.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created UnifiedOnbording resource with stackName: %v\n", resp.StackName)

	AddOnbordingIdAsSchemaId(d, resp)

	return resourceUnifiedOnbordingRead(d, meta)
}

func AddOnbordingIdAsSchemaId(d *schema.ResourceData, resp *awsUnifiedOnbording.UnifiedOnbordingConfigurationResponse) {
	var p = resp.Parameters
	var schemaId string
	for _, value := range p {
		if value.Key == OnboardingId {
			schemaId = value.Value
		}
	}

	if len(schemaId) > 0 {
		d.SetId(schemaId)
	}
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
	intelligenceConfigurations.Enabled = d.Get(Enabled).(bool)
	intelligenceConfigurations.Rulesets = *getRulesets(d)

	return intelligenceConfigurations
}

func expendServerlessConfiguration(d *schema.ResourceData) awsUnifiedOnbording.ServerlessConfiguration {
	var serverlessConfiguration awsUnifiedOnbording.ServerlessConfiguration
	item := d.Get(Enabled).(bool)
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

func resourceContinuousCompliancePolicyCreate2(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandContinuousCompliancePolicyRequest(d)
	log.Printf("[INFO] Creating compliance policy request %+v\n", req)
	resp, _, err := d9Client.continuousCompliancePolicy.Create(&req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created compliance policy with ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceContinuousCompliancePolicyRead(d, meta)
}
