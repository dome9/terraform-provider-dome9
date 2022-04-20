package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceAdmissionControlPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAdmissionControlPolicyRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_id": {
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
			"action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ruleset_platform": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAdmissionControlPolicyRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	policyID := d.Get("id").(string)
	log.Printf("Getting data for Admission Control Policy id: %s\n", policyID)

	resp, _, err := d9Client.admissionControlPolicy.GetAdmissionControlPolicy(policyID)
	if err != nil {
		return err
	}

	d.SetId(resp.ID)
	_ = d.Set("id", resp.ID)
	_ = d.Set("target_id", resp.TargetId)
	_ = d.Set("target_type", resp.TargetType)
	_ = d.Set("ruleset_id", resp.RulesetId)
	_ = d.Set("action", resp.Action)
	if err := d.Set("notification_ids", resp.NotificationIds); err != nil {
		return err
	}

	return nil
}
