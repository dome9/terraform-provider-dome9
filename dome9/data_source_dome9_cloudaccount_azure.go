package dome9

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func dataSourceCloudAccountAzure() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAzureRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operation_mode": {
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
			"vendor": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAzureRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := d.Get("id").(string)
	log.Printf("Getting data for cloud account %s with id %s\n", variable.CloudAccountAzureVendor, id)

	azureCloudAccount, _, err := d9Client.cloudaccountAzure.Get(cloudaccounts.QueryParameters{ID: id})
	if err != nil {
		return err
	}

	d.SetId(azureCloudAccount.ID)
	_ = d.Set("name", azureCloudAccount.Name)
	_ = d.Set("subscription_id", azureCloudAccount.SubscriptionID)
	_ = d.Set("tenant_id", azureCloudAccount.TenantID)
	_ = d.Set("operation_mode", azureCloudAccount.OperationMode)
	// Converting the timestamp to string in the format yyyy-MM-dd HH:mm:ss
	_ = d.Set("creation_date", azureCloudAccount.CreationDate.Format("2006-01-02 15:04:05"))
	_ = d.Set("organizational_unit_id", azureCloudAccount.OrganizationalUnitID)
	_ = d.Set("organizational_unit_path", azureCloudAccount.OrganizationalUnitPath)
	_ = d.Set("organizational_unit_name", azureCloudAccount.OrganizationalUnitName)
	_ = d.Set("vendor", azureCloudAccount.Vendor)

	return nil
}
