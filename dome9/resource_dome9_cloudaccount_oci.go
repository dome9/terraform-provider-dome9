package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/oci"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"

	"github.com/dome9/dome9-sdk-go/dome9/client"
)

func resourceCloudAccountOCI() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudAccountOciCreate,
		Read:   resourceCloudAccountOciRead,
		Update: resourceCloudAccountOciUpdate,
		Delete: resourceCloudAccountOciDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"tenancy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_ocid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"organizational_unit_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "00000000-0000-0000-0000-000000000000",
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

func resourceCloudAccountOciCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandCloudAccountOciRequest(d)
	log.Printf("[INFO] Creating Oci Cloud Account with request %+v\n", req)

	resp, _, err := d9Client.cloudaccountOci.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created Oci CloudAccount. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceCloudAccountOciRead(d, meta)
}

func resourceCloudAccountOciRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.cloudaccountOci.Get(d.Id())

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing Oci cloud account %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	d.SetId(resp.ID)
	_ = d.Set("name", resp.Name)
	_ = d.Set("creation_date", resp.CreationDate)
	_ = d.Set("tenancy_id", resp.TenancyId)
	_ = d.Set("home_region", resp.HomeRegion)
	_ = d.Set("organizational_unit_id", resp.OrganizationalUnitID)
	_ = d.Set("organizational_unit_path", resp.OrganizationalUnitPath)
	_ = d.Set("organizational_unit_name", resp.OrganizationalUnitName)
	_ = d.Set("vendor", resp.Vendor)

	if err := d.Set("credentials", flattenOciCredentials(resp.Credentials)); err != nil {
		return err
	}

	return nil
}

func resourceCloudAccountOciDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting Oci Cloud Account ID: %v\n", d.Id())
	if _, err := d9Client.cloudaccountOci.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceCloudAccountOciUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Println("An updated occurred")

	if d.HasChange("organizational_unit_id") {
		log.Println("The organizational unit id has been changed")

		if resp, _, err := d9Client.cloudaccountOci.UpdateOrganizationalID(d.Id(), oci.CloudAccountUpdateOrganizationalIDRequest{
			OrganizationalUnitID: d.Get("organizational_unit_id").(string),
		}); err != nil {
			return err
		} else {
			log.Printf("resourceCloudAccountOciUpdate response is: %+v\n", resp)
		}
	}

	return nil
}

func expandCloudAccountOciRequest(d *schema.ResourceData) oci.CloudAccountRequest {
	req := oci.CloudAccountRequest{
		UserOcid:             d.Get("user_ocid").(string),
		TenancyId:            d.Get("tenancy_id").(string),
		OrganizationalUnitID: d.Get("organizational_unit_id").(string),
	}
	return req
}
