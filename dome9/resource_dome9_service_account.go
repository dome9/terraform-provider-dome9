package dome9

import (
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/serviceaccounts"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceServiceAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceAccountCreate,
		Read:   resourceServiceAccountRead,
		Update: resourceServiceAccountUpdate,
		Delete: resourceServiceAccountDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"api_key_secret": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role_ids": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func resourceServiceAccountCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	var roleIds []int64
	for _, i := range d.Get("role_ids").(*schema.Set).List() {
		roleIds = append(roleIds, int64(i.(int)))
	}
	req := serviceaccounts.ServiceAccountRequest{
		Name: d.Get("name").(string),
		RoleIds: roleIds,
	}
	log.Printf("[INFO] Creating service account request\n%+v\n", req)
	resp, _, err := d9Client.serviceAccounts.Create(&req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created service account request. ID: %v\n", resp.Id)
	d.SetId(resp.Id)
	_ = d.Set("api_key_id", resp.ApiKeyId)
	_ = d.Set("api_key_secret", resp.ApiKeySecret)

	return resourceServiceAccountRead(d, meta)
}

func resourceServiceAccountRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	resp, _, err := d9Client.serviceAccounts.Get(d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing service account %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting service account:\n%+v\n", resp)
	d.SetId(resp.Id)
	_ = d.Set("name", resp.Name)
	_ = d.Set("api_key_id", resp.ApiKeyId)
	_ = d.Set("role_ids", resp.RoleIds)
	return nil
}

func resourceServiceAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Updating service account ID: %v\n", d.Id())

	var roleIds []int64
	for _, i := range d.Get("role_ids").(*schema.Set).List() {
		roleIds = append(roleIds, int64(i.(int)))
	}

	req := serviceaccounts.UpdateServiceAccountRequest{
		Name: d.Get("name").(string),
		Id: d.Id(),
		RoleIds: roleIds,
	}

	_, _, err := d9Client.serviceAccounts.Update(&req)
	if err != nil {
		return err
	}
	return nil
}

func resourceServiceAccountDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting service account ID: %v", d.Id())

	_, err := d9Client.serviceAccounts.Delete(d.Id())
	if err != nil {
		return err
	}
	return nil
}