package dome9

import (
	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/assessment"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
	"log"
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
			"request_id": {
				Type:     schema.TypeString,
				Optional: true,
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"relevant_count": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"non_complying_count": {
							Type:     schema.TypeString,
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
									"complianceTag": {
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
									"controlTitle": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ruleId": {
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
				Type:     schema.TypeSet,
				Computed: true,
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"general_fetch_permission_issues": {
							Type:     schema.TypeString,
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
			"assessment_passed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"has_errors": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"assessment_id": {
				Type:     schema.TypeString,
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
	d.SetId(string(rune(resp.ID)))

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

	d.SetId(string(rune(resp.ID)))
	_ = d.Set("test_entities", resp.TestEntities)
	_ = d.Set("assessment_passed", resp.AssessmentPassed)
	_ = d.Set("has_errors", resp.HasErrors)
	_ = d.Set("assessment_id", resp.AssessmentId)

	if err := d.Set("request", flattenAssessmentRequest(resp.Request)); err != nil {
		return err
	}

	if err := d.Set("tests", flattenAssessmentTests(resp.Tests)); err != nil {
		return err
	}

	if err := d.Set("location_metadata", flattenAssessmentLocationMetadata(resp.LocationMetadata)); err != nil {
		return err
	}

	if err := d.Set("data_sync_status", flattenAssessmentDataSyncStatus(resp.DataSyncStatus)); err != nil {
		return err
	}

	return nil
}

func resourceAssessmentDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting assessment ID: %v\n", d.Id())
	if _, err := d9Client.cloudaccountAlibaba.Delete(d.Id()); err != nil {
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
			"test_obj":          val.TestObj, //TODO: can cause a problem: interface{}
		}
	}

	return allEntityResults
}

func flattenAssessmentTestsRule(Request assessment.Rule) []interface{} {
	m := map[string]interface{}{
		"name":          Request.Name,
		"severity":      Request.Severity,
		"logic":         Request.Logic,
		"description":   Request.Description,
		"remediation":   Request.Remediation,
		"cloudbots":     Request.Cloudbots,
		"complianceTag": Request.ComplianceTag,
		"domain":        Request.Domain,
		"priority":      Request.Priority,
		"controlTitle":  Request.ControlTitle,
		"ruleId":        Request.RuleID,
		"category":      Request.Category,
		"labels":        Request.Labels, //TODO: can cause a problem: []string
		"logic_hash":    Request.LogicHash,
		"is_default":    Request.IsDefault,
	}

	return []interface{}{m}
}

func flattenAssessmentLocationMetadata(locationMetadata assessment.LocationMetadata) []interface{} {
	m := map[string]interface{}{
		"account": flattenAssessmentLocationMetadataAccount(locationMetadata.Account),
	}

	return []interface{}{m}
}

func flattenAssessmentLocationMetadataAccount(Request assessment.Account) []interface{} {
	m := map[string]interface{}{
		"srl":        Request.Srl,
		"name":       Request.Name,
		"id":         Request.ID,
		"externalId": Request.ExternalID,
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
