package aws

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
)

type CloudAccountRequest struct {
	Name                   string                  `json:"name"`
	Credentials            CloudAccountCredentials `json:"credentials"`
	FullProtection         bool                    `json:"fullProtection,omitempty"`
	AllowReadOnly          bool                    `json:"allowReadOnly,omitempty"`
	OrganizationalUnitID   string                  `json:"organizationalUnitId,omitempty"`
	OrganizationalUnitPath string                  `json:"organizationalUnitPath,omitempty"`
	OrganizationalUnitName string                  `json:"organizationalUnitName,omitempty"`
	LambdaScanner          bool                    `json:"lambdaScanner,omitempty"`
}

type CloudAccountResponse struct {
	ID                     string                  `json:"id"`
	Vendor                 string                  `json:"vendor"`
	Name                   string                  `json:"name"`
	ExternalAccountNumber  string                  `json:"externalAccountNumber"`
	Error                  string                  `json:"error,omitempty"`
	IsFetchingSuspended    bool                    `json:"isFetchingSuspended"`
	CreationDate           time.Time               `json:"creationDate"`
	Credentials            CloudAccountCredentials `json:"credentials"`
	IamSafe                CloudAccountIamSafe     `json:"iamSafe,omitempty"`
	NetSec                 CloudAccountNetSec      `json:"netSec,omitempty"`
	Magellan               bool                    `json:"magellan"`
	FullProtection         bool                    `json:"fullProtection"`
	AllowReadOnly          bool                    `json:"allowReadOnly"`
	OrganizationalUnitID   string                  `json:"organizationalUnitId,omitempty"`
	OrganizationalUnitPath string                  `json:"organizationalUnitPath"`
	OrganizationalUnitName string                  `json:"organizationalUnitName"`
	LambdaScanner          bool                    `json:"lambdaScanner"`
}

type CloudAccountUpdateRegionConfigRequest struct {
	CloudAccountID        string                   `json:"cloudAccountId,omitempty"`
	ExternalAccountNumber string                   `json:"externalAccountNumber,omitempty"`
	Data                  CloudAccountNetSecRegion `json:"data,omitempty"`
}

type CloudAccountUpdateOrganizationalIDRequest struct {
	OrganizationalUnitId string `json:"organizationalUnitId,omitempty"`
}

type CloudAccountUpdateCredentialsRequest struct {
	CloudAccountID        string                  `json:"cloudAccountId,omitempty"`
	ExternalAccountNumber string                  `json:"externalAccountNumber,omitempty"`
	Data                  CloudAccountCredentials `json:"data,omitempty"`
}

type CloudAccountUpdateNameRequest struct {
	CloudAccountID        string `json:"cloudAccountId,omitempty"`
	ExternalAccountNumber string `json:"externalAccountNumber,omitempty"`
	Data                  string `json:"data,omitempty"`
}

type CloudAccountCredentials struct {
	ApiKey     string `json:"apikey,omitempty"`
	Arn        string `json:"arn,omitempty"`
	Secret     string `json:"secret,omitempty"`
	IamUser    string `json:"iamUser,omitempty"`
	Type       string `json:"type,omitempty"`
	IsReadOnly bool   `json:"isReadOnly,omitempty"`
}

type CloudAccountNetSec struct {
	Regions []CloudAccountNetSecRegion `json:"regions,omitempty"`
}

type CloudAccountNetSecRegion struct {
	Region           string `json:"region,omitempty"`
	Name             string `json:"name,omitempty"`
	Hidden           bool   `json:"hidden,omitempty"`
	NewGroupBehavior string `json:"newGroupBehavior,omitempty"`
}

type CloudAccountIamSafe struct {
	AwsGroupArn           string                  `json:"awsGroupArn,omitempty"`
	AwsPolicyArn          string                  `json:"awsPolicyArn,omitempty"`
	Mode                  string                  `json:"mode,omitempty"`
	State                 string                  `json:"state,omitempty"`
	ExcludedIamEntities   CloudAccountIamEntities `json:"excludedIamEntities,omitempty"`
	RestrictedIamEntities CloudAccountIamEntities `json:"restrictedIamEntities,omitempty"`
}

type CloudAccountIamEntities struct {
	RolesArns []string `json:"rolesArns,omitempty"`
	UsersArns []string `json:"usersArns,omitempty"`
}

func (service *Service) Get(options interface{}) (*CloudAccountResponse, *http.Response, error) {
	if options == nil {
		return nil, nil, fmt.Errorf("options parameter must be passed")
	}

	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathAWS, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]CloudAccountResponse, *http.Response, error) {
	v := new([]CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulPathAWS, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body CloudAccountRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("POST", cloudaccounts.RESTfulPathAWS, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAWS, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) UpdateName(body CloudAccountUpdateNameRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAWS, cloudaccounts.RESTfulServicePathAWSName)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateRegionConfig(body CloudAccountUpdateRegionConfigRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAWS, cloudaccounts.RESTfulServicePathAWSRegionConfig)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateOrganizationalID(id string, body CloudAccountUpdateOrganizationalIDRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathAWS, id, cloudaccounts.RESTfulServicePathAWSOrganizationalUnit)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateCredentials(body CloudAccountUpdateCredentialsRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulPathAWS, cloudaccounts.RESTfulServicePathAWSCredentials)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
