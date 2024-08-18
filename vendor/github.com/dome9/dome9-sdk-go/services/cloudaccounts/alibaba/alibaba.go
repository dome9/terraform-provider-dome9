package alibaba

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"net/http"
	"time"
)

type CloudAccountRequest struct {
	Name                 string                         `json:"name,omitempty"`
	Credentials          CloudAccountCredentialsRequest `json:"credentials,omitempty"`
	OrganizationalUnitID string                         `json:"organizationalUnitId,omitempty"`
}

type CloudAccountResponse struct {
	ID                     string                          `json:"id"`
	Name                   string                          `json:"name"`
	CreationDate           time.Time                       `json:"creationDate"`
	AlibabaAccountId       string                          `json:"alibabaAccountId"`
	Credentials            CloudAccountCredentialsResponse `json:"credentials"`
	OrganizationalUnitID   string                          `json:"organizationalUnitId,omitempty"`
	OrganizationalUnitPath string                          `json:"organizationalUnitPath"`
	OrganizationalUnitName string                          `json:"organizationalUnitName"`
	Vendor                 string                          `json:"vendor"`
}

type CloudAccountCredentialsRequest struct {
	AccessKey    string `json:"accessKey,omitempty"`
	AccessSecret string `json:"accessSecret,omitempty"`
}

type CloudAccountCredentialsResponse struct {
	AccessKey string `json:"accessKey,omitempty"`
}

type CloudAccountUpdateNameRequest struct {
	Name string `json:"name,omitempty"`
}

type CloudAccountUpdateOrganizationalIDRequest struct {
	OrganizationalUnitID string `json:"organizationalUnitId,omitempty"`
}

func (service *Service) GetAll() (*[]CloudAccountResponse, *http.Response, error) {
	v := new([]CloudAccountResponse)
	resp, err := service.Client.NewRequestDoRetry("GET", cloudaccounts.RESTfulPathAlibaba, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Get(id string) (*CloudAccountResponse, *http.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id parameter must be passed")
	}

	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAlibaba, id)
	resp, err := service.Client.NewRequestDoRetry("GET", relativeURL, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body CloudAccountRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDoRetry("POST", cloudaccounts.RESTfulPathAlibaba, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAlibaba, id)
	resp, err := service.Client.NewRequestDoRetry("DELETE", relativeURL, nil, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) UpdateName(id string, body CloudAccountUpdateNameRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathAlibaba, id, cloudaccounts.RESTfulServicePathAlibabaName)
	resp, err := service.Client.NewRequestDoRetry("PUT", relativeURL, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateOrganizationalID(id string, body CloudAccountUpdateOrganizationalIDRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathAlibaba, id, cloudaccounts.RESTfulServicePathAlibabaOrganizationalUnit)
	resp, err := service.Client.NewRequestDoRetry("PUT", relativeURL, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateCredentials(id string, body CloudAccountCredentialsRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathAlibaba, id, cloudaccounts.RESTfulServicePathAlibabaCredentials)
	resp, err := service.Client.NewRequestDoRetry("PUT", relativeURL, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
