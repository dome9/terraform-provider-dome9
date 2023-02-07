package securitygroupazure

import (
	"fmt"
	"net/http"
)

const (
	azureSgResourcePath = "AzureSecurityGroupPolicy"
)

type AzureSecurityGroupRequest struct {
	Name              string         `json:"name"`
	Region            string         `json:"region"`
	ResourceGroup     string         `json:"resourceGroup"`
	CloudAccountID    string         `json:"cloudAccountId"`
	Description       string         `json:"description"`
	IsTamperProtected bool           `json:"isTamperProtected"`
	Tags              []Tags         `json:"tags,omitempty"`
	InboundServices   []BoundService `json:"inboundServices,omitempty"`
	OutboundServices  []BoundService `json:"outboundServices,omitempty"`
}

type AzureSecurityGroupResponse struct {
	ID                string         `json:"id"`
	Name              string         `json:"name"`
	Region            string         `json:"region"`
	ResourceGroup     string         `json:"resourceGroup"`
	CloudAccountID    string         `json:"cloudAccountId"`
	Description       string         `json:"description"`
	IsTamperProtected bool           `json:"isTamperProtected"`
	Tags              []Tags         `json:"tags"`
	InboundServices   []BoundService `json:"inboundServices"`
	OutboundServices  []BoundService `json:"outboundServices"`

	ExternalSecurityGroupID string `json:"externalSecurityGroupId"`
	AccountID               int    `json:"accountId"`
	CloudAccountName        string `json:"cloudAccountName"`
	LastUpdatedByDome9      bool   `json:"lastUpdatedByDome9"`
	Error                   Error  `json:"error"`
}

type Tags struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Scope struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type BoundService struct {
	Name                  string   `json:"name"`
	Description           string   `json:"description"`
	Priority              int      `json:"priority"`
	Access                string   `json:"access"`
	Protocol              string   `json:"protocol"`
	Direction             string   `json:"direction"`
	SourcePortRanges      []string `json:"sourcePortRanges"`
	SourceScopes          []Scope  `json:"sourceScopes"`
	DestinationPortRanges []string `json:"destinationPortRanges"`
	DestinationScopes     []Scope  `json:"destinationScopes"`
	IsDefault             bool     `json:"isDefault"`
}

type Error struct {
	Action       string `json:"action"`
	ErrorMessage string `json:"errorMessage"`
}

func (service *Service) Get(id string) (*AzureSecurityGroupResponse, *http.Response, error) {
	v := new(AzureSecurityGroupResponse)
	relativeURL := fmt.Sprintf("%s/%s", azureSgResourcePath, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]AzureSecurityGroupResponse, *http.Response, error) {
	v := new([]AzureSecurityGroupResponse)
	resp, err := service.Client.NewRequestDo("GET", azureSgResourcePath, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body AzureSecurityGroupRequest) (*AzureSecurityGroupResponse, *http.Response, error) {
	v := new(AzureSecurityGroupResponse)
	resp, err := service.Client.NewRequestDo("POST", azureSgResourcePath, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", azureSgResourcePath, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) Update(id string, body AzureSecurityGroupRequest) (*AzureSecurityGroupResponse, *http.Response, error) {
	v := new(AzureSecurityGroupResponse)
	relativeURL := fmt.Sprintf("%s/%s", azureSgResourcePath, id)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
