package dome9

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"

	"github.com/dome9/terraform-provider-dome9/dome9/common/testing/variable"
)

func dataSourceCloudAccountGCP() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGCPRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_date": {
				Type:     schema.TypeString,
				Computed: true,
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
			"gsuite_user": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_name": {
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

func dataSourceGCPRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	id := d.Get("account_id").(string)
	log.Printf("Getting data for %s cloud account with id %s\n", variable.CloudAccountGCPVendor, id)

	GCPCloudAccount, _, err := client.cloudaccountGCP.Get(cloudaccounts.QueryParameters{ID: id})
	if err != nil {
		return err
	}

	d.SetId(GCPCloudAccount.ID)
	_ = d.Set("name", GCPCloudAccount.Name)
	_ = d.Set("project_id", GCPCloudAccount.ProjectID)
	// Converting the timestamp to string in the format yyyy-MM-dd HH:mm:ss
	_ = d.Set("creation_date", GCPCloudAccount.CreationDate.Format("2006-01-02 15:04:05"))
	_ = d.Set("organizational_unit_id", GCPCloudAccount.OrganizationalUnitID)
	_ = d.Set("organizational_unit_path", GCPCloudAccount.OrganizationalUnitPath)
	_ = d.Set("organizational_unit_name", GCPCloudAccount.OrganizationalUnitName)
	_ = d.Set("gsuite_user", GCPCloudAccount.GSuite.GSuiteUser)
	_ = d.Set("domain_name", GCPCloudAccount.GSuite.DomainName)
	_ = d.Set("vendor", GCPCloudAccount.Vendor)

	return nil
}
