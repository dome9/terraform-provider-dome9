package dome9

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
			"permit_rulesets": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"permit_notifications": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"permit_policies": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"permit_alert_actions": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"permit_on_boarding": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"cross_account_access": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"create": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"access": srlDescriptorDataSchema(),
			"view":   srlDescriptorDataSchema(),
			"manage": srlDescriptorDataSchema(),
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
	_ = d.Set("access", breakSRL(resp.Permissions.Access))
	_ = d.Set("manage", breakSRL(resp.Permissions.Manage))
	_ = d.Set("view", breakSRL(resp.Permissions.View))
	_ = d.Set("permit_rulesets", isEmpty(resp.Permissions.Rulesets))
	_ = d.Set("permit_notifications", isEmpty(resp.Permissions.Notifications))
	_ = d.Set("permit_policies", isEmpty(resp.Permissions.Policies))
	_ = d.Set("permit_alert_actions", isEmpty(resp.Permissions.AlertActions))
	_ = d.Set("permit_on_boarding", isEmpty(resp.Permissions.OnBoarding))
	_ = d.Set("create", resp.Permissions.Create)
	_ = d.Set("cross_account_access", resp.Permissions.CrossAccountAccess)

	return nil
}
