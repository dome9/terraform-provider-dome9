package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceContinuousComplianceFinding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceContinuousComplianceFindingRead,

		Schema: map[string]*schema.Schema{
			"page_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"sorting": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"direction": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"multi_sorting": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"direction": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"filter": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"free_text_phrase": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"fields": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Optional: false,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: false,
									},
								},
							},
						},
						"only_ciem": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"included_features": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"creation_time": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from": {
										Type:     schema.TypeString,
										Optional: false,
									},
									"to": {
										Type:     schema.TypeString,
										Optional: false,
									},
								},
							},
						},
					},
				},
			},
			"search_after": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"data_source": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"search_request": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"page_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"sorting": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"field_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"direction": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"multi_sorting": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"field_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"direction": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"filter": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"free_text_phrase": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fields": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Computed: false,
												},
												"value": {
													Type:     schema.TypeString,
													Computed: false,
												},
											},
										},
									},
									"only_ciem": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"included_features": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"creation_time": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"from": {
													Type:     schema.TypeString,
													Computed: false,
												},
												"to": {
													Type:     schema.TypeString,
													Computed: false,
												},
											},
										},
									},
								},
							},
						},
						"search_after": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"data_source": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"findings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"finding_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cloud_account_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"comments": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"text": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"timestamp": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"cloud_account_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cloud_account_external_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organizational_unit_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organizational_unit_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bundle_id": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"alert_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rule_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rule_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rule_logic": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"entity_dome9_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"entity_external_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"entity_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"entity_type_by_environment_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"entity_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"entity_network": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"entity_tags": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"severity": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remediation": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bundle_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"acknowledged": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"origin": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_seen_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"owner_user_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"magellan": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alert_window_start_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"alert_window_end_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"is_excluded": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"webhook_responses": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"request_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"response_content": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"remediation_actions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"additional_fields": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"occurrences": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"scan_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceContinuousComplianceFindingRead(d *schema.ResourceData, meta interface{}) error {
	//d9Client := meta.(*Client)
	//req := expandContinuousComplianceFindingRequest(d)
	//log.Printf("[INFO] Creating compliance policy request %+v\n", req)
	//resp, _, err := d9Client.continuousCompliancePolicy.Create(&req)
	//if err != nil {
	//	return err
	//}
	//
	//log.Printf("[INFO] Created compliance policy with ID: %v\n", resp.ID)
	//d.SetId(resp.ID)
	//
	//return resourceContinuousCompliancePolicyRead(d, meta)
}

//func dataSourceContinuousComplianceFindingRead(d *schema.ResourceData, meta interface{}) error {
//	d9Client := meta.(*Client)
//
//	policyID := d.Get("id").(string)
//	log.Printf("Getting data for Continuous Compliance Policy id: %s\n", policyID)
//
//	resp, _, err := d9Client.continuousCompliancePolicy.Get(policyID)
//	if err != nil {
//		return err
//	}
//
//	d.SetId(resp.ID)
//	_ = d.Set("target_internal_id", resp.TargetInternalId)
//	_ = d.Set("target_external_id", resp.TargetExternalId)
//	_ = d.Set("target_type", resp.TargetType)
//	_ = d.Set("ruleset_id", resp.RulesetId)
//	if err := d.Set("notification_ids", resp.NotificationIds); err != nil {
//		return err
//	}
//
//	return nil
//}
