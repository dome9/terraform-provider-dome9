package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceAssessment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAssessmentRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"bundle_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dome9_cloud_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_account_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"should_minimize_result": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_cloud_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"request": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_template": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dome9_cloud_account_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cloud_account_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cloud_account_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"should_minimize_result": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"external_cloud_account_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"request_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tests": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tested_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"relevant_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"non_complying_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"exclusion_stats": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tested_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"relevant_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"non_complying_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"entity_results": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"validation_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_relevant": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_valid": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_excluded": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"exclusion_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remediation_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"error": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"test_obj": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"dome9_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"entity_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"entity_index": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"custom_entity_comparison_hash": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"rule": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"logic": {
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
									"cloudbots": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compliance_tag": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"priority": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"control_title": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"rule_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"category": {
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
									"logic_hash": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_default": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"test_passed": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"location_metadata": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"srl": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"external_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"exclusions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"platform": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rules": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"logic_hash": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"logic_expressions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ruleset_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"cloud_account_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"comment": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organizational_unit_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"date_range": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"to": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"remediations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"platform": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rules": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"logic_hash": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"logic_expressions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ruleset_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"cloud_account_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"comment": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cloud_bots": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"organizational_unit_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"date_range": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"to": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"data_sync_status": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entity_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"recently_successful_sync": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"general_fetch_permission_issues": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"entities_with_permission_issues": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"external_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cloud_vendor_identifier": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"created_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assessment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"triggered_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assessment_passed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"has_errors": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"stats": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"passed": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"passed_rules_by_severity": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"informational": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"low": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"medium": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"high": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"critical": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"failed": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"failed_rules_by_severity": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"informational": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"low": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"medium": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"high": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"critical": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"error": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"failed_tests": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"logically_tested": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"failed_entities": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"excluded_tests": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"excluded_failed_tests": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"excluded_rules": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"excluded_rules_by_severity": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"informational": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"low": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"medium": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"high": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"critical": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"has_data_sync_status_issues": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"comparison_custom_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"additional_fields": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func dataSourceAssessmentRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := strconv.Itoa(d.Get("id").(int))
	log.Printf("Getting data for assessment with id %s\n", id)

	assessmentData, _, err := d9Client.assessment.Get(id)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(assessmentData.ID))
	_ = d.Set("request", flattenAssessmentRequest(assessmentData.Request))
	_ = d.Set("tests", flattenAssessmentTests(assessmentData.Tests))
	_ = d.Set("test_entities", assessmentData.TestEntities)
	_ = d.Set("exclusions", flattenAssessmentExclusion(assessmentData.Exclusions))
	_ = d.Set("remediations", flattenAssessmentRemediation(assessmentData.Remediations))
	_ = d.Set("data_sync_status", flattenAssessmentDataSyncStatus(assessmentData.DataSyncStatus))
	_ = d.Set("created_time", assessmentData.CreatedTime)
	_ = d.Set("assessment_id", assessmentData.AssessmentId)
	_ = d.Set("triggered_by", assessmentData.TriggeredBy)
	_ = d.Set("assessment_passed", assessmentData.AssessmentPassed)
	_ = d.Set("has_errors", assessmentData.HasErrors)
	_ = d.Set("stats", flattenAssessmentStats(assessmentData.Stats))
	_ = d.Set("has_data_sync_status_issues", assessmentData.HasDataSyncStatusIssues)
	_ = d.Set("comparison_custom_id", assessmentData.ComparisonCustomId)
	_ = d.Set("additional_fields", assessmentData.AdditionalFields)

	return nil
}
