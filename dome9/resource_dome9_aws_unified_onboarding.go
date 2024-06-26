package dome9

import (
	"encoding/json"
	"github.com/dome9/dome9-sdk-go/services/unifiedonboarding/aws_unified_onboarding"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
	"log"
	"strconv"
)

func resourceAwsUnifiedOnboarding() *schema.Resource {
	return &schema.Resource{
		Create: resourceUnifiedOnboardingCreate,
		Read:   unifiedOnboardingResourceRead,
		Update: resourceUnifiedOnboardingUpdate,
		Delete: resourceUnifiedOnboardingDelete,
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
							Elem:     &schema.Schema{Type: schema.TypeInt},
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
							Elem:     &schema.Schema{Type: schema.TypeInt},
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
	req := expandAwsUnifiedOnboardingRequest(d)
	resp, _, err := d9Client.awsUnifiedOnboarding.Create(req)
	if err != nil {
		return err
	}

	addOnboardingIdAsSchemaId(d, resp)

	_ = d.Set(providerconst.StackName, resp.StackName)
	_ = d.Set(providerconst.Parameters, convertParametersFromListToMap(resp))
	_ = d.Set(providerconst.IamCapabilities, resp.IamCapabilities)
	_ = d.Set(providerconst.TemplateUrl, resp.TemplateUrl)

	return nil
}

func convertParametersFromListToMap(responce *aws_unified_onboarding.UnifiedOnboardingConfigurationResponse) map[string]string {

	parameters := responce.Parameters
	mapOfParameters := map[string]string{}
	for _, v := range parameters {
		mapOfParameters[v.Key] = v.Value
	}

	return mapOfParameters
}

func expandAwsUnifiedOnboardingRequest(d *schema.ResourceData) aws_unified_onboarding.UnifiedOnboardingRequest {

	return aws_unified_onboarding.UnifiedOnboardingRequest{
		CloudVendor:                    d.Get(providerconst.CloudVendor).(string),
		OnboardType:                    d.Get(providerconst.OnboardType).(string),
		EnableStackModify:              d.Get(providerconst.EnableStackModify).(bool),
		FullProtection:                 d.Get(providerconst.FullProtection).(bool),
		PostureManagementConfiguration: expendPostureManagementConfiguration(d),
		ServerlessConfiguration:        expendServerlessConfiguration(d),
		IntelligenceConfigurations:     expendIntelligenceConfigurations(d),
	}
}

func expendIntelligenceConfigurations(d *schema.ResourceData) aws_unified_onboarding.IntelligenceConfigurations {
	var intelligenceConfigurations aws_unified_onboarding.IntelligenceConfigurations
	configuration := d.Get("intelligence_configurations").(map[string]interface{})
	intelligenceConfigurations.Enabled = getEnabledFromMap(configuration)
	intelligenceConfigurations.Rulesets = *getRulesetsFromMap(configuration)

	return intelligenceConfigurations
}

func getEnabledFromMap(configurations map[string]interface{}) bool {
	b := false
	if len(configurations) > 0 {
		enabled := configurations[providerconst.Enabled]

		if enabled != "" && enabled != nil {
			v := enabled.(string)
			b, _ = strconv.ParseBool(v)
		}
	}
	return b
}

func expendServerlessConfiguration(d *schema.ResourceData) aws_unified_onboarding.ServerlessConfiguration {
	var serverlessConfiguration aws_unified_onboarding.ServerlessConfiguration
	serverlessConfiguration.Enabled = getEnabledFromMap(d.Get("serverless_configuration").(map[string]interface{}))

	return serverlessConfiguration
}

func expendPostureManagementConfiguration(d *schema.ResourceData) aws_unified_onboarding.PostureManagementConfiguration {
	var postureManagementConfiguration aws_unified_onboarding.PostureManagementConfiguration
	postureManagementConfiguration.Rulesets = *getRulesetsFromMap(d.Get("posture_management_configuration").(map[string]interface{}))
	return postureManagementConfiguration
}

func getRulesetsFromMap(m map[string]interface{}) *[]int {
	var rulesets []int

	if m == nil || m[providerconst.Rulesets] == nil {
		rulesets = make([]int, 0)
		return &rulesets
	}

	RulesetsAsString := m[providerconst.Rulesets].(string)
	err := json.Unmarshal([]byte(RulesetsAsString), &rulesets)
	if err != nil {
		rulesets = make([]int, 0)
	}

	return &rulesets
}

func addOnboardingIdAsSchemaId(d *schema.ResourceData, resp *aws_unified_onboarding.UnifiedOnboardingConfigurationResponse) {
	var p = resp.Parameters
	var schemaId string
	for _, value := range p {
		if value.Key == "OnboardingId" {
			schemaId = value.Value
		}
	}

	if len(schemaId) > 0 {
		d.SetId(schemaId)
	}
}

func resourceUnifiedOnboardingDelete(data *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	receivedAwsUnifiedOnboardingResponse, _, err := d9Client.awsUnifiedOnboarding.Get(data.Id())
	if err != nil {
		return err
	}

	// in case of a failing onboarding we won't actually run the deletion 
	if receivedAwsUnifiedOnboardingResponse.EnvironmentId == "00000000-0000-0000-0000-000000000000" {
		return nil
	}

	log.Printf("[INFO] Deleting AWS Cloud Account ID: %v\n", data.Id())
	if _, err := d9Client.awsUnifiedOnboarding.ForceDelete(receivedAwsUnifiedOnboardingResponse.EnvironmentId); err != nil {
		return err
	}

	return nil
}

func resourceUnifiedOnboardingUpdate(data *schema.ResourceData, i interface{}) error {
	return nil
}

func unifiedOnboardingResourceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
