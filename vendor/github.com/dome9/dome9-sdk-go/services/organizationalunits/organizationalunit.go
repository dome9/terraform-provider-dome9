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
		AccountID                         int       `json:"accountId"`
		ID                                string    `json:"id"`
		Name                              string    `json:"name"`
		Path                              string    `json:"path"`
		ParentID                          string    `json:"parentId"`
		Created                           time.Time `json:"created"`
		Updated                           time.Time `json:"updated"`
		AwsCloudAcountsCount              int       `json:"awsCloudAcountsCount"`
		AzureCloudAccountsCount           int       `json:"azureCloudAccountsCount"`
		GoogleCloudAccountsCount          int       `json:"googleCloudAccountsCount"`
		AwsAggregatedCloudAcountsCount    int       `json:"awsAggregatedCloudAcountsCount"`
		AzureAggregateCloudAccountsCount  int       `json:"azureAggregateCloudAccountsCount"`
		GoogleAggregateCloudAccountsCount int       `json:"googleAggregateCloudAccountsCount"`
		SubOrganizationalUnitsCount       int       `json:"subOrganizationalUnitsCount"`
		IsRoot                            bool      `json:"isRoot"`
		IsParentRoot                      bool      `json:"isParentRoot"`
		PathStr                           string    `json:"pathStr"`
	} `json:"item"`
	ParentID string        `json:"parentId"`
	Children []interface{} `json:"children"`
}

func (service *Service) Get(ouId string) (*OUResponse, *http.Response, error) {
	v := new(OUResponse)
	path := fmt.Sprintf("%s/%s", ouResourcePath, ouId)
	resp, err := service.Client.NewRequestDo("GET", path, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) GetAll() (*[]OUResponse, *http.Response, error) {
	v := new([]OUResponse)
	resp, err := service.Client.NewRequestDo("GET", ouResourcePath, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(ou *OURequest) (*OUResponse, *http.Response, error) {
	v := new(OUResponse)
	resp, err := service.Client.NewRequestDo("POST", ouResourcePath, nil, ou, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(ouId string, ou *OURequest) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", ouResourcePath, ouId)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, ou, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(ouId string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", ouResourcePath, ouId)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
