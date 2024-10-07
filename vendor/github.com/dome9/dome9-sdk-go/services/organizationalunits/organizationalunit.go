package organizationalunits

import (
	"fmt"
	"net/http"
	"time"
)

const (
	ouResourcePath = "organizationalunit"
)

type OURequest struct {
	Name     string `json:"name"`
	ParentID string `json:"parentId,omitempty"`
}

type OUResponse struct {
	Item struct {
		AccountID                                    int       `json:"accountId"`
		ID                                           string    `json:"id"`
		Name                                         string    `json:"name"`
		Path                                         string    `json:"path"`
		ParentID                                     string    `json:"parentId"`
		Created                                      time.Time `json:"created"`
		Updated                                      time.Time `json:"updated"`
		AwsCloudAcountsCount                         int       `json:"awsCloudAcountsCount"`
		AzureCloudAccountsCount                      int       `json:"azureCloudAccountsCount"`
		OciCloudAccountsCount                        int       `json:"ociCloudAccountsCount"`
		GoogleCloudAccountsCount                     int       `json:"googleCloudAccountsCount"`
		K8sCloudAccountsCount                        int       `json:"k8sCloudAccountsCount"`
		ShiftLeftCloudAccountsCount                  int       `json:"shiftLeftCloudAccountsCount"`
		AlibabaCloudAccountsCount                    int       `json:"alibabaCloudAccountsCount"`
		ContainerRegistryAccountsCount               int       `json:"containerRegistryAccountsCount"`
		AwsAggregatedCloudAcountsCount               int       `json:"awsAggregatedCloudAcountsCount"`
		AzureAggregateCloudAccountsCount             int       `json:"azureAggregateCloudAccountsCount"`
		OciAggregateCloudAccountsCount               int       `json:"ociAggregateCloudAccountsCount"`
		GoogleAggregateCloudAccountsCount            int       `json:"googleAggregateCloudAccountsCount"`
		K8sAggregateCloudAccountsCount               int       `json:"k8sAggregateCloudAccountsCount"`
		ShiftLeftAggregateCloudAccountsCount         int       `json:"shiftLeftAggregateCloudAccountsCount"`
		AlibabaAggregateCloudAccountsCount           int       `json:"alibabaAggregateCloudAccountsCount"`
		ContainerRegistryAggregateCloudAccountsCount int       `json:"containerRegistryAggregateCloudAccountsCount"`
		SubOrganizationalUnitsCount                  int       `json:"subOrganizationalUnitsCount"`
		IsRoot                                       bool      `json:"isRoot"`
		IsParentRoot                                 bool      `json:"isParentRoot"`
		PathStr                                      string    `json:"pathStr"`
	} `json:"item"`
	ParentID string       `json:"parentId"`
	Children []OUResponse `json:"children"`
}

func (service *Service) Get(ouId string) (*OUResponse, *http.Response, error) {
	v := new(OUResponse)
	path := fmt.Sprintf("%s/%s", ouResourcePath, ouId)
	resp, err := service.Client.NewRequestDoRetry("GET", path, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) GetAll() (*[]OUResponse, *http.Response, error) {
	v := new([]OUResponse)
	resp, err := service.Client.NewRequestDoRetry("GET", ouResourcePath, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(ou *OURequest) (*OUResponse, *http.Response, error) {
	v := new(OUResponse)
	resp, err := service.Client.NewRequestDoRetry("POST", ouResourcePath, nil, ou, &v, nil)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(ouId string, ou *OURequest) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", ouResourcePath, ouId)
	resp, err := service.Client.NewRequestDoRetry("PUT", path, nil, ou, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(ouId string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", ouResourcePath, ouId)
	resp, err := service.Client.NewRequestDoRetry("DELETE", path, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
