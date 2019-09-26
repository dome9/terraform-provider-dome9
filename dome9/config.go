package dome9

import (
	"log"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/azure"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/gcp"
	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_policy"
	"github.com/dome9/dome9-sdk-go/services/iplist"
)

type Client struct {
	iplist                     iplist.Service
	cloudaccountAWS            aws.Service
	cloudaccountAzure          azure.Service
	cloudaccountGCP            gcp.Service
	continuousCompliancePolicy continuous_compliance_policy.Service
}

type Config struct {
	AccessID  string
	SecretKey string
	BaseURL   string
}

func (c *Config) Client() (*Client, error) {

	config, err := dome9.NewConfig(c.AccessID, c.SecretKey, c.BaseURL)

	if err != nil {
		return nil, err
	}

	client := &Client{
		iplist:                     *iplist.New(config),
		cloudaccountAWS:            *aws.New(config),
		cloudaccountAzure:          *azure.New(config),
		cloudaccountGCP:            *gcp.New(config),
		continuousCompliancePolicy: *continuous_compliance_policy.New(config),
	}

	log.Println("initialized client")
	return client, nil
}
