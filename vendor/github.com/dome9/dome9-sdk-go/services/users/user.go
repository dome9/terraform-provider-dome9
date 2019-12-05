package users

import (
	"fmt"
	"net/http"
	"time"
)

const (
	userResourcePath = "user"
)

type UserRequest struct {
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	SsoEnabled bool   `json:"ssoEnabled"`
}

type UserResponse struct {
	ID                    int         `json:"id"`
	Name                  string      `json:"name"`
	IsSuspended           bool        `json:"isSuspended"`
	IsOwner               bool        `json:"isOwner"`
	IsSuperUser           bool        `json:"isSuperUser"`
	IsAuditor             bool        `json:"isAuditor"`
	HasAPIKey             bool        `json:"hasApiKey"`
	HasAPIKeyV1           bool        `json:"hasApiKeyV1"`
	HasAPIKeyV2           bool        `json:"hasApiKeyV2"`
	IsMfaEnabled          bool        `json:"isMfaEnabled"`
	SsoEnabled            bool        `json:"ssoEnabled"`
	RoleIds               []int       `json:"roleIds"`
	IamSafe               IamSafe     `json:"iamSafe"`
	CanSwitchRole         bool        `json:"canSwitchRole"`
	IsLocked              bool        `json:"isLocked"`
	LastLogin             time.Time   `json:"lastLogin"`
	Permissions           Permissions `json:"permissions"`
	CalculatedPermissions Permissions `json:"calculatedPermissions"`
	IsMobileDevicePaired  bool        `json:"isMobileDevicePaired"`
}

type CloudAccounts struct {
	CloudAccountID           string                     `json:"cloudAccountId"`
	Name                     string                     `json:"name"`
	ExternalAccountNumber    string                     `json:"externalAccountNumber"`
	LastLeaseTime            time.Time                  `json:"lastLeaseTime"`
	State                    string                     `json:"state"`
	IamEntities              []string                   `json:"iamEntities"`
	IamEntitiesLastLeaseTime []IamEntitiesLastLeaseTime `json:"iamEntitiesLastLeaseTime"`
	CloudAccountState        string                     `json:"cloudAccountState"`
	IamEntity                string                     `json:"iamEntity"`
}

type IamEntitiesLastLeaseTime struct {
	IamEntity     string    `json:"iamEntity"`
	LastLeaseTime time.Time `json:"lastLeaseTime"`
}

type IamSafe struct {
	CloudAccounts []CloudAccounts `json:"cloudAccounts"`
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

func (service *Service) Get(userId string) (*UserResponse, *http.Response, error) {
	v := new(UserResponse)
	path := fmt.Sprintf("%s/%s", userResourcePath, userId)
	resp, err := service.Client.NewRequestDo("GET", path, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) GetAll() (*[]UserResponse, *http.Response, error) {
	v := new([]UserResponse)
	path := fmt.Sprintf("%s", userResourcePath)
	resp, err := service.Client.NewRequestDo("GET", path, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(user UserRequest) (*UserResponse, *http.Response, error) {
	v := new(UserResponse)
	resp, err := service.Client.NewRequestDo("POST", userResourcePath, nil, user, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

// blocked by bug: https://dome9-security.atlassian.net/browse/DOME-12720
func (service *Service) Update(userId string, user UserRequest) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", userResourcePath, userId)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, user, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(userId string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", userResourcePath, userId)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
