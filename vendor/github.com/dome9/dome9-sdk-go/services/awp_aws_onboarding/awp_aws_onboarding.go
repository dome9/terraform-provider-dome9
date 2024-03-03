package awp_aws_onboarding

import (
	"fmt"
	"net/http"
)

const (
	awpAWSGetOnboardingDataPath = "workload/agentless/aws/terraform/onboarding"
	cloudAccountsPath           = "cloudaccounts/"
)

type AgentlessAwsTerraformOnboardingDataResponse struct {
	Stage                               string `json:"stage"`
	Region                              string `json:"region"`
	CloudGuardBackendAccountId          string `json:"cloudGuardBackendAccountId"`
	AgentlessBucketName                 string `json:"agentlessBucketName"`
	RemoteFunctionsPrefixKey            string `json:"remoteFunctionsPrefixKey"`
	RemoteSnapshotsUtilsFunctionName    string `json:"remoteSnapshotsUtilsFunctionName"`
	RemoteSnapshotsUtilsFunctionRunTime string `json:"remoteSnapshotsUtilsFunctionRunTime"`
	RemoteSnapshotsUtilsFunctionTimeOut int    `json:"remoteSnapshotsUtilsFunctionTimeOut"`
	AwpClientSideSecurityGroupName      string `json:"awpClientSideSecurityGroupName"`
}

type CloudAccountResponse struct {
	ID string `json:"id"`
}

func (service *Service) Get() (*AgentlessAwsTerraformOnboardingDataResponse, *http.Response, error) {
	v := new(AgentlessAwsTerraformOnboardingDataResponse)
	resp, err := service.Client.NewRequestDo("GET", awpAWSGetOnboardingDataPath, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetCloudAccountId(externalAccountId string) (string, *http.Response, error) {
	path := fmt.Sprintf("%s%s", cloudAccountsPath, externalAccountId)
	respData := new(CloudAccountResponse)
	resp, err := service.Client.NewRequestDo("GET", path, nil, nil, respData)
	if err != nil {
		return "", nil, err
	}
	return respData.ID, resp, nil
}
