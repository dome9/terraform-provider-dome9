package gcp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
)

// refer to API type: GoogleCloudAccountPost
type CloudAccountRequest struct {
	Name                      string                    `json:"name,omitempty"`
	ServiceAccountCredentials ServiceAccountCredentials `json:"serviceAccountCredentials,omitempty"`
	GsuiteUser                string                    `json:"gsuiteUser,omitempty"`
	DomainName                string                    `json:"domainName,omitempty"`
	OrganizationalUnitID      string                    `json:"organizationalUnitId,omitempty"`
}

type CloudAccountResponse struct {
	ID                     string    `json:"id"`
	Name                   string    `json:"name"`
	ProjectID              string    `json:"projectId"`
	CreationDate           time.Time `json:"creationDate"`
	OrganizationalUnitID   string    `json:"organizationalUnitId"`
	OrganizationalUnitPath string    `json:"organizationalUnitPath"`
	OrganizationalUnitName string    `json:"organizationalUnitName"`
	GSuite                 GSuite    `json:"gSuite,omitempty"`
	Vendor                 string    `json:"vendor"`
}

type ServiceAccountCredentials struct {
	Type                    string `json:"type,omitempty"`
	ProjectID               string `json:"project_id,omitempty"`
	PrivateKeyID            string `json:"private_key_id,omitempty"`
	PrivateKey              string `json:"private_key,omitempty"`
	ClientEmail             string `json:"client_email,omitempty"`
	ClientID                string `json:"client_id,omitempty"`
	AuthURI                 string `json:"auth_uri,omitempty"`
	TokenURI                string `json:"token_uri,omitempty"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url,omitempty"`
	ClientX509CertURL       string `json:"client_x509_cert_url,omitempty"`
}

type GSuite struct {
	GSuiteUser string `json:"gSuiteUser"`
	DomainName string `json:"domainName"`
}

type CloudAccountUpdateNameRequest struct {
	Name string `json:"name,omitempty"`
}

type CloudAccountUpdateCredentialsRequest struct {
	Name                      string                    `json:"name,omitempty"`
	ServiceAccountCredentials ServiceAccountCredentials `json:"serviceAccountCredentials,omitempty"`
}

type CloudAccountUpdateOrganizationalIDRequest struct {
	OrganizationalUnitID string `json:"organizationalUnitId"`
}

func (service *Service) Get(options interface{}) (*CloudAccountResponse, *http.Response, error) {
	if options == nil {
		return nil, nil, fmt.Errorf("options parameter must be passed")
	}
	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathGCP, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]CloudAccountResponse, *http.Response, error) {
	v := new([]CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathGCP, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body CloudAccountRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("POST", cloudaccounts.RESTfulPathGCP, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathGCP, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) UpdateName(id string, body CloudAccountUpdateNameRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathGCP, id, cloudaccounts.RESTfulServicePathGCPName)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateAccountGSuite(id string, body GSuite) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathGCP, id, cloudaccounts.RESTfulServicePathGCPCredentialsGSuite)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateCredentials(id string, body CloudAccountUpdateCredentialsRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathGCP, id, cloudaccounts.RESTfulServicePathGCPCredentials)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateOrganizationalID(id string, body CloudAccountUpdateOrganizationalIDRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathGCP, id, cloudaccounts.RESTfulServicePathGCPOrganizationalUnit)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
