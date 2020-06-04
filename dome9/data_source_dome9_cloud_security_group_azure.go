package dome9

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSecurityGroupAzure() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecurityGroupAzureRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dome9_security_group_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dome9_cloud_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_tamper_protected": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"tags": {
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
			"inbound": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"access": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_port_ranges": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"source_scopes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"data": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{},
										},
									},
								},
							},
						},
						"destination_port_ranges": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"destination_scopes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"data": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{},
										},
									},
								},
							},
						},
						"direction": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"outbound": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"access": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_port_ranges": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"source_scopes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"data": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{},
										},
									},
								},
							},
						},
						"destination_port_ranges": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"destination_scopes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"data": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{},
										},
									},
								},
							},
						},
						"direction": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"external_security_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_account_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by_dome9": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSecurityGroupAzureRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for azure security group with id %s\n", id)

	resp, _, err := d9Client.azureSecurityGroup.Get(id)
	if err != nil {
		return err
	}

	d.SetId(resp.ID)
	_ = d.Set("dome9_security_group_name", resp.Name)
	_ = d.Set("region", resp.Region)
	_ = d.Set("resource_group", resp.ResourceGroup)
	_ = d.Set("dome9_cloud_account_id", resp.CloudAccountID)
	_ = d.Set("description", resp.Description)
	_ = d.Set("is_tamper_protected", resp.IsTamperProtected)
	_ = d.Set("external_security_group_id", resp.ExternalSecurityGroupID)
	_ = d.Set("cloud_account_name", resp.CloudAccountName)
	_ = d.Set("last_updated_by_dome9", resp.LastUpdatedByDome9)

	if err := d.Set("tags", flattenSecurityGroupAzureTags(resp.Tags)); err != nil {
		return err
	}

	if err := d.Set("inbound", flattenSecurityGroupAzureServices(resp.InboundServices)); err != nil {
		return err
	}

	if err := d.Set("outbound", flattenSecurityGroupAzureServices(resp.OutboundServices)); err != nil {
		return err
	}

	return nil
}
