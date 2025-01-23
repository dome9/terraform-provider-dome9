package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceNotification() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNotificationRead,
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

func dataSourceNotificationRead(d *schema.ResourceData, m interface{}) error {
	d.SetId(d.Get("id").(string))
	return resourceNotificationRead(d, m)
}
