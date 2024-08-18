package continuous_compliance_policy

import (
	"fmt"
	"net/http"
)

const (
	continuousComplianceResourcePath = "ContinuousCompliancePolicyV2"
)

type ContinuousCompliancePolicyRequest struct {
	TargetId        string   `json:"targetId"`
	TargetType      string   `json:"targetType,omitempty"`
	RulesetId       int      `json:"rulesetId"`
	NotificationIds []string `json:"notificationIds"`
}

type ContinuousCompliancePolicyResponse struct {
	ID               string   `json:"id"`
	TargetType       string   `json:"targetType"`
	TargetInternalId string   `json:"targetInternalId"`
	TargetExternalId string   `json:"targetExternalId"`
	RulesetId        int      `json:"rulesetId"`
	NotificationIds  []string `json:"notificationIds"`
	ErrorMessage     string   `json:"errorMessage"`
}

func (service *Service) Get(id string) (*ContinuousCompliancePolicyResponse, *http.Response, error) {
	v := new(ContinuousCompliancePolicyResponse)
	path := fmt.Sprintf("%s/%s", continuousComplianceResourcePath, id)
	resp, err := service.Client.NewRequestDoRetry("GET", path, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]ContinuousCompliancePolicyResponse, *http.Response, error) {
	v := new([]ContinuousCompliancePolicyResponse)
	resp, err := service.Client.NewRequestDoRetry("GET", continuousComplianceResourcePath, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body *ContinuousCompliancePolicyRequest) (*ContinuousCompliancePolicyResponse, *http.Response, error) {
	v := new([]ContinuousCompliancePolicyResponse)
	resp, err := service.Client.NewRequestDoRetry("POST", continuousComplianceResourcePath, nil, []*ContinuousCompliancePolicyRequest{body}, v, nil)
	if err != nil {
		return nil, nil, err
	}
	policy := new(ContinuousCompliancePolicyResponse)
	if len(*v) > 0 {
		policy = &(*v)[0]
	}
	return policy, resp, nil
}

func (service *Service) Update(body *ContinuousCompliancePolicyRequest) (*ContinuousCompliancePolicyResponse, *http.Response, error) {
	v := new([]ContinuousCompliancePolicyResponse)
	resp, err := service.Client.NewRequestDoRetry("PUT", continuousComplianceResourcePath, nil, []*ContinuousCompliancePolicyRequest{body}, v, nil)
	if err != nil {
		return nil, nil, err
	}
	policy := new(ContinuousCompliancePolicyResponse)
	if len(*v) > 0 {
		policy = &(*v)[0]
	}
	return policy, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", continuousComplianceResourcePath, id)
	resp, err := service.Client.NewRequestDoRetry("DELETE", path, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
