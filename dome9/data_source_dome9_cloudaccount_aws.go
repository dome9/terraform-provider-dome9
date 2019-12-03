package dome9

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func dataSourceCloudAccountAWS() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAWSRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vendor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
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
				Computed: true,
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
		},
	}
}

func dataSourceAWSRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for cloud account %s with id %s\n", variable.CloudAccountAWSVendor, id)

	resp, _, err := d9Client.cloudaccountAWS.Get(cloudaccounts.QueryParameters{ID: id})
	if err != nil {
		return err
	}

	d.SetId(resp.ID)
	_ = d.Set("vendor", resp.Vendor)
	_ = d.Set("name", resp.Name)
	_ = d.Set("external_account_number", resp.ExternalAccountNumber)
	_ = d.Set("is_fetching_suspended", resp.IsFetchingSuspended)
	// Converting the timestamp to string in the format yyyy-MM-dd HH:mm:ss
	_ = d.Set("creation_date", resp.CreationDate.Format("2006-01-02 15:04:05"))
	_ = d.Set("full_protection", resp.FullProtection)
	_ = d.Set("allow_read_only", resp.AllowReadOnly)
	_ = d.Set("organizational_unit_id", resp.OrganizationalUnitID)
	if err := d.Set("net_sec", flattenCloudAccountAWSNetSec(resp.NetSec)); err != nil {
		return err
	}

	return nil
}
