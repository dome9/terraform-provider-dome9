package dome9

import (
	"log"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_policy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceContinuousCompliancePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceContinuousCompliancePolicyCreate,
		Read:   resourceContinuousCompliancePolicyRead,
		Update: resourceContinuousCompliancePolicyUpdate,
		Delete: resourceContinuousCompliancePolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"target_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"Aws", "Azure", "Gcp", "Kubernetes", "OrganizationalUnit"}, false),
			},
			"ruleset_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"notification_ids": {
				Type:     schema.TypeSet,
				Required: true,
				// ForceNew: true,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceContinuousCompliancePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandContinuousCompliancePolicyRequest(d)
	log.Printf("[INFO] Creating compliance policy request %+v\n", req)
	resp, _, err := d9Client.continuousCompliancePolicy.Create(&req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created compliance policy with ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceContinuousCompliancePolicyRead(d, meta)
}

func resourceContinuousCompliancePolicyRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.continuousCompliancePolicy.Get(d.Id())

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing continuous compliance policy %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting continuous compliance policy: %+v\n", resp)
	d.SetId(resp.ID)
	_ = d.Set("target_id", resp.TargetInternalId)
	_ = d.Set("target_type", resp.TargetType)
	_ = d.Set("ruleset_id", resp.RulesetId)
	if err := d.Set("notification_ids", resp.NotificationIds); err != nil {
		return err
	}

	return nil
}

func resourceContinuousCompliancePolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Updating continuous compliance policy ID: %v\n", d.Id())
	req := expandContinuousCompliancePolicyRequest(d)

	if _, _, err := d9Client.continuousCompliancePolicy.Update(&req); err != nil {
		return err
	}

	return nil
}

func resourceContinuousCompliancePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting continuous compliance policy ID: %v\n", d.Id())

	if _, err := d9Client.continuousCompliancePolicy.Delete(d.Id()); err != nil {
		return err
	}
	return nil
}

func expandContinuousCompliancePolicyRequest(d *schema.ResourceData) continuous_compliance_policy.ContinuousCompliancePolicyRequest {
	return continuous_compliance_policy.ContinuousCompliancePolicyRequest{
		TargetId:        d.Get("target_id").(string),
		RulesetId:       d.Get("ruleset_id").(int),
		NotificationIds: expandNotificationIDs(d, "notification_ids"),
		TargetType:      d.Get("target_type").(string),
	}
}
