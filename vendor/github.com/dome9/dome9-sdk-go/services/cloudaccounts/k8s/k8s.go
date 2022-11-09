package k8s

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"net/http"
	"strconv"
	"time"
)

type CloudAccountRequest struct {
	Name                 string `json:"name"`
	OrganizationalUnitID string `json:"organizationalUnitId,omitempty"`
}

type CloudAccountResponse struct {
	ID                        string    `json:"id"` //The k8s cluster ID
	Name                      string    `json:"name"`
	CreationDate              time.Time `json:"creationDate"`
	Vendor                    string    `json:"vendor"`
	OrganizationalUnitID      string    `json:"organizationalUnitId,omitempty"`
	OrganizationalUnitPath    string    `json:"organizationalUnitPath,omitempty"`
	OrganizationalUnitName    string    `json:"organizationalUnitName,omitempty"`
	ClusterVersion            string    `json:"clusterVersion"`
	RuntimeProtectionEnabled  bool      `json:"runtimeProtection"`
	AdmissionControlEnabled   bool      `json:"admissionControl"`
	ImageAssuranceEnabled     bool      `json:"vulnerabilityAssessment"`
	ThreatIntelligenceEnabled bool      `json:"magellan"`
}

type CloudAccountUpdateNameRequest struct {
	Name string `json:"name"`
}

type CloudAccountUpdateOrganizationalIDRequest struct {
	OrganizationalUnitId string `json:"organizationalUnitId,omitempty"`
}

type RuntimeProtectionEnableRequest struct {
	CloudAccountId string `json:"k8sAccountId"`
	Enabled        bool   `json:"enabled"`
}

type AdmissionControlEnableRequest struct {
	CloudAccountId      string `json:"k8sAccountId"`
	Enabled             bool   `json:"enabled"`
	CreateDefaultPolicy bool   `json:"create_default_policy"`
}

type ImageAssuranceEnableRequest struct {
	CloudAccountId string `json:"cloudAccountId"`
	Enabled        bool   `json:"enabled"`
}

type ThreatIntelligenceEnableRequest struct {
	CloudAccountId string `json:"k8sAccountId"`
	Enabled        bool   `json:"enabled"`
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

/*
	runtime-protection
*/

func (service *Service) EnableRuntimeProtection(body RuntimeProtectionEnableRequest) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathK8S, cloudaccounts.RESTfulPathK8SRuntimeProtection, cloudaccounts.RESTfulPathK8sEnable)
	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, body, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
	admission-control
*/

func (service *Service) EnableAdmissionControl(req AdmissionControlEnableRequest) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s/%s/%s", cloudaccounts.RESTfulPathK8SNew, req.CloudAccountId, cloudaccounts.RESTfulPathK8SAdmissionControl,
		IfThenElse(req.Enabled, cloudaccounts.RESTfulPathK8sEnable, cloudaccounts.RESTfulPathK8sDisable))
	if req.Enabled {
		headers := make(http.Header)
		headers.Add(cloudaccounts.CreateDefaultACPolicyHeader, strconv.FormatBool(req.CreateDefaultPolicy))
		service.Client.Config.Headers = headers
	}
	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
	image-assurance
*/

func (service *Service) EnableImageAssurance(body ImageAssuranceEnableRequest) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathK8S, cloudaccounts.RESTfulPathK8SImageAssurance, cloudaccounts.RESTfulPathK8sEnable)
	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, body, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
threat-intelligence
*/
func (service *Service) EnableThreatIntelligence(body ThreatIntelligenceEnableRequest) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathK8S, cloudaccounts.RESTfulPathK8SThreatIntelligence, cloudaccounts.RESTfulPathK8sEnable)
	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, body, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
