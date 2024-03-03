package dome9

import (
	"encoding/base64"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAwpAwsOnboardingData() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwpAwsOnboardingDataRead,

		Schema: map[string]*schema.Schema{
			"externalAwsAccountId": {
				Type:     schema.TypeString,
				Required: true,
			},
			"stage": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloudGuardBackendAccountId": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"agentlessBucketName": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remoteFunctionsPrefixKey": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remoteSnapshotsUtilsFunctionName": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remoteSnapshotsUtilsFunctionRunTime": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remoteSnapshotsUtilsFunctionTimeOut": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"awpClientSideSecurityGroupName": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"crossAccountRoleExternalId": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAwpAwsOnboardingDataRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	resp, _, err := d9Client.awpAwsOnboarding.Get()
	if err != nil {
		return err
	}

	d.SetId(resp.CloudGuardBackendAccountId)
	_ = d.Set("stage", resp.Stage)
	_ = d.Set("region", resp.Region)
	_ = d.Set("cloudGuardBackendAccountId", resp.CloudGuardBackendAccountId)
	_ = d.Set("agentlessBucketName", resp.AgentlessBucketName)
	_ = d.Set("remoteFunctionsPrefixKey", resp.RemoteFunctionsPrefixKey)
	_ = d.Set("remoteSnapshotsUtilsFunctionName", resp.RemoteSnapshotsUtilsFunctionName)
	_ = d.Set("remoteSnapshotsUtilsFunctionRunTime", resp.RemoteSnapshotsUtilsFunctionRunTime)
	_ = d.Set("remoteSnapshotsUtilsFunctionTimeOut", resp.RemoteSnapshotsUtilsFunctionTimeOut)
	_ = d.Set("awpClientSideSecurityGroupName", resp.AwpClientSideSecurityGroupName)
	cloudAccountID, _, err := d9Client.awpAwsOnboarding.GetCloudAccountId(d.Get("externalAwsAccountId").(string))
	if err != nil {
		return err
	}
	combinedString := resp.CloudGuardBackendAccountId + "-" + cloudAccountID
	encodedString := base64.StdEncoding.EncodeToString([]byte(combinedString))
	_ = d.Set("crossAccountRoleExternalId", encodedString)

	return nil
}
