package serviceaccounts

import (
	"fmt"
	"net/http"
	"time"
)

const (
	serviceAccountResourcePath            = "service-account"
	updateServiceAccountResourcePath      = "update"
	generateKeyServiceAccountResourcePath = "generate-key"
)

type ServiceAccountRequest struct {
	Name    string  `json:"name"`
	RoleIds []int64 `json:"roleIds"`
}

type ServiceAccountResponse struct {
	Name         string  `json:"name"`
	Id           string  `json:"id"`
	ApiKeyId     string  `json:"apiKeyId"`
	ApiKeySecret string  `json:"apiKeySecret"`
	RoleIds      []int64 `json:"roleIds"`
}

type UpdateServiceAccountRequest struct {
	Name    string  `json:"name"`
	Id      string  `json:"id"`
	RoleIds []int64 `json:"roleIds"`
}

type GetServiceAccountResponse struct {
	Name        string    `json:"name"`
	Id          string    `json:"id"`
	ApiKeyId    string    `json:"apiKeyId"`
	RoleIds     []int64   `json:"roleIds"`
	DateCreated time.Time `json:"dateCreated"`
	LastUsed    time.Time `json:"lastUsed"`
}

type GenerateKeyRequest struct {
	Id string `json:"id"`
}

type GenerateKeyResponse struct {
	Name         string `json:"name"`
	Id           string `json:"id"`
	ApiKeySecret string `json:"apiKeySecret"`
}

func (service *Service) Create(serviceAccount *ServiceAccountRequest) (*ServiceAccountResponse, *http.Response, error) {
	v := new(ServiceAccountResponse)
	resp, err := service.Client.NewRequestDo("POST", serviceAccountResourcePath, nil, serviceAccount, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(serviceAccount *UpdateServiceAccountRequest) (*ServiceAccountResponse, *http.Response, error) {
	v := new(ServiceAccountResponse)
	path := fmt.Sprintf("%s/%s", serviceAccountResourcePath, updateServiceAccountResourcePath)
	resp, err := service.Client.NewRequestDo("POST", path, nil, serviceAccount, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) GetAll() (*[]GetServiceAccountResponse, *http.Response, error) {
	v := new([]GetServiceAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", serviceAccountResourcePath, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Get(id string) (*GetServiceAccountResponse, *http.Response, error) {
	v := new(GetServiceAccountResponse)
	path := fmt.Sprintf("%s/%s", serviceAccountResourcePath, id)
	resp, err := service.Client.NewRequestDo("GET", path, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", serviceAccountResourcePath, id)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (service *Service) DeleteAll() (*http.Response, error) {
	resp, err := service.Client.NewRequestDo("DELETE", serviceAccountResourcePath, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (service *Service) GenerateKey(req *GenerateKeyRequest) (*GenerateKeyResponse, *http.Response, error) {
	v := new(GenerateKeyResponse)
	path := fmt.Sprintf("%s/%s", serviceAccountResourcePath, generateKeyServiceAccountResourcePath)
	resp, err := service.Client.NewRequestDo("POST", path, nil, req, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}
