package imageassurance_policy

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"net/http"
)

const (
	imageAssurancePolicyResourcePath = "kubernetes/imageAssurance/policy"
)

type Service struct {
	Client *client.Client
}

func New(c *dome9.Config) *Service {
	return &Service{Client: client.NewClient(c)}
}

type ImageAssurancePolicyRequest struct {
	TargetId                        string   `json:"targetId"`
	TargetType                      string   `json:"targetType"`
	NotificationIds                 []string `json:"notificationIds"`
	RulesetId                       int      `json:"rulesetId"`
	AdmissionControllerAction       string   `json:"admissionControllerAction,omitempty"`
	AdmissionControlUnScannedAction string   `json:"admissionControlUnScannedAction,omitempty"`
}

type ImageAssurancePolicyResponse struct {
	ID                              string   `json:"id"`
	TargetId                        string   `json:"targetId"`
	TargetType                      string   `json:"targetType"`
	NotificationIds                 []string `json:"notificationIds"`
	RulesetId                       int      `json:"rulesetId"`
	AdmissionControllerAction       string   `json:"admissionControllerAction"`
	AdmissionControlUnScannedAction string   `json:"admissionControlUnScannedAction"`
	ErrorMessage                    string   `json:"errorMessage"`
}

func (service *Service) Get(id string) (*ImageAssurancePolicyResponse, *http.Response, error) {
	v := new(ImageAssurancePolicyResponse)
	path := fmt.Sprintf("%s/%s", imageAssurancePolicyResourcePath, id)
	resp, err := service.Client.NewRequestDoRetry("GET", path, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]ImageAssurancePolicyResponse, *http.Response, error) {
	v := new([]ImageAssurancePolicyResponse)
	resp, err := service.Client.NewRequestDoRetry("GET", imageAssurancePolicyResourcePath, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body *ImageAssurancePolicyRequest) (*ImageAssurancePolicyResponse, *http.Response, error) {
	v := new([]ImageAssurancePolicyResponse)
	resp, err := service.Client.NewRequestDoRetry("POST", imageAssurancePolicyResourcePath, nil, []*ImageAssurancePolicyRequest{body}, v, nil)
	if err != nil {
		return nil, nil, err
	}
	policy := new(ImageAssurancePolicyResponse)
	if len(*v) > 0 {
		policy = &(*v)[0]
	}
	return policy, resp, nil
}

func (service *Service) Update(body *ImageAssurancePolicyRequest) (*ImageAssurancePolicyResponse, *http.Response, error) {
	v := new([]ImageAssurancePolicyResponse)
	resp, err := service.Client.NewRequestDoRetry("PUT", imageAssurancePolicyResourcePath, nil, []*ImageAssurancePolicyRequest{body}, v, nil)
	if err != nil {
		return nil, nil, err
	}
	policy := new(ImageAssurancePolicyResponse)
	if len(*v) > 0 {
		policy = &(*v)[0]
	}
	return policy, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", imageAssurancePolicyResourcePath, id)
	resp, err := service.Client.NewRequestDoRetry("DELETE", path, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
