package dome9

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"

	"github.com/dome9/terraform-provider-dome9/dome9/common/structservers"
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
				Type:     schema.TypeSet,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"regions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": {
										Type:     schema.TypeString,
										Computed: true,
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
		},
	}
}

func expandingCloudAccountRequestAWS(d *schema.ResourceData) *aws.CloudAccountRequest {
	var req aws.CloudAccountRequest
	req.Name = d.Get("name").(string)
	req.Credentials.ApiKey = d.Get("credentials.0.api_key").(string)
	req.Credentials.Arn = d.Get("credentials.0.arn").(string)
	req.Credentials.Secret = d.Get("credentials.0.secret").(string)
	req.Credentials.IamUser = d.Get("credentials.0.iam_user").(string)
	req.Credentials.IamUser = d.Get("credentials.0.iam_user").(string)
	req.Credentials.Type = d.Get("credentials.0.type").(string)
	req.Credentials.IsReadOnly = d.Get("credentials.0.is_read_only").(bool)
	log.Printf("[INFO] Creating AWS Cloud Account request: %+v\n", req)

	return &req
}

func resourceCloudAccountAWSCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	req := expandingCloudAccountRequestAWS(d)
	log.Printf("[INFO] Creating AWS Cloud Account with request %+v\n", req)
	resp, _, err := client.cloudaccountAWS.Create(*req)
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
	_ = d.Set("net_sec", structservers.FlattenNetSec(resp))

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
			Data: struct {
				Apikey     string `json:"apikey,omitempty"`
				Arn        string `json:"arn,omitempty"`
				Secret     string `json:"secret,omitempty"`
				IamUser    string `json:"iamUser,omitempty"`
				Type       string `json:"type,omitempty"`
				IsReadOnly bool   `json:"isReadOnly,omitempty"`
			}{
				Apikey:  d.Get("credentials.0.api_key").(string),
				Arn:     d.Get("credentials.0.arn").(string),
				Secret:  d.Get("credentials.0.secret").(string),
				IamUser: d.Get("credentials.0.iam_user").(string),
				Type:    d.Get("credentials.0.type").(string),
			},
		}); err != nil {
			return err
		}
	}

	return nil
}
