package users

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"
)

const (
	userResourcePath      = "user"
	userResourceOwnerPath = "account/owner"
	userIAMSAfe           = "iam-safe"
	userAccounts          = "accounts"
	userIAMEntities       = "iamEntities"
)

var gUserEmailID = map[string]string{}
var onlyOnce sync.Once

type UserRequest struct {
	Email       string      `json:"email"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	SsoEnabled  bool        `json:"ssoEnabled"`
	Permissions Permissions `json:"permissions"`
}

type UserResponse struct {
	ID int `json:"id,omitempty"`
	// The name response is users email
	Email                 string      `json:"name"`
	IsSuspended           bool        `json:"isSuspended"`
	IsOwner               bool        `json:"isOwner"`
	IsSuperUser           bool        `json:"isSuperUser"`
	IsAuditor             bool        `json:"isAuditor"`
	HasAPIKey             bool        `json:"hasApiKey"`
	HasAPIKeyV1           bool        `json:"hasApiKeyV1"`
	HasAPIKeyV2           bool        `json:"hasApiKeyV2"`
	IsMfaEnabled          bool        `json:"isMfaEnabled,omitempty"`
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

type IAMSafeEntitiesResponse struct {
	CloudAccountID        string   `json:"cloudAccountId"`
	CloudAccountName      string   `json:"cloudAccountName"`
	ExternalAccountNumber string   `json:"externalAccountNumber"`
	IamEntities           []string `json:"iamEntities"`
	FailedIamEntities     []string `json:"failedIamEntities"`
}

type UserUpdate struct {
	RoleIds     []int       `json:"roleIds"`
	Permissions Permissions `json:"permissions"`
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

type SetOwnerQueryParameters struct {
	UserID string `json:"userId"`
}

type IAMSafeEntitiesBody struct {
	IAMEntities []string `json:"iamEntities"`
}

func (service *Service) Get(userId string) (*UserResponse, *http.Response, error) {
	v := new(UserResponse)
	path := fmt.Sprintf("%s/%s", userResourcePath, userId)
	resp, err := service.Client.NewRequestDoRetry("GET", path, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) GetAll() (*[]UserResponse, *http.Response, error) {
	v := new([]UserResponse)
	resp, err := service.Client.NewRequestDoRetry("GET", userResourcePath, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(user UserRequest) (*UserResponse, *http.Response, error) {
	v := new(UserResponse)
	var err error

	onlyOnce.Do(func() {
		err = service.refreshUserEmailIDMap()
	})
	if err != nil {
		return nil, nil, err
	}

	resp, err := service.Client.NewRequestDoRetry("POST", userResourcePath, nil, user, &v, nil)
	if err != nil {
		return nil, nil, err
	}

	gUserEmailID[v.Email] = strconv.Itoa(v.ID)
	return v, resp, nil
}

// blocked by bug: https://dome9-security.atlassian.net/browse/DOME-12720
func (service *Service) Update(userId string, user *UserUpdate) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", userResourcePath, userId)
	resp, err := service.Client.NewRequestDoRetry("PUT", path, nil, user, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) SetUserAsOwner(userId string) (*http.Response, error) {
	body := SetOwnerQueryParameters{
		UserID: userId,
	}
	resp, err := service.Client.NewRequestDoRetry("PUT", userResourceOwnerPath, nil, body, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(userID string) (*http.Response, error) {
	var err error

	onlyOnce.Do(func() {
		err = service.refreshUserEmailIDMap()
	})
	if err != nil {
		return nil, err
	}

	user, _, err := service.Get(userID)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s", userResourcePath, userID)
	resp, err := service.Client.NewRequestDoRetry("DELETE", path, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	delete(gUserEmailID, user.Email)
	return resp, err
}

/*
	iam safe entities
*/

func (service *Service) ProtectWithElevationIAMSafeEntity(d9CloudAccountID, entityName, entityType string, d9UsersIDToProtect []string) (*[]IAMSafeEntitiesResponse, error) {
	var awsCloudAccountService = aws.New(service.Client.Config)
	v := make([]IAMSafeEntitiesResponse, len(d9UsersIDToProtect))
	var err error

	if len(d9UsersIDToProtect) == 0 {
		return nil, errors.New("you must specify at least one user in protect with elevation mode")
	}

	iamEntityArn, err := awsCloudAccountService.GetProtectIAMSafeEntityStatusByName(d9CloudAccountID, entityName, entityType)
	if err != nil {
		return nil, err
	}

	body := IAMSafeEntitiesBody{
		IAMEntities: []string{iamEntityArn.Arn},
	}
	for i, userID := range d9UsersIDToProtect {
		relativeURL := fmt.Sprintf("%s/%s/%s/%s/%s/%s", userResourcePath, userID, userIAMSAfe, userAccounts, d9CloudAccountID, userIAMEntities)
		_, err = service.Client.NewRequestDoRetry("POST", relativeURL, nil, body, &v[i], nil)
		if err != nil {
			return nil, err
		}
	}

	return &v, nil
}

func (service *Service) ProtectWithElevationIAMSafeEntityUpdate(d9CloudAccountID, entityType, entityName string, d9UsersIDToProtect []string) (*[]IAMSafeEntitiesResponse, error) {
	var err error
	var awsCloudAccountService = aws.New(service.Client.Config)
	v := make([]IAMSafeEntitiesResponse, len(d9UsersIDToProtect))

	onlyOnce.Do(func() {
		err = service.refreshUserEmailIDMap()
	})
	if err != nil {
		return nil, err
	}

	iamEntityArn, err := awsCloudAccountService.GetProtectIAMSafeEntityStatusByName(d9CloudAccountID, entityName, entityType)
	if err != nil {
		return nil, err
	}
	currProtectedDome9UsersID := getUsersIDsAccordingToEmails(iamEntityArn.AttachedDome9Users)
	// create map where the key is the user id and the value is true or false, where true indicates to protect the user and false to unprotect.
	// if the value is true then call update api func with aws iam user role arn (to protect) otherwise call with empty sting (to unprotect).
	protectedUnprotectedMap := generateProtectUnprotectMap(currProtectedDome9UsersID, d9UsersIDToProtect)

	unprotectIAMEntitiesBody := IAMSafeEntitiesBody{
		IAMEntities: []string{},
	}
	protectIAMEntitiesBody := IAMSafeEntitiesBody{
		IAMEntities: []string{iamEntityArn.Arn},
	}

	i := 0
	for userID, toProtect := range protectedUnprotectedMap {
		relativeURL := fmt.Sprintf("%s/%s/%s/%s/%s/%s", userResourcePath, userID, userIAMSAfe, userAccounts, d9CloudAccountID, userIAMEntities)
		if toProtect {
			_, err = service.Client.NewRequestDoRetry("PUT", relativeURL, nil, protectIAMEntitiesBody, &v[i], nil)
		} else {
			_, err = service.Client.NewRequestDoRetry("PUT", relativeURL, nil, unprotectIAMEntitiesBody, &v[i], nil)
		}

		if err != nil {
			return nil, err
		}
		i++
	}

	return &v, nil
}

func (service *Service) UnprotectWithElevationIAMSafeEntity(d9CloudAccountID, entityName, entityType string) (*http.Response, error) {
	awsCloudAccountService := aws.New(service.Client.Config)
	req := aws.RestrictedIamEntitiesRequest{
		EntityName: entityName,
		EntityType: entityType,
	}

	_, _, err := awsCloudAccountService.ProtectIAMSafeEntity(d9CloudAccountID, req)
	if err != nil {
		return nil, err
	}

	_, err = awsCloudAccountService.UnprotectIAMSafeEntity(d9CloudAccountID, entityName, entityType)
	return nil, err
}

func getUsersIDsAccordingToEmails(emailsForProtectedD9Users []string) []string {
	usersIDs := make([]string, len(emailsForProtectedD9Users))

	for i, userEmail := range emailsForProtectedD9Users {
		usersIDs[i] = gUserEmailID[userEmail]
	}
	return usersIDs
}

// This function return a map where the key is the user id and the value is true or false, where true indicates to protect the user and false to unprotect the user.
func generateProtectUnprotectMap(currProtectedUsersID []string, d9UsersIDToProtect []string) map[string]bool {
	protectUnprotectMap := map[string]bool{}

	for _, currProtectedUserID := range currProtectedUsersID {
		protectUnprotectMap[currProtectedUserID] = false
	}

	for _, d9UserIDToProtect := range d9UsersIDToProtect {
		// if the user already protected then there is to need to protect him.
		if _, ok := protectUnprotectMap[d9UserIDToProtect]; ok {
			delete(protectUnprotectMap, d9UserIDToProtect)
		} else {
			protectUnprotectMap[d9UserIDToProtect] = true
		}
	}

	return protectUnprotectMap
}

func (service *Service) refreshUserEmailIDMap() error {
	users, _, err := service.GetAll()
	if err != nil {
		return err
	}

	for _, user := range *users {
		gUserEmailID[user.Email] = strconv.Itoa(user.ID)
	}
	return nil
}
