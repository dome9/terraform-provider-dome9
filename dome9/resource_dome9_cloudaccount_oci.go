package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/oci"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceCloudAccountOciTempData() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudAccountOciCreate,
		Read:   resourceCloudAccountOciRead,
		Update: resourceCloudAccountOciUpdate,
		Delete: resourceCloudAccountOciDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenancyId": {
				Type:     schema.TypeString,
				Required: true,
			},
			"homeRegion": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
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
	log.Printf("[INFO] Creating oci Cloud Account with request %+v\n", req)

	resp, _, err := d9Client.cloudaccountOci.CreateTempData(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created oci Temp Data information. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceCloudAccountOciRead(d, meta)
}

func resourceCloudAccountOciRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCloudAccountOciDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCloudAccountOciUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func expandCloudAccountOciRequest(d *schema.ResourceData) oci.CloudAccountRequestTempData {
	req := oci.CloudAccountRequestTempData{
		Name:       d.Get("name").(string),
		TenancyId:  d.Get("tenancy_id").(string),
		HomeRegion: d.Get("home_region").(string),
		UserName:   d.Get("user_name").(string),
		GroupName:  d.Get("group_name").(string),
		PolicyName: d.Get("policy_name").(string),
	}
	return req
}
