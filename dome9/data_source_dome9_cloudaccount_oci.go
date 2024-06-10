package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/oci"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
	"log"
)

func dataSourceCloudAccountOCI() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOciRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"home_region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"organizational_unit_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"credentials": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fingerprint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"creation_date": {
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

func dataSourceOciRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := d.Get("id").(string)
	log.Printf("Getting data for cloud account %s with id %s\n", variable.CloudAccountOciVendor, id)

	ociCloudAccount, _, err := d9Client.cloudaccountOci.Get(id)
	if err != nil {
		return err
	}

	d.SetId(ociCloudAccount.ID)
	_ = d.Set("name", ociCloudAccount.Name)
	// Converting the timestamp to string in the format yyyy-MM-dd HH:mm:ss
	_ = d.Set("creation_date", ociCloudAccount.CreationDate.Format("2006-01-02 15:04:05"))
	_ = d.Set("tenancy_id", ociCloudAccount.TenancyId)
	_ = d.Set("home_region", ociCloudAccount.HomeRegion)
	_ = d.Set("organizational_unit_id", ociCloudAccount.OrganizationalUnitID)
	_ = d.Set("credentials", setOciCredentials(ociCloudAccount.Credentials))
	_ = d.Set("organizational_unit_path", ociCloudAccount.OrganizationalUnitPath)
	_ = d.Set("organizational_unit_name", ociCloudAccount.OrganizationalUnitName)
	_ = d.Set("vendor", ociCloudAccount.Vendor)

	return nil
}

func setOciCredentials(credentials oci.CloudAccountCredentialsResponse) map[string]interface{} {
	return map[string]interface{}{
		"user":        credentials.User,
		"fingerprint": credentials.Fingerprint,
		"publicKey":   credentials.PublicKey,
	}
}
