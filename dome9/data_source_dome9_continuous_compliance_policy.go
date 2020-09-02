package dome9

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceContinuousCompliancePolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceContinuousCompliancePolicyRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_internal_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ruleset_id": {
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
	d9Client := meta.(*Client)

	policyID := d.Get("id").(string)
	log.Printf("Getting data for Continuous Compliance Policy id: %s\n", policyID)

	resp, _, err := d9Client.continuousCompliancePolicy.Get(policyID)
	if err != nil {
		return err
	}

	d.SetId(resp.ID)
	_ = d.Set("target_internal_id", resp.TargetInternalId)
	_ = d.Set("target_external_id", resp.TargetExternalId)
	_ = d.Set("target_type", resp.TargetType)
	_ = d.Set("ruleset_id", resp.RulesetId)
	if err := d.Set("notification_ids", resp.NotificationIds); err != nil {
		return err
	}

	return nil
}
