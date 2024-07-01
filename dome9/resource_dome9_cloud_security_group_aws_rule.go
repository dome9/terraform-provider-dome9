package dome9

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/cloudsecuritygroup/securitygroupaws"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
)

func resourceCloudSecurityGroupAWSRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudSecurityGroupAWSRuleCreate,
		Read:   resourceCloudSecurityGroupAWSRuleRead,
		Update: resourceCloudSecurityGroupAWSRuleUpdate,
		Delete: resourceCloudSecurityGroupAWSRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"dome9_security_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"services": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inbound": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// required to create inbound
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									// required to create inbound
									"protocol_type": {
										Optional:     true,
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice(providerconst.ProtocolTypes, true),
									},
									"port": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"open_for_all": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"scope": {
										Type:     schema.TypeSet,
										Optional: true,

										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:     schema.TypeString,
													Required: true,
												},
												"data": {
													Type:     schema.TypeMap,
													Required: true,
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
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// required to create outbound
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									// required to create inbound
									"protocol_type": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice(providerconst.ProtocolTypes, true),
									},
									"port": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"open_for_all": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"scope": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"data": {
													Type:     schema.TypeMap,
													Optional: true,
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

func resourceCloudSecurityGroupAWSRuleCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandBoundServices(d)
	log.Printf("[INFO] Bounding service to AWS security group request:%+v\n", req)
	cloudAccountID := d.Get("dome9_security_group_id").(string)
	resp, _, err := d9Client.awsSecurityGroup.UpdateBoundService(cloudAccountID, req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Bounding service to AWS security group. ID: %v\n", resp.ID)
	d.SetId(strconv.Itoa(resp.ID))

	return resourceCloudSecurityGroupAWSRuleRead(d, meta)
}

func resourceCloudSecurityGroupAWSRuleRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.awsSecurityGroup.Get(d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing AWS cloud account security group %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting AWS cloud account security group:\n%+v\n", resp)
	_ = d.Set("dome9_security_group_id", resp.ID)

	if err := d.Set("services", flattenCloudSecurityGroupAWSServices(resp.Services)); err != nil {
		return err
	}

	return nil
}

func resourceCloudSecurityGroupAWSRuleDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Dettach all the inbounds and outbounds from AWS security group ID: %v", d.Id())
	cloudAccountID := d.Get("dome9_security_group_id").(string)

	// detach all the input and outbound
	req := securitygroupaws.UpdateBoundServiceRequest{
		Services: securitygroupaws.ServicesRequest{
			Inbound:  []securitygroupaws.BoundServicesRequest{},
			Outbound: []securitygroupaws.BoundServicesRequest{},
		},
	}

	_, _, err := d9Client.awsSecurityGroup.UpdateBoundService(cloudAccountID, req)
	if err != nil {
		return err
	}

	return nil
}

func resourceCloudSecurityGroupAWSRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	cloudAccountID := d.Get("dome9_security_group_id").(string)
	if _, _, err := d9Client.awsSecurityGroup.UpdateBoundService(cloudAccountID, expandBoundServices(d)); err != nil {
		return err
	}

	return nil
}

func expandBoundServices(d *schema.ResourceData) securitygroupaws.UpdateBoundServiceRequest {
	return securitygroupaws.UpdateBoundServiceRequest{
		Services: *expandServices(d),
	}
}
