package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_finding"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
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

	//d.SetId(resp.ID)
	return flattenContinuousComplianceFindingSearchResponse(resp)
}

func flattenContinuousComplianceFindingSearchResponse(resp continuous_compliance_finding.ContinuousComplianceFindingResponse) []interface{} {
	m := map[string]interface{}{
		"is_template": resp,
	}

	return []interface{}{m}
}

func flattenContinuousComplianceFindingSearchResponse(resp continuous_compliance_finding.ContinuousComplianceFindingResponse) error {
	d.SetId(strconv.Itoa(resp.ID))
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

func expandContinuousComplianceFindingIncludedFeatures(d *schema.ResourceData) *[]string {
	includedFeatures := d.Get("filter.included_features").([]interface{})
	features := make([]string, len(includedFeatures))
	for i, v := range includedFeatures {
		features[i] = v.(string)
	}
	return &features
}

func expandContinuousComplianceFindingFilterFields(d *schema.ResourceData) *[]continuous_compliance_finding.FieldFilter {
	fields := d.Get("filter.fields").([]interface{})
	fieldFilters := make([]continuous_compliance_finding.FieldFilter, len(fields))
	for _, v := range fields {
		field := v.(map[string]interface{})
		fieldFilters = append(fieldFilters, continuous_compliance_finding.FieldFilter{
			Name:  field["name"].(string),
			Value: field["value"].(string),
		})
	}
	return &fieldFilters
}

func expandContinuousComplianceFindingMultiSorting(d *schema.ResourceData) *[]continuous_compliance_finding.Sorting {
	multiSorting := d.Get("multi_sorting").([]interface{})
	sorting := make([]continuous_compliance_finding.Sorting, len(multiSorting))
	for _, v := range multiSorting {
		m := v.(map[string]interface{})
		sorting = append(sorting, continuous_compliance_finding.Sorting{
			FieldName: m["field_name"].(string),
			Direction: m["direction"].(int),
		})
	}
	return &sorting
}

func expandContinuousComplianceFindingSorting(d *schema.ResourceData) *continuous_compliance_finding.Sorting {
	sorting := continuous_compliance_finding.Sorting{
		FieldName: d.Get("sorting.field_name").(string),
		Direction: d.Get("sorting.direction").(int),
	}
	return &sorting
}
