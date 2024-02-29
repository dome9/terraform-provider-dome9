package awp_aws_onboarding

import (
	"net/http"
)

const (
	awpAWSGetOnboardingDataPath = "workload/agentless/aws/terraform/onboarding"
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

func (service *Service) Get() (*AgentlessAwsTerraformOnboardingDataResponse, *http.Response, error) {
	v := new(AgentlessAwsTerraformOnboardingDataResponse)
	resp, err := service.Client.NewRequestDo("GET", awpAWSGetOnboardingDataPath, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
