package roles

import (
	"fmt"
	"net/http"
)

const (
	roleResourcePath = "role"
)

type RoleRequest struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Permissions Permissions `json:"permissions"`
}

type RoleResponse struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Permissions Permissions `json:"permissions"`
}

type Permissions struct {
	Access             []string `json:"access"`
	Manage             []string `json:"manage"`
	Rulesets           []string `json:"rulesets"`
	Notifications      []string `json:"notifications"`
	Policies           []string `json:"policies"`
	AlertActions       []string `json:"alertActions"`
	Create             []string `json:"create"`
	View               []string `json:"view"`
	OnBoarding         []string `json:"onBoarding"`
	CrossAccountAccess []string `json:"crossAccountAccess"`
}

func (service *Service) GetAll() (*[]RoleResponse, *http.Response, error) {
	v := new([]RoleResponse)
	resp, err := service.Client.NewRequestDoRetry("GET", roleResourcePath, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Get(id string) (*RoleResponse, *http.Response, error) {
	v := new(RoleResponse)
	relativeURL := fmt.Sprintf("%s/%s", roleResourcePath, id)
	resp, err := service.Client.NewRequestDoRetry("GET", relativeURL, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(role RoleRequest) (*RoleResponse, *http.Response, error) {
	v := new(RoleResponse)
	resp, err := service.Client.NewRequestDoRetry("POST", roleResourcePath, nil, role, &v, nil)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(id string, roleRequest RoleRequest) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", roleResourcePath, id)
	resp, err := service.Client.NewRequestDoRetry("PUT", path, nil, roleRequest, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", roleResourcePath, id)
	resp, err := service.Client.NewRequestDoRetry("DELETE", path, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
