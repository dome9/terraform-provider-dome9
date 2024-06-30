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

func resourceCloudSecurityGroupAWS() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudSecurityGroupAWSCreate,
		Read:   resourceCloudSecurityGroupAWSRead,
		Update: resourceCloudSecurityGroupAWSUpdate,
		Delete: resourceCloudSecurityGroupAWSDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"dome9_security_group_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dome9_cloud_account_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"aws_region_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "us_east_1",
				ValidateFunc: validation.StringInSlice(providerconst.AllAWSRegions, true),
			},
			// Always true in creation.
			"is_protected": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cloud_account_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"vpc_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
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

func resourceCloudSecurityGroupAWSCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandCloudSecurityGroupRequest(d)
	log.Printf("[INFO] Creating AWS security group request:%+v\n", req)
	resp, _, err := d9Client.awsSecurityGroup.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created AWS security group. ID: %v\n", resp.ID)
	d.SetId(strconv.Itoa(resp.ID))

	return resourceCloudSecurityGroupAWSRead(d, meta)
}

func resourceCloudSecurityGroupAWSRead(d *schema.ResourceData, meta interface{}) error {
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
	d.SetId(strconv.Itoa(resp.ID))
	_ = d.Set("dome9_security_group_name", resp.SecurityGroupName)
	_ = d.Set("cloud_account_id", resp.CloudAccountID)
	_ = d.Set("description", resp.Description)
	_ = d.Set("aws_region_id", resp.RegionID)
	_ = d.Set("is_protected", resp.IsProtected)
	_ = d.Set("cloud_account_name", resp.CloudAccountName)
	_ = d.Set("vpc_id", resp.VpcID)
	_ = d.Set("external_id", resp.ExternalID)
	_ = d.Set("tags", resp.Tags)

	if err := d.Set("services", flattenCloudSecurityGroupAWSServices(resp.Services)); err != nil {
		return err
	}

	if resp.VpcName != nil {
		_ = d.Set("vpc_name", *resp.VpcName)
	}

	return nil
}

func resourceCloudSecurityGroupAWSDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting AWS security group ID: %v", d.Id())

	if _, err := d9Client.awsSecurityGroup.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceCloudSecurityGroupAWSUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	if d.HasChange("is_protected") {
		protectionMode := getProtectionMode(d.Get("is_protected").(bool))
		log.Printf("[INFO] Updating security group protection mode to: %s", protectionMode)
		if _, _, err := d9Client.awsSecurityGroup.UpdateProtectionMode(d.Id(), protectionMode); err != nil {
			return err
		}
	}
	if d.HasChange("tags") || d.HasChange("services") {
		log.Println("[INFO] Tags or services has been changed")

		if _, _, err := d9Client.awsSecurityGroup.Update(d.Id(), expandCloudSecurityGroupRequest(d)); err != nil {
			return err
		}
	}

	return nil
}

func getProtectionMode(isProtected bool) (protectionMode string) {
	if isProtected {
		protectionMode = providerconst.FullManage
	} else {
		protectionMode = providerconst.ReadOnly
	}

	return
}

func expandCloudSecurityGroupRequest(d *schema.ResourceData) securitygroupaws.CloudSecurityGroupRequest {
	cloudSecurityGroupRequest := securitygroupaws.CloudSecurityGroupRequest{
		SecurityGroupName: d.Get("dome9_security_group_name").(string),
		CloudAccountID:    d.Get("dome9_cloud_account_id").(string),
		Services:          expandServices(d),
	}

	if description, ok := d.GetOk("description"); ok {
		cloudSecurityGroupRequest.Description = description.(string)
	}

	if awsRegionID, ok := d.GetOk("aws_region_id"); ok {
		cloudSecurityGroupRequest.RegionID = awsRegionID.(string)
	}

	if isProtected, ok := d.GetOk("is_protected"); ok {
		cloudSecurityGroupRequest.IsProtected = isProtected.(bool)
	}

	if vpcID, ok := d.GetOk("vpc_id"); ok {
		cloudSecurityGroupRequest.VpcId = vpcID.(string)
	}

	if vpcName, ok := d.GetOk("vpc_name"); ok {
		cloudSecurityGroupRequest.VpcName = vpcName.(string)
	}

	if tags, ok := d.GetOk("tags"); ok {
		cloudSecurityGroupRequest.Tags = tags.(map[string]interface{})
	}

	return cloudSecurityGroupRequest
}

func expandServices(d *schema.ResourceData) *securitygroupaws.ServicesRequest {
	if services, ok := d.GetOk("services"); ok {
		servicesItem := services.(*schema.Set).List()[0]
		service := servicesItem.(map[string]interface{})

		return &securitygroupaws.ServicesRequest{
			Inbound:  expandBoundServicesRequest(service["inbound"].(*schema.Set)),
			Outbound: expandBoundServicesRequest(service["outbound"].(*schema.Set)),
		}
	}

	return nil
}

func expandBoundServicesRequest(boundServicesRequest *schema.Set) []securitygroupaws.BoundServicesRequest {
	boundServices := make([]securitygroupaws.BoundServicesRequest, boundServicesRequest.Len())

	for i, boundService := range boundServicesRequest.List() {
		boundServiceItem := boundService.(map[string]interface{})

		boundServices[i] = securitygroupaws.BoundServicesRequest{
			Name:         boundServiceItem["name"].(string),
			Description:  boundServiceItem["description"].(string),
			ProtocolType: boundServiceItem["protocol_type"].(string),
			Port:         boundServiceItem["port"].(string),
			OpenForAll:   boundServiceItem["open_for_all"].(bool),
			Scope:        expandScope(boundServiceItem["scope"].(*schema.Set)),
		}
	}

	return boundServices
}

func expandScope(scopeRequest *schema.Set) []securitygroupaws.Scope {
	scopes := make([]securitygroupaws.Scope, scopeRequest.Len())
	for i, scope := range scopeRequest.List() {
		scopeItem := scope.(map[string]interface{})
		scopes[i] = securitygroupaws.Scope{
			Type: scopeItem["type"].(string),
			Data: scopeItem["data"].(map[string]interface{}),
		}
	}

	return scopes
}

func flattenCloudSecurityGroupAWSServices(servicesResponse securitygroupaws.ServicesResponse) []interface{} {
	m := map[string]interface{}{
		"inbound":  flattenBoundServicesResponse(servicesResponse.Inbound),
		"outbound": flattenBoundServicesResponse(servicesResponse.Outbound),
	}

	return []interface{}{m}
}

func flattenBoundServicesResponse(boundServicesResponse []securitygroupaws.BoundServicesResponse) []interface{} {
	boundServices := make([]interface{}, len(boundServicesResponse))
	for i, boundServiceItem := range boundServicesResponse {
		boundServices[i] = map[string]interface{}{
			"name":          boundServiceItem.Name,
			"description":   boundServiceItem.Description,
			"protocol_type": boundServiceItem.ProtocolType,
			"port":          boundServiceItem.Port,
			"open_for_all":  boundServiceItem.OpenForAll,
			"scope":         flattenScope(boundServiceItem.Scope),
		}
	}
	return boundServices
}

func flattenScope(scopeBoundServices []securitygroupaws.Scope) []interface{} {
	scopes := make([]interface{}, len(scopeBoundServices))
	for i, scopeBoundServicesItem := range scopeBoundServices {
		scopes[i] = map[string]interface{}{
			"type": scopeBoundServicesItem.Type,
			"data": scopeBoundServicesItem.Data,
		}
	}

	return scopes
}
