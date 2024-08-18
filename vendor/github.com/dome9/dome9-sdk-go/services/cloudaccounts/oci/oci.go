package oci

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"net/http"
	"time"
)

type CloudAccountRequestTempData struct {
	Name                            string `json:"name"`
	TenancyId                       string `json:"tenancyId"`
	HomeRegion                      string `json:"homeRegion"`
	TenantAdministratorEmailAddress string `json:"tenantAdministratorEmailAddress"`
}

type CloudAccountRequest struct {
	UserOcid             string `json:"userOcid"`
	TenancyId            string `json:"tenancyId"`
	OrganizationalUnitID string `json:"organizationalUnitId"`
}

type CloudAccountResponse struct {
	ID                     string                          `json:"id"`
	Name                   string                          `json:"name"`
	CreationDate           time.Time                       `json:"creationDate"`
	TenancyId              string                          `json:"tenancyId"`
	HomeRegion             string                          `json:"homeRegion"`
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
	User        string `json:"user,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	PublicKey   string `json:"publicKey,omitempty"`
}

type CloudAccountUpdateNameRequest struct {
	Name string `json:"name,omitempty"`
}

type CloudAccountUpdateOrganizationalIDRequest struct {
	OrganizationalUnitID string `json:"organizationalUnitId,omitempty"`
}

func (service *Service) GetAll() (*[]CloudAccountResponse, *http.Response, error) {
	v := new([]CloudAccountResponse)
	resp, err := service.Client.NewRequestDoRetry("GET", cloudaccounts.RESTfulPathOci, nil, nil, v, nil)
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
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathOci, id)
	resp, err := service.Client.NewRequestDoRetry("GET", relativeURL, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body CloudAccountRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDoRetry("POST", cloudaccounts.RESTfulPathOci, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) CreateTempData(body CloudAccountRequestTempData) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathOci, cloudaccounts.RESTfulServicePathOciTempData)
	resp, err := service.Client.NewRequestDoRetry("POST", relativeURL, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathOci, id)
	resp, err := service.Client.NewRequestDoRetry("DELETE", relativeURL, nil, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) UpdateOrganizationalID(id string, body CloudAccountUpdateOrganizationalIDRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathOci, id, cloudaccounts.RESTfulServicePathOciOrganizationalUnit)
	resp, err := service.Client.NewRequestDoRetry("PUT", relativeURL, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
