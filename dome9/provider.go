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
			resourcetype.CloudAccountOCI:                  resourceCloudAccountOCI(),
			resourcetype.CloudAccountOCITempData:          resourceCloudAccountOciTempData(),
			resourcetype.CloudAccountGCP:                  resourceCloudAccountGCP(),
			resourcetype.CloudAccountAzure:                resourceCloudAccountAzure(),
			resourcetype.CloudAccountKubernetes:           resourceCloudAccountKubernetes(),
			resourcetype.AwsUnifiedOnboarding:             resourceAwsUnifiedOnboarding(),
			resourcetype.ContinuousCompliancePolicy:       resourceContinuousCompliancePolicy(),
			resourcetype.ContinuousComplianceNotification: resourceContinuousComplianceNotification(),
			resourcetype.Notification:                     resourceNotification(),
			resourcetype.Integration:                      resourceIntegration(),
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
			resourcetype.AdmissionControlPolicy:           resourceAdmissionPolicy(),
			resourcetype.Assessment:                       resourceAssessment(),
			resourcetype.ImageAssurancePolicy:             resourceImageAssurancePolicy(),
			resourcetype.AwpAwsOnboarding:                 resourceAwpAwsOnboarding(),
			resourcetype.AWSOrganizationOnboarding:        resourceAwsOrganizationOnboarding(),
			resourcetype.AzureOrganizationOnboarding:      resourceAzureOrganizationOnboarding(),
			resourcetype.AwpAzureOnboarding:               resourceAwpAzureOnboarding(),
			resourcetype.VulnerabilityPolicy:              resourceVulnerabilityPolicy(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			// terraform date source name: data source schema
			resourcetype.IPList:                                       dataSourceIpList(),
			resourcetype.CloudAccountAlibaba:                          dataSourceCloudAccountAlibaba(),
			resourcetype.CloudAccountAWS:                              dataSourceCloudAccountAWS(),
			resourcetype.CloudAccountOCI:                              dataSourceCloudAccountOCI(),
			resourcetype.AwsUnifiedOnboardingUpdateVersionStackConfig: dataSourceAwsUnifiedOnboardingUpdateVersionStackConfig(),
			resourcetype.AwsUnifiedOnboarding:                         dataSourceAwsUnifiedOnboarding(),
			resourcetype.CloudAccountGCP:                              dataSourceCloudAccountGCP(),
			resourcetype.CloudAccountAzure:                            dataSourceCloudAccountAzure(),
			resourcetype.CloudAccountKubernetes:                       dataSourceCloudAccountKubernetes(),
			resourcetype.ContinuousCompliancePolicy:                   dataSourceContinuousCompliancePolicy(),
			resourcetype.ContinuousComplianceNotification:             dataSourceContinuousComplianceNotification(),
			resourcetype.Notification:                                 dataSourceNotification(),
			resourcetype.Integration:                                  dataSourceIntegration(),
			resourcetype.RuleSet:                                      dataSourceRuleSet(),
			resourcetype.CloudAccountAWSSecurityGroup:                 dataSourceCloudSecurityGroupAWS(),
			resourcetype.CloudAccountAWSSecurityGroupRule:             dataSourceCloudSecurityGroupAWSRule(),
			resourcetype.Role:                                         dataSourceRole(),
			resourcetype.OrganizationalUnit:                           dataSourceOrganizationalUnit(),
			resourcetype.OrganizationalUnitAll:                        dataSourceOrganizationalUnitAll(),
			resourcetype.CloudAccountAzureSecurityGroup:               dataSourceSecurityGroupAzure(),
			resourcetype.User:                                         dataSourceUser(),
			resourcetype.ServiceAccount:                               dataSourceServiceAccount(),
			resourcetype.AdmissionControlPolicy:                       dataSourceAdmissionControlPolicy(),
			resourcetype.Assessment:                                   dataSourceAssessment(),
			resourcetype.ImageAssurancePolicy:                         dataSourceImageAssurancePolicy(),
			resourcetype.AwpAwsOnboardingData:                         dataSourceAwpAwsOnboardingData(),
			resourcetype.AwpAwsOnboarding:                             dataSourceAwpAwsOnboarding(),
			resourcetype.AWSOrganizationOnboarding:                    dataSourceAwsOrganizationOnboarding(),
			resourcetype.AWSOrganizationOnboardingManagementStack:     dataSourceAwsOrganizationOnboardingManagementStack(),
			resourcetype.AWSOrganizationOnboardingMemberAccountConfig: dataSourceAwsOrganizationOnboardingMemberAccountConfiguration(),
			resourcetype.AwpAzureOnboardingData:                       dataSourceAwpAzureOnboardingData(),
			resourcetype.AwpAzureOnboarding:                           dataSourceAwpAzureOnboarding(),
			resourcetype.VulnerabilityPolicy:                          dataSourceVulnerabilityPolicy(),
			resourcetype.AzureOrganizationOnboarding:                  dataSourceAzureOrganizationOnboarding(),
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
