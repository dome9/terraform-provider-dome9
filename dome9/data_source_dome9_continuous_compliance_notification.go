package dome9

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceContinuousComplianceNotification() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceContinuousComplianceNotificationRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"alerts_console": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"scheduled_report": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email_sending_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"schedule_data": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cron_expression": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"recipients": {
										Type:     schema.TypeList,
										Computed: true,
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
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email_sending_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email_per_finding_sending_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sns_sending_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"external_ticket_creating_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_security_hub_integration_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"webhook_integration_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"slack_integration_state": {
							Type:         schema.TypeString,
							Computed: true,
						},
						"teams_integration_state": {
							Type:         schema.TypeString,
							Computed: true,
						},
						"email_data": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"recipients": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"email_per_finding_data": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"recipients": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"notification_output_format": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"sns_data": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sns_topic_arn": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sns_output_format": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"ticketing_system_data": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"system_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"should_close_tickets": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"user": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pass": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"project_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"issue_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"aws_security_hub_integration": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"external_account_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"region": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"webhook_data": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"http_method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"auth_method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"username": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"password": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"format_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"payload_format": {
										Type:     schema.TypeMap,
										Computed: true,
									},
									"ignore_certificate": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"advanced_url": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"slack_data": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"teams_data": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"gcp_security_command_center_integration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceContinuousComplianceNotificationRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for continuous compliance notification with id %s\n", id)

	resp, _, err := d9Client.continuousComplianceNotification.Get(id)
	if err != nil {
		return err
	}

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
