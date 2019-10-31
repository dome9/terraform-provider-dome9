package dome9

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_notification"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
)

func resourceContinuousComplianceNotification() *schema.Resource {
	return &schema.Resource{
		Create: resourceContinuousComplianceNotificationCreate,
		Read:   resourceContinuousComplianceNotificationRead,
		Update: resourceContinuousComplianceNotificationUpdate,
		Delete: resourceContinuousComplianceNotificationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alerts_console": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"scheduled_report": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email_sending_state": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "Disabled",
							ValidateFunc: validation.StringInSlice([]string{"Disabled", "Enabled"}, true),
						},
						"schedule_data": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cron_expression": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"Detailed", "Summary", "FullCsv", "FullCsvZip"}, true),
									},
									"recipients": {
										Type:     schema.TypeList,
										Required: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},

			"change_detection": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email_sending_state": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "Disabled",
							ValidateFunc: validation.StringInSlice([]string{"Disabled", "Enabled"}, true),
						},
						"email_per_finding_sending_state": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "Disabled",
							ValidateFunc: validation.StringInSlice([]string{"Disabled", "Enabled"}, true),
						},
						"sns_sending_state": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "Disabled",
							ValidateFunc: validation.StringInSlice([]string{"Disabled", "Enabled"}, true),
						},
						"external_ticket_creating_state": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "Disabled",
							ValidateFunc: validation.StringInSlice([]string{"Disabled", "Enabled"}, true),
						},
						"aws_security_hub_integration_state": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "Disabled",
							ValidateFunc: validation.StringInSlice([]string{"Disabled", "Enabled"}, true),
						},
						"webhook_integration_state": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "Disabled",
							ValidateFunc: validation.StringInSlice([]string{"Disabled", "Enabled"}, true),
						},
						"email_data": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"recipients": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"email_per_finding_data": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"recipients": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"notification_output_format": {
										Type:         schema.TypeString,
										Optional:     true,
										Default:      "JsonWithFullEntity",
										ValidateFunc: validation.StringInSlice([]string{"JsonWithFullEntity", "JsonWithBasicEntity", "PlainText"}, true),
									},
								},
							},
						},
						"sns_data": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sns_topic_arn": {
										Type:     schema.TypeString,
										Required: true,
									},
									"sns_output_format": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"JsonWithFullEntity", "JsonWithBasicEntity", "PlainText"}, true),
									},
								},
							},
						},
						"ticketing_system_data": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"system_type": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "PagerDuty",
									},
									"should_close_tickets": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"domain": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"user": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"pass": {
										Type:     schema.TypeString,
										Required: true,
									},
									"project_key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"issue_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"aws_security_hub_integration": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"external_account_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"region": {
										Type:         schema.TypeString,
										Required:     true,
										ForceNew:     true,
										ValidateFunc: validation.StringInSlice(providerconst.AWSRegions, true),
									},
								},
							},
						},
						"webhook_data": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": {
										Type:     schema.TypeString,
										Required: true,
									},
									"http_method": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "Post",
										ForceNew: true,
									},
									"auth_method": {
										Type:         schema.TypeString,
										Optional:     true,
										Default:      "NoAuth",
										ValidateFunc: validation.StringInSlice([]string{"NoAuth", "BasicAuth"}, true),
									},
									"username": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"password": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"format_type": {
										Type:         schema.TypeString,
										Optional:     true,
										Default:      "JsonWithFullEntity",
										ValidateFunc: validation.StringInSlice([]string{"JsonWithFullEntity", "SplunkBasic", "ServiceNow"}, true),
									},
								},
							},
						},
					},
				},
			},
			"gcp_security_command_center_integration": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "Disabled",
							ValidateFunc: validation.StringInSlice([]string{"Disabled", "Enabled"}, true),
						},
						"project_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceContinuousComplianceNotificationCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandContinuousComplianceNotificationRequest(d)
	log.Printf("[INFO] Creating continuous compliance notification request\n%+v\n", req)
	resp, _, err := d9Client.continuousComplianceNotification.Create(&req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created continuous compliance notification request. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceContinuousComplianceNotificationRead(d, meta)
}

func resourceContinuousComplianceNotificationRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.continuousComplianceNotification.Get(d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing continuous compliance notification %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting continuous compliance notification:\n%+v\n", resp)
	d.SetId(resp.ID)
	_ = d.Set("name", resp.Name)
	_ = d.Set("description", resp.Description)
	_ = d.Set("alerts_console", resp.AlertsConsole)

	if resp.ScheduledReport != nil {
		if err := d.Set("scheduled_report", flattenScheduledReport(resp.ScheduledReport)); err != nil {
			return err
		}
	}

	if err := d.Set("change_detection", flattenChangeDetection(&resp.ChangeDetection)); err != nil {
		return err
	}

	if resp.GCPSecurityCommandCenterIntegration != nil {
		if err = d.Set("gcp_security_command_center_integration", flattenGCPSecurityCommandCenterIntegration(resp.GCPSecurityCommandCenterIntegration)); err != nil {
			return err
		}
	}

	return nil
}

func resourceContinuousComplianceNotificationDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting continuous compliance notification ID: %v", d.Id())

	if _, err := d9Client.continuousComplianceNotification.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceContinuousComplianceNotificationUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Updating continuous compliance notification ID: %v\n", d.Id())
	req := expandContinuousComplianceNotificationRequest(d)

	if _, _, err := d9Client.continuousComplianceNotification.Update(d.Id(), &req); err != nil {
		return err
	}

	return nil
}

func expandContinuousComplianceNotificationRequest(d *schema.ResourceData) continuous_compliance_notification.ContinuousComplianceNotificationRequest {
	return continuous_compliance_notification.ContinuousComplianceNotificationRequest{
		Name:                                d.Get("name").(string),
		Description:                         d.Get("description").(string),
		AlertsConsole:                       d.Get("alerts_console").(bool),
		ScheduledReport:                     expandScheduledReport(d),
		ChangeDetection:                     *expandChangeDetection(d),
		GCPSecurityCommandCenterIntegration: expandGCPSecurityCommandCenterIntegration(d),
	}
}

func expandScheduledReport(d *schema.ResourceData) *continuous_compliance_notification.ScheduledReport {
	if scheduledReports, ok := d.GetOk("scheduled_report"); ok {
		scheduledReportItem := scheduledReports.(*schema.Set).List()[0]
		scheduledReport := scheduledReportItem.(map[string]interface{})

		return &continuous_compliance_notification.ScheduledReport{
			EmailSendingState: scheduledReport["email_sending_state"].(string),
			ScheduleData:      expandScheduleData(scheduledReport["schedule_data"].(*schema.Set)),
		}
	}

	return nil
}

func expandScheduleData(scheduleData *schema.Set) *continuous_compliance_notification.ScheduleData {
	scheduledDataLst := scheduleData.List()
	if len(scheduledDataLst) > 0 {
		scheduledDataItem := scheduledDataLst[0]
		scheduledData := scheduledDataItem.(map[string]interface{})

		return &continuous_compliance_notification.ScheduleData{
			CronExpression: scheduledData["cron_expression"].(string),
			Type:           scheduledData["type"].(string),
			Recipients:     expandRecipients(scheduledData["recipients"].([]interface{})),
		}
	}

	return nil
}

func expandChangeDetection(d *schema.ResourceData) *continuous_compliance_notification.ChangeDetection {
	changeDetectionItem := d.Get("change_detection").(*schema.Set).List()[0]
	changeDetection := changeDetectionItem.(map[string]interface{})

	return &continuous_compliance_notification.ChangeDetection{
		EmailSendingState:              changeDetection["email_sending_state"].(string),
		EmailPerFindingSendingState:    changeDetection["email_per_finding_sending_state"].(string),
		SNSSendingState:                changeDetection["sns_sending_state"].(string),
		ExternalTicketCreatingState:    changeDetection["external_ticket_creating_state"].(string),
		AWSSecurityHubIntegrationState: changeDetection["aws_security_hub_integration_state"].(string),
		WebhookIntegrationState:        changeDetection["webhook_integration_state"].(string),
		EmailData:                      expandEmailData(changeDetection["email_data"].(*schema.Set)),
		EmailPerFindingData:            expandEmailPerFindingData(changeDetection["email_per_finding_data"].(*schema.Set)),
		SNSData:                        expandSNSData(changeDetection["sns_data"].(*schema.Set)),
		TicketingSystemData:            expandTicketingSystemData(changeDetection["ticketing_system_data"].(*schema.Set)),
		AWSSecurityHubIntegration:      expandAWSSecurityHubIntegration(changeDetection["aws_security_hub_integration"].(*schema.Set)),
		WebhookData:                    expandWebhookData(changeDetection["webhook_data"].(*schema.Set)),
	}
}

func expandEmailData(emailData *schema.Set) *continuous_compliance_notification.EmailData {
	emailDataLst := emailData.List()
	if len(emailDataLst) > 0 {
		emailDataItem := emailDataLst[0]
		emailData := emailDataItem.(map[string]interface{})

		return &continuous_compliance_notification.EmailData{
			Recipients: expandRecipients(emailData["recipients"].([]interface{})),
		}
	}

	return nil
}

func expandEmailPerFindingData(emailPerFindingData *schema.Set) *continuous_compliance_notification.EmailPerFindingData {
	emailPerFindingDataLst := emailPerFindingData.List()
	if len(emailPerFindingDataLst) > 0 {
		emailPerFindingDataItem := emailPerFindingDataLst[0]
		emailPerFindingData := emailPerFindingDataItem.(map[string]interface{})

		return &continuous_compliance_notification.EmailPerFindingData{
			Recipients:               expandRecipients(emailPerFindingData["recipients"].([]interface{})),
			NotificationOutputFormat: emailPerFindingData["notification_output_format"].(string),
		}
	}

	return nil
}

func expandSNSData(snsData *schema.Set) *continuous_compliance_notification.SNSData {
	snsDataLst := snsData.List()
	if len(snsDataLst) > 0 {
		snsDataItem := snsDataLst[0]
		snsData := snsDataItem.(map[string]interface{})

		return &continuous_compliance_notification.SNSData{
			SNSTopicArn:     snsData["sns_topic_arn"].(string),
			SNSOutputFormat: snsData["sns_output_format"].(string),
		}
	}

	return nil
}

func expandTicketingSystemData(ticketingSystemData *schema.Set) *continuous_compliance_notification.TicketingSystemData {
	ticketingSystemDataLst := ticketingSystemData.List()
	if len(ticketingSystemDataLst) > 0 {
		ticketingSystemDataItem := ticketingSystemDataLst[0]
		ticketingSystemData := ticketingSystemDataItem.(map[string]interface{})

		return &continuous_compliance_notification.TicketingSystemData{
			SystemType:         ticketingSystemData["system_type"].(string),
			ShouldCloseTickets: ticketingSystemData["should_close_tickets"].(bool),
			Domain:             ticketingSystemData["domain"].(string),
			User:               ticketingSystemData["user"].(string),
			Pass:               ticketingSystemData["pass"].(string),
			ProjectKey:         ticketingSystemData["project_key"].(string),
			IssueType:          ticketingSystemData["issue_type"].(string),
		}
	}

	return nil
}

func expandAWSSecurityHubIntegration(awsSecurityHubIntegration *schema.Set) *continuous_compliance_notification.AWSSecurityHubIntegration {
	awsSecurityHubIntegrationLst := awsSecurityHubIntegration.List()
	if len(awsSecurityHubIntegrationLst) > 0 {
		awsSecurityHubIntegrationItem := awsSecurityHubIntegrationLst[0]
		awsSecurityHubIntegration := awsSecurityHubIntegrationItem.(map[string]interface{})

		return &continuous_compliance_notification.AWSSecurityHubIntegration{
			ExternalAccountID: awsSecurityHubIntegration["external_account_id"].(string),
			Region:            awsSecurityHubIntegration["region"].(string),
		}
	}

	return nil
}

func expandWebhookData(webhookData *schema.Set) *continuous_compliance_notification.WebhookData {
	webhookDataLst := webhookData.List()
	if len(webhookDataLst) > 0 {
		webhookDataItem := webhookDataLst[0]
		webhookData := webhookDataItem.(map[string]interface{})

		return &continuous_compliance_notification.WebhookData{
			URL:        webhookData["url"].(string),
			HTTPMethod: webhookData["http_method"].(string),
			AuthMethod: webhookData["auth_method"].(string),
			Username:   webhookData["username"].(string),
			Password:   webhookData["password"].(string),
			FormatType: webhookData["format_type"].(string),
		}
	}

	return nil
}

func expandRecipients(generalRecipients []interface{}) []string {
	recipients := make([]string, len(generalRecipients))
	for i, recipient := range generalRecipients {
		recipients[i] = recipient.(string)
	}

	return recipients
}

func expandGCPSecurityCommandCenterIntegration(d *schema.ResourceData) *continuous_compliance_notification.GCPSecurityCommandCenterIntegration {
	if gcpSecurityCommandCenterIntegration, ok := d.GetOk("gcp_security_command_center_integration"); ok {
		gcpSecurityCommandCenterIntegrationItem := gcpSecurityCommandCenterIntegration.(*schema.Set).List()[0]
		gcpSecurityCommandCenterIntegration := gcpSecurityCommandCenterIntegrationItem.(map[string]interface{})

		return &continuous_compliance_notification.GCPSecurityCommandCenterIntegration{
			State:     gcpSecurityCommandCenterIntegration["state"].(string),
			ProjectID: gcpSecurityCommandCenterIntegration["project_id"].(string),
			SourceID:  gcpSecurityCommandCenterIntegration["source_id"].(string),
		}
	}

	return nil
}

func flattenGCPSecurityCommandCenterIntegration(respGCPSecurityCommandCenterIntegration *continuous_compliance_notification.GCPSecurityCommandCenterIntegration) []interface{} {
	m := map[string]interface{}{
		"state":      respGCPSecurityCommandCenterIntegration.State,
		"project_id": respGCPSecurityCommandCenterIntegration.ProjectID,
		"source_id":  respGCPSecurityCommandCenterIntegration.SourceID,
	}

	return []interface{}{m}
}

func flattenScheduledReport(respScheduledReport *continuous_compliance_notification.ScheduledReport) []interface{} {
	m := map[string]interface{}{
		"email_sending_state": respScheduledReport.EmailSendingState,
		"schedule_data":       flattenScheduleData(respScheduledReport.ScheduleData),
	}

	return []interface{}{m}
}

func flattenScheduleData(respScheduleData *continuous_compliance_notification.ScheduleData) []interface{} {
	if respScheduleData == nil {
		return nil
	}

	m := map[string]interface{}{
		"cron_expression": respScheduleData.CronExpression,
		"type":            respScheduleData.Type,
		"recipients":      flattenRecipients(respScheduleData.Recipients),
	}

	return []interface{}{m}
}

func flattenRecipients(generalRecipients []string) []string {
	recipients := make([]string, len(generalRecipients))
	for i, val := range generalRecipients {
		recipients[i] = val
	}

	return recipients
}

func flattenChangeDetection(respChangeDetection *continuous_compliance_notification.ChangeDetection) []interface{} {
	m := map[string]interface{}{
		"email_sending_state":                respChangeDetection.EmailSendingState,
		"email_per_finding_sending_state":    respChangeDetection.EmailPerFindingSendingState,
		"sns_sending_state":                  respChangeDetection.SNSSendingState,
		"external_ticket_creating_state":     respChangeDetection.ExternalTicketCreatingState,
		"aws_security_hub_integration_state": respChangeDetection.AWSSecurityHubIntegrationState,
		"webhook_integration_state":          respChangeDetection.WebhookIntegrationState,
	}

	if respChangeDetection.EmailData != nil {
		m["email_data"] = flattenEmailData(respChangeDetection.EmailData)
	}

	if respChangeDetection.EmailPerFindingData != nil {
		m["email_per_finding_data"] = flattenEmailPerFindingData(respChangeDetection.EmailPerFindingData)
	}

	if respChangeDetection.SNSData != nil {
		m["sns_data"] = flattenSnsData(respChangeDetection.SNSData)
	}

	if respChangeDetection.TicketingSystemData != nil {
		m["ticketing_system_data"] = flattenTicketingSystemData(respChangeDetection.TicketingSystemData)
	}

	if respChangeDetection.AWSSecurityHubIntegration != nil {
		m["aws_security_hub_integration"] = flattenAWSSecurityHubIntegration(respChangeDetection.AWSSecurityHubIntegration)
	}

	if respChangeDetection.WebhookData != nil {
		m["webhook_data"] = flattenWebhookData(respChangeDetection.WebhookData)
	}

	return []interface{}{m}
}

func flattenAWSSecurityHubIntegration(respAWSSecurityHubIntegration *continuous_compliance_notification.AWSSecurityHubIntegration) []interface{} {
	m := map[string]interface{}{
		"external_account_id": respAWSSecurityHubIntegration.ExternalAccountID,
		"region":              respAWSSecurityHubIntegration.Region,
	}

	return []interface{}{m}
}

func flattenWebhookData(respWebhookData *continuous_compliance_notification.WebhookData) []interface{} {
	m := map[string]interface{}{
		"url":         respWebhookData.URL,
		"http_method": respWebhookData.HTTPMethod,
		"auth_method": respWebhookData.AuthMethod,
		"username":    respWebhookData.Username,
		"password":    respWebhookData.Password,
		"format_type": respWebhookData.FormatType,
	}

	return []interface{}{m}
}

func flattenTicketingSystemData(respTicketingSystemData *continuous_compliance_notification.TicketingSystemData) []interface{} {
	m := map[string]interface{}{
		"system_type":          respTicketingSystemData.SystemType,
		"should_close_tickets": respTicketingSystemData.ShouldCloseTickets,
		"domain":               respTicketingSystemData.Domain,
		"user":                 respTicketingSystemData.User,
		"pass":                 respTicketingSystemData.Pass,
		"project_key":          respTicketingSystemData.ProjectKey,
		"issue_type":           respTicketingSystemData.IssueType,
	}

	return []interface{}{m}
}

func flattenSnsData(respSNSData *continuous_compliance_notification.SNSData) []interface{} {
	m := map[string]interface{}{
		"sns_topic_arn":     respSNSData.SNSTopicArn,
		"sns_output_format": respSNSData.SNSOutputFormat,
	}

	return []interface{}{m}
}

func flattenEmailData(respEmailData *continuous_compliance_notification.EmailData) []interface{} {
	m := map[string]interface{}{
		"recipients": flattenRecipients(respEmailData.Recipients),
	}

	return []interface{}{m}
}

func flattenEmailPerFindingData(respEmailPerFindingData *continuous_compliance_notification.EmailPerFindingData) []interface{} {
	m := map[string]interface{}{
		"recipients":                 flattenRecipients(respEmailPerFindingData.Recipients),
		"notification_output_format": respEmailPerFindingData.NotificationOutputFormat,
	}

	return []interface{}{m}
}
