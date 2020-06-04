package dome9

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCloudSecurityGroupAWS() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecurityGroupAWSRead,
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
			"dome9_cloud_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_protected": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"cloud_account_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpc_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			"services": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceSecurityGroupAWSRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for aws security group with id %s\n", id)

	resp, _, err := d9Client.awsSecurityGroup.Get(id)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(resp.ID))
	_ = d.Set("dome9_security_group_name", resp.SecurityGroupName)
	_ = d.Set("cloud_account_id", resp.CloudAccountID)
	_ = d.Set("description", resp.Description)
	_ = d.Set("aws_region_id", resp.RegionID)
	_ = d.Set("is_protected", resp.IsProtected)
	_ = d.Set("cloud_account_name", resp.CloudAccountName)
	_ = d.Set("vpc_id", resp.VpcID)
	_ = d.Set("external_id", resp.VpcID)
	_ = d.Set("tags", resp.Tags)
	_ = d.Set("services", flattenCloudSecurityGroupAWSServices(resp.Services))

	if resp.VpcName != nil {
		_ = d.Set("vpc_name", *resp.VpcName)
	}
	return nil
}
