package dome9

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCloudSecurityGroupAWSRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecurityGroupAWSRuleRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"services": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inbound": {
							Type:     schema.TypeSet,
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
									"protocol_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"open_for_all": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"scope": {
										Type:     schema.TypeSet,
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
								},
							},
						},
						"outbound": {
							Type:     schema.TypeSet,
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
									"protocol_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"open_for_all": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"scope": {
										Type:     schema.TypeSet,
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceSecurityGroupAWSRuleRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting inbounds and outbounds for aws security group with id %s\n", id)

	resp, _, err := d9Client.awsSecurityGroup.Get(id)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(resp.ID))
	_ = d.Set("dome9_security_group_id", id)
	_ = d.Set("services", flattenCloudSecurityGroupAWSServices(resp.Services))

	return nil
}
