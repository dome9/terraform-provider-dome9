package dome9

import (
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/admissioncontrol/admission_policy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"log"
)

func resourceAdmissionPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAdmissionControlPolicyCreate,
		Read:   resourceAdmissionControlPolicyRead,
		Update: resourceAdmissionControlPolicyUpdate,
		Delete: resourceAdmissionControlPolicyDelete,
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
			"action": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"Prevention", "Detection"}, false),
			},
		},
	}
}

func resourceAdmissionControlPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandAdmissionControlPolicyRequest(d)
	log.Printf("[INFO] Creating Admission Control policy request %+v\n", req)
	resp, _, err := d9Client.admissionControlPolicy.Create(&req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created admission control Policy with ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceAdmissionControlPolicyRead(d, meta)
}

func expandAdmissionControlPolicyRequest(d *schema.ResourceData) admission_policy.AdmissionControlPolicyRequest {
	targetType, isExists := d.GetOk("target_type")
	if !isExists {
		targetType = variable.AdmissionControlPolicyTargetType
	}
	return admission_policy.AdmissionControlPolicyRequest{
		TargetId:        d.Get("target_id").(string),
		RulesetId:       d.Get("ruleset_id").(int),
		NotificationIds: expandNotificationIDs(d, "notification_ids"),
		TargetType:      targetType.(string),
		Action:          d.Get("action").(string),
	}
}

func resourceAdmissionControlPolicyRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.admissionControlPolicy.Get(d.Id())

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing admission control policy %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting admission control policy: %+v\n", resp)
	d.SetId(resp.ID)
	_ = d.Set("target_id", resp.TargetId)
	_ = d.Set("target_type", resp.TargetType)
	_ = d.Set("ruleset_id", resp.RulesetId)
	_ = d.Set("action", resp.Action)
	if err := d.Set("notification_ids", resp.NotificationIds); err != nil {
		return err
	}

	return nil
}

func resourceAdmissionControlPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Updating admission control policy ID: %v\n", d.Id())
	req := expandAdmissionControlPolicyRequest(d)

	if _, _, err := d9Client.admissionControlPolicy.Update(&req); err != nil {
		return err
	}

	return nil
}

func resourceAdmissionControlPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting admission control policy ID: %v\n", d.Id())

	if _, err := d9Client.admissionControlPolicy.Delete(d.Id()); err != nil {
		return err
	}
	return nil
}
