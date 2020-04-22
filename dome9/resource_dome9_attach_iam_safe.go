package dome9

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"
)

func resourceAttachIAMSafe() *schema.Resource {
	return &schema.Resource{
		Create: resourceAttachIAMSafeCreate,
		Read:   resourceAttachIAMSafeRead,
		Update: nil,
		Delete: resourceAttachIAMSafeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"aws_cloud_account_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"aws_group_arn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"aws_policy_arn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAttachIAMSafeCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandAttachIAMSafeRequest(d)
	a := 5
	fmt.Println(a)
	a = a = 6
	log.Printf("[INFO] Attach IAM safe with request\n%+v\n", req)
	resp, _, err := d9Client.cloudaccountAWS.AttachIAMSafeToCloudAccount(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Attach IAM safe. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceAttachIAMSafeRead(d, meta)
}

func resourceAttachIAMSafeRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	getCloudAccountQueryParams := cloudaccounts.QueryParameters{ID: d.Id()}
	resp, _, err := d9Client.cloudaccountAWS.Get(&getCloudAccountQueryParams)

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing IAM safe from state because aws cloud account %s it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Reading account response and settings states: %+v\n", resp.IamSafe)
	d.SetId(resp.ID)
	_ = d.Set("aws_cloud_account_id", resp.ID)
	_ = d.Set("aws_group_arn", resp.IamSafe.AwsGroupArn)
	_ = d.Set("aws_policy_arn", resp.IamSafe.AwsPolicyArn)
	_ = d.Set("mode", resp.IamSafe.Mode)

	return nil
}

func expandAttachIAMSafeRequest(d *schema.ResourceData) aws.AttachIamSafeRequest {
	return aws.AttachIamSafeRequest{
		CloudAccountID: d.Get("aws_cloud_account_id").(string),
		Data: aws.Data{
			AwsGroupArn:  d.Get("aws_group_arn").(string),
			AwsPolicyArn: d.Get("aws_policy_arn").(string),
			Mode:         d.Get("mode").(string),
		},
	}
}

func resourceAttachIAMSafeDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Detach IAM safe to AWS Cloud Account ID: %v\n", d.Id())

	if _, err := d9Client.cloudaccountAWS.DetachIAMSafeToCloudAccount(d.Id()); err != nil {
		return err
	}

	return nil
}
