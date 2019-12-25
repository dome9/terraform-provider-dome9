package dome9

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/roles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
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
			"permit_rulesets": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"permit_notifications": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"permit_policies": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"permit_alert_actions": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"permit_on_boarding": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cross_account_access": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"create": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"access": srlDescriptorSchema(),
			"view":   srlDescriptorSchema(),
			"manage": srlDescriptorSchema(),
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
	permissions := roles.Permissions{
		Access:             generateSRL(d.Get("access").([]interface{})),
		Manage:             generateSRL(d.Get("manage").([]interface{})),
		Create:             expandList(d.Get("create").([]interface{})),
		View:               generateSRL(d.Get("view").([]interface{})),
		CrossAccountAccess: expandList(d.Get("cross_account_access").([]interface{})),
	}

	if permitRulesets, ok := d.GetOk("permit_rulesets"); ok {
		permissions.Rulesets = convertBoolToSRL(permitRulesets.(bool))
	}
	if permitNotifications, ok := d.GetOk("permit_notifications"); ok {
		permissions.Notifications = convertBoolToSRL(permitNotifications.(bool))
	}
	if permitPolicies, ok := d.GetOk("permit_policies"); ok {
		permissions.Policies = convertBoolToSRL(permitPolicies.(bool))
	}
	if permitAlertActions, ok := d.GetOk("permit_alert_actions"); ok {
		permissions.AlertActions = convertBoolToSRL(permitAlertActions.(bool))
	}
	if permitOnBoarding, ok := d.GetOk("permit_on_boarding"); ok {
		permissions.OnBoarding = convertBoolToSRL(permitOnBoarding.(bool))
	}

	return permissions
}

func expandList(permissionsLst []interface{}) []string {
	permissions := make([]string, len(permissionsLst))
	for i, permission := range permissionsLst {
		permissions[i] = permission.(string)
	}

	return permissions
}

func convertBoolToSRL(status bool) []string {
	if status {
		return []string{""}
	}
	return []string{}
}

func generateSRL(attributes []interface{}) []string {
	srlList := make([]string, len(attributes))
	for i, attr := range attributes {
		if attr != nil {
			dict := attr.(map[string]interface{})

			// Checking value not empty since d.Get() returns empty strings as default for un given optional fields
			if val := dict["type"].(string); val != "" {
				srlList[i] = providerconst.SRlType[val]
			}
			if val := dict["main_id"].(string); val != "" {
				appendSRLMember(&srlList[i], val)
			}
			if val := dict["region"].(string); val != "" {
				appendSRLMember(&srlList[i], "rg")
				appendSRLMember(&srlList[i], providerconst.AWSRegionsEnum[val])
			}
			if val := dict["security_group_id"].(string); val != "" {
				appendSRLMember(&srlList[i], "sg")
				appendSRLMember(&srlList[i], val)
			}
			if val := dict["traffic"].(string); val != "" {
				appendSRLMember(&srlList[i], providerconst.PermissionTrafficType[val])
			}
		}
	}
	log.Printf("[DEBUG] SRL list %v generated", srlList)
	return srlList
}

func appendSRLMember(srl *string, addition string) {
	if addition != "" {
		*srl = fmt.Sprintf("%s%s%s", *srl, "|", addition)
	}
}

func breakSRL(srlList []string) []map[string]string {
	// this function is used for breaking SRL string into data structure to fit resource schema
	ret := make([]map[string]string, len(srlList))
	for i, val := range srlList {
		log.Printf("[DEBUG] SRL List Lenght %v ", len(srlList))
		srlSplit := strings.Split(val, "|")
		tempMap := make(map[string]string)

		for _, srlMember := range providerconst.SRLStructure {
			if len(srlSplit) == 0 {
				break

			} else {
				log.Printf("[DEBUG] SRL split list %v generated", srlSplit)
				x := srlSplit[0]

				switch srlMember {
				case "type":
					tempMap[srlMember] = getKeyByValue(providerconst.SRlType, x)

				case "main_id", "security_group_id":
					tempMap[srlMember] = x

				case "rg", "sg":
					srlSplit = srlSplit[1:] // chopping srl
					continue

				case "region":
					tempMap[srlMember] = getKeyByValue(providerconst.AWSRegionsEnum, x)

				case "traffic":
					tempMap[srlMember] = getKeyByValue(providerconst.PermissionTrafficType, x)

				default:
					log.Printf("unrecognized SRL member")
				}
				srlSplit = srlSplit[1:] // chopping srl
			}
		}
		ret[i] = tempMap
	}

	log.Printf("[DEBUG] RET: %#v generated", ret)
	return ret
}

func getKeyByValue(m map[string]string, value string) (key string) {
	key, ok := mapKey(m, value)
	if !ok {
		log.Printf("value %s does not exist in map", key)
	}
	return key
}

func mapKey(m map[string]string, value string) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func isEmpty(lst []string) bool {
	return len(lst) != 0
}
