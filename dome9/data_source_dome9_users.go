package dome9

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUsersRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_sso_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_suspended": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_owner": {
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
			"role_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
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
			"calculated_permissions": {
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
			"is_mobile_device_paired": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceUsersRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data user with id %s\n", id)

	resp, _, err := d9Client.users.Get(id)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(resp.ID))
	_ = d.Set("email", resp.Name)
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
	_ = d.Set("permissions", flattenPermissions(resp.Permissions))
	_ = d.Set("calculated_permissions", flattenPermissions(resp.Permissions))
	_ = d.Set("is_mobile_device_paired", resp.IsMobileDevicePaired)

	return nil
}
