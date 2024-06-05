package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/oci"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceCloudAccountOciTempData() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudAccountOciTempDataCreate,
		Read:   resourceCloudAccountOciTempDataRead,
		Update: resourceCloudAccountOciTempDataUpdate,
		Delete: resourceCloudAccountOciTempDataDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"home_region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_administrator_email_address": {
				Type:     schema.TypeString,
				Required: true,
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

func resourceCloudAccountOciTempDataCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandCloudAccountOciTempDataRequest(d)
	log.Printf("[INFO] Creating oci Cloud Account with request %+v\n", req)

	resp, _, err := d9Client.cloudaccountOci.CreateTempData(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created oci Temp Data information. ID: %v\n", resp.ID)

	d.SetId(resp.ID)
	_ = d.Set("name", resp.Name)
	_ = d.Set("creation_date", resp.CreationDate)
	_ = d.Set("tenancy_id", resp.TenancyId)
	_ = d.Set("home_region", resp.HomeRegion)
	_ = d.Set("organizational_unit_id", resp.OrganizationalUnitID)
	_ = d.Set("organizational_unit_path", resp.OrganizationalUnitPath)
	_ = d.Set("organizational_unit_name", resp.OrganizationalUnitName)
	_ = d.Set("vendor", resp.Vendor)

	if err := d.Set("credentials", flattenOciCredentialsResponse(resp.Credentials)); err != nil {
		return err
	}

	return nil
}

func resourceCloudAccountOciTempDataRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCloudAccountOciTempDataDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCloudAccountOciTempDataUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func expandCloudAccountOciTempDataRequest(d *schema.ResourceData) oci.CloudAccountRequestTempData {
	req := oci.CloudAccountRequestTempData{
		Name:                            d.Get("name").(string),
		TenancyId:                       d.Get("tenancy_id").(string),
		HomeRegion:                      d.Get("home_region").(string),
		TenantAdministratorEmailAddress: d.Get("tenant_administrator_email_address").(string),
	}
	return req
}

func flattenOciCredentialsResponse(CredentialsResponse oci.CloudAccountCredentialsResponse) map[string]interface{} {
	m := map[string]interface{}{
		"user":        CredentialsResponse.User,
		"fingerprint": CredentialsResponse.Fingerprint,
		"public_key":  CredentialsResponse.PublicKey,
	}

	return m
}
