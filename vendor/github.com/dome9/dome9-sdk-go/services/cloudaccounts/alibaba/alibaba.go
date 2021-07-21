package alibaba

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"net/http"
	"time"
)

type CloudAccountRequest struct {
	Name                                   string             `json:"name,omitempty"`
	AlibabaAccountCredentialsViewPostModel AccountCredentials `json:"alibaba_account_credentials,omitempty"`
}

type AccountCredentials struct {
	AccessKey    string `json:"access_key,omitempty"`
	AccessSecret string `json:"access_secret,omitempty"`
}

type AccountName struct {
	Name string `json:"name,omitempty"`
}

type CloudAccountResponse struct {
	ID                     string             `json:"id"`
	Name                   string             `json:"name"`
	CreationDate           time.Time          `json:"creation_date"`
	AlibabaAccountId       string             `json:"alibaba_account_id"`
	Credentials            AccountCredentials `json:"credentials"`
	OrganizationalUnitID   string             `json:"organizationalUnitId"`
	OrganizationalUnitPath string             `json:"organizationalUnitPath"`
	OrganizationalUnitName string             `json:"organizationalUnitName"`
	Vendor                 string             `json:"vendor"`
}

type OrganizationalUnitId struct {
	OrganizationalUnitId string `json:"organizational_unit_id"`
}

type ID struct {
	id int `json:"id"`
}

func (service *Service) Get(options interface{}) (*CloudAccountResponse, *http.Response, error) {
	if options == nil {
		return nil, nil, fmt.Errorf("options parameter must be passed")
	}
	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathAlibaba, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]CloudAccountResponse, *http.Response, error) {
	v := new([]CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathAlibaba, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body CloudAccountRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("POST", cloudaccounts.RESTfulPathAlibaba, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAlibaba, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) UpdateName(id string, body AccountName) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathAlibaba, id, cloudaccounts.RESTfulServicePathAliName)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateCredentials(id string, body AccountCredentials) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathAlibaba, id, cloudaccounts.RESTfulServicePathAliCredentials)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateOrganizationalID(id string, body OrganizationalUnitId) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathAlibaba, id, cloudaccounts.RESTfulServicePathAliOrganizationalUnit)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
