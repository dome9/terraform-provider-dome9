package dome9

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceRole() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRoleRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permissions": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"manage": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"rulesets": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"notifications": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"policies": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"alert_actions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"create": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"view": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"on_boarding": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"cross_account_access": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func dataSourceRoleRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for role %s\n", id)

	resp, _, err := d9Client.role.Get(id)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(resp.ID))
	_ = d.Set("name", resp.Name)
	_ = d.Set("description", resp.Description)
	_ = d.Set("permissions", flattenPermission(resp.Permissions))

	return nil
}
