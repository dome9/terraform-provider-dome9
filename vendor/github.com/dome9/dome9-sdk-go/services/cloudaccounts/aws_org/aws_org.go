package aws_org

import (
	_ "encoding/json"
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"net/http"
)

type CloudCredentialsType string

const (
	UserBased CloudCredentialsType = "UserBased"
	RoleBased CloudCredentialsType = "RoleBased"
)

type OnboardingPermissionRequest struct {
	RoleArn string               `json:"roleArn" validate:"required,roleArn"`
	Secret  string               `json:"secret" validate:"required,secret"`
	ApiKey  string               `json:"apiKey,omitempty"`
	Type    CloudCredentialsType `json:"type" validate:"required,oneof=UserBased RoleBased"`
}

type ValidateStackSetArnRequest struct {
	OnboardingPermissionRequest
	StackSetArn string `json:"stackSetArn" validate:"required,stackSetArn"`
}

type OnboardingRequest struct {
	ValidateStackSetArnRequest
	AwsOrganizationName string `json:"awsOrganizationName,omitempty"`
	EnableStackModify   bool   `json:"enableStackModify" validate:"required"`
}

type OnboardingUpdateRequest struct {
	AwsOrganizationName string `json:"awsOrganizationName,omitempty"`
	EnableStackModify   bool   `json:"enableStackModify"`
}

type UpdateConfigurationRequest struct {
	OrganizationRootOuId string                         `json:"organizationRootOuId" validate:"required"`
	MappingStrategy      MappingStrategyType            `json:"mappingStrategy" validate:"required"`
	PostureManagement    PostureManagementConfiguration `json:"postureManagement" validate:"required"`
}

type MappingStrategyType string
type OnboardingMode string

const (
	Flat  MappingStrategyType = "Flat"
	Clone MappingStrategyType = "Clone"

	Read   OnboardingMode = "Read"
	Manage OnboardingMode = "Manage"
)

type PostureManagementConfiguration struct {
	RulesetsIds    []int64        `json:"rulesetsIds"`
	OnboardingMode OnboardingMode `json:"onboardingMode"`
}

type UpdateStackSetArnRequest struct {
	StackSetArn string `json:"stackSetArn" validate:"required,stackSetArn"`
}

type OrganizationOnboardingConfigurationBase struct {
	OrganizationRootOuId string                         `json:"organizationRootOuId,omitempty"`
	MappingStrategy      MappingStrategyType            `json:"mappingStrategy"`
	PostureManagement    PostureManagementConfiguration `json:"postureManagement"`
}

type AwsOrganizationOnboardingConfiguration struct {
	OrganizationOnboardingConfigurationBase
}

type OnboardingCftBase struct {
	ExternalId string `json:"externalId"`
	Content    string `json:"content"`
}

type ManagementCftConfiguration struct {
	OnboardingCftBase
	ManagementCftUrl      string `json:"managementCftUrl"`
	IsManagementOnboarded bool   `json:"isManagementOnboarded"`
}

type OnboardingMemberCft struct {
	OnboardingCftBase
	OnboardingCftUrl string `json:"onboardingCftUrl"`
}

type OrganizationManagementViewModel struct {
	Id                            string                                 `json:"id"`
	AccountId                     int64                                  `json:"accountId"`
	ExternalOrganizationId        string                                 `json:"externalOrganizationId"`
	ExternalManagementAccountId   string                                 `json:"externalManagementAccountId"`
	ManagementAccountStackId      string                                 `json:"managementAccountStackId"`
	ManagementAccountStackRegion  string                                 `json:"managementAccountStackRegion"`
	OnboardingConfiguration       AwsOrganizationOnboardingConfiguration `json:"onboardingConfiguration"`
	UserId                        int                                    `json:"userId"`
	EnableStackModify             bool                                   `json:"enableStackModify"`
	StackSetArn                   string                                 `json:"stackSetArn"`
	OrganizationName              string                                 `json:"organizationName"`
	UpdateTime                    string                                 `json:"updateTime"`
	CreationTime                  string                                 `json:"creationTime"`
	StackSetRegions               []string                               `json:"stackSetRegions"`
	StackSetOrganizationalUnitIds []string                               `json:"stackSetOrganizationalUnitIds"`
}

type OnboardingConfigurationOptions struct {
	AwsAccountId string `json:"awsAccountId"`
}

func (service *Service) Create(body OnboardingRequest) (*OrganizationManagementViewModel, *http.Response, error) {
	v := new(OrganizationManagementViewModel)
	resp, err := service.Client.NewRequestDo("POST", cloudaccounts.RESTfulServicePathAwsOrgMgmt, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) UpdateStackSetArn(id string, body UpdateStackSetArnRequest) (*http.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id parameter must be passed")
	}

	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulServicePathAwsOrgMgmt, id, cloudaccounts.RESTfulServicePathAwsOrgMgmtStacksetArn)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) UpdateConfiguration(id string, body UpdateConfigurationRequest) (*http.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id parameter must be passed")
	}

	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulServicePathAwsOrgMgmt, id, cloudaccounts.RESTfulServicePathAwsOrgMgmtConfiguration)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulServicePathAwsOrgMgmt, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) Get(id string) (*OrganizationManagementViewModel, *http.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id parameter must be passed")
	}

	v := new(OrganizationManagementViewModel)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulServicePathAwsOrgMgmt, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]OrganizationManagementViewModel, *http.Response, error) {
	v := new([]OrganizationManagementViewModel)
	resp, err := service.Client.NewRequestDo("GET", cloudaccounts.RESTfulServicePathAwsOrgMgmt, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetOnboardingConfiguration(awsAccountId string) (*ManagementCftConfiguration, *http.Response, error) {
	if awsAccountId == "" {
		return nil, nil, fmt.Errorf("awsAccountId parameter must be passed")
	}

	v := new(ManagementCftConfiguration)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulServicePathAwsOrgMgmtOnboarding, cloudaccounts.RESTfulServicePathAwsOrgMgmtOnboardingMgmtStack)
	onboardingConfigurationOptions := OnboardingConfigurationOptions{
		AwsAccountId: awsAccountId,
	}

	resp, err := service.Client.NewRequestDo("GET", relativeURL, onboardingConfigurationOptions, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetMemberAccountConfiguration() (*OnboardingMemberCft, *http.Response, error) {
	v := new(OnboardingMemberCft)
	relativeURL := fmt.Sprintf("%s/%s", cloudaccounts.RESTfulServicePathAwsOrgMgmtOnboarding, cloudaccounts.RESTfulServicePathAwsOrgMgmtOnboardingMemberAccountStack)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
