package awp_aws_onboarding

import (
	"fmt"
	"log"
	"net/http"
)

const (
	awpAWSGetOnboardingDataPath = "workload/agentless/aws/terraform/onboarding"
	awsOnboardingResourcePath   = "workload/agentless/aws/accounts"
	cloudAccountsPath           = "cloudaccounts/"
)

type AgentlessAwsTerraformOnboardingDataResponse struct {
	Stage                               string `json:"stage"`
	Region                              string `json:"region"`
	CloudGuardBackendAccountId          string `json:"cloudGuardBackendAccountId"`
	AgentlessBucketName                 string `json:"agentlessBucketName"`
	RemoteFunctionsPrefixKey            string `json:"remoteFunctionsPrefixKey"`
	RemoteSnapshotsUtilsFunctionName    string `json:"remoteSnapshotsUtilsFunctionName"`
	RemoteSnapshotsUtilsFunctionRunTime string `json:"remoteSnapshotsUtilsFunctionRunTime"`
	RemoteSnapshotsUtilsFunctionTimeOut int    `json:"remoteSnapshotsUtilsFunctionTimeOut"`
	AwpClientSideSecurityGroupName      string `json:"awpClientSideSecurityGroupName"`
}

type CloudAccountResponse struct {
	ID                     string      `json:"id"`
	Vendor                 string      `json:"vendor"`
	Name                   string      `json:"name"`
	ExternalAccountNumber  string      `json:"externalAccountNumber"`
	Error                  interface{} `json:"error"`
	IsFetchingSuspended    bool        `json:"isFetchingSuspended"`
	CreationDate           string      `json:"creationDate"`
	Credentials            Credentials `json:"credentials"`
	IamSafe                interface{} `json:"iamSafe"`
	NetSec                 NetSec      `json:"netSec"`
	Magellan               bool        `json:"magellan"`
	FullProtection         bool        `json:"fullProtection"`
	AllowReadOnly          bool        `json:"allowReadOnly"`
	OrganizationId         string      `json:"organizationId"`
	OrganizationalUnitId   interface{} `json:"organizationalUnitId"`
	OrganizationalUnitPath string      `json:"organizationalUnitPath"`
	OrganizationalUnitName string      `json:"organizationalUnitName"`
	LambdaScanner          bool        `json:"lambdaScanner"`
	Serverless             Serverless  `json:"serverless"`
	OnboardingMode         string      `json:"onboardingMode"`
}

type Credentials struct {
	Apikey     interface{} `json:"apikey"`
	Arn        string      `json:"arn"`
	Secret     interface{} `json:"secret"`
	IamUser    interface{} `json:"iamUser"`
	Type       string      `json:"type"`
	IsReadOnly bool        `json:"isReadOnly"`
}

type NetSec struct {
	Regions []Region `json:"regions"`
}

type Region struct {
	Region           string `json:"region"`
	Name             string `json:"name"`
	Hidden           bool   `json:"hidden"`
	NewGroupBehavior string `json:"newGroupBehavior"`
}

type Serverless struct {
	CodeAnalyzerEnabled           bool `json:"codeAnalyzerEnabled"`
	CodeDependencyAnalyzerEnabled bool `json:"codeDependencyAnalyzerEnabled"`
}

type AgentlessAccountSettings struct {
	DisabledRegions              []string          `json:"disabledRegions"`
	ScanMachineIntervalInHours   int               `json:"scanMachineIntervalInHours"`
	MaxConcurrenceScansPerRegion int               `json:"maxConcurrenceScansPerRegion"`
	SkipFunctionAppsScan         bool              `json:"skipFunctionAppsScan"`
	CustomTags                   map[string]string `json:"customTags"`
}

type CreateAWPOnboardingRequest struct {
	CrossAccountRoleName       string                   `json:"crossAccountRoleName"`
	CrossAccountRoleExternalId string                   `json:"crossAccountRoleExternalId"`
	ScanMode                   string                   `json:"scanMode"`
	IsTerraform                bool                     `json:"isTerraform"`
	AgentlessAccountSettings   AgentlessAccountSettings `json:"agentlessAccountSettings"`
}

type AccountIssues struct {
	Regions map[string]string `json:"regions"`
	Account map[string]string `json:"account"`
}

type GetAWPOnboardingResponse struct {
	AgentlessAccountSettings        *AgentlessAccountSettings `json:"agentlessAccountSettings"`
	MissingAwpPrivateNetworkRegions []string                  `json:"missingAwpPrivateNetworkRegions"`
	AccountIssues                   *AccountIssues            `json:"accountIssues"`
	CloudAccountId                  string                    `json:"cloudAccountId"`
	AgentlessProtectionEnabled      bool                      `json:"agentlessProtectionEnabled"`
	ScanMode                        string                    `json:"scanMode"`
	Provider                        string                    `json:"provider"`
	ShouldUpdate                    bool                      `json:"shouldUpdate"`
	IsOrgOnboarding                 bool                      `json:"isOrgOnboarding"`
	CentralizedCloudAccountId       string                    `json:"centralizedCloudAccountId"`
}

func (service *Service) CreateAWPOnboarding(id string, req CreateAWPOnboardingRequest) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s/enable", awsOnboardingResourcePath, id)
	resp, err := service.Client.NewRequestDo("POST", path, nil, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (service *Service) GetAWPOnboarding(cloudProvider, id string) (*GetAWPOnboardingResponse, *http.Response, error) {
	v := new(GetAWPOnboardingResponse)
	path := fmt.Sprintf("workload/agentless/%s/accounts/%s", cloudProvider, id)
	resp, err := service.Client.NewRequestDo("GET", path, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) DeleteAWPOnboarding(id string, forceDelete bool) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s?forceDelete=%t", awsOnboardingResourcePath, id, forceDelete)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (service *Service) Get() (*AgentlessAwsTerraformOnboardingDataResponse, *http.Response, error) {
	v := new(AgentlessAwsTerraformOnboardingDataResponse)
	resp, err := service.Client.NewRequestDo("GET", awpAWSGetOnboardingDataPath, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetCloudAccountId(externalAccountId string) (string, *http.Response, error) {
	path := fmt.Sprintf("%s%s", cloudAccountsPath, externalAccountId)
	respData := new(CloudAccountResponse)
	log.Printf("[DEBUG] GetCloudAccountId Path: %s", path)
	resp, err := service.Client.NewRequestDo("GET", path, nil, nil, respData)
	if err != nil {
		return "", nil, err
	}
	return respData.ID, resp, nil
}
