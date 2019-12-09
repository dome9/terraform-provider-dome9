package securitygroupaws

import (
	"fmt"
	"net/http"
)

const (
	awsSgResourcePath           = "CloudSecurityGroup"
	awsSgResourceProtectionMode = "protection-mode"
	awsSgResourceServices       = "services"
	awsSgResourceTags           = "tags"
)

// There is a bug when we pass nil inbound or outbound, ticket link: https://dome9-security.atlassian.net/browse/DOME-12727
type CloudSecurityGroupRequest struct {
	SecurityGroupName string                 `json:"securityGroupName"`
	CloudAccountID    string                 `json:"cloudAccountId"`
	Description       string                 `json:"description,omitempty"`
	RegionID          string                 `json:"regionId,omitempty"`
	IsProtected       bool                   `json:"isProtected,omitempty"`
	VpcId             string                 `json:"VpcId,omitempty"`
	VpcName           string                 `json:"VpcName,omitempty"`
	Services          ServicesRequest        `json:"services,omitempty"`
	Tags              map[string]interface{} `json:"tags,omitempty"`
}

type CloudSecurityGroupResponse struct {
	ID                int               `json:"securityGroupId"`
	SecurityGroupName string            `json:"securityGroupName"`
	Description       string            `json:"description"`
	RegionID          string            `json:"regionId"`
	IsProtected       bool              `json:"isProtected"`
	CloudAccountID    string            `json:"cloudAccountId"`
	ExternalID        string            `json:"externalId"`
	VpcID             string            `json:"vpcId"`
	VpcName           *string           `json:"vpcName"`
	CloudAccountName  string            `json:"cloudAccountName"`
	Services          ServicesResponse  `json:"services,omitempty"`
	Tags              map[string]string `json:"tags,omitempty"`
}

type ServicesRequest struct {
	Inbound  []BoundServicesRequest `json:"inbound"`
	Outbound []BoundServicesRequest `json:"outbound"`
}

type ServicesResponse struct {
	Inbound  []BoundServicesResponse `json:"inbound"`
	Outbound []BoundServicesResponse `json:"outbound"`
}

type BoundServicesRequest struct {
	Name         string  `json:"name"`
	Description  string  `json:"description,omitempty"`
	ProtocolType string  `json:"protocolType,omitempty"`
	Port         string  `json:"port,omitempty"`
	OpenForAll   bool    `json:"openForAll,omitempty"`
	Scope        []Scope `json:"scope,omitempty"`
}

type BoundServicesResponse struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ProtocolType string  `json:"protocolType"`
	Port         string  `json:"port"`
	OpenForAll   bool    `json:"openForAll"`
	Scope        []Scope `json:"scope"`
	Inbound      bool    `json:"inbound"`
	ICMPType     string  `json:"icmpType"`
	ICMPv6Type   string  `json:"icmpv6Type,omitempty"`
}

type Scope struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type GetSecurityGroupQueryParameters struct {
	CloudAccountID string
	RegionID       string
}

type UpdateProtectionModeQueryParameters struct {
	ProtectionMode string `json:"protectionMode"`
}

type UpdateTagsQueryParameters struct {
	Tags map[string]interface{} `json:"tags"`
}

func (service *Service) Get(d9SecurityGroupID string) (*CloudSecurityGroupResponse, *http.Response, error) {
	v := new(CloudSecurityGroupResponse)
	relativeURL := fmt.Sprintf("%s/%s", awsSgResourcePath, d9SecurityGroupID)

	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// GetAllInRegion will return list of all the security groups in specific region
func (service *Service) GetAllInRegion(d9CloudAccountID, awsRegionName string) (*[]CloudSecurityGroupResponse, *http.Response, error) {
	if d9CloudAccountID == "" && awsRegionName == "" {
		return nil, nil, fmt.Errorf("d9 cloud account id and aws region name must be passed")
	}

	options := GetSecurityGroupQueryParameters{
		CloudAccountID: d9CloudAccountID,
		RegionID:       awsRegionName,
	}

	v := new([]CloudSecurityGroupResponse)
	resp, err := service.Client.NewRequestDo("GET", awsSgResourcePath, options, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// GetAll will return list of all the security groups in the whole regions
func (service *Service) GetAll() (*[]CloudSecurityGroupResponse, *http.Response, error) {
	v := new([]CloudSecurityGroupResponse)
	resp, err := service.Client.NewRequestDo("GET", awsSgResourcePath, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body CloudSecurityGroupRequest) (*CloudSecurityGroupResponse, *http.Response, error) {
	v := new(CloudSecurityGroupResponse)
	resp, err := service.Client.NewRequestDo("POST", awsSgResourcePath, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// update aws security group, in order to update the field isProtected we should call the function UpdateProtectionMode
func (service *Service) Update(d9SecurityGroupID string, body CloudSecurityGroupRequest) (*CloudSecurityGroupResponse, *http.Response, error) {
	v := new(CloudSecurityGroupResponse)
	relativeURL := fmt.Sprintf("%s/%s", awsSgResourcePath, d9SecurityGroupID)

	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// update protection mode is post api call, mode can be one of the following "FullManage","ReadOnly"
func (service *Service) UpdateProtectionMode(d9SecurityGroupID, protectionMode string) (*CloudSecurityGroupResponse, *http.Response, error) {
	if protectionMode != "FullManage" && protectionMode != "ReadOnly" {
		return nil, nil, fmt.Errorf("protection mode can be FullManage or ReadOnly")
	}

	v := new(CloudSecurityGroupResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s", awsSgResourcePath, d9SecurityGroupID, awsSgResourceProtectionMode)
	body := UpdateProtectionModeQueryParameters{
		ProtectionMode: protectionMode,
	}

	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// update tags is post api call
func (service *Service) UpdateTags(d9SecurityGroupID string, tags map[string]interface{}) (*map[string]string, *http.Response, error) {
	updatedTags := new(map[string]string)
	relativeURL := fmt.Sprintf("%s/%s/%s", awsSgResourcePath, d9SecurityGroupID, awsSgResourceTags)
	body := UpdateTagsQueryParameters{
		Tags: tags,
	}

	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, body, updatedTags)
	if err != nil {
		return nil, nil, err
	}

	return updatedTags, resp, nil
}

// create and attach or update bound service (post api call)
// There is a bug when we trying to update read only security group, get 500 (internal error) ticket link: https://dome9-security.atlassian.net/browse/DOME-12737
func (service *Service) HandelBoundServices(d9SecurityGroupID, policyType string, boundService BoundServicesRequest) (*CloudSecurityGroupResponse, *http.Response, error) {
	v := new(CloudSecurityGroupResponse)
	relativeURL := fmt.Sprintf("%s/%s/%s/%s", awsSgResourcePath, d9SecurityGroupID, awsSgResourceServices, policyType)

	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, boundService, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(d9SecurityGroupID string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", awsSgResourcePath, d9SecurityGroupID)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
