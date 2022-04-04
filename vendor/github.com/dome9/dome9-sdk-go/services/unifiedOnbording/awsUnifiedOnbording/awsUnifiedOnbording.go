package awsUnifiedOnbording

import (
	"fmt"
	"net/http"
)

const (
	UnifiedOnboardingResourcePath = "AwsUnifiedOnboarding"
	UpdateVersion                 = "UpdateVersion"
	StackConfig                   = "StackConfig"
)

type PostureManagementConfiguration struct {
	Rulesets []int `json:"rulesets"`
}

type ServerlessConfiguration struct {
	Enabled bool `json:"enabled"`
}
type IntelligenceConfigurations struct {
	Enabled  bool  `json:"enabled"`
	Rulesets []int `json:"rulesets"`
}

type UnifiedOnbordingRequest struct {
	OnboardType                    string                         `json:"onboardType"`
	FullProtection                 bool                           `json:"fullProtection"`
	CloudVendor                    string                         `json:"cloudVendor"`
	EnableStackModify              bool                           `json:"enableStackModify"`
	PostureManagementConfiguration PostureManagementConfiguration `json:"postureManagementConfiguration"`
	ServerlessConfiguration        ServerlessConfiguration        `json:"serverlessConfiguration"`
	IntelligenceConfigurations     IntelligenceConfigurations     `json:"intelligenceConfigurations"`
}

type UnifiedOnbordingConfigurationResponse struct {
	StackName       string     `json:"stackName"`
	TemplateUrl     string     `json:"templateUrl"`
	Parameters      []Parameter `json:"parameters"`
	IamCapabilities []string   `json:"iamCapabilities"`
}

type Parameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UnifiedOnbordingInformationResponse struct {
	OnboardingId            string                  `json:"onboardingId"`
	InitiatedUserName       string                  `json:"initiatedUserName"`
	InitiatedUserId         int                     `json:"initiatedUserId"`
	EnvironmentId           string                  `json:"environmentId"`
	EnvironmentName         string                  `json:"environmentName"`
	EnvironmentExternalId   string                  `json:"environmentExternalId"`
	RootStackId             string                  `json:"rootStackId"`
	CftVersion              string                  `json:"cftVersion"`
	UnifiedOnbordingRequest UnifiedOnbordingRequest `json:"onboardingRequest"`
	Statuses                Statuses                `json:"statuses"`
}

type Statuses []struct {
	Module                    string `json:"module"`
	Feature                   string `json:"feature"`
	Status                    string `json:"status"`
	StatusMessage             string `json:"statusMessage"`
	StackStatus               string `json:"stackStatus"`
	StackMessage              string `json:"stackMessage"`
	RemediationRecommendation string `json:"remediationRecommendation"`
}

func (service *Service) Get(id string) (*UnifiedOnbordingConfigurationResponse, *http.Response, error) {
	v := new(UnifiedOnbordingConfigurationResponse)
	relativeURL := fmt.Sprintf("%s/%s", UnifiedOnboardingResourcePath, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetUpdateStackConfig(id string) (*UnifiedOnbordingConfigurationResponse, *http.Response, error) {
	v := new(UnifiedOnbordingConfigurationResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s/%s", UnifiedOnboardingResourcePath, UpdateVersion, StackConfig, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(onbordingRequest UnifiedOnbordingRequest) (*UnifiedOnbordingConfigurationResponse, *http.Response, error) {
	v := new(UnifiedOnbordingConfigurationResponse)
	relativeURL := fmt.Sprintf("%s/%s", UnifiedOnboardingResourcePath, StackConfig)
	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, onbordingRequest, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}
