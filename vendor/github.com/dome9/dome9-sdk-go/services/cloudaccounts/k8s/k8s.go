package k8s

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"net/http"
	"time"
)

type CloudAccountRequest struct {
	Name                 string `json:"name"`
	OrganizationalUnitID string `json:"organizationalUnitId,omitempty"`
}

type CloudAccountResponse struct {
	ID                              string    `json:"id"` //The k8s cluster ID
	Name                            string    `json:"name"`
	CreationDate                    time.Time `json:"creationDate"`
	Vendor                          string    `json:"vendor"`
	OrganizationalUnitID            string    `json:"organizationalUnitId,omitempty"`
	OrganizationalUnitPath          string    `json:"organizationalUnitPath,omitempty"`
	OrganizationalUnitName          string    `json:"organizationalUnitName,omitempty"`
	ClusterVersion                  string    `json:"clusterVersion"`
	RuntimeProtectionEnabled        bool      `json:"runtimeProtectionEnabled"`
	RuntimeProtectionNetwork        bool      `json:"runtimeProtectionNetwork"`
	RuntimeProtectionProfiling      bool      `json:"runtimeProtectionProfiling"`
	RuntimeProtectionFileReputation bool      `json:"runtimeProtectionFileReputation"`
	AdmissionControlEnabled         bool      `json:"admissionControlEnabled"`
	AdmissionControlFailOpen        bool      `json:"admissionControlFailOpen"`
	ImageAssuranceEnabled           bool      `json:"imageAssuranceEnabled"`
	ThreatIntelligenceEnabled       bool      `json:"threatIntelligenceEnabled"`
	Description                     string    `json:"description"`
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
	CloudAccountId string `json:"k8sAccountId"`
	Enabled        bool   `json:"enabled"`
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
	resp, err := service.Client.NewRequestDoRetry("POST", cloudaccounts.RESTfulPathK8S, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Get(id string) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathK8S, id)
	resp, err := service.Client.NewRequestDoRetry("GET", relativeURL, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathK8S, id)
	resp, err := service.Client.NewRequestDoRetry("DELETE", relativeURL, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) UpdateName(id string, newNameParam CloudAccountUpdateNameRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathK8S, id, cloudaccounts.RESTfulServicePathK8SName)
	resp, err := service.Client.NewRequestDoRetry("PUT", relativeURL, newNameParam, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateOrganizationalID(id string, body CloudAccountUpdateOrganizationalIDRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathK8S, id, cloudaccounts.RESTfulServicePathK8SOrganizationalUnit)
	resp, err := service.Client.NewRequestDoRetry("PUT", relativeURL, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*
	runtime-protection
*/

func GetEnableDisablePath(enabled bool) string {
	if enabled {
		return cloudaccounts.RESTfulPathK8sEnable
	}
	return cloudaccounts.RESTfulPathK8sDisable
}

func (service *Service) EnableRuntimeProtection(body RuntimeProtectionEnableRequest) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s/%s/%s", cloudaccounts.RESTfulPathK8S, body.CloudAccountId, cloudaccounts.RESTfulPathK8SRuntimeProtection, GetEnableDisablePath(body.Enabled))
	resp, err := service.Client.NewRequestDoRetry("POST", relativeURL, nil, body, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
	admission-control
*/

func (service *Service) EnableAdmissionControl(body AdmissionControlEnableRequest) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s/%s/%s", cloudaccounts.RESTfulPathK8S, body.CloudAccountId, cloudaccounts.RESTfulPathK8SAdmissionControl, GetEnableDisablePath(body.Enabled))
	resp, err := service.Client.NewRequestDoRetry("POST", relativeURL, nil, body, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
	image-assurance
*/

func (service *Service) EnableImageAssurance(body ImageAssuranceEnableRequest) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s/%s/%s", cloudaccounts.RESTfulPathK8S, body.CloudAccountId, cloudaccounts.RESTfulPathK8SImageAssurance, GetEnableDisablePath(body.Enabled))
	resp, err := service.Client.NewRequestDoRetry("POST", relativeURL, nil, body, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
threat-intelligence
*/
func (service *Service) EnableThreatIntelligence(body ThreatIntelligenceEnableRequest) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s/%s/%s", cloudaccounts.RESTfulPathK8S, body.CloudAccountId, cloudaccounts.RESTfulPathK8SThreatIntelligence, GetEnableDisablePath(body.Enabled))
	resp, err := service.Client.NewRequestDoRetry("POST", relativeURL, nil, body, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
