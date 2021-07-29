package dome9

import (
	"github.com/dome9/terraform-provider-dome9/dome9/common/providerconst"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/cloudsecuritygroup/securitygroupazure"

	//"github.com/dome9/terraform-provider-dome9/dome9/common/providerconst"
)

func resourceAzureSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityGroupAzureCreate,
		Read:   resourceSecurityGroupAzureRead,
		Update: resourceSecurityGroupAzureUpdate,
		Delete: resourceSecurityGroupAzureDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"dome9_security_group_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_group": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dome9_cloud_account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_tamper_protected": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tags": {
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
			"inbound": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"access": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "Allow",
							ValidateFunc: validation.StringInSlice(providerconst.AzureSecurityGroupAccess, true),
						},
						"protocol": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice(providerconst.AzureSecurityGroupProtocol, true),
						},
						"source_port_ranges": {
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"source_scopes": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice(providerconst.AzureSecurityGroupSourceScopeTypes, true),
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
						"destination_port_ranges": {
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"destination_scopes": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice(providerconst.AzureSecurityGroupSourceScopeTypes, true),
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
						"direction": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_default": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"outbound": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"access": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "Allow",
							ValidateFunc: validation.StringInSlice(providerconst.AzureSecurityGroupAccess, true),
						},
						"protocol": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice(providerconst.AzureSecurityGroupProtocol, true),
						},
						"source_port_ranges": {
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"source_scopes": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice(providerconst.AzureSecurityGroupSourceScopeTypes, true),
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
						"destination_port_ranges": {
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"destination_scopes": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice(providerconst.AzureSecurityGroupSourceScopeTypes, true),
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
						"direction": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_default": {
							Type:     schema.TypeBool,
							Optional: true,
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

func resourceSecurityGroupAzureCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandSecurityGroupAzureRequest(d)
	log.Printf("[INFO] Creating Azure security group request:%+v\n", req)
	resp, _, err := d9Client.azureSecurityGroup.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created Azure security group. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceSecurityGroupAzureRead(d, meta)
}

func resourceSecurityGroupAzureRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.azureSecurityGroup.Get(d.Id())
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing Azure security group %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting Azure security group %+v\n", resp)
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

func resourceSecurityGroupAzureDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting Azure security group ID: %v\n", d.Id())
	if _, err := d9Client.azureSecurityGroup.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceSecurityGroupAzureUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Updating Azure sequrity group ID: %v\n", d.Id())
	req := expandSecurityGroupAzureRequest(d)
	if _, _, err := d9Client.azureSecurityGroup.Update(d.Id(), req); err != nil {
		return err
	}

	return nil
}

func expandSecurityGroupAzureRequest(d *schema.ResourceData) securitygroupazure.AzureSecurityGroupRequest {
	req := securitygroupazure.AzureSecurityGroupRequest{
		Name:             d.Get("dome9_security_group_name").(string),
		Region:           d.Get("region").(string),
		ResourceGroup:    d.Get("resource_group").(string),
		CloudAccountID:   d.Get("dome9_cloud_account_id").(string),
		Tags:             expandSecurityGroupAzureTags(d.Get("tags").([]interface{})),
		InboundServices:  expandSecurityGroupAzureServices(d.Get("inbound").([]interface{})),
		OutboundServices: expandSecurityGroupAzureServices(d.Get("outbound").([]interface{})),
	}

	if description, ok := d.GetOk("description"); ok {
		req.Description = description.(string)
	}

	if isTamperProtected, ok := d.GetOk("is_tamper_protected"); ok {
		req.IsTamperProtected = isTamperProtected.(bool)
	}

	return req
}

func expandSecurityGroupAzureServices(boundServicesRequest []interface{}) []securitygroupazure.BoundService {
	boundServices := make([]securitygroupazure.BoundService, len(boundServicesRequest))

	for i, boundService := range boundServicesRequest {
		boundServiceItem := boundService.(map[string]interface{})
		boundServices[i] = securitygroupazure.BoundService{
			Name:                  boundServiceItem["name"].(string),
			Description:           boundServiceItem["description"].(string),
			Priority:              boundServiceItem["priority"].(int),
			Access:                boundServiceItem["access"].(string),
			Protocol:              boundServiceItem["protocol"].(string),
			SourcePortRanges:      expandSecurityGroupAzurePortRanges(boundServiceItem["source_port_ranges"].([]interface{})),
			SourceScopes:          expandSecurityGroupAzureScope(boundServiceItem["source_scopes"].([]interface{})),
			DestinationPortRanges: expandSecurityGroupAzurePortRanges(boundServiceItem["destination_port_ranges"].([]interface{})),
			DestinationScopes:     expandSecurityGroupAzureScope(boundServiceItem["destination_scopes"].([]interface{})),
			Direction:             boundServiceItem["direction"].(string),
			IsDefault:             boundServiceItem["is_default"].(bool),
		}
	}

	return boundServices
}

func expandSecurityGroupAzureTags(tags []interface{}) []securitygroupazure.Tags {
	securityGroupTags := make([]securitygroupazure.Tags, len(tags))

	for i, tag := range tags {
		tagItem := tag.(map[string]interface{})
		securityGroupTags[i] = securitygroupazure.Tags{
			Key:   tagItem["key"].(string),
			Value: tagItem["value"].(string),
		}
	}

	return securityGroupTags
}

func expandSecurityGroupAzurePortRanges(generalRecipients []interface{}) []string {
	recipients := make([]string, len(generalRecipients))

	for i, recipient := range generalRecipients {
		recipients[i] = recipient.(string)
	}

	return recipients
}

func expandSecurityGroupAzureScope(scopeRequest []interface{}) []securitygroupazure.Scope {
	scopes := make([]securitygroupazure.Scope, len(scopeRequest))

	for i, scope := range scopeRequest {
		scopeItem := scope.(map[string]interface{})
		scopes[i] = securitygroupazure.Scope{
			Type: scopeItem["type"].(string),
			Data: scopeItem["data"].(map[string]interface{}),
		}
	}

	return scopes
}

func flattenSecurityGroupAzureServices(boundServicesResponse []securitygroupazure.BoundService) []interface{} {
	boundServices := make([]interface{}, len(boundServicesResponse))

	for i, boundServiceItem := range boundServicesResponse {
		boundServices[i] = map[string]interface{}{
			"name":                    boundServiceItem.Name,
			"description":             boundServiceItem.Description,
			"priority":                boundServiceItem.Priority,
			"access":                  boundServiceItem.Access,
			"protocol":                boundServiceItem.Protocol,
			"source_port_ranges":      boundServiceItem.SourcePortRanges,
			"source_scopes":           flattenSecurityGroupAzureScope(boundServiceItem.SourceScopes),
			"destination_port_ranges": boundServiceItem.DestinationPortRanges,
			"destination_scopes":      flattenSecurityGroupAzureScope(boundServiceItem.DestinationScopes),
			"direction":               boundServiceItem.Direction,
			"is_default":              boundServiceItem.IsDefault,
		}
	}

	return boundServices
}

func flattenSecurityGroupAzureScope(scopeBoundServices []securitygroupazure.Scope) []interface{} {
	scopes := make([]interface{}, len(scopeBoundServices))

	for i, scopeBoundServicesItem := range scopeBoundServices {
		scopes[i] = map[string]interface{}{
			"type": scopeBoundServicesItem.Type,
			"data": scopeBoundServicesItem.Data,
		}
	}

	return scopes
}

func flattenSecurityGroupAzureTags(azureSecurityGroupTags []securitygroupazure.Tags) []interface{} {
	tags := make([]interface{}, len(azureSecurityGroupTags))

	for i, tag := range azureSecurityGroupTags {
		tags[i] = map[string]interface{}{
			"key":   tag.Key,
			"value": tag.Value,
		}
	}

	return tags
}
