package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/notifications"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceNotification() *schema.Resource {
	return &schema.Resource{
		Create: resourceNotificationCreate,
		Read:   resourceNotificationRead,
		Update: resourceNotificationUpdate,
		Delete: resourceNotificationDelete,
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
				Optional: true,
			},
			"alerts_console": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"send_on_each_occurrence": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"origin": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ComplianceEngine",
			},
			"integration_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reports_integration_settings": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"integration_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"output_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"filter": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"severities": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"rule_entity_types": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"entity_tags": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:     schema.TypeString,
																Required: true,
															},
															"value": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},
												"entity_names": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"entity_ids": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"entity_categories": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"single_notification_integration_settings": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"integration_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"output_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"payload": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"filter": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"severities": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"rule_entity_types": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"entity_tags": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:     schema.TypeString,
																Required: true,
															},
															"value": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},
												"entity_names": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"entity_ids": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"entity_categories": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"scheduled_integration_settings": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"integration_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"output_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"cron_expression": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"filter": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"severities": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"rule_entity_types": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"entity_tags": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:     schema.TypeString,
																Required: true,
															},
															"value": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},
												"entity_names": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"entity_ids": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"entity_categories": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func expandNotificationCreateRequest(d *schema.ResourceData) (notifications.PostNotificationViewModel, error) {
	integrationSettings, _ := expandIntegrationSettings(d)

	postModel := notifications.PostNotificationViewModel{
		BaseNotificationViewModel: notifications.BaseNotificationViewModel{
			Name:                 d.Get("name").(string),
			Description:          d.Get("description").(string),
			AlertsConsole:        d.Get("alerts_console").(bool),
			SendOnEachOccurrence: d.Get("send_on_each_occurrence").(bool),
			Origin:               d.Get("origin").(string), //origin,
			IntegrationSettings:  integrationSettings,
		},
	}

	return postModel, nil
}

func expandNotificationUpdateRequest(id string, d *schema.ResourceData) (notifications.PutNotificationViewModel, error) {
	integrationSettings, _ := expandIntegrationSettings(d)

	postModel := notifications.PutNotificationViewModel{
		BaseNotificationViewModel: notifications.BaseNotificationViewModel{
			Name:                 d.Get("name").(string),
			Description:          d.Get("description").(string),
			AlertsConsole:        d.Get("alerts_console").(bool),
			SendOnEachOccurrence: d.Get("send_on_each_occurrence").(bool),
			Origin:               d.Get("origin").(string),
			IntegrationSettings:  integrationSettings,
		},
		Id: id,
	}

	return postModel, nil
}

func expandIntegrationSettings(d *schema.ResourceData) (notifications.NotificationIntegrationSettingsModel, error) {
	integrationSettingsRaw := d.Get("integration_settings").([]interface{})

	var notificationSettings = notifications.NotificationIntegrationSettingsModel{
		SingleNotificationIntegrationSettings: []notifications.SingleNotificationIntegrationSettings{},
		ReportsIntegrationSettings:            []notifications.ReportNotificationIntegrationSettings{},
		ScheduledIntegrationSettings:          []notifications.ScheduledNotificationIntegrationSettings{},
	}

	// Check if integrationSettings is not empty
	if len(integrationSettingsRaw) == 0 {
		return notificationSettings, nil
	}

	integrationSettings := integrationSettingsRaw[0].(map[string]interface{})

	var singleNotificationIntegrationSettings, reportsIntegrationSettings, scheduledIntegrationSettings []interface{}
	var ok bool

	if singleNotificationIntegrationSettings, ok = integrationSettings["single_notification_integration_settings"].([]interface{}); ok {
		notificationSettings.SingleNotificationIntegrationSettings, _ = expandSingleNotificationIntegrationSettings(singleNotificationIntegrationSettings)
	}

	if reportsIntegrationSettings, ok = integrationSettings["reports_integration_settings"].([]interface{}); ok {
		notificationSettings.ReportsIntegrationSettings, _ = expandReportsIntegrationSettings(reportsIntegrationSettings)
	}

	if scheduledIntegrationSettings, ok = integrationSettings["scheduled_integration_settings"].([]interface{}); ok {
		notificationSettings.ScheduledIntegrationSettings, _ = expandScheduledIntegrationSettings(scheduledIntegrationSettings)
	}

	return notificationSettings, nil
}

func expandFilterSettings(filter []interface{}) *notifications.ComplianceNotificationFilter {
	if len(filter) == 0 {
		return nil
	}

	filterMap, ok := filter[0].(map[string]interface{})
	if !ok {
		log.Printf("[WARN] Unable to process filter: not a valid map")
		return nil
	}

	entityTags := []notifications.TagRuleEntity{}
	if tags, ok := filterMap["entity_tags"].([]interface{}); ok {
		for _, tag := range tags {
			if tagMap, ok := tag.(map[string]interface{}); ok {
				entityTags = append(entityTags, notifications.TagRuleEntity{
					Key:   tagMap["key"].(string),
					Value: tagMap["value"].(string),
				})
			} else {
				log.Printf("[WARN] Skipping invalid tag in entity_tags")
			}
		}
	}

	return &notifications.ComplianceNotificationFilter{
		Severities:       expandStringList(filterMap["severities"]),
		RuleEntityTypes:  expandStringList(filterMap["rule_entity_types"]),
		EntityTags:       entityTags,
		EntityNames:      expandStringList(filterMap["entity_names"]),
		EntityIds:        expandStringList(filterMap["entity_ids"]),
		EntityCategories: expandStringList(filterMap["entity_categories"]),
	}
}

func expandStringList(raw interface{}) []string {
	if raw == nil {
		return nil
	}

	rawList := raw.([]interface{})
	result := make([]string, len(rawList))
	for i, v := range rawList {
		result[i] = v.(string)
	}
	return result
}

func createBaseNotification(itemMap map[string]interface{}) notifications.BaseNotificationIntegrationSettings {
	return notifications.BaseNotificationIntegrationSettings{
		IntegrationId: itemMap["integration_id"].(string),
		OutputType:    itemMap["output_type"].(string),
		Filter:        nil,
	}
}

func expandSingleNotificationIntegrationSettings(singleNotificationIntegrationSettings []interface{}) ([]notifications.SingleNotificationIntegrationSettings, error) {
	settings := []notifications.SingleNotificationIntegrationSettings{}

	if singleNotificationIntegrationSettings == nil {
		return settings, nil
	}

	for _, item := range singleNotificationIntegrationSettings {
		itemMap := item.(map[string]interface{})

		singleSetting := notifications.SingleNotificationIntegrationSettings{
			BaseNotificationIntegrationSettings: createBaseNotification(itemMap),
			Payload:                             itemMap["payload"].(string),
		}

		// Handle filter for Single Notification Integration Settings
		if filterRaw, ok := itemMap["filter"].([]interface{}); ok && len(filterRaw) > 0 {
			singleSetting.Filter = expandFilterSettings(filterRaw)
		}

		settings = append(settings, singleSetting)
	}

	return settings, nil
}

func expandReportsIntegrationSettings(reportsIntegrationSettings []interface{}) ([]notifications.ReportNotificationIntegrationSettings, error) {
	settings := []notifications.ReportNotificationIntegrationSettings{}

	if reportsIntegrationSettings == nil {
		return settings, nil
	}

	for _, item := range reportsIntegrationSettings {
		itemMap := item.(map[string]interface{})

		reportSetting := notifications.ReportNotificationIntegrationSettings{
			BaseNotificationIntegrationSettings: createBaseNotification(itemMap),
		}

		// Handle filter for Reports Integration Settings
		if filterRaw, ok := itemMap["filter"].([]interface{}); ok && len(filterRaw) > 0 {
			reportSetting.Filter = expandFilterSettings(filterRaw)
		}

		settings = append(settings, reportSetting)
	}

	return settings, nil
}

func expandScheduledIntegrationSettings(scheduledIntegrationSettings []interface{}) ([]notifications.ScheduledNotificationIntegrationSettings, error) {
	settings := []notifications.ScheduledNotificationIntegrationSettings{}

	if scheduledIntegrationSettings == nil {
		return settings, nil
	}

	for _, item := range scheduledIntegrationSettings {
		itemMap := item.(map[string]interface{})

		scheduledSetting := notifications.ScheduledNotificationIntegrationSettings{
			BaseNotificationIntegrationSettings: createBaseNotification(itemMap),
			CronExpression:                      itemMap["cron_expression"].(string),
		}

		// Handle filter for Scheduled Integration Settings
		if filterRaw, ok := itemMap["filter"].([]interface{}); ok && len(filterRaw) > 0 {
			scheduledSetting.Filter = expandFilterSettings(filterRaw)
		}

		settings = append(settings, scheduledSetting)
	}

	return settings, nil
}

func expandIntegrationSettingsForRead(settings notifications.NotificationIntegrationSettingsModel) ([]interface{}, error) {
	result := make(map[string]interface{})

	// Convert Reports Integration Settings
	reportsSettings := make([]interface{}, len(settings.ReportsIntegrationSettings))
	for i, setting := range settings.ReportsIntegrationSettings {
		reportSetting := map[string]interface{}{
			"integration_id": setting.IntegrationId,
			"output_type":    setting.OutputType,
		}

		if setting.Filter != nil {
			reportSetting["filter"] = []interface{}{
				map[string]interface{}{
					"severities":        setting.Filter.Severities,
					"rule_entity_types": setting.Filter.RuleEntityTypes,
					"entity_tags":       flattenEntityTags(setting.Filter.EntityTags),
					"entity_names":      setting.Filter.EntityNames,
					"entity_ids":        setting.Filter.EntityIds,
					"entity_categories": setting.Filter.EntityCategories,
				},
			}
		}
		reportsSettings[i] = reportSetting
	}
	result["reports_integration_settings"] = reportsSettings

	// Convert Single Notification Integration Settings
	singleSettings := make([]interface{}, len(settings.SingleNotificationIntegrationSettings))
	for i, setting := range settings.SingleNotificationIntegrationSettings {
		singleSetting := map[string]interface{}{
			"integration_id": setting.IntegrationId,
			"output_type":    setting.OutputType,
			"payload":        setting.Payload,
		}

		if setting.Filter != nil {
			singleSetting["filter"] = []interface{}{
				map[string]interface{}{
					"severities":        setting.Filter.Severities,
					"rule_entity_types": setting.Filter.RuleEntityTypes,
					"entity_tags":       flattenEntityTags(setting.Filter.EntityTags),
					"entity_names":      setting.Filter.EntityNames,
					"entity_ids":        setting.Filter.EntityIds,
					"entity_categories": setting.Filter.EntityCategories,
				},
			}
		}
		singleSettings[i] = singleSetting
	}
	result["single_notification_integration_settings"] = singleSettings

	// Convert Scheduled Integration Settings
	scheduledSettings := make([]interface{}, len(settings.ScheduledIntegrationSettings))
	for i, setting := range settings.ScheduledIntegrationSettings {
		scheduledSetting := map[string]interface{}{
			"integration_id":  setting.IntegrationId,
			"output_type":     setting.OutputType,
			"cron_expression": setting.CronExpression,
		}

		if setting.Filter != nil {
			scheduledSetting["filter"] = []interface{}{
				map[string]interface{}{
					"severities":        setting.Filter.Severities,
					"rule_entity_types": setting.Filter.RuleEntityTypes,
					"entity_tags":       flattenEntityTags(setting.Filter.EntityTags),
					"entity_names":      setting.Filter.EntityNames,
					"entity_ids":        setting.Filter.EntityIds,
					"entity_categories": setting.Filter.EntityCategories,
				},
			}
		}
		scheduledSettings[i] = scheduledSetting
	}
	result["scheduled_integration_settings"] = scheduledSettings

	// Wrap the result in a slice since the Terraform schema expects a TypeList
	return []interface{}{result}, nil
}

func flattenEntityTags(tags []notifications.TagRuleEntity) []interface{} {
	if tags == nil {
		return nil
	}

	flattened := make([]interface{}, len(tags))
	for i, tag := range tags {
		flattened[i] = map[string]interface{}{
			"key":   tag.Key,
			"value": tag.Value,
		}
	}
	return flattened
}

// CRUD Functions

func resourceNotificationCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req, err := expandNotificationCreateRequest(d)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Creating notification request\n%+v\n", req)
	resp, _, err := d9Client.notifications.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created notification. ID: %v\n", resp.Id)
	d.SetId(resp.Id)

	return resourceNotificationRead(d, meta)
}

func resourceNotificationRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Reading notification ID: %v", d.Id())

	resp, _, err := d9Client.notifications.GetById(d.Id())
	if err != nil {
		return err
	}

	d.SetId(resp.Id)
	_ = d.Set("name", resp.Name)
	_ = d.Set("description", resp.Description)
	_ = d.Set("alerts_console", resp.AlertsConsole)
	_ = d.Set("send_on_each_occurrence", resp.SendOnEachOccurrence)
	_ = d.Set("origin", resp.Origin)

	integrationSettings, err := expandIntegrationSettingsForRead(resp.IntegrationSettings)
	if err != nil {
		return err
	}
	_ = d.Set("integration_settings", integrationSettings)

	return nil
}

func resourceNotificationUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req, err := expandNotificationUpdateRequest(d.Id(), d)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Updating notification request\n%+v\n", req)
	resp, _, err := d9Client.notifications.Update(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Updated notification. ID: %v\n", resp.Id)

	return resourceNotificationRead(d, meta)
}

func resourceNotificationDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting notification ID: %v", d.Id())

	if _, err := d9Client.notifications.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}
