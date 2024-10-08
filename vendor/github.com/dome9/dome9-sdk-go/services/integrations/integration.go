package integrations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	RESTfulServicePathIntegration = "integration"
)

type IntegrationType string

const (
	IntegrationTypeSNS                      IntegrationType = "SNS"
	IntegrationTypeEmail                    IntegrationType = "Email"
	IntegrationTypePagerDuty                IntegrationType = "PagerDuty"
	IntegrationTypeAwsSecurityHub           IntegrationType = "AwsSecurityHub"
	IntegrationTypeAzureDefender            IntegrationType = "AzureDefender"
	IntegrationTypeGcpSecurityCommandCenter IntegrationType = "GcpSecurityCommandCenter"
	IntegrationTypeWebhook                  IntegrationType = "Webhook"
	IntegrationTypeServiceNow               IntegrationType = "ServiceNow"
	IntegrationTypeSplunk                   IntegrationType = "Splunk"
	IntegrationTypeJira                     IntegrationType = "Jira"
	IntegrationTypeSumoLogic                IntegrationType = "SumoLogic"
	IntegrationTypeQRadar                   IntegrationType = "QRadar"
	IntegrationTypeSlack                    IntegrationType = "Slack"
	IntegrationTypeTeams                    IntegrationType = "Teams"
)

type IntegrationPostRequestModel struct {
	Name          string          `json:"name" validate:"required"`
	Type          IntegrationType `json:"type" validate:"required"`
	Configuration json.RawMessage `json:"configuration" validate:"required"`
}

func (m IntegrationPostRequestModel) String() string {
	return fmt.Sprintf("Name: %s, Type: %d, Configuration: %s", m.Name, m.Type, string(m.Configuration))
}

type IntegrationUpdateRequestModel struct {
	Id            string          `json:"id" validate:"required"`
	Name          string          `json:"name" validate:"required"`
	Type          IntegrationType `json:"type" validate:"required"`
	Configuration json.RawMessage `json:"configuration" validate:"required"`
}

func (m IntegrationUpdateRequestModel) String() string {
	return fmt.Sprintf("Id: %s, Name: %s, Type: %d, Configuration: %s", m.Id, m.Name, m.Type, string(m.Configuration))
}

type IntegrationViewModel struct {
	Id            string          `json:"id" validate:"required"`
	Name          string          `json:"name" validate:"required"`
	Type          IntegrationType `json:"type" validate:"required"`
	CreatedAt     string          `json:"createdAt"`
	Configuration json.RawMessage `json:"configuration" validate:"required"`
}

func (m IntegrationViewModel) String() string {
	return fmt.Sprintf("Id: %s, Name: %s, Type: %d, CreatedAt: %s, Configuration: %s", m.Id, m.Name, m.Type, m.CreatedAt, string(m.Configuration))
}

// APIs

func (service *Service) Create(body IntegrationPostRequestModel) (*IntegrationViewModel, *http.Response, error) {
	v := new(IntegrationViewModel)
	resp, err := service.Client.NewRequestDoRetry("POST", RESTfulServicePathIntegration, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]IntegrationViewModel, *http.Response, error) {
	v := new([]IntegrationViewModel)
	resp, err := service.Client.NewRequestDoRetry("GET", RESTfulServicePathIntegration, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetById(id string) (*IntegrationViewModel, *http.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id parameter must be passed")
	}

	var resp *http.Response
	var err error

	v := new(IntegrationViewModel)
	relativeURL := fmt.Sprintf("%s/%s", RESTfulServicePathIntegration, id)
	
	resp, err = service.Client.NewRequestDoRetry("GET", relativeURL, nil, nil, v, nil)

	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetByType(integrationType string) (*IntegrationViewModel, *http.Response, error) {
	if integrationType == "" {
		return nil, nil, fmt.Errorf("integrationType parameter must be passed")
	}

	v := new(IntegrationViewModel)
	relativeURL := fmt.Sprintf("%s?type=%s", RESTfulServicePathIntegration, integrationType)
	resp, err := service.Client.NewRequestDoRetry("GET", relativeURL, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(body IntegrationUpdateRequestModel) (*IntegrationViewModel, *http.Response, error) {
	if body.Id == "" {
		return nil, nil, fmt.Errorf("id parameter must be passed")
	}

	v := new(IntegrationViewModel)
	resp, err := service.Client.NewRequestDoRetry("PUT", RESTfulServicePathIntegration, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", RESTfulServicePathIntegration, id)
	var resp *http.Response
	var err error

	resp, err = service.Client.NewRequestDoRetry("DELETE", relativeURL, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
