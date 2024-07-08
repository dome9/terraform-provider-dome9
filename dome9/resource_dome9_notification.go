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
				MaxItems: 1,
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

var notificationOutputTypeMap = map[string]notifications.NotificationOutputType{
	"Default":            notifications.Default,
	"Detailed":           notifications.Detailed,
	"Summary":            notifications.Summary,
	"FullCsv":            notifications.FullCsv,
	"FullCsvZip":         notifications.FullCsvZip,
	"ExecutivePlatform":  notifications.ExecutivePlatform,
	"JsonFullEntity":     notifications.JsonFullEntity,
	"JsonSimpleEntity":   notifications.JsonSimpleEntity,
	"PlainText":          notifications.PlainText,
	"TemplateBased":      notifications.TemplateBased,
	"CustomOutputFormat": notifications.CustomOutputFormat,
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

func expandOutputType(outputType string) notifications.NotificationOutputType {
	// Convert string to appropriate NotificationOutputType
	return notificationOutputTypeMap[outputType]
}

func expandNotificationRequest(d *schema.ResourceData) (notifications.PostNotificationViewModel, error) {
	//originStr := d.Get("origin").(string)
	//origin, err := getAssessmentFindingOrigin(originStr)
	//if err != nil {
	//	return notifications.PostNotificationViewModel{}, err
	//}

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

func expandIntegrationSettings(d *schema.ResourceData) (notifications.NotificationIntegrationSettingsModel, error) {
	integrationSettings := d.Get("integration_settings").([]interface{})[0].(map[string]interface{})

	singleNotificationIntegrationSettings, _ := integrationSettings["single_notification_integration_settings"].([]interface{})
	//reportsIntegrationSettings, _ := integrationSettings["reports_integration_settings"].([]interface{})
	//scheduledIntegrationSettings, _ := integrationSettings["scheduled_integration_settings"].([]interface{})

	SingleNotificationIntegrationSettingsData, _ := expandSingleNotificationIntegrationSettings(singleNotificationIntegrationSettings)
	//reportsIntegrationSettingsData := expandReportsIntegrationSettings(reportsIntegrationSettings)
	//scheduledIntegrationSettingsData := expandScheduledIntegrationSettings(scheduledIntegrationSettings)

	// Assuming functions to expand these settings are implemented correctly
	return notifications.NotificationIntegrationSettingsModel{
		ReportsIntegrationSettings:            []notifications.ReportNotificationIntegrationSettings{}, //expandReportsIntegrationSettings(reportsIntegrationSettings),
		SingleNotificationIntegrationSettings: SingleNotificationIntegrationSettingsData,
		ScheduledIntegrationSettings:          []notifications.ScheduledNotificationIntegrationSettings{}, //expandScheduledIntegrationSettings(scheduledIntegrationSettings),
	}, nil
}

func expandScheduledIntegrationSettings(scheduledIntegrationSettings []interface{}) ([]notifications.ScheduledNotificationIntegrationSettings, error) {
	var settings []notifications.ScheduledNotificationIntegrationSettings

	// Process scheduled_integration_settings
	fmt.Println("Scheduled Integration Settings:")
	for i, item := range scheduledIntegrationSettings {
		itemMap := item.(map[string]interface{})
		fmt.Printf("  Item %d: integration_id=%s, output_type=%s, cron_expression=%s\n", i, itemMap["integration_id"].(string), itemMap["output_type"].(string), itemMap["cron_expression"].(string))
	}

	return settings, nil
}

func expandReportsIntegrationSettings(reportsIntegrationSettings []interface{}) ([]notifications.ReportNotificationIntegrationSettings, error) {
	var settings []notifications.ReportNotificationIntegrationSettings

	// Process reports_integration_settings
	fmt.Println("Reports Integration Settings:")
	for i, item := range reportsIntegrationSettings {
		itemMap := item.(map[string]interface{})
		fmt.Printf("  Item %d: integration_id=%s, output_type=%s\n", i, itemMap["integration_id"].(string), itemMap["output_type"].(string))
	}

	return settings, nil
}

func expandSingleNotificationIntegrationSettings(singleNotificationIntegrationSettings []interface{}) ([]notifications.SingleNotificationIntegrationSettings, error) {
	var settings []notifications.SingleNotificationIntegrationSettings

	// Process single_notification_integration_settings
	fmt.Println("Single Notification Integration Settings:")
	for i, item := range singleNotificationIntegrationSettings {
		itemMap := item.(map[string]interface{})
		fmt.Printf("  Item %d: integration_id=%s, output_type=%s, payload=%s\n", i, itemMap["integration_id"].(string), itemMap["output_type"].(string), itemMap["payload"].(string))

		settings = append(settings, notifications.SingleNotificationIntegrationSettings{
			BaseNotificationIntegrationSettings: notifications.BaseNotificationIntegrationSettings{
				IntegrationId: itemMap["integration_id"].(string),
				OutputType:    itemMap["output_type"].(string), //expandOutputType(itemMap["output_type"].(string)),
			},
			Payload: itemMap["payload"].(string),
		})
	}

	return settings, nil
}

// CRUD Functions

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
	//d9Client := meta.(*Client)
	//req, err := expandNotificationRequest(d)
	//if err != nil {
	//	return err
	//}
	//log.Printf("[INFO] Updating notification request\n%+v\n", req)
	//_, _, err = d9Client.notifications.Update(d.Id(), req)
	//if err != nil {
	//	return err
	//}
	//
	//return resourceNotificationRead(d, meta)
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
