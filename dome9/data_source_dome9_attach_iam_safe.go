package dome9

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
)

func dataSourceAttachIAMSafe() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAttachIAMSafeRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"aws_cloud_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_group_arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_policy_arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAttachIAMSafeRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for AWS cloud account attach IAM safe %s\n", id)

	resp, _, err := d9Client.cloudaccountAWS.Get(cloudaccounts.QueryParameters{ID: id})
	if err != nil {
		return err
	}

	d.SetId(resp.ID)
	_ = d.Set("aws_cloud_account_id", resp.ID)
	_ = d.Set("aws_group_arn", resp.IamSafe.AwsGroupArn)
	_ = d.Set("aws_policy_arn", resp.IamSafe.AwsPolicyArn)
	_ = d.Set("mode", resp.IamSafe.Mode)

	return nil
}
