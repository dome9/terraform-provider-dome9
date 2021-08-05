package dome9

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/dome9/dome9-sdk-go/dome9/client"
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
				Optional: true,
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
							Optional: true,
						},
						"api_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"secret": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"iam_user": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The cloud account onboarding method. Should be set to 'RoleBased' for AWS and 'UserBased' forAWS_GOV.",
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
										Type:         schema.TypeString,
										Computed:     true,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"ReadOnly", "FullManage", "Reset"}, true),
									},
								},
							},
						},
					},
				},
			},
			"iam_safe": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aws_group_arn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aws_policy_arn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"restricted_iam_entities": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"roles_arns": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"users_arns": {
										Type:     schema.TypeList,
										Computed: true,
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
			"full_protection": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_read_only": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"organizational_unit_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCloudAccountAWSCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandCloudAccountAWSRequest(d)
	log.Printf("[INFO] Creating AWS Cloud Account with request\n%+v\n", req)
	resp, _, err := d9Client.cloudaccountAWS.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created AWS CloudAccount. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceCloudAccountAWSRead(d, meta)
}

func resourceCloudAccountAWSRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	getCloudAccountQueryParams := cloudaccounts.QueryParameters{ID: d.Id()}
	resp, _, err := d9Client.cloudaccountAWS.Get(&getCloudAccountQueryParams)

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing AWS cloud account %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
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
	_ = d.Set("organizational_unit_id", resp.OrganizationalUnitID)

	if err := d.Set("net_sec", flattenCloudAccountAWSNetSec(resp.NetSec)); err != nil {
		return err
	}

	if resp.IamSafe != nil {
		if err := d.Set("iam_safe", flattenCloudAccountIAMSafe(*resp.IamSafe)); err != nil {
			return err
		}
	}

	return nil
}

func resourceCloudAccountAWSDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting AWS Cloud Account ID: %v\n", d.Id())

	if _, err := d9Client.cloudaccountAWS.ForceDelete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceCloudAccountAWSUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Println("An updated occurred")

	if d.HasChange("name") {
		log.Println("The name has been changed")

		if _, _, err := d9Client.cloudaccountAWS.UpdateName(aws.CloudAccountUpdateNameRequest{
			CloudAccountID:        d.Id(),
			ExternalAccountNumber: d.Get("external_account_number").(string),
			Data:                  d.Get("name").(string),
		}); err != nil {
			return err
		}
	}

	if d.HasChange("organizational_unit_id") {
		log.Println("The Organizational Unit ID has been changed")

		if _, _, err := d9Client.cloudaccountAWS.UpdateOrganizationalID(d.Id(), aws.CloudAccountUpdateOrganizationalIDRequest{
			OrganizationalUnitId: d.Get("organizational_unit_id").(string),
		}); err != nil {
			return err
		}
	}

	if d.HasChange("credentials.0") {
		log.Println("credentials has been changed")

		if _, _, err := d9Client.cloudaccountAWS.UpdateCredentials(aws.CloudAccountUpdateCredentialsRequest{
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
				if _, _, err := d9Client.cloudaccountAWS.UpdateRegionConfig(aws.CloudAccountUpdateRegionConfigRequest{
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
	credentials := expandCloudAccountAWSCredentials(d)
	vendor := "aws"
	if credentials.Type == "UserBased" {
		vendor = "awsgov"
	}
	return aws.CloudAccountRequest{
		Name:                 d.Get("name").(string),
		Credentials:          credentials,
		OrganizationalUnitID: d.Get("organizational_unit_id").(string),
		Vendor:               vendor,
	}
}

func expandCloudAccountAWSCredentials(d *schema.ResourceData) aws.CloudAccountCredentials {
	onbordingType := d.Get("credentials.0.type").(string)
	if onbordingType == "UserBased" {
		return aws.CloudAccountCredentials{
			ApiKey:     d.Get("credentials.0.api_key").(string),
			Secret:     d.Get("credentials.0.secret").(string),
			Type:       onbordingType,
			IsReadOnly: d.Get("credentials.0.is_read_only").(bool),
		}
	}

	return aws.CloudAccountCredentials{
		Arn:        d.Get("credentials.0.arn").(string),
		Secret:     d.Get("credentials.0.secret").(string),
		Type:       onbordingType,
		IsReadOnly: d.Get("credentials.0.is_read_only").(bool),
	}
}

func flattenCloudAccountAWSNetSec(responseNetSec aws.CloudAccountNetSec) []interface{} {
	m := map[string]interface{}{
		"regions": flattenCloudAccountAWSNetSecRegions(responseNetSec.Regions),
	}

	return []interface{}{m}
}

func flattenCloudAccountIAMSafe(responseIAMSafe aws.CloudAccountIamSafe) []interface{} {
	m := map[string]interface{}{
		"aws_group_arn":           responseIAMSafe.AwsGroupArn,
		"aws_policy_arn":          responseIAMSafe.AwsPolicyArn,
		"mode":                    responseIAMSafe.Mode,
		"restricted_iam_entities": flattenRestrictedIamEntities(responseIAMSafe.RestrictedIamEntities),
	}

	return []interface{}{m}
}

func flattenRestrictedIamEntities(restrictedIamEntities aws.CloudAccountIamEntities) []interface{} {
	m := map[string]interface{}{
		"roles_arns": restrictedIamEntities.RolesArn,
		"users_arns": restrictedIamEntities.UsersArn,
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
