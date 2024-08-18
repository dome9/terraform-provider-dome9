package aws

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
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
	Vendor                 string                  `json:"vendor,omitempty"`
}

type AttachIamSafeRequest struct {
	CloudAccountID string `json:"cloudAccountId"`
	Data           Data   `json:"data"`
}

type RestrictedIamEntitiesRequest struct {
	EntityName string `json:"entityName"` // aws iam user name or aws role
	EntityType string `json:"entityType"` // must be one of the following Role or User
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
	IamSafe                *CloudAccountIamSafe    `json:"iamSafe"`
	NetSec                 CloudAccountNetSec      `json:"netSec,omitempty"`
	Magellan               bool                    `json:"magellan"`
	FullProtection         bool                    `json:"fullProtection"`
	AllowReadOnly          bool                    `json:"allowReadOnly"`
	OrganizationalUnitID   string                  `json:"organizationalUnitId,omitempty"`
	OrganizationalUnitPath string                  `json:"organizationalUnitPath"`
	OrganizationalUnitName string                  `json:"organizationalUnitName"`
	LambdaScanner          bool                    `json:"lambdaScanner"`
}

type ProtectIAMEntitiesResponse struct {
	RolesArn []IAMSafeEntityResponse `json:"rolesArns"`
	UsersArn []IAMSafeEntityResponse `json:"usersArns"`
}

type IAMSafeEntityResponse struct {
	State              string   `json:"state"`
	AttachedDome9Users []string `json:"attachedDome9Users"`
	IsUsedByDome9      bool     `json:"isUsedByDome9"`
	ExistsInAws        bool     `json:"existsInAws"`
	Arn                string   `json:"arn"`
	Name               string   `json:"name"`
	Type               *string  `json:"type"`
}

type CloudAccountCredentials struct {
	ApiKey     string `json:"apikey,omitempty"`
	Arn        string `json:"arn,omitempty"`
	Secret     string `json:"secret,omitempty"`
	IamUser    string `json:"iamUser,omitempty"`
	Type       string `json:"type,omitempty"`
	IsReadOnly bool   `json:"isReadOnly,omitempty"`
}

type Data struct {
	AwsGroupArn  string `json:"AwsGroupArn"`
	AwsPolicyArn string `json:"AwsPolicyArn"`
	Mode         string `json:"Mode,omitempty"`
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

type UnprotectAWSIAMEntityOptions struct {
	EntityName string `json:"entityName"`
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
	RolesArn []string `json:"rolesArns,omitempty"`
	UsersArn []string `json:"usersArns,omitempty"`
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
	ven := body.Vendor
	if ven != "aws" && ven != "awsgov" && ven != "awschina" {
		return nil, nil, errors.New("vendor must be aws/awsgov/awschina")
	}
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

func (service *Service) ForceDelete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathAWS, id, cloudaccounts.DeleteForce)
	var resp *http.Response
	var err error

	for i := 1; i <= 3; i++ {
		resp, err = service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)
		if err == nil || resp == nil || resp.StatusCode <= 400 || resp.StatusCode >= 500 || i == 3 {
			break
		}
		time.Sleep(time.Duration(i) * 2 * time.Second)
	}

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

/*
	attach iam safe to cloud account
*/

func (service *Service) AttachIAMSafeToCloudAccount(body AttachIamSafeRequest) (*CloudAccountResponse, *http.Response, error) {
	v := new(CloudAccountResponse)
	path := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulServicePathAWSCloudAccounts, cloudaccounts.RESTfulServicePathAWSIAMSafe)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, body, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, err
}

func (service *Service) DetachIAMSafeToCloudAccount(id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulServicePathAWSCloudAccounts, id, cloudaccounts.RESTfulServicePathAWSIAMSafe)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

/*
	iam protect (restrict) entity
*/

func (service *Service) ProtectIAMSafeEntity(d9CloudAccountID string, body RestrictedIamEntitiesRequest) (*string, *http.Response, error) {
	// iam entity can be aws iam user or aws role, according the type of the field EntityType inside the body
	var arn string
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathAWS, d9CloudAccountID, cloudaccounts.RESTfulPathRestrictedIamEntities)

	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, body, &arn)
	if err != nil {
		return nil, nil, err
	}

	return &arn, resp, nil
}

func (service *Service) GetAllProtectIAMSafeEntityStatus(d9CloudAccountID string) (*ProtectIAMEntitiesResponse, *http.Response, error) {
	v := new(ProtectIAMEntitiesResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathAWS, d9CloudAccountID, cloudaccounts.RESTfulPathIAM)

	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetProtectIAMSafeEntityStatusByName(d9CloudAccountID, entityName, entityType string) (*IAMSafeEntityResponse, error) {
	var iamEntities []IAMSafeEntityResponse

	protectAWSIAMEntitiesStatus, _, err := service.GetAllProtectIAMSafeEntityStatus(d9CloudAccountID)
	if err != nil {
		return nil, err
	}

	if strings.EqualFold(entityType, cloudaccounts.RESTfulPathUser) {
		iamEntities = protectAWSIAMEntitiesStatus.UsersArn
	} else {
		iamEntities = protectAWSIAMEntitiesStatus.RolesArn
	}

	for _, arn := range iamEntities {
		if arn.Name == entityName {
			return &arn, nil
		}
	}

	errMsg := fmt.Sprintf("There is no aws IAM entity with %s name %s", entityType, entityName)
	return nil, errors.New(errMsg)
}

func (service *Service) UnprotectIAMSafeEntity(d9CloudAccountID, entityName, entityType string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s/%s/%s", cloudaccounts.RESTfulPathAWS, d9CloudAccountID, cloudaccounts.RESTfulPathRestrictedIamEntities, entityType)
	unprotectAWSIAMEntityOptions := UnprotectAWSIAMEntityOptions{
		EntityName: entityName,
	}

	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, unprotectAWSIAMEntityOptions, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
