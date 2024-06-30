package awp_onboarding

import (
	"fmt"
	"net/http"
	"time"
	"github.com/dome9/dome9-sdk-go/dome9/client"
)

const (
	OnboardingResourcePath = "workload/agentless/%s/accounts/%s"
	ScanModeInAccountSub   = "inAccountSub"
    ScanModeInAccountHub   = "inAccountHub"
)

type CreateOptions struct {
	ShouldCreatePolicy string `url:"shouldCreatePolicy"`
}

type DeleteOptions struct {
	ForceDelete string `url:"forceDelete"`
}

type AgentlessAccountSettings struct {
	DisabledRegions              []string          `json:"disabledRegions"`
	ScanMachineIntervalInHours   int               `json:"scanMachineIntervalInHours"`
	MaxConcurrenceScansPerRegion int               `json:"maxConcurrenceScansPerRegion"`
	SkipFunctionAppsScan         bool              `json:"skipFunctionAppsScan"`
	CustomTags                   map[string]string `json:"customTags"`
}

type AccountIssues struct {
	Regions map[string]interface{}  `json:"regions"`
	Account *map[string]interface{} `json:"account"`
}

type GetAWPOnboardingResponse struct {
	AgentlessAccountSettings        *AgentlessAccountSettings `json:"agentlessAccountSettings"`
	MissingAwpPrivateNetworkRegions *[]string                 `json:"missingAwpPrivateNetworkRegions"`
	AccountIssues                   *AccountIssues            `json:"accountIssues"`
	CloudAccountId                  string                    `json:"cloudAccountId"`
	AgentlessProtectionEnabled      bool                      `json:"agentlessProtectionEnabled"`
	ScanMode                        string                    `json:"scanMode"`
	Provider                        string                    `json:"provider"`
	ShouldUpdate                    bool                      `json:"shouldUpdate"`
	IsOrgOnboarding                 bool                      `json:"isOrgOnboarding"`
	CentralizedCloudAccountId       string                    `json:"centralizedCloudAccountId"`
}

// Common functionality

func CreateAWPOnboarding(client *client.Client, req interface{}, path string, queryParams CreateOptions) (*http.Response, error) {
	// Define the maximum number of retries and the interval between retries
	maxRetries := 3
	retryInterval := time.Second * 5

	// Initialize the response and error variables outside the loop
	var resp *http.Response
	var err error

	// Attempt the request up to maxRetries times
	for i := 0; i < maxRetries; i++ {
		// Make the request
		resp, err = client.NewRequestDo("POST", path, queryParams, req, nil)
		if err == nil {
			// If the request was successful, return the response
			return resp, nil
		}

		// If the request failed with a 404 status code, wait for the retry interval before trying again
		if resp != nil && resp.StatusCode == 404 {
			time.Sleep(retryInterval)
		} else {
			// If the status code is not 404, return the response and error immediately
			return resp, err
		}
	}

	// If the function hasn't returned after maxRetries, return an error
	return nil, fmt.Errorf("failed to create AWP Onboarding after %d attempts: %w", maxRetries, err)
}

func GetAWPOnboarding(client *client.Client, cloudProvider string, id string) (*GetAWPOnboardingResponse, *http.Response, error) {
	v := new(GetAWPOnboardingResponse)
	path := fmt.Sprintf(OnboardingResourcePath, cloudProvider, id)
	resp, err := client.NewRequestDo("GET", path, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func DeleteAWPOnboarding(client *client.Client, cloudProvider string, id string, queryParams DeleteOptions) (*http.Response, error) {
	path := fmt.Sprintf(OnboardingResourcePath, cloudProvider, id)
	resp, err := client.NewRequestDo("DELETE", path, queryParams, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func UpdateAWPSettings(client *client.Client, cloudProvider string, id string, req AgentlessAccountSettings) (*http.Response, error) {
	// Construct the URL path
	path := fmt.Sprintf(OnboardingResourcePath, cloudProvider, id)
	// Make a PATCH request with the JSON body
	resp, err := client.NewRequestDo("PATCH", fmt.Sprintf("%s/settings", path), nil, req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}