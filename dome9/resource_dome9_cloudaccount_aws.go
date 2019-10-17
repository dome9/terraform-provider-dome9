package dome9

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"

	"github.com/dome9/terraform-provider-dome9/dome9/common/providerconst"
)

func resourceCloudAccountAWS() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudAccountAWSCreate,
		Read:   resourceCloudAccountAWSRead,
		Update: resourceCloudAccountAWSUpdate,
		Delete: resourceCloudAccountAWSDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vendor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_account_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_fetching_suspended": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"creation_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"credentials": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:     schema.TypeString,
							Required: true,
						},
						"api_key": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "Deprecated",
						},
						"secret": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"iam_user": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "Deprecated",
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The cloud account onboarding method. Should be set to 'RoleBased' as other methods are deprecated",
						},
						"is_read_only": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"net_sec": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"regions": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice(providerconst.AWSRegions, true),
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"hidden": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"new_group_behavior": {
										Type:     schema.TypeString,
										Computed: true,
										Optional: true,
										// Reset == regionlock
										ValidateFunc: validation.StringInSlice([]string{"ReadOnly", "FullManage", "Reset"}, true),
									},
								},
							},
						},
					},
				},
			},
			// TODO: full_protection and allow_read_only are currently computed only since the server always returns false
			"full_protection": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_read_only": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceCloudAccountAWSCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	req := expandCloudAccountAWSRequest(d)
	log.Printf("[INFO] Creating AWS Cloud Account with request\n%+v\n", req)
	resp, _, err := client.cloudaccountAWS.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created AWS CloudAccount. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceCloudAccountAWSRead(d, meta)
}

func resourceCloudAccountAWSRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	getCloudAccountQueryParams := cloudaccounts.QueryParameters{ID: d.Id()}
	resp, _, err := client.cloudaccountAWS.Get(&getCloudAccountQueryParams)
	if err != nil {
		return nil
	}

	log.Printf("[INFO] Reading account response and settings states: %+v\n", resp)
	d.SetId(resp.ID)
	_ = d.Set("vendor", resp.Vendor)
	_ = d.Set("name", resp.Name)
	_ = d.Set("external_account_number", resp.ExternalAccountNumber)
	_ = d.Set("is_fetching_suspended", resp.IsFetchingSuspended)
	_ = d.Set("creation_date", resp.CreationDate.Format("2006-01-02 15:04:05"))
	_ = d.Set("full_protection", resp.FullProtection)
	_ = d.Set("allow_read_only", resp.AllowReadOnly)
	if err := d.Set("net_sec", flattenCloudAccountAWSNetSec(resp.NetSec)); err != nil {
		return err
	}

	return nil
}

func resourceCloudAccountAWSDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	log.Printf("[INFO] Deleting AWS Cloud Account ID: %v\n", d.Id())

	if _, err := client.cloudaccountAWS.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceCloudAccountAWSUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	log.Println("An updated occurred")

	if d.HasChange("name") {
		log.Println("The name has been changed")

		if _, _, err := client.cloudaccountAWS.UpdateName(aws.CloudAccountUpdateNameRequest{
			CloudAccountID:        d.Id(),
			ExternalAccountNumber: d.Get("external_account_number").(string),
			Data:                  d.Get("name").(string),
		}); err != nil {
			return err
		}
	}

	if d.HasChange("credentials.0") {
		log.Println("credentials has been changed")

		if _, _, err := client.cloudaccountAWS.UpdateCredentials(aws.CloudAccountUpdateCredentialsRequest{
			CloudAccountID: d.Id(),
			Data:           expandCloudAccountAWSCredentials(d),
		}); err != nil {
			return err
		}
	}

	if d.HasChange("net_sec.0.regions") {
		log.Println("NetSec regions has been changed")

		netSecRegionsIterator := d.Get("net_sec.0.regions").([]interface{})
		for i, val := range netSecRegionsIterator {
			regionObject := val.(map[string]interface{})
			newGroupBehaviorKeyFormat := fmt.Sprintf("net_sec.0.regions.%d.new_group_behavior", i)
			if d.HasChange(newGroupBehaviorKeyFormat) {
				if _, _, err := client.cloudaccountAWS.UpdateRegionConfig(aws.CloudAccountUpdateRegionConfigRequest{
					CloudAccountID: d.Id(),
					Data: aws.CloudAccountNetSecRegion{
						Region:           regionObject["region"].(string),
						NewGroupBehavior: regionObject["new_group_behavior"].(string),
					},
				}); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func expandCloudAccountAWSRequest(d *schema.ResourceData) aws.CloudAccountRequest {
	return aws.CloudAccountRequest{
		Name:        d.Get("name").(string),
		Credentials: expandCloudAccountAWSCredentials(d),
	}
}

func expandCloudAccountAWSCredentials(d *schema.ResourceData) aws.CloudAccountCredentials {
	return aws.CloudAccountCredentials{
		ApiKey:     d.Get("credentials.0.api_key").(string),
		Arn:        d.Get("credentials.0.arn").(string),
		Secret:     d.Get("credentials.0.secret").(string),
		IamUser:    d.Get("credentials.0.iam_user").(string),
		Type:       d.Get("credentials.0.type").(string),
		IsReadOnly: d.Get("credentials.0.is_read_only").(bool),
	}
}

func flattenCloudAccountAWSNetSec(responseNetSec aws.CloudAccountNetSec) []interface{} {
	m := map[string]interface{}{
		"regions": flattenCloudAccountAWSNetSecRegions(responseNetSec.Regions),
	}

	return []interface{}{m}
}

func flattenCloudAccountAWSNetSecRegions(responseNetSecRegions []aws.CloudAccountNetSecRegion) []interface{} {
	netSecRegions := make([]interface{}, len(responseNetSecRegions))
	for i, val := range responseNetSecRegions {
		netSecRegions[i] = map[string]interface{}{
			"region":             val.Region,
			"name":               val.Name,
			"hidden":             val.Hidden,
			"new_group_behavior": val.NewGroupBehavior,
		}
	}

	return netSecRegions
}
