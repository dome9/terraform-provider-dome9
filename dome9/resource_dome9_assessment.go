package dome9

import (
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/assessment"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
	"log"
	"strconv"
)

func resourceAssessment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAssessmentCreate,
		Read:   resourceAssessmentRead,
		Update: resourceAssessmentUpdate,
		Delete: resourceAssessmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"bundle_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"dome9_cloud_account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_account_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(providerconst.AssessmentCloudAccountType, false),
			},
			"should_minimize_result": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_cloud_account_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			"test_entities": {
				Type:     schema.TypeMap,
				Computed: true,
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

func resourceAssessmentCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandAssessmentRequest(d)
	log.Printf("[INFO] Creating assessment with request %+v\n", req)

	resp, _, err := d9Client.assessment.Run(&req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created assessment. ID: %v\n", resp.ID)
	d.SetId(strconv.Itoa(resp.ID))

	return resourceAssessmentRead(d, meta)
}

func resourceAssessmentRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.assessment.Get(d.Id())

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing assessment %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(strconv.Itoa(resp.ID))
	_ = d.Set("test_entities", resp.TestEntities)
	_ = d.Set("created_time", resp.CreatedTime)
	_ = d.Set("assessment_id", resp.AssessmentId)
	_ = d.Set("triggered_by", resp.TriggeredBy)
	_ = d.Set("assessment_passed", resp.AssessmentPassed)
	_ = d.Set("has_errors", resp.HasErrors)
	_ = d.Set("has_data_sync_status_issues", resp.HasDataSyncStatusIssues)
	_ = d.Set("comparison_custom_id", resp.ComparisonCustomId)
	_ = d.Set("additional_fields", resp.AdditionalFields)

	if err := d.Set("request", flattenAssessmentRequest(resp.Request)); err != nil {
		return err
	}

	if err := d.Set("tests", flattenAssessmentTests(resp.Tests)); err != nil {
		return err
	}

	if err := d.Set("exclusions", flattenAssessmentExclusion(resp.Exclusions)); err != nil {
		return err
	}

	if err := d.Set("remediations", flattenAssessmentRemediation(resp.Remediations)); err != nil {
		return err
	}

	if err := d.Set("data_sync_status", flattenAssessmentDataSyncStatus(resp.DataSyncStatus)); err != nil {
		return err
	}

	if err := d.Set("stats", flattenAssessmentStats(resp.Stats)); err != nil {
		return err
	}

	return nil
}

func resourceAssessmentDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting assessment ID: %v\n", d.Id())
	if _, err := d9Client.assessment.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceAssessmentUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Println("[WARN] An update can not be made to an assessment")
	return nil
}

func expandAssessmentRequest(d *schema.ResourceData) assessment.RunBundleRequest {
	req := assessment.RunBundleRequest{
		BundleID:               d.Get("bundle_id").(int),
		Dome9CloudAccountID:    d.Get("dome9_cloud_account_id").(string),
		CloudAccountID:         d.Get("cloud_account_id").(string),
		CloudAccountType:       d.Get("cloud_account_type").(string),
		ShouldMinimizeResult:   d.Get("should_minimize_result").(bool),
		Name:                   d.Get("name").(string),
		Description:            d.Get("description").(string),
		ExternalCloudAccountID: d.Get("external_cloud_account_id").(string),
		RequestID:              d.Get("request_id").(string),
	}
	return req
}

func flattenAssessmentRequest(Request assessment.Request) []interface{} {
	m := map[string]interface{}{
		"is_template":               Request.IsTemplate,
		"id":                        Request.BundleID,
		"name":                      Request.Name,
		"description":               Request.Description,
		"dome9_cloud_account_id":    Request.Dome9CloudAccountID,
		"external_cloud_account_id": Request.ExternalCloudAccountID,
		"cloud_account_id":          Request.CloudAccountID,
		"cloud_account_type":        Request.CloudAccountType,
		"request_id":                Request.RequestID,
		"should_minimize_result":    Request.ShouldMinimizeResult,
	}

	return []interface{}{m}
}

func flattenAssessmentTests(tests []assessment.Test) []interface{} {
	allTests := make([]interface{}, len(tests))
	for i, val := range tests {
		allTests[i] = map[string]interface{}{
			"error":               val.Error,
			"tested_count":        val.TestedCount,
			"relevant_count":      val.RelevantCount,
			"non_complying_count": val.NonComplyingCount,
			"exclusion_stats":     flattenAssessmentTestsExclusionStats(val.ExclusionStats),
			"entity_results":      flattenAssessmentTestsEntityResults(val.EntityResults),
			"rule":                flattenAssessmentTestsRule(val.Rule),
			"test_passed":         val.TestPassed,
		}
	}

	return allTests
}

func flattenAssessmentTestsExclusionStats(Request assessment.ExclusionStats) []interface{} {
	m := map[string]interface{}{
		"tested_count":        Request.TestedCount,
		"relevant_count":      Request.RelevantCount,
		"non_complying_count": Request.NonComplyingCount,
	}

	return []interface{}{m}
}

func flattenAssessmentTestsEntityResults(entityResults []assessment.EntityResult) []interface{} {
	allEntityResults := make([]interface{}, len(entityResults))
	for i, val := range entityResults {
		allEntityResults[i] = map[string]interface{}{
			"validation_status": val.ValidationStatus,
			"is_relevant":       val.IsRelevant,
			"is_valid":          val.IsValid,
			"is_excluded":       val.IsExcluded,
			"exclusion_id":      val.ExclusionID,
			"remediation_id":    val.RemediationID,
			"error":             val.Error,
			"test_obj":          flattenAssessmentTestsEntityResultsTestObj(val.TestObj),
		}
	}

	return allEntityResults
}

func flattenAssessmentTestsEntityResultsTestObj(Request assessment.RuleEngineFailedEntityReference) []interface{} {
	m := map[string]interface{}{
		"id":                            Request.Id,
		"dome9_id":                      Request.Dome9Id,
		"entity_type":                   Request.EntityType,
		"entity_index":                  Request.EntityIndex,
		"custom_entity_comparison_hash": Request.CustomEntityComparisonHash,
	}

	return []interface{}{m}
}

func flattenAssessmentTestsRule(Request assessment.Rule) []interface{} {
	m := map[string]interface{}{
		"name":           Request.Name,
		"severity":       Request.Severity,
		"logic":          Request.Logic,
		"description":    Request.Description,
		"remediation":    Request.Remediation,
		"cloudbots":      Request.Cloudbots,
		"compliance_tag": Request.ComplianceTag,
		"domain":         Request.Domain,
		"priority":       Request.Priority,
		"control_title":  Request.ControlTitle,
		"rule_id":        Request.RuleID,
		"category":       Request.Category,
		"labels":         Request.Labels,
		"logic_hash":     Request.LogicHash,
		"is_default":     Request.IsDefault,
	}

	return []interface{}{m}
}

func flattenAssessmentExclusion(Exclusions []assessment.Exclusion) []interface{} {
	allExclusions := make([]interface{}, len(Exclusions))
	for i, val := range Exclusions {
		allExclusions[i] = map[string]interface{}{
			"platform":                val.Platform,
			"id":                      val.ID,
			"rules":                   flattenAssessmentExclusionOrRemediationRule(val.Rules),
			"logic_expressions":       val.LogicExpressions,
			"ruleset_id":              val.RulesetId,
			"cloud_account_ids":       val.CloudAccountIds,
			"comment":                 val.Comment,
			"organizational_unit_ids": val.OrganizationalUnitIds,
			"date_range":              flattenAssessmentDateRange(val.DateRange),
		}
	}

	return allExclusions
}

func flattenAssessmentRemediation(Exclusions []assessment.Remediation) []interface{} {
	allExclusions := make([]interface{}, len(Exclusions))
	for i, val := range Exclusions {
		allExclusions[i] = map[string]interface{}{
			"platform":                val.Platform,
			"id":                      val.ID,
			"rules":                   flattenAssessmentExclusionOrRemediationRule(val.Rules),
			"logic_expressions":       val.LogicExpressions,
			"ruleset_id":              val.RulesetId,
			"cloud_account_ids":       val.CloudAccountIds,
			"comment":                 val.Comment,
			"cloudBots":               val.CloudBots,
			"organizational_unit_ids": val.OrganizationalUnitIds,
			"date_range":              flattenAssessmentDateRange(val.DateRange),
		}
	}

	return allExclusions
}

func flattenAssessmentExclusionOrRemediationRule(ExclusionOrRemediationRule []assessment.ExclusionOrRemediationRule) []interface{} {
	allExclusionOrRemediationRule := make([]interface{}, len(ExclusionOrRemediationRule))
	for i, val := range ExclusionOrRemediationRule {
		allExclusionOrRemediationRule[i] = map[string]interface{}{
			"logic_hash": val.LogicHash,
			"id":         val.ID,
			"name":       val.Name,
		}
	}

	return allExclusionOrRemediationRule
}

func flattenAssessmentDateRange(Request assessment.Date) []interface{} {
	m := map[string]interface{}{
		"from": Request.From,
		"to":   Request.To,
	}

	return []interface{}{m}
}

func flattenAssessmentDataSyncStatus(dataSyncStatus []assessment.DataSyncStatus) []interface{} {
	allDataSyncStatus := make([]interface{}, len(dataSyncStatus))
	for i, val := range dataSyncStatus {
		allDataSyncStatus[i] = map[string]interface{}{
			"entity_type":                     val.EntityType,
			"recently_successful_sync":        val.RecentlySuccessfulSync,
			"general_fetch_permission_issues": val.GeneralFetchPermissionIssues,
			"entities_with_permission_issues": flattenAssessmentDataSyncStatusEntitiesWithPermissionIssues(val.EntitiesWithPermissionIssues),
		}
	}

	return allDataSyncStatus
}

func flattenAssessmentDataSyncStatusEntitiesWithPermissionIssues(entitiesWithPermissionIssues []assessment.EntitiesWithPermissionIssues) []interface{} {
	allTests := make([]interface{}, len(entitiesWithPermissionIssues))
	for i, val := range entitiesWithPermissionIssues {
		allTests[i] = map[string]interface{}{
			"external_id":             val.ExternalID,
			"name":                    val.Name,
			"cloud_vendor_identifier": val.CloudVendorIdentifier,
		}
	}

	return allTests
}

func flattenAssessmentStats(Request assessment.Stats) []interface{} {
	m := map[string]interface{}{
		"passed":                     Request.Passed,
		"passed_rules_by_severity":   flattenAssessmentStatsRulesSeverity(Request.PassedRulesBySeverity),
		"failed":                     Request.Failed,
		"failed_rules_by_severity":   flattenAssessmentStatsRulesSeverity(Request.FailedRulesBySeverity),
		"error":                      Request.Error,
		"failed_tests":               Request.FailedTests,
		"logically_tested":           Request.LogicallyTested,
		"failed_entities":            Request.FailedEntities,
		"excluded_tests":             Request.ExcludedTests,
		"excluded_failed_tests":      Request.ExcludedFailedTests,
		"excluded_rules":             Request.ExcludedRules,
		"excluded_rules_by_severity": flattenAssessmentStatsRulesSeverity(Request.ExcludedRulesBySeverity),
	}

	return []interface{}{m}
}

func flattenAssessmentStatsRulesSeverity(Request assessment.RulesSeverity) []interface{} {
	m := map[string]interface{}{
		"informational": Request.Informational,
		"low":           Request.Low,
		"medium":        Request.Medium,
		"high":          Request.High,
		"critical":      Request.Critical,
	}

	return []interface{}{m}
}

/*
func flattenAssessment(Request assessment.Request) []interface{} {
	m := map[string]interface{}{
		"":    Request.,

	}

	return []interface{}{m}
}

func flattenAssessment(Tests []assessment.Test) []interface{} {
	allTests := make([]interface{}, len(Tests))
	for i, val := range Tests {
		allTests[i] = map[string]interface{}{
			"":  val.,

		}
	}

	return allTests
}

*/
