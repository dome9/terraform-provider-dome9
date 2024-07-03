package dome9

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/notifications"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "ComplianceEngine",
				ValidateFunc: validateAssessmentFindingOrigin,
			},
			"integration_settings": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reports_integration_settings": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"integration_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"output_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"single_notification_integration_settings": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"integration_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"output_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"payload": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"scheduled_integration_settings": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"integration_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"output_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"cron_expression": {
										Type:     schema.TypeString,
										Required: true,
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

var assessmentFindingOriginMap = map[string]notifications.AssessmentFindingOrigin{
	"ComplianceEngine":            notifications.ComplianceEngine,
	"Magellan":                    notifications.Magellan,
	"MagellanAwsGuardDuty":        notifications.MagellanAwsGuardDuty,
	"Serverless":                  notifications.Serverless,
	"AwsInspector":                notifications.AwsInspector,
	"ServerlessSecurityAnalyzer":  notifications.ServerlessSecurityAnalyzer,
	"ExternalFindingSource":       notifications.ExternalFindingSource,
	"Qualys":                      notifications.Qualys,
	"Tenable":                     notifications.Tenable,
	"AwsGuardDuty":                notifications.AwsGuardDuty,
	"KubernetesImageScanning":     notifications.KubernetesImageScanning,
	"KubernetesRuntimeAssurance":  notifications.KubernetesRuntimeAssurance,
	"ContainersRuntimeProtection": notifications.ContainersRuntimeProtection,
	"WorkloadChangeMonitoring":    notifications.WorkloadChangeMonitoring,
	"ImageAssurance":              notifications.ImageAssurance,
	"SourceCodeAssurance":         notifications.SourceCodeAssurance,
	"InfrastructureAsCode":        notifications.InfrastructureAsCode,
	"CIEM":                        notifications.CIEM,
	"Incident":                    notifications.Incident,
}

func validateAssessmentFindingOrigin(val interface{}, key string) ([]string, []error) {
	validOrigins := []string{
		"ComplianceEngine",
		"Magellan",
		"MagellanAwsGuardDuty",
		"Serverless",
		"AwsInspector",
		"ServerlessSecurityAnalyzer",
		"ExternalFindingSource",
		"Qualys",
		"Tenable",
		"AwsGuardDuty",
		"KubernetesImageScanning",
		"KubernetesRuntimeAssurance",
		"ContainersRuntimeProtection",
		"WorkloadChangeMonitoring",
		"ImageAssurance",
		"SourceCodeAssurance",
		"InfrastructureAsCode",
		"CIEM",
		"Incident",
	}
	return validation.StringInSlice(validOrigins, true)(val, key)
}

func getAssessmentFindingOrigin(originStr string) (notifications.AssessmentFindingOrigin, error) {
	if origin, exists := assessmentFindingOriginMap[originStr]; exists {
		return origin, nil
	}
	return -1, fmt.Errorf("unknown AssessmentFindingOrigin: %s", originStr)
}

func expandReportsIntegrationSettings(settingsSlice []interface{}) ([]notifications.ReportNotificationIntegrationSettings, error) {
	var settings []notifications.ReportNotificationIntegrationSettings

	for _, item := range settingsSlice {
		settingMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid format for reports_integration_settings item")
		}

		integrationID, ok := settingMap["integration_id"].(string)
		if !ok {
			return nil, fmt.Errorf("integration_id is required and must be a string")
		}

		outputTypeStr, ok := settingMap["output_type"].(string)
		if !ok {
			return nil, fmt.Errorf("output_type is required and must be a string")
		}

		settings = append(settings, notifications.ReportNotificationIntegrationSettings{
			BaseNotificationIntegrationSettings: notifications.BaseNotificationIntegrationSettings{
				IntegrationId: integrationID,
				OutputType:    expandOutputType(outputTypeStr),
				Filter:        notifications.ComplianceNotificationFilter{},
			},
		})
	}

	return settings, nil
}

func expandSingleNotificationIntegrationSettings(settingsSlice []interface{}) ([]notifications.SingleNotificationIntegrationSettings, error) {
	var settings []notifications.SingleNotificationIntegrationSettings

	for _, item := range settingsSlice {
		settingMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid format for single_notification_integration_settings item")
		}

		integrationID, ok := settingMap["integration_id"].(string)
		if !ok {
			return nil, fmt.Errorf("integration_id is required and must be a string")
		}

		outputTypeStr, ok := settingMap["output_type"].(string)
		if !ok {
			return nil, fmt.Errorf("output_type is required and must be a string")
		}

		payload, _ := settingMap["payload"].(string)

		settings = append(settings, notifications.SingleNotificationIntegrationSettings{
			BaseNotificationIntegrationSettings: notifications.BaseNotificationIntegrationSettings{
				IntegrationId: integrationID,
				OutputType:    expandOutputType(outputTypeStr),
				Filter:        notifications.ComplianceNotificationFilter{},
			},
			Payload: payload,
		})
	}

	return settings, nil
}

func expandScheduledIntegrationSettings(settingsSlice []interface{}) ([]notifications.ScheduledNotificationIntegrationSettings, error) {
	var settings []notifications.ScheduledNotificationIntegrationSettings

	for _, item := range settingsSlice {
		settingMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid format for scheduled_integration_settings item")
		}

		integrationID, ok := settingMap["integration_id"].(string)
		if !ok {
			return nil, fmt.Errorf("integration_id is required and must be a string")
		}

		outputTypeStr, ok := settingMap["output_type"].(string)
		if !ok {
			return nil, fmt.Errorf("output_type is required and must be a string")
		}

		cronExpression, ok := settingMap["cron_expression"].(string)
		if !ok {
			return nil, fmt.Errorf("cron_expression is required and must be a string")
		}

		settings = append(settings, notifications.ScheduledNotificationIntegrationSettings{
			BaseNotificationIntegrationSettings: notifications.BaseNotificationIntegrationSettings{
				IntegrationId: integrationID,
				OutputType:    expandOutputType(outputTypeStr),
				Filter:        notifications.ComplianceNotificationFilter{},
			},
			CronExpression: cronExpression,
		})
	}

	return settings, nil
}

func expandOutputType(str string) notifications.NotificationOutputType {
	return 0
}

func expandNotificationRequest(d *schema.ResourceData) (notifications.PostNotificationViewModel, error) {

	//convert origin field to AssessmentFindingOrigin
	originStr := d.Get("origin").(string)
	origin, err := getAssessmentFindingOrigin(originStr)
	if err != nil {
		return notifications.PostNotificationViewModel{}, err
	}

	integrationSettingsItem := d.Get("integration_settings").(*schema.Set).List()[0]
	integrationSettings := integrationSettingsItem.(map[string]interface{})
	log.Printf("[INFO] Integration settings: %+v\n", integrationSettings)
	fmt.Printf("[INFO] Integration settings: %+v\n", integrationSettings)

	reportsSettings, _ := expandReportsIntegrationSettings(integrationSettings)
	singleNotificationSettings, _ := expandSingleNotificationIntegrationSettings(integrationSettings)
	scheduledSettings, _ := expandScheduledIntegrationSettings(integrationSettings)

	postModel := notifications.PostNotificationViewModel{
		BaseNotificationViewModel: notifications.BaseNotificationViewModel{
			Name:                 d.Get("name").(string),
			Description:          d.Get("description").(string),
			AlertsConsole:        d.Get("alerts_console").(bool),
			SendOnEachOccurrence: d.Get("send_on_each_occurrence").(bool),
			Origin:               origin,
			IntegrationSettings: notifications.NotificationIntegrationSettingsModel{
				ReportsIntegrationSettings:            reportsSettings,
				SingleNotificationIntegrationSettings: singleNotificationSettings,
				ScheduledIntegrationSettings:          scheduledSettings,
			},
		},
	}

	return postModel, nil
}

func resourceNotificationCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req, err := expandNotificationRequest(d)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Creating notification request\n%+v\n", req)
	resp, _, err := d9Client.notifications.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created notification request. ID: %v\n", resp.Id)
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

	_ = d.Set("name", resp.Name)
	_ = d.Set("description", resp.Description)
	_ = d.Set("alerts_console", resp.AlertsConsole)
	_ = d.Set("send_on_each_occurrence", resp.SendOnEachOccurrence)

	return nil
}

func resourceNotificationUpdate(d *schema.ResourceData, meta interface{}) error {
	// Implementation for Update
	return nil
}

func resourceNotificationDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting notification ID: %v", d.Id())

	if _, err := d9Client.notifications.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}
