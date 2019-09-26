package dome9

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/dome9/terraform-provider-dome9/dome9/common/providerconst"
	"github.com/dome9/terraform-provider-dome9/dome9/common/resourcetype"
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
			resourcetype.IPList:                     resourceIpList(),
			resourcetype.CloudAccountAWS:            resourceCloudAccountAWS(),
			resourcetype.CloudAccountGCP:            resourceCloudAccountGCP(),
			resourcetype.CloudAccountAzure:          resourceCloudAccountAzure(),
			resourcetype.ContinuousCompliancePolicy: resourceContinuousCompliancePolicy(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			// terraform date source name: data source schema
			resourcetype.IPList:                     dataSourceIpList(),
			resourcetype.CloudAccountAWS:            dataSourceCloudAccountAWS(),
			resourcetype.CloudAccountGCP:            dataSourceCloudAccountGCP(),
			resourcetype.CloudAccountAzure:          dataSourceCloudAccountAzure(),
			resourcetype.ContinuousCompliancePolicy: dataSourceContinuousCompliancePolicy(),
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

	log.Println("initializing dome9 client with config:", config)
	return config.Client()
}
