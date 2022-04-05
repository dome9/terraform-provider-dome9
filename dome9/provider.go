package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/resourcetype"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			providerconst.ProviderAccessID: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(providerconst.ProviderAccessIDEnvVariable, nil),
				Description: "dome9 access id",
			},
			providerconst.ProviderSecretKey: {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc(providerconst.ProviderSecretKeyEnvVariable, nil),
				Description: "dome9 api secret",
			},
			providerconst.ProviderBaseURL: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(providerconst.ProviderBaseURL, nil),
				Description: "dome9 base url",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			/*
				terraform resource name: resource schema
				resource formation: provider-resourcename-subresource
			*/
			resourcetype.IPList:                           resourceIpList(),
			resourcetype.CloudAccountAlibaba:              resourceCloudAccountAlibaba(),
			resourcetype.CloudAccountAWS:                  resourceCloudAccountAWS(),
			resourcetype.CloudAccountGCP:                  resourceCloudAccountGCP(),
			resourcetype.CloudAccountAzure:                resourceCloudAccountAzure(),
			resourcetype.CloudAccountKubernetes:           resourceCloudAccountKubernetes(),
			resourcetype.AwsUnifiedOnbording:		       resourceUnifiedOnbording(),
			resourcetype.ContinuousCompliancePolicy:       resourceContinuousCompliancePolicy(),
			resourcetype.ContinuousComplianceNotification: resourceContinuousComplianceNotification(),
			resourcetype.RuleSet:                          resourceRuleSet(),
			resourcetype.CloudAccountAWSSecurityGroup:     resourceCloudSecurityGroupAWS(),
			resourcetype.CloudAccountAWSSecurityGroupRule: resourceCloudSecurityGroupAWSRule(),
			resourcetype.Role:                             resourceRole(),
			resourcetype.OrganizationalUnit:               resourceOrganizationalUnit(),
			resourcetype.CloudAccountAzureSecurityGroup:   resourceAzureSecurityGroup(),
			resourcetype.AttachIAMSafeToAwsCloudAccount:   resourceAttachIAMSafe(),
			resourcetype.User:                             resourceUser(),
			resourcetype.IAMSafeEntity:                    resourceIAMSafeEntity(),
			resourcetype.ServiceAccount:                   resourceServiceAccount(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			// terraform date source name: data source schema
			resourcetype.IPList:                           dataSourceIpList(),
			resourcetype.CloudAccountAlibaba:              dataSourceCloudAccountAlibaba(),
			resourcetype.CloudAccountAWS:                  dataSourceCloudAccountAWS(),
			resourcetype.CloudAccountGCP:                  dataSourceCloudAccountGCP(),
			resourcetype.CloudAccountAzure:                dataSourceCloudAccountAzure(),
			resourcetype.CloudAccountKubernetes:           dataSourceCloudAccountKubernetes(),
			resourcetype.ContinuousCompliancePolicy:       dataSourceContinuousCompliancePolicy(),
			resourcetype.ContinuousComplianceNotification: dataSourceContinuousComplianceNotification(),
			resourcetype.RuleSet:                          dataSourceRuleSet(),
			resourcetype.CloudAccountAWSSecurityGroup:     dataSourceCloudSecurityGroupAWS(),
			resourcetype.CloudAccountAWSSecurityGroupRule: dataSourceCloudSecurityGroupAWSRule(),
			resourcetype.Role:                             dataSourceRole(),
			resourcetype.OrganizationalUnit:               dataSourceOrganizationalUnit(),
			resourcetype.CloudAccountAzureSecurityGroup:   dataSourceSecurityGroupAzure(),
			resourcetype.User:                             dataSourceUser(),
			resourcetype.ServiceAccount:                   dataSourceServiceAccount(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccessID:  d.Get(providerconst.ProviderAccessID).(string),
		SecretKey: d.Get(providerconst.ProviderSecretKey).(string),
		BaseURL:   d.Get(providerconst.ProviderBaseURL).(string),
	}

	return config.Client()
}
