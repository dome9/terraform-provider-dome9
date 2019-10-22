package dome9

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceContinuousCompliancePolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceContinuousCompliancePolicyRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_account_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bundle_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"notification_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceContinuousCompliancePolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	policyID := d.Get("id").(string)
	log.Printf("Getting data for Continuous Compliance Policy id: %s\n", policyID)

	resp, _, err := client.continuousCompliancePolicy.Get(policyID)
	if err != nil {
		return err
	}

	d.SetId(resp.ID)
	_ = d.Set("cloud_account_id", resp.CloudAccountID)
	_ = d.Set("external_account_id", resp.ExternalAccountID)
	_ = d.Set("cloud_account_type", resp.CloudAccountType)
	_ = d.Set("bundle_id", resp.BundleID)
	if err := d.Set("notification_ids", flattenNotificationIDs(resp)); err != nil {
		return err
	}

	return nil
}
