package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceServiceAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServiceAccountRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role_ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"date_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_used": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceServiceAccountRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for service account id %s\n", id)

	resp, _, err := d9Client.serviceAccounts.Get(id)
	if err != nil {
		return err
	}

	d.SetId(resp.Id)
	_ = d.Set("name", resp.Name)
	_ = d.Set("api_key_id", resp.ApiKeyId)
	_ = d.Set("role_ids", resp.RoleIds)
	_ = d.Set("date_created", resp.DateCreated.Format("2006-01-02 15:04:05"))
	_ = d.Set("last_used", resp.LastUsed.Format("2006-01-02 15:04:05"))

	return nil
}
