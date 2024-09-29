package azure_org

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws_org"
	"net/http"
)

type CloudVendor string

const (
	CloudVendorAzure      CloudVendor = "azure"
	CloudVendorAzureChina CloudVendor = "azurechina"
	CloudVendorAzureGov   CloudVendor = "azuregov"
)

type Blades struct {
	Awp               AwpConfiguration        `json:"awp" validate:"required"`
	Serverless        ServerlessConfiguration `json:"serverless" validate:"required"`
	Cdr               CdrConfiguration        `json:"cdr" validate:"required"`
	PostureManagement PostureManagement       `json:"postureManagement" validate:"required"`
}

type AwpOnboardingMode string

const (
	AwpOnboardingModeSaas         AwpOnboardingMode = "saas"
	AwpOnboardingModeInAccount    AwpOnboardingMode = "inAccount"
	AwpOnboardingModeInAccountHub AwpOnboardingMode = "inAccountHub"
)

type BladeConfiguration struct {
	IsEnabled bool `json:"isEnabled"`
}

type AwpConfiguration struct {
	BladeConfiguration
	OnboardingMode               AwpOnboardingMode `json:"onboardingMode"`
	CentralizedSubscriptionId    string            `json:"centralizedSubscriptionId,omitempty"`
	WithFunctionAppsScan         bool              `json:"withFunctionAppsScan"`
	WithSseCmkEncryptedDisksScan bool              `json:"withSseCmkEncryptedDisksScan"`
}

type ServerlessConfiguration struct {
	BladeConfiguration
}

type StorageAccount struct {
	StorageId string   `json:"storageId"`
	LogTypes  []string `json:"logTypes"`
}

type CdrConfiguration struct {
	BladeConfiguration
	Accounts []StorageAccount `json:"accounts"`
}

type PostureManagement struct {
	OnboardingMode aws_org.OnboardingMode `json:"onboardingMode"`
}

type OnboardingUpdateRequest struct {
	OrganizationName string `json:"organizationName"`
}

type OnboardingRequest struct {
	WorkflowId              string      `json:"workflowId,omitempty"`
	TenantId                string      `json:"tenantId" validate:"required"`
	ManagementGroupId       string      `json:"managementGroupId,omitempty"`
	OrganizationName        string      `json:"organizationName,omitempty"`
	AppRegistrationName     string      `json:"appRegistrationName,omitempty"`
	ClientId                string      `json:"clientId,omitempty"`
	ClientSecret            string      `json:"clientSecret,omitempty"`
	ActiveBlades            Blades      `json:"activeBlades" validate:"required"`
	Vendor                  CloudVendor `json:"vendor" validate:"required,oneof=azure azurechina azuregov"`
	UseCloudGuardManagedApp bool        `json:"useCloudGuardManagedApp"`
	IsAutoOnboarding        bool        `json:"isAutoOnboarding"`
}

type OrganizationManagementViewModel struct {
	Id                      string                                   `json:"id"`
	AccountId               int64                                    `json:"accountId"`
	UserId                  int                                      `json:"userId"`
	OrganizationName        string                                   `json:"organizationName"`
	TenantId                string                                   `json:"tenantId"`
	ManagementGroupId       string                                   `json:"managementGroupId"`
	AppRegistrationName     string                                   `json:"appRegistrationName"`
	OnboardingConfiguration AzureOrganizationOnboardingConfiguration `json:"onboardingConfiguration"`
	UpdateTime              string                                   `json:"updateTime"`
	CreationTime            string                                   `json:"creationTime"`
	IsAutoOnboarding        bool                                     `json:"isAutoOnboarding"`
}

type AzureOrganizationOnboardingConfiguration struct {
	aws_org.OrganizationOnboardingConfigurationBase
	AwpConfiguration        *AwpConfiguration        `json:"awpConfiguration,omitempty"`
	ServerlessConfiguration *ServerlessConfiguration `json:"serverlessConfiguration,omitempty"`
	CdrConfiguration        *CdrConfiguration        `json:"cdrConfiguration,omitempty"`
	IsAutoOnboarding        bool                     `json:"isAutoOnboarding"`
}

type AzureSimplifiedOnboardingExecCmdRequest struct {
	WorkflowId                  string      `json:"workflowId"`
	SubscriptionId              string      `json:"subscriptionId,omitempty"`
	ManagementGroupIdOrTenantId string      `json:"managementGroupIdOrTenantId,omitempty"`
	UseCloudGuardManagedApp     bool        `json:"useCloudGuardManagedApp" validate:"required"`
	ValidatePermission          bool        `json:"validatePermission"`
	AppId                       string      `json:"appId,omitempty"`
	AppName                     string      `json:"appName,omitempty"`
	AccountType                 CloudVendor `json:"accountType" validate:"required,oneof=azure azuregov azurechina"`
	ActiveBlades                Blades      `json:"activeBlades" validate:"required"`
}

func (service *Service) Create(body OnboardingRequest) (*OrganizationManagementViewModel, *http.Response, error) {
	v := new(OrganizationManagementViewModel)
	resp, err := service.Client.NewRequestDoRetry("POST", cloudaccounts.RESTfulServicePathAzureOrgMgmt, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateOrganizationManagementAsync(id string, body OnboardingUpdateRequest) (*http.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id parameter must be passed")
	}

	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulServicePathAzureOrgMgmt, id)
	resp, err := service.Client.NewRequestDoRetry("PUT", relativeURL, nil, body, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulServicePathAzureOrgMgmt, id)
	resp, err := service.Client.NewRequestDoRetry("DELETE", relativeURL, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) Get(id string) (*OrganizationManagementViewModel, *http.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id parameter must be passed")
	}

	v := new(OrganizationManagementViewModel)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulServicePathAzureOrgMgmt, id)
	resp, err := service.Client.NewRequestDoRetry("GET", relativeURL, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]OrganizationManagementViewModel, *http.Response, error) {
	v := new([]OrganizationManagementViewModel)
	resp, err := service.Client.NewRequestDoRetry("GET", cloudaccounts.RESTfulServicePathAzureOrgMgmt, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GenerateOnboardingExecutionCommand(body AzureSimplifiedOnboardingExecCmdRequest) (*string, *http.Response, error) {
	v := new(string)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAzure, cloudaccounts.RESTfulServicePathAzureOnboardingExecutionCommand)

	resp, err := service.Client.NewRequestDoRetry("POST", relativeURL, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
