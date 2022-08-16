package dome9

import (
	"encoding/json"
	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_finding"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
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
	d9Client := meta.(*Client)
	req := expandContinuousComplianceFindingRequest(d)
	log.Printf("[INFO] Executing continuous compliance finding search with request %+v\n", req)
	resp, _, err := d9Client.continuousComplianceFinding.Search(&req)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Successfully executed continuous compliance finding search with response %+v\n", resp)

	log.Printf("[INFO] Start flattening continuous compliance finding search response\n")
	if err := d.Set("search_request", flattenFindingResponseSearchRequest(resp.SearchRequest)); err != nil {
		return err
	}

	if err := d.Set("findings", flattenFindingResponseFindings(resp.Findings)); err != nil {
		return err
	}

	_ = d.Set("total_findings_count", resp.TotalFindingsCount)

	if err := d.Set("aggregations", flattenFindingResponseAggregations(resp.Aggregations)); err != nil {
		return err
	}

	_ = d.Set("search_after", resp.SearchAfter)
	log.Printf("[INFO] Successfuly finished flattening continuous compliance finding search response\n")

	return nil
}

func flattenFindingResponseAggregations(aggregations map[string][]continuous_compliance_finding.FieldAggregation) []interface{} {
	if aggregations == nil {
		return nil
	}
	allAggregations := make([]interface{}, len(aggregations))
	for _, aggregation := range aggregations {
		agg := make([]interface{}, len(aggregation))
		for _, fieldAggregation := range aggregation {
			agg = append(agg, map[string]interface{}{
				"value": fieldAggregation.Value,
				"count": fieldAggregation.Count,
			})
		}
		allAggregations = append(allAggregations, agg)
	}
	return allAggregations
}

func flattenFindingResponseFindings(findings []continuous_compliance_finding.Finding) []interface{} {
	if findings == nil {
		return nil
	}
	allFindings := make([]interface{}, len(findings))
	for i, val := range findings {
		allFindings[i] = map[string]interface{}{
			"id":                          val.Id,
			"FindingKey":                  val.FindingKey,
			"CreatedTime":                 val.CreatedTime,
			"UpdatedTime":                 val.UpdatedTime,
			"CloudAccountType":            val.CloudAccountType,
			"Comments":                    flattenFindingResponseFindingsComments(val.Comments),
			"CloudAccountId":              val.CloudAccountId,
			"CloudAccountExternalId":      val.CloudAccountExternalId,
			"OrganizationalUnitId":        val.OrganizationalUnitId,
			"OrganizationalUnitPath":      val.OrganizationalUnitPath,
			"BundleId":                    val.BundleId,
			"AlertType":                   val.AlertType,
			"RuleId":                      val.RuleId,
			"RuleName":                    val.RuleName,
			"RuleLogic":                   val.RuleLogic,
			"EntityDome9Id":               val.EntityDome9Id,
			"EntityExternalId":            val.EntityExternalId,
			"EntityType":                  val.EntityType,
			"EntityTypeByEnvironmentType": val.EntityTypeByEnvironmentType,
			"EntityName":                  val.EntityName,
			"EntityNetwork":               val.EntityNetwork,
			"EntityTags":                  flattenFindingResponseFindingsEntityTags(val.EntityTags),
			"Severity":                    val.Severity,
			"Description":                 val.Description,
			"Remediation":                 val.Remediation,
			"Tag":                         val.Tag,
			"Region":                      val.Region,
			"BundleName":                  val.BundleName,
			"Acknowledged":                val.Acknowledged,
			"Origin":                      val.Origin,
			"LastSeenTime":                val.LastSeenTime,
			"OwnerUserName":               val.OwnerUserName,
			"Magellan":                    flattenFindingResponseFindingsMagellan(val.Magellan),
			"IsExcluded":                  val.IsExcluded,
			"WebhookResponses":            flattenFindingResponseFindingsWebhookResponses(val.WebhookResponses),
			"RemediationActions":          val.RemediationActions,
			"AdditionalFields":            flattenFindingResponseFindingsAdditionalFields(val.AdditionalFields),
			"Occurrences":                 val.Occurrences,
			"ScanId":                      val.ScanId,
			"Status":                      val.Status,
			"Category":                    val.Category,
			"Action":                      val.Action,
			"Labels":                      val.Labels,
		}
	}

	return allFindings
}

func flattenFindingResponseFindingsAdditionalFields(fields []continuous_compliance_finding.AdditionalField) []interface{} {
	if fields == nil {
		return nil
	}
	allComments := make([]interface{}, len(fields))
	for i, val := range fields {
		allComments[i] = map[string]interface{}{
			"name":  val.Name,
			"value": val.Value,
		}
	}

	return allComments
}

func flattenFindingResponseFindingsWebhookResponses(responses map[string]continuous_compliance_finding.WebhookResponse) ([]interface{}, error) {
	if responses == nil {
		return nil, nil
	}

	allResponses := make([]interface{}, len(responses))
	for _, val := range responses {
		responseContent, err := flattenFindingResponseFindingsWebhookResponsesResponseContent(val.ResponseContent)
		if err != nil {
			return nil, err
		}
		allResponses = append(allResponses, map[string]interface{}{
			"request_time":     val.RequestTime,
			"response_content": responseContent,
		})
	}

	return allResponses, nil
}

func flattenFindingResponseFindingsWebhookResponsesResponseContent(content map[string]interface{}) (string, error) {
	responseContentBytes, err := json.Marshal(content)
	if err != nil {
		return "", err
	}
	return string(responseContentBytes), nil
}

func flattenFindingResponseFindingsMagellan(magellan continuous_compliance_finding.Magellan) []interface{} {
	m := map[string]interface{}{
		"alert_window_start_time": magellan.AlertWindowStartTime,
		"alert_window_end_time":   magellan.AlertWindowEndTime,
	}

	return []interface{}{m}
}

func flattenFindingResponseFindingsEntityTags(tags []continuous_compliance_finding.TagRule) []interface{} {
	if tags == nil {
		return nil
	}
	allTags := make([]interface{}, len(tags))
	for i, val := range tags {
		allTags[i] = map[string]interface{}{
			"key":   val.Key,
			"value": val.Value,
		}
	}

	return allTags
}

func flattenFindingResponseFindingsComments(comments []continuous_compliance_finding.FindingComment) []interface{} {
	if comments == nil {
		return nil
	}
	allComments := make([]interface{}, len(comments))
	for i, val := range comments {
		allComments[i] = map[string]interface{}{
			"text":      val.Text,
			"timestamp": val.Timestamp,
			"user_name": val.UserName,
		}
	}

	return allComments
}

func flattenFindingResponseSearchRequest(request continuous_compliance_finding.ContinuousComplianceFindingRequest) []interface{} {
	m := map[string]interface{}{
		"page_size":     request.PageSize,
		"sorting":       flattenFindingResponseSearchRequestSorting(request.Sorting),
		"multi_sorting": flattenFindingResponseSearchRequestMultiSorting(request.MultiSorting),
		"filter":        flattenFindingResponseSearchRequestFilter(request.Filter),
		"search_after":  request.SearchAfter,
		"data_source":   request.DataSource,
	}

	return []interface{}{m}
}

func flattenFindingResponseSearchRequestSearchAfter(after *[]string) []interface{} {
	if after == nil {
		return nil
	}

	m := map[string]interface{}{}
	return []interface{}{m}
}

func flattenFindingResponseSearchRequestFilter(filter *continuous_compliance_finding.Filter) []interface{} {
	if filter == nil {
		return nil
	}

	m := map[string]interface{}{
		"free_text_phrase":  filter.FreeTextPhrase,
		"fields":            flattenFindingResponseSearchRequestFilterFields(filter.Fields),
		"only_ciem":         filter.OnlyCIEM,
		"included_features": filter.IncludedFeatures,
		"creation_time":     flattenFindingResponseSearchRequestFilterCreationTime(filter.CreationTime),
	}
	return []interface{}{m}
}

func flattenFindingResponseSearchRequestFilterCreationTime(time *continuous_compliance_finding.DateRange) []interface{} {
	if time == nil {
		return nil
	}

	m := map[string]interface{}{
		"from": time.From,
		"to":   time.To,
	}
	return []interface{}{m}
}

func flattenFindingResponseSearchRequestFilterFields(fields []continuous_compliance_finding.FieldFilter) []interface{} {
	if fields == nil {
		return nil
	}
	allFields := make([]interface{}, len(fields))
	for i, val := range fields {
		allFields[i] = map[string]interface{}{
			"name":  val.Name,
			"value": val.Value,
		}
	}

	return allFields
}

func flattenFindingResponseSearchRequestMultiSorting(sorting []continuous_compliance_finding.Sorting) []interface{} {
	if sorting == nil {
		return nil
	}
	allSorting := make([]interface{}, len(sorting))
	for i, val := range sorting {
		allSorting[i] = map[string]interface{}{
			"field_name": val.FieldName,
			"direction":  val.Direction,
		}
	}

	return allSorting
}

func flattenFindingResponseSearchRequestSorting(sorting *continuous_compliance_finding.Sorting) []interface{} {
	if sorting == nil {
		return nil
	}

	m := map[string]interface{}{
		"field_name": sorting.FieldName,
		"direction":  sorting.Direction,
	}
	return []interface{}{m}
}

func expandContinuousComplianceFindingRequest(d *schema.ResourceData) continuous_compliance_finding.ContinuousComplianceFindingRequest {
	req := continuous_compliance_finding.ContinuousComplianceFindingRequest{
		PageSize:     d.Get("page_size").(int),
		Sorting:      expandContinuousComplianceFindingSorting(d),
		MultiSorting: expandContinuousComplianceFindingMultiSorting(d),
		Filter:       expandContinuousComplianceFindingFilter(d),
		SearchAfter:  expandContinuousComplianceFindingSearchAfter(d),
		DataSource:   d.Get("data_source").(string),
	}
	return req
}

func expandContinuousComplianceFindingSearchAfter(d *schema.ResourceData) *[]string {
	searchAfter := d.Get("filter.search_after").([]interface{})
	search := make([]string, len(searchAfter))
	for i, v := range searchAfter {
		search[i] = v.(string)
	}
	return &search
}

func expandContinuousComplianceFindingFilter(d *schema.ResourceData) *continuous_compliance_finding.Filter {
	filter := continuous_compliance_finding.Filter{
		FreeTextPhrase:   d.Get("filter.free_text_phrase").(string),
		Fields:           expandContinuousComplianceFindingFilterFields(d),
		OnlyCIEM:         d.Get("filter.only_ciem").(bool),
		IncludedFeatures: expandContinuousComplianceFindingIncludedFeatures(d),
		CreationTime:     expandContinuousComplianceFindingCreationTime(d),
	}
	return &filter
}

func expandContinuousComplianceFindingCreationTime(d *schema.ResourceData) *continuous_compliance_finding.DateRange {
	creationTime := continuous_compliance_finding.DateRange{
		From: d.Get("filter.creation_time.from").(string),
		To:   d.Get("filter.creation_time.to").(string),
	}
	return &creationTime
}

func expandContinuousComplianceFindingIncludedFeatures(d *schema.ResourceData) []string {
	includedFeatures := d.Get("filter.included_features").([]interface{})
	features := make([]string, len(includedFeatures))
	for i, v := range includedFeatures {
		features[i] = v.(string)
	}
	return features
}

func expandContinuousComplianceFindingFilterFields(d *schema.ResourceData) []continuous_compliance_finding.FieldFilter {
	fields := d.Get("filter.fields").([]interface{})
	fieldFilters := make([]continuous_compliance_finding.FieldFilter, len(fields))
	for _, v := range fields {
		field := v.(map[string]interface{})
		fieldFilters = append(fieldFilters, continuous_compliance_finding.FieldFilter{
			Name:  field["name"].(string),
			Value: field["value"].(string),
		})
	}
	return fieldFilters
}

func expandContinuousComplianceFindingMultiSorting(d *schema.ResourceData) []continuous_compliance_finding.Sorting {
	multiSorting := d.Get("multi_sorting").([]interface{})
	sorting := make([]continuous_compliance_finding.Sorting, len(multiSorting))
	for _, v := range multiSorting {
		m := v.(map[string]interface{})
		sorting = append(sorting, continuous_compliance_finding.Sorting{
			FieldName: m["field_name"].(string),
			Direction: m["direction"].(int),
		})
	}
	return sorting
}

func expandContinuousComplianceFindingSorting(d *schema.ResourceData) *continuous_compliance_finding.Sorting {
	sorting := continuous_compliance_finding.Sorting{
		FieldName: d.Get("sorting.field_name").(string),
		Direction: d.Get("sorting.direction").(int),
	}
	return &sorting
}
