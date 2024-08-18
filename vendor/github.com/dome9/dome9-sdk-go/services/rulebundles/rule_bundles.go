package rulebundles

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	ruleBundleResourcePath = "Compliance/Ruleset"
)

type RuleBundleRequest struct {
	Name             string  `json:"name,omitempty"`
	Description      string  `json:"description,omitempty"`
	Rules            *[]Rule `json:"rules,omitempty"`
	ID               int     `json:"id,omitempty"`
	HideInCompliance bool    `json:"hideInCompliance,omitempty"`
	MinFeatureTier   string  `json:"minFeatureTier,omitempty"`
	CloudVendor      string  `json:"cloudVendor,omitempty"`
	Language         string  `json:"language,omitempty"`
}

type RuleBundleResponse struct {
	Rules            []Rule `json:"rules"`
	AccountID        int    `json:"accountId"`
	CreatedTime      string `json:"createdTime"`
	UpdatedTime      string `json:"updatedTime"`
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	IsTemplate       bool   `json:"isTemplate"`
	HideInCompliance bool   `json:"hideInCompliance"`
	MinFeatureTier   string `json:"minFeatureTier"`
	Section          int    `json:"section"`
	TooltipText      string `json:"tooltipText"`
	ShowBundle       bool   `json:"showBundle"`
	SystemBundle     bool   `json:"systemBundle"`
	CloudVendor      string `json:"cloudVendor"`
	Version          int    `json:"version"`
	Language         string `json:"language"`
	RulesCount       int    `json:"rulesCount"`
}

type Rule struct {
	Name          string   `json:"name,omitempty"`
	Severity      string   `json:"severity,omitempty"`
	Logic         string   `json:"logic,omitempty"`
	Description   string   `json:"description,omitempty"`
	Remediation   string   `json:"remediation,omitempty"`
	ComplianceTag string   `json:"complianceTag,omitempty"`
	Domain        string   `json:"domain,omitempty"`
	Priority      string   `json:"priority,omitempty"`
	ControlTitle  string   `json:"controlTitle,omitempty"`
	RuleID        string   `json:"ruleId,omitempty"`
	Category      string   `json:"category,omitempty"`
	LogicHash     string   `json:"logicHash,omitempty"`
	IsDefault     bool     `json:"isDefault,omitempty"`
}

func (service *Service) Get(id string) (*RuleBundleResponse, *http.Response, error) {
	v := new(RuleBundleResponse)
	relativeURL := fmt.Sprintf("%s/%s", ruleBundleResourcePath, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAccountRuleBundles() (*[]RuleBundleResponse, *http.Response, error) {
	v := new([]RuleBundleResponse)
	resp, err := service.Client.NewRequestDo("GET", ruleBundleResourcePath, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body *RuleBundleRequest) (*RuleBundleResponse, *http.Response, error) {
	v := new(RuleBundleResponse)
	resp, err := service.Client.NewRequestDo("POST", ruleBundleResourcePath, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(body *RuleBundleRequest) (*RuleBundleResponse, *http.Response, error) {
	// Rule bundle ID passed within the request body
	v := new(RuleBundleResponse)
	relativeURL := fmt.Sprintf("%s/%s", ruleBundleResourcePath, strconv.Itoa(body.ID))
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", ruleBundleResourcePath, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
