package awp_aws_onboarding

import (
	"fmt"
	"log"
	"net/http"
	"time"
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
	CrossAccountRoleName       string                    `json:"crossAccountRoleName"`
	CrossAccountRoleExternalId string                    `json:"crossAccountRoleExternalId"`
	ScanMode                   string                    `json:"scanMode"`
	IsTerraform                bool                      `json:"isTerraform"`
	AgentlessAccountSettings   *AgentlessAccountSettings `json:"agentlessAccountSettings"`
}

type AccountIssues struct {
	Regions map[string]interface{}  `json:"regions"`
	Account *map[string]interface{} `json:"account"`
}

type GetAWPOnboardingResponse struct {
	AgentlessAccountSettings        *AgentlessAccountSettings `json:"agentlessAccountSettings"`
	MissingAwpPrivateNetworkRegions *[]string                 `json:"missingAwpPrivateNetworkRegions"`
	AccountIssues                   *AccountIssues            `json:"accountIssues"`
	CloudAccountId                  string                    `json:"cloudAccountId"`
	AgentlessProtectionEnabled      bool                      `json:"agentlessProtectionEnabled"`
	ScanMode                        string                    `json:"scanMode"`
	Provider                        string                    `json:"provider"`
	ShouldUpdate                    bool                      `json:"shouldUpdate"`
	IsOrgOnboarding                 bool                      `json:"isOrgOnboarding"`
	CentralizedCloudAccountId       string                    `json:"centralizedCloudAccountId"`
}

type CreateOptions struct {
	ShouldCreatePolicy string `url:"shouldCreatePolicy"`
}

type DeleteOptions struct {
	ForceDelete string `url:"forceDelete"`
}

func (service *Service) CreateAWPOnboarding(id string, req CreateAWPOnboardingRequest, queryParams CreateOptions) (*http.Response, error) {
	// Define the maximum number of retries and the interval between retries
	maxRetries := 3
	retryInterval := time.Second * 5

	// Create the base path
	basePath := fmt.Sprintf("%s/%s/enable", awsOnboardingResourcePath, id)

	// Initialize the response and error variables outside the loop
	var resp *http.Response
	var err error

	// Attempt the request up to maxRetries times
	for i := 0; i < maxRetries; i++ {
		// Make the request
		resp, err = service.Client.NewRequestDo("POST", basePath, queryParams, req, nil)
		if err == nil {
			// If the request was successful, return the response
			return resp, nil
		}

		// If the request failed with a 404 status code, wait for the retry interval before trying again
		if resp != nil && resp.StatusCode == 404 {
			time.Sleep(retryInterval)
		} else {
			// If the status code is not 404, return the response and error immediately
			return resp, err
		}
	}

	// If the function hasn't returned after maxRetries, return an error
	return nil, fmt.Errorf("failed to create AWP Onboarding after %d attempts: %w", maxRetries, err)
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

func (service *Service) DeleteAWPOnboarding(id string, queryParams DeleteOptions) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", awsOnboardingResourcePath, id)
	resp, err := service.Client.NewRequestDo("DELETE", path, queryParams, nil, nil)
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

func (service *Service) UpdateAWPSettings(cloudProvider, id string, req AgentlessAccountSettings) (*http.Response, error) {
	// Construct the URL path
	path := fmt.Sprintf("workload/agentless/%s/accounts/%s/settings", cloudProvider, id)
	// Make a PATCH request with the JSON body
	resp, err := service.Client.NewRequestDo("PATCH", path, nil, req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
