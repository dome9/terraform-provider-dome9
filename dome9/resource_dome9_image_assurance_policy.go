package dome9

import (
	"log"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/imageassurance/imageassurance_policy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceImageAssurancePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceImageAssurancePolicyCreate,
		Read:   resourceImageAssurancePolicyRead,
		Update: resourceImageAssurancePolicyUpdate,
		Delete: resourceImageAssurancePolicyDelete,
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
				ValidateFunc: validation.StringInSlice([]string{"Environment", "OrganizationalUnit"}, false),
			},
			"ruleset_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"notification_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"admission_control_action": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Prevention", "Detection"}, false),
			},
			"admission_control_unscanned_action": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Prevention", "Detection"}, false),
			},
		},
	}
}

func resourceImageAssurancePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandImageAssurancePolicyRequest(d)
	log.Printf("[INFO] Creating ImageAssurance policy request %+v\n", req)
	resp, _, err := d9Client.imageAssurancePolicy.Create(&req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created ImageAssurance Policy with ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceImageAssurancePolicyRead(d, meta)
}

func expandImageAssurancePolicyRequest(d *schema.ResourceData) imageassurance_policy.ImageAssurancePolicyRequest {
	return imageassurance_policy.ImageAssurancePolicyRequest{
		TargetId:                        d.Get("target_id").(string),
		RulesetId:                       d.Get("ruleset_id").(int),
		NotificationIds:                 expandNotificationIDs(d, "notification_ids"),
		TargetType:                      d.Get("target_type").(string),
		AdmissionControllerAction:       d.Get("admission_control_action").(string),
		AdmissionControlUnScannedAction: d.Get("admission_control_unscanned_action").(string),
	}
}

func resourceImageAssurancePolicyRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.imageAssurancePolicy.Get(d.Id())

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing ImageAssurance policy %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting ImageAssurance policy: %+v\n", resp)
	d.SetId(resp.ID)
	_ = d.Set("target_id", resp.TargetId)
	_ = d.Set("target_type", resp.TargetType)
	_ = d.Set("ruleset_id", resp.RulesetId)
	_ = d.Set("admission_control_action", resp.AdmissionControllerAction)
	_ = d.Set("admission_control_unscanned_action", resp.AdmissionControlUnScannedAction)
	if err := d.Set("notification_ids", resp.NotificationIds); err != nil {
		return err
	}

	return nil
}

func resourceImageAssurancePolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Updating ImageAssurance policy ID: %v\n", d.Id())
	req := expandImageAssurancePolicyRequest(d)

	if _, _, err := d9Client.imageAssurancePolicy.Update(&req); err != nil {
		return err
	}

	return nil
}

func resourceImageAssurancePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting ImageAssurance policy ID: %v\n", d.Id())

	if _, err := d9Client.imageAssurancePolicy.Delete(d.Id()); err != nil {
		return err
	}
	return nil
}
