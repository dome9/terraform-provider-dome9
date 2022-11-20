package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/assessment"
	"log"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/admissioncontrol/admission_policy"
	"github.com/dome9/dome9-sdk-go/services/imageassurance/imageassurance_policy"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/alibaba"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/azure"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/gcp"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/k8s"
	"github.com/dome9/dome9-sdk-go/services/cloudsecuritygroup/securitygroupaws"
	"github.com/dome9/dome9-sdk-go/services/cloudsecuritygroup/securitygroupazure"
	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_notification"
	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_policy"
	"github.com/dome9/dome9-sdk-go/services/iplist"
	"github.com/dome9/dome9-sdk-go/services/organizationalunits"
	"github.com/dome9/dome9-sdk-go/services/roles"
	"github.com/dome9/dome9-sdk-go/services/rulebundles"
	"github.com/dome9/dome9-sdk-go/services/serviceaccounts"
	"github.com/dome9/dome9-sdk-go/services/unifiedonboarding/aws_unified_onboarding"
	"github.com/dome9/dome9-sdk-go/services/users"
)

func init() {
	// remove timestamp from Dome9 provider logger, use the timestamp from the default terraform logger
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

type Client struct {
	iplist                           iplist.Service
	cloudaccountAWS                  aws.Service
	cloudaccountAzure                azure.Service
	cloudaccountGCP                  gcp.Service
	cloudaccountKubernetes           k8s.Service
	cloudaccountAlibaba              alibaba.Service
	continuousCompliancePolicy       continuous_compliance_policy.Service
	continuousComplianceNotification continuous_compliance_notification.Service
	ruleSet                          rulebundles.Service
	awsSecurityGroup                 securitygroupaws.Service
	role                             roles.Service
	organizationalUnit               organizationalunits.Service
	azureSecurityGroup               securitygroupazure.Service
	users                            users.Service
	serviceAccounts                  serviceaccounts.Service
	awsUnifiedOnboarding             aws_unified_onboarding.Service
	admissionControlPolicy           admission_policy.Service
	imageAssurancePolicy             imageassurance_policy.Service
	assessment                       assessment.Service
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
		iplist:                           *iplist.New(config),
		cloudaccountAlibaba:              *alibaba.New(config),
		cloudaccountAWS:                  *aws.New(config),
		cloudaccountAzure:                *azure.New(config),
		cloudaccountGCP:                  *gcp.New(config),
		cloudaccountKubernetes:           *k8s.New(config),
		continuousCompliancePolicy:       *continuous_compliance_policy.New(config),
		continuousComplianceNotification: *continuous_compliance_notification.New(config),
		ruleSet:                          *rulebundles.New(config),
		awsSecurityGroup:                 *securitygroupaws.New(config),
		role:                             *roles.New(config),
		organizationalUnit:               *organizationalunits.New(config),
		azureSecurityGroup:               *securitygroupazure.New(config),
		users:                            *users.New(config),
		serviceAccounts:                  *serviceaccounts.New(config),
		awsUnifiedOnboarding:             *aws_unified_onboarding.New(config),
		admissionControlPolicy:           *admission_policy.New(config),
		imageAssurancePolicy:             *imageassurance_policy.New(config),
		assessment:                       *assessment.New(config),
	}

	log.Println("[INFO] initialized Dome9 client")
	return client, nil
}
