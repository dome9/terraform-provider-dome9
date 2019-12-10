package dome9

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/roles"
)

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleCreate,
		Read:   resourceRoleRead,
		Update: resourceRoleUpdate,
		Delete: resourceRoleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"permissions": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"manage": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"rulesets": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"notifications": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"policies": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"alert_actions": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"create": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"view": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"on_boarding": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"cross_account_access": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func resourceRoleCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandRoleCreateRequest(d)
	log.Printf("[INFO] Creating dome9 role with request\n%+v\n", req)

	role, _, err := d9Client.role.Create(req)
	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(role.ID))

	return resourceRoleRead(d, meta)
}

func resourceRoleRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.role.Get(d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing role %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting role:\n%+v\n", resp)
	d.SetId(strconv.Itoa(resp.ID))
	_ = d.Set("name", resp.Name)
	_ = d.Set("description", resp.Description)
	_ = d.Set("permissions", flattenPermission(resp.Permissions))

	return nil
}

func resourceRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	id := d.Id()
	log.Printf("[INFO] Updating role ID: %v\n", id)
	req := expandRoleCreateRequest(d)

	if _, err := d9Client.role.Update(id, req); err != nil {
		return err
	}

	return resourceRoleRead(d, meta)
}

func resourceRoleDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting role with id %v\n", d.Id())

	if _, err := d9Client.role.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func expandRoleCreateRequest(d *schema.ResourceData) roles.RoleRequest {
	return roles.RoleRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Permissions: expandPermissions(d),
	}
}

func expandPermissions(d *schema.ResourceData) roles.Permissions {
	if permissions, ok := d.GetOk("permissions"); ok {
		permissionItem := permissions.(*schema.Set).List()[0]
		permission := permissionItem.(map[string]interface{})

		return roles.Permissions{
			Access:             expandList(permission["access"].([]interface{})),
			Manage:             expandList(permission["manage"].([]interface{})),
			Rulesets:           expandList(permission["rulesets"].([]interface{})),
			Notifications:      expandList(permission["notifications"].([]interface{})),
			Policies:           expandList(permission["policies"].([]interface{})),
			AlertActions:       expandList(permission["alert_actions"].([]interface{})),
			Create:             expandList(permission["create"].([]interface{})),
			View:               expandList(permission["view"].([]interface{})),
			OnBoarding:         expandList(permission["on_boarding"].([]interface{})),
			CrossAccountAccess: expandList(permission["cross_account_access"].([]interface{})),
		}
	}
	return roles.Permissions{}
}

func expandList(permissionsLst []interface{}) []string {
	permissions := make([]string, len(permissionsLst))
	for i, permission := range permissionsLst {
		permissions[i] = permission.(string)
	}

	return permissions
}

func flattenPermission(respPermission roles.Permissions) []interface{} {
	m := map[string]interface{}{
		"access":               respPermission.Access,
		"manage":               respPermission.Manage,
		"rulesets":             respPermission.Rulesets,
		"notifications":        respPermission.Notifications,
		"policies":             respPermission.Policies,
		"alert_actions":        respPermission.AlertActions,
		"create":               respPermission.Create,
		"view":                 respPermission.View,
		"on_boarding":          respPermission.OnBoarding,
		"cross_account_access": respPermission.CrossAccountAccess,
	}

	return []interface{}{m}
}
