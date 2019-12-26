package dome9

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/users"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_sso_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"role_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"is_owner": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
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
			// computed fields
			"is_suspended": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_super_user": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_auditor": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"has_api_key": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"has_api_key_v1": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"has_api_key_v2": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_mfa_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"iam_safe": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloud_accounts": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cloud_account_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"external_account_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_lease_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"iam_entities": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"iam_entities_last_lease_time": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"iam_entity": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"last_lease_time": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"cloud_account_state": {
										Computed: true,
										Type:     schema.TypeString,
									},
									"iam_entity": {
										Computed: true,
										Type:     schema.TypeString,
									},
								},
							},
						},
					},
				},
			},
			"can_switch_role": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_locked": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_login": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_mobile_device_paired": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandUserRequest(d)
	log.Printf("[INFO] Creating user with request\n%+v\n", req)
	resp, _, err := d9Client.users.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created user. ID: %v\n", resp.ID)
	d.SetId(strconv.Itoa(resp.ID))

	return resourceUserRead(d, meta)
}

func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.users.Get(d.Id())

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing user %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Reading user and settings states: %+v\n", resp)
	d.SetId(strconv.Itoa(resp.ID))
	_ = d.Set("email", resp.Email)
	_ = d.Set("is_sso_enabled", resp.SsoEnabled)
	_ = d.Set("is_suspended", resp.IsSuspended)
	_ = d.Set("is_owner", resp.IsOwner)
	_ = d.Set("is_super_user", resp.IsSuperUser)
	_ = d.Set("is_auditor", resp.IsAuditor)
	_ = d.Set("has_api_key", resp.HasAPIKey)
	_ = d.Set("has_api_key_v1", resp.HasAPIKeyV1)
	_ = d.Set("has_api_key_v2", resp.HasAPIKeyV2)
	_ = d.Set("is_mfa_enabled", resp.IsMfaEnabled)
	_ = d.Set("role_ids", resp.RoleIds)
	_ = d.Set("iam_safe", flattenIamSafe(resp.IamSafe))
	_ = d.Set("can_switch_role", resp.CanSwitchRole)
	_ = d.Set("is_locked", resp.IsLocked)
	_ = d.Set("last_login", resp.LastLogin.Format("2006-01-02 15:04:05"))
	_ = d.Set("is_mobile_device_paired", resp.IsMobileDevicePaired)
	// set permissions
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

func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting user ID: %v\n", d.Id())

	if _, err := d9Client.users.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Updating user with ID: %v\n", d.Id())

	if d.HasChange("is_owner") {
		if d.Get("is_owner").(bool) {
			if _, err := d9Client.users.SetUserAsOwner(d.Id()); err != nil {
				return err
			}

			log.Printf("User with ID %s now is owner", d.Id())
		} else {
			// to drop ownership from user, we must set another user to be owner, so is_owner field in the tf state must stay true
			_ = d.Set("is_owner", true)
		}
	} else {
		log.Println("[INFO] Roles id's or permissions has been changed")
		req := expandUpdateRequest(d)
		if _, err := d9Client.users.Update(d.Id(), &req); err != nil {
			return err
		}
	}

	return nil
}

func expandUserRequest(d *schema.ResourceData) users.UserRequest {
	return users.UserRequest{
		Email:       d.Get("email").(string),
		FirstName:   d.Get("first_name").(string),
		LastName:    d.Get("last_name").(string),
		SsoEnabled:  d.Get("is_sso_enabled").(bool),
		Permissions: expandUsersPermissions(d),
	}
}

func expandUpdateRequest(d *schema.ResourceData) users.UserUpdate {
	return users.UserUpdate{
		RoleIds:     expandRoles(d.Get("role_ids").([]interface{})),
		Permissions: expandUsersPermissions(d),
	}
}

func expandUsersPermissions(d *schema.ResourceData) users.Permissions {
	permissions := users.Permissions{
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

func expandRoles(generalRoles []interface{}) []int {
	roles := make([]int, len(generalRoles))
	for i, role := range generalRoles {
		roles[i] = role.(int)
	}

	return roles
}

func flattenIamSafe(iamSafe users.IamSafe) []interface{} {
	m := map[string]interface{}{
		"cloud_accounts": flattenIamSafeCloudAccounts(iamSafe.CloudAccounts),
	}

	return []interface{}{m}
}

func flattenIamSafeCloudAccounts(iamSafeCloudAccounts []users.CloudAccounts) []interface{} {
	cloudAccounts := make([]interface{}, len(iamSafeCloudAccounts))
	for i, val := range iamSafeCloudAccounts {
		cloudAccounts[i] = map[string]interface{}{
			"cloud_account_id":             val.CloudAccountID,
			"name":                         val.Name,
			"external_account_number":      val.ExternalAccountNumber,
			"last_lease_time":              val.LastLeaseTime.Format("2006-01-02 15:04:05"),
			"state":                        val.State,
			"iam_entities":                 val.IamEntities,
			"iam_entities_last_lease_time": flattenIamEntitiesLastLeaseTime(val.IamEntitiesLastLeaseTime),
			"cloud_account_state":          val.CloudAccountState,
			"iam_entity":                   val.IamEntity,
		}
	}

	return cloudAccounts
}

func flattenIamEntitiesLastLeaseTime(iamEntitiesLastLeaseTime []users.IamEntitiesLastLeaseTime) []interface{} {
	iamEntitiesLastLeaseTimes := make([]interface{}, len(iamEntitiesLastLeaseTime))
	for i, val := range iamEntitiesLastLeaseTime {
		iamEntitiesLastLeaseTimes[i] = map[string]interface{}{
			"iam_entity":          val.IamEntity,
			"cloud_account_state": val.LastLeaseTime.Format("2006-01-02 15:04:05"),
		}
	}

	return iamEntitiesLastLeaseTimes
}
