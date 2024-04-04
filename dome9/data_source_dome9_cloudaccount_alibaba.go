package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/alibaba"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func dataSourceCloudAccountAlibaba() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlibabaRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"alibaba_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"credentials": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"organizational_unit_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"organizational_unit_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"organizational_unit_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vendor": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAlibabaRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := d.Get("id").(string)
	log.Printf("Getting data for cloud account %s with id %s\n", variable.CloudAccountAlibabaVendor, id)

	alibabaCloudAccount, _, err := d9Client.cloudaccountAlibaba.Get(id)
	if err != nil {
		return err
	}

	d.SetId(alibabaCloudAccount.ID)
	_ = d.Set("name", alibabaCloudAccount.Name)
	_ = d.Set("alibaba_account_id", alibabaCloudAccount.AlibabaAccountId)
	// Converting the timestamp to string in the format yyyy-MM-dd HH:mm:ss
	_ = d.Set("creation_date", alibabaCloudAccount.CreationDate.Format("2006-01-02 15:04:05"))
	_ = d.Set("organizational_unit_id", alibabaCloudAccount.OrganizationalUnitID)
	_ = d.Set("credentials", setCredentials(alibabaCloudAccount.Credentials))
	_ = d.Set("organizational_unit_path", alibabaCloudAccount.OrganizationalUnitPath)
	_ = d.Set("organizational_unit_name", alibabaCloudAccount.OrganizationalUnitName)
	_ = d.Set("vendor", alibabaCloudAccount.Vendor)

	return nil
}

func setCredentials(credentials alibaba.CloudAccountCredentialsResponse) map[string]interface{} {
	return map[string]interface{}{
		"access_key": credentials.AccessKey,
	}
}
