package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"

	"github.com/dome9/dome9-sdk-go/dome9/client"
)

func resourceAssessment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAssessmentCreate,
		Read:   resourceAssessmentRead,
		Update: resourceAssessmentUpdate,
		Delete: resourceAssessmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"credentials": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"access_secret": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
					},
				},
			},
			"vendor": {
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
		},
	}
}

func resourceAssessmentCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandCloudAccountAlibabaRequest(d)
	log.Printf("[INFO] Creating Alibaba Cloud Account with request %+v\n", req)

	resp, _, err := d9Client.cloudaccountAlibaba.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created Alibaba CloudAccount. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceAssessmentRead(d, meta)
}

func resourceAssessmentRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.assessment.RunBundle(d.Id())

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing Alibaba cloud account %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	d.SetId(resp.ID)
	_ = d.Set("name", resp.Name)
	_ = d.Set("alibaba_account_id", resp.AlibabaAccountId)
	_ = d.Set("vendor", resp.Vendor)
	_ = d.Set("creation_date", resp.CreationDate)
	_ = d.Set("organizational_unit_id", resp.OrganizationalUnitID)
	_ = d.Set("organizational_unit_path", resp.OrganizationalUnitPath)
	_ = d.Set("organizational_unit_name", resp.OrganizationalUnitName)

	return nil
}

func resourceAssessmentDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting Alibaba Cloud Account ID: %v\n", d.Id())
	if _, err := d9Client.cloudaccountAlibaba.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceAssessmentUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Println("An update can not be made to an assessment")
	return nil
}

