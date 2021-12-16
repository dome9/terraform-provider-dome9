package dome9

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

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
						"slack_integration_state": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "Disabled",
							ValidateFunc: validation.StringInSlice([]string{"Disabled", "Enabled"}, true),
						},
						"teams_integration_state": {
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
										ValidateFunc: validation.StringInSlice(providerconst.AllAWSRegions, true),
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
									"payload_format": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: ValidatePayloadFormatJSON,
										Default:      "{}",
									},
									"ignore_certificate": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"advanced_url": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"slack_data": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"teams_data": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": {
										Type:     schema.TypeString,
										Required: true,
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
	req, err := expandContinuousComplianceNotificationRequest(d)
	if err != nil {
		return err
	}
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
			resourceContinuousComplianceNotificationDelete(d, meta)
			return err
		}
	}

	flattenChangeDetection, err:= flattenChangeDetection(&resp.ChangeDetection)
	if err != nil {
		return err
	}
	if err := d.Set("change_detection", flattenChangeDetection); err != nil {
		resourceContinuousComplianceNotificationDelete(d, meta)
		return err
	}

	if resp.GCPSecurityCommandCenterIntegration != nil {
		if err = d.Set("gcp_security_command_center_integration", flattenGCPSecurityCommandCenterIntegration(resp.GCPSecurityCommandCenterIntegration)); err != nil {
			resourceContinuousComplianceNotificationDelete(d, meta)
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
	req, err := expandContinuousComplianceNotificationRequest(d)
	if err != nil {
		return err
	}

	if _, _, err := d9Client.continuousComplianceNotification.Update(d.Id(), &req); err != nil {
		return err
	}

	return nil
}

func expandContinuousComplianceNotificationRequest(d *schema.ResourceData) (continuous_compliance_notification.ContinuousComplianceNotificationRequest, error) {
	ChangeDetectionData, err := expandChangeDetection(d)
	if err != nil {
		return continuous_compliance_notification.ContinuousComplianceNotificationRequest{}, err
	}

	return continuous_compliance_notification.ContinuousComplianceNotificationRequest{
		Name:                                d.Get("name").(string),
		Description:                         d.Get("description").(string),
		AlertsConsole:                       d.Get("alerts_console").(bool),
		ScheduledReport:                     expandScheduledReport(d),
		ChangeDetection:                     *ChangeDetectionData,
		GCPSecurityCommandCenterIntegration: expandGCPSecurityCommandCenterIntegration(d),
	}, nil
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

func expandChangeDetection(d *schema.ResourceData) (*continuous_compliance_notification.ChangeDetection, error) {
	changeDetectionItem := d.Get("change_detection").(*schema.Set).List()[0]
	changeDetection := changeDetectionItem.(map[string]interface{})

	webhookData, err := expandWebhookData(changeDetection["webhook_data"].(*schema.Set))
	if err != nil {
		return nil, err
	}

	return &continuous_compliance_notification.ChangeDetection{
		EmailSendingState:              changeDetection["email_sending_state"].(string),
		EmailPerFindingSendingState:    changeDetection["email_per_finding_sending_state"].(string),
		SNSSendingState:                changeDetection["sns_sending_state"].(string),
		ExternalTicketCreatingState:    changeDetection["external_ticket_creating_state"].(string),
		AWSSecurityHubIntegrationState: changeDetection["aws_security_hub_integration_state"].(string),
		WebhookIntegrationState:        changeDetection["webhook_integration_state"].(string),
		SlackIntegrationState:          changeDetection["slack_integration_state"].(string),
		TeamsIntegrationState:          changeDetection["teams_integration_state"].(string),
		EmailData:                      expandEmailData(changeDetection["email_data"].(*schema.Set)),
		EmailPerFindingData:            expandEmailPerFindingData(changeDetection["email_per_finding_data"].(*schema.Set)),
		SNSData:                        expandSNSData(changeDetection["sns_data"].(*schema.Set)),
		TicketingSystemData:            expandTicketingSystemData(changeDetection["ticketing_system_data"].(*schema.Set)),
		AWSSecurityHubIntegration:      expandAWSSecurityHubIntegration(changeDetection["aws_security_hub_integration"].(*schema.Set)),
		WebhookData:                    webhookData,
		SlackData:                      expandSlackData(changeDetection["slack_data"].(*schema.Set)),
		TeamsData:                      expandTeamsData(changeDetection["teams_data"].(*schema.Set)),
	}, nil
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

func expandWebhookData(webhookData *schema.Set) (*continuous_compliance_notification.WebhookData, error) {
	webhookDataLst := webhookData.List()
	if len(webhookDataLst) > 0 {
		webhookDataItem := webhookDataLst[0]
		webhookData := webhookDataItem.(map[string]interface{})


		PayloadFormatJson := make(map[string]interface{})
		err := json.Unmarshal([]byte(webhookData["payload_format"].(string)), &PayloadFormatJson)
		if err != nil {
			return nil, err
		}

		return &continuous_compliance_notification.WebhookData{
			URL:               webhookData["url"].(string),
			HTTPMethod:        webhookData["http_method"].(string),
			AuthMethod:        webhookData["auth_method"].(string),
			Username:          webhookData["username"].(string),
			Password:          webhookData["password"].(string),
			FormatType:        webhookData["format_type"].(string),
			PayloadFormat:     PayloadFormatJson,
			IgnoreCertificate: webhookData["ignore_certificate"].(bool),
			AdvancedUrl:       webhookData["advanced_url"].(bool),
		}, nil
	}

	return nil, nil
}

func expandSlackData(slackData *schema.Set) *continuous_compliance_notification.SlackData {
	slackDataLst := slackData.List()
	if len(slackDataLst) > 0 {
		slackDataItem := slackDataLst[0]
		slackData := slackDataItem.(map[string]interface{})

		return &continuous_compliance_notification.SlackData{
			URL: slackData["url"].(string),
		}
	}

	return nil
}

func expandTeamsData(teamsData *schema.Set) *continuous_compliance_notification.TeamsData {
	teamsDataLst := teamsData.List()
	if len(teamsDataLst) > 0 {
		teamsDataItem := teamsDataLst[0]
		teamsData := teamsDataItem.(map[string]interface{})

		return &continuous_compliance_notification.TeamsData{
			URL: teamsData["url"].(string),
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
		"recipients":      respScheduleData.Recipients,
	}

	return []interface{}{m}
}

func flattenChangeDetection(respChangeDetection *continuous_compliance_notification.ChangeDetection) ([]interface{}, error) {
	m := map[string]interface{}{
		"email_sending_state":                respChangeDetection.EmailSendingState,
		"email_per_finding_sending_state":    respChangeDetection.EmailPerFindingSendingState,
		"sns_sending_state":                  respChangeDetection.SNSSendingState,
		"external_ticket_creating_state":     respChangeDetection.ExternalTicketCreatingState,
		"aws_security_hub_integration_state": respChangeDetection.AWSSecurityHubIntegrationState,
		"webhook_integration_state":          respChangeDetection.WebhookIntegrationState,
		"slack_integration_state":            respChangeDetection.SlackIntegrationState,
		"teams_integration_state":            respChangeDetection.TeamsIntegrationState,
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
		var err error
		m["webhook_data"], err = flattenWebhookData(respChangeDetection.WebhookData)
		if err != nil {
			return nil, err
		}
	}

	if respChangeDetection.SlackData != nil {
		m["slack_data"] = flattenSlackData(respChangeDetection.SlackData)
	}

	if respChangeDetection.TeamsData != nil {
		m["teams_data"] = flattenTeamsData(respChangeDetection.TeamsData)
	}

	return []interface{}{m}, nil
}

func flattenAWSSecurityHubIntegration(respAWSSecurityHubIntegration *continuous_compliance_notification.AWSSecurityHubIntegration) []interface{} {
	m := map[string]interface{}{
		"external_account_id": respAWSSecurityHubIntegration.ExternalAccountID,
		"region":              respAWSSecurityHubIntegration.Region,
	}

	return []interface{}{m}
}

func flattenWebhookData(respWebhookData *continuous_compliance_notification.WebhookData) ([]interface{}, error) {

	PayloadFormatBytes, err := json.Marshal(respWebhookData.PayloadFormat)
	if err != nil {
		return nil, err
	}
	PayloadFormatStr := string(PayloadFormatBytes)

	m := map[string]interface{}{
		"url":                respWebhookData.URL,
		"http_method":        respWebhookData.HTTPMethod,
		"auth_method":        respWebhookData.AuthMethod,
		"username":           respWebhookData.Username,
		"password":           respWebhookData.Password,
		"format_type":        respWebhookData.FormatType,
		"payload_format":     PayloadFormatStr,
		"ignore_certificate": respWebhookData.IgnoreCertificate,
		"advanced_url":       respWebhookData.AdvancedUrl,
	}

	return []interface{}{m}, nil
}

func flattenSlackData(respWebhookData *continuous_compliance_notification.SlackData) []interface{} {
	m := map[string]interface{}{
		"url": respWebhookData.URL,
	}

	return []interface{}{m}
}

func flattenTeamsData(respWebhookData *continuous_compliance_notification.TeamsData) []interface{} {
	m := map[string]interface{}{
		"url": respWebhookData.URL,
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
		"recipients": respEmailData.Recipients,
	}

	return []interface{}{m}
}

func flattenEmailPerFindingData(respEmailPerFindingData *continuous_compliance_notification.EmailPerFindingData) []interface{} {
	m := map[string]interface{}{
		"recipients":                 respEmailPerFindingData.Recipients,
		"notification_output_format": respEmailPerFindingData.NotificationOutputFormat,
	}

	return []interface{}{m}
}

func ValidatePayloadFormatJSON(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) < 1 || value[:1] != "{" {
		errors = append(errors, fmt.Errorf("%q Contains an invalid JSON policy", k))
		return
	}
	if _, err := structure.NormalizeJsonString(v); err != nil {
		errors = append(errors, fmt.Errorf("%q Contains an invalid JSON: %s", k, err))
		return
	}
	return
}