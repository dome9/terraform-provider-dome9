package dome9

import (
	"encoding/base64"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAwpAwsOnboardingData() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwpAwsOnboardingDataRead,

		Schema: map[string]*schema.Schema{
			"cloudguard_account_id": {
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
			"cloud_guard_backend_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"agentless_bucket_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_functions_prefix_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_snapshots_utils_function_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_snapshots_utils_function_run_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_snapshots_utils_function_time_out": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"awp_client_side_security_group_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cross_account_role_external_id": {
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
	_ = d.Set("cloud_guard_backend_account_id", resp.CloudGuardBackendAccountId)
	_ = d.Set("agentless_bucket_name", resp.AgentlessBucketName)
	_ = d.Set("remote_functions_prefix_key", resp.RemoteFunctionsPrefixKey)
	_ = d.Set("remote_snapshots_utils_function_name", resp.RemoteSnapshotsUtilsFunctionName)
	_ = d.Set("remote_snapshots_utils_function_run_time", resp.RemoteSnapshotsUtilsFunctionRunTime)
	_ = d.Set("remote_snapshots_utils_function_time_out", resp.RemoteSnapshotsUtilsFunctionTimeOut)
	_ = d.Set("awp_client_side_security_group_name", resp.AwpClientSideSecurityGroupName)
	cloudAccountID, _, err := d9Client.awpAwsOnboarding.GetCloudAccountId(d.Get("cloudguard_account_id").(string))
	if err != nil {
		return err
	}
	combinedString := resp.CloudGuardBackendAccountId + "-" + cloudAccountID
	encodedString := base64.StdEncoding.EncodeToString([]byte(combinedString))
	_ = d.Set("cross_account_role_external_id", encodedString)

	return nil
}
