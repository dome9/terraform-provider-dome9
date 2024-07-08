package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNotification() *schema.Resource {
	return &schema.Resource{
		Read: resourceNotificationRead,
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
			"send_on_each_occurrence": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"origin": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"integration_settings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reports_integration_settings": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"integration_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"output_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"single_notification_integration_settings": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"integration_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"output_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"payload": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"scheduled_integration_settings": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"integration_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"output_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cron_expression": {
										Type:     schema.TypeString,
										Computed: true,
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
