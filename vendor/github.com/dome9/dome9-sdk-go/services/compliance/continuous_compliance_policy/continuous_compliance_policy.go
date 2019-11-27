package continuous_compliance_policy

import (
	"fmt"
	"net/http"
)

const (
	continuousComplianceResourcePath = "Compliance/ContinuousCompliancePolicy"
)

type ContinuousCompliancePolicyRequest struct {
	CloudAccountID    string   `json:"cloudAccountId"`
	ExternalAccountID string   `json:"externalAccountId"`
	CloudAccountType  string   `json:"cloudAccountType,omitempty"`
	BundleID          int      `json:"bundleId"`
	NotificationIds   []string `json:"notificationIds"`
}

type ContinuousCompliancePolicyResponse struct {
	ID                string   `json:"id"`
	CloudAccountID    string   `json:"cloudAccountId"`
	ExternalAccountID string   `json:"externalAccountId"`
	CloudAccountType  string   `json:"cloudAccountType"`
	BundleID          int      `json:"bundleId"`
	NotificationIds   []string `json:"notificationIds"`
}

func (service *Service) Get(id string) (*ContinuousCompliancePolicyResponse, *http.Response, error) {
	v := new(ContinuousCompliancePolicyResponse)
	path := fmt.Sprintf("%s/%s", continuousComplianceResourcePath, id)
	resp, err := service.Client.NewRequestDo("GET", path, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]ContinuousCompliancePolicyResponse, *http.Response, error) {
	v := new([]ContinuousCompliancePolicyResponse)
	resp, err := service.Client.NewRequestDo("GET", continuousComplianceResourcePath, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body *ContinuousCompliancePolicyRequest) (*ContinuousCompliancePolicyResponse, *http.Response, error) {
	v := new(ContinuousCompliancePolicyResponse)
	resp, err := service.Client.NewRequestDo("POST", continuousComplianceResourcePath, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(id string, body *ContinuousCompliancePolicyRequest) (*ContinuousCompliancePolicyResponse, *http.Response, error) {
	v := new(ContinuousCompliancePolicyResponse)
	path := fmt.Sprintf("%s/%s", continuousComplianceResourcePath, id)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", continuousComplianceResourcePath, id)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
