package k8s

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"net/http"
	k8s "github.com/dome9/dome9-sdk-go/services/cloudaccounts/k8s"
)

type Service k8s.Service


func (service *Service) Create(body k8s.CloudAccountRequest) (*k8s.CloudAccountResponse, *http.Response, error) {
	v := new(k8s.CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("POST", cloudaccounts.RESTfulPathK8S, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Get(id string) (*k8s.CloudAccountResponse, *http.Response, error) {
	v := new(k8s.CloudAccountResponse)
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

func (service *Service) UpdateName(id string, newNameParam k8s.CloudAccountUpdateNameRequest) (*k8s.CloudAccountResponse, *http.Response, error) {
	v := new(k8s.CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathK8S, id, cloudaccounts.RESTfulServicePathK8SName)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, newNameParam, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateOrganizationalID(id string, body k8s.CloudAccountUpdateOrganizationalIDRequest) (*k8s.CloudAccountResponse, *http.Response, error) {
	v := new(k8s.CloudAccountResponse)
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

func (service *Service) EnableRuntimeProtection(body k8s.RuntimeProtectionEnableRequest) (*http.Response, error) {
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

func (service *Service) EnableAdmissionControl(body k8s.AdmissionControlEnableRequest) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathK8S, cloudaccounts.RESTfulPathK8SAdmissionControl, cloudaccounts.RESTfulPathK8sEnable)
	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, body, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*
	image-assurance
*/



/*
threat-intelligence
*/
func (service *Service) EnableThreatIntelligence(body k8s.ThreatIntelligenceEnableRequest) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathK8S, cloudaccounts.RESTfulPathK8SThreatIntelligence, cloudaccounts.RESTfulPathK8sEnable)
	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, body, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
