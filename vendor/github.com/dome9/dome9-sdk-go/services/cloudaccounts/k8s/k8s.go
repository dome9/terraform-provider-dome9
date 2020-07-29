package k8s

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"net/http"
	"time"
)

type CloudAccountRequest struct {
	Name                   string                  `json:"name"`
	OrganizationalUnitID   string                  `json:"organizationalUnitId,omitempty"`
}

type CloudAccountResponse struct {
	ID                     string                  `json:"id"` //The k8s cluster ID
	Name                   string                  `json:"name"`
	CreationDate           time.Time               `json:"creationDate"`
	Vendor                 string                  `json:"vendor"`
	OrganizationalUnitID   string                  `json:"organizationalUnitId,omitempty"`
	OrganizationalUnitPath string                  `json:"organizationalUnitPath,omitempty"`
	OrganizationalUnitName string                  `json:"organizationalUnitName,omitempty"`
	ClusterVersion         string                  `json:"clusterVersion"`
}

type CloudAccountUpdateNameRequest struct {
	Name                   string                  `json:"name"`
}

type CloudAccountUpdateOrganizationalIDRequest struct {
	OrganizationalUnitId   string                  `json:"organizationalUnitId,omitempty"`
}

func (service *Service) Create(body CloudAccountRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("POST", cloudaccounts.RESTfulPathK8S, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Get(id string) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathK8S, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathK8S, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) UpdateName(id string, newNameParam CloudAccountUpdateNameRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathK8S, id, cloudaccounts.RESTfulServicePathK8SName)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, newNameParam, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateOrganizationalID(id string, body CloudAccountUpdateOrganizationalIDRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathK8S, id, cloudaccounts.RESTfulServicePathK8SOrganizationalUnit)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
