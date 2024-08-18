package admission_policy

import (
	"fmt"
	"net/http"
)

const (
	admissionControlPolicyResourcePath = "kubernetes/admissionControl/policy"
)

type AdmissionControlPolicyRequest struct {
	TargetId        string   `json:"targetId"`
	TargetType      string   `json:"targetType,omitempty"`
	RulesetId       int      `json:"rulesetId"`
	NotificationIds []string `json:"notificationIds"`
	Action          string   `json:"action"`
}

type AdmissionControlPolicyResponse struct {
	ID              string   `json:"id"`
	TargetId        string   `json:"targetId"`
	TargetType      string   `json:"targetType"`
	RulesetId       int      `json:"rulesetId"`
	Action          string   `json:"action"`
	NotificationIds []string `json:"notificationIds"`
	ErrorMessage    string   `json:"errorMessage"`
}

func (service *Service) Get(id string) (*AdmissionControlPolicyResponse, *http.Response, error) {
	v := new(AdmissionControlPolicyResponse)
	path := fmt.Sprintf("%s/%s", admissionControlPolicyResourcePath, id)
	resp, err := service.Client.NewRequestDoRetry("GET", path, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]AdmissionControlPolicyResponse, *http.Response, error) {
	v := new([]AdmissionControlPolicyResponse)
	resp, err := service.Client.NewRequestDoRetry("GET", admissionControlPolicyResourcePath, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body *AdmissionControlPolicyRequest) (*AdmissionControlPolicyResponse, *http.Response, error) {
	v := new([]AdmissionControlPolicyResponse)
	resp, err := service.Client.NewRequestDoRetry("POST", admissionControlPolicyResourcePath, nil, []*AdmissionControlPolicyRequest{body}, v, nil)
	if err != nil {
		return nil, nil, err
	}
	policy := new(AdmissionControlPolicyResponse)
	if len(*v) > 0 {
		policy = &(*v)[0]
	}
	return policy, resp, nil
}

func (service *Service) Update(body *AdmissionControlPolicyRequest) (*AdmissionControlPolicyResponse, *http.Response, error) {
	v := new([]AdmissionControlPolicyResponse)
	resp, err := service.Client.NewRequestDoRetry("PUT", admissionControlPolicyResourcePath, nil, []*AdmissionControlPolicyRequest{body}, v, nil)
	if err != nil {
		return nil, nil, err
	}
	policy := new(AdmissionControlPolicyResponse)
	if len(*v) > 0 {
		policy = &(*v)[0]
	}
	return policy, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", admissionControlPolicyResourcePath, id)
	resp, err := service.Client.NewRequestDoRetry("DELETE", path, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
