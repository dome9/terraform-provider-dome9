package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/k8s"
	"log"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudAccountK8S() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudAccountK8SCreate,
		Read:   resourceCloudAccountK8SRead,
		Update: resourceCloudAccountK8SUpdate,
		Delete: resourceCloudAccountK8SDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"creation_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vendor": {
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
			"cluster_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCloudAccountK8SCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := createK8SCloudAccountRequest(d)
	log.Printf("[INFO] Creating K8S Cloud Account with request\n%+v\n", req)
	resp, _, err := d9Client.cloudaccountK8S.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created K8S CloudAccount. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceCloudAccountK8SRead(d, meta)
}

func resourceCloudAccountK8SRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	resp, _, err := d9Client.cloudaccountK8S.Get(d.Id())

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() { // 404 response code
			log.Printf("[WARN] Removing K8S cloud account %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Reading K8S account response and settings states: %+v\n", resp)
	d.SetId(resp.ID)
	_ = d.Set("name", resp.Name)
	_ = d.Set("creation_date", resp.CreationDate.Format("2006-01-02 15:04:05"))
	_ = d.Set("vendor", resp.Vendor)
	_ = d.Set("organizational_unit_id", resp.OrganizationalUnitID)
	_ = d.Set("organizational_unit_path", resp.OrganizationalUnitPath)
	_ = d.Set("organizational_unit_name", resp.OrganizationalUnitName)
	_ = d.Set("cluster_version", resp.ClusterVersion)

	return nil
}

func resourceCloudAccountK8SDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting K8S Cloud Account ID: %v\n", d.Id())

	if _, err := d9Client.cloudaccountK8S.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceCloudAccountK8SUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Println("An update occurred for K8S account")

	if d.HasChange("name") {
		log.Println("The name has been changed")

		if _, _, err := d9Client.cloudaccountK8S.UpdateName(d.Id(), k8s.CloudAccountUpdateNameRequest{
			Name: d.Get("name").(string),
		}); err != nil {
			return err
		}
	}

	if d.HasChange("organizational_unit_id") {
		log.Println("The Organizational Unit ID has been changed")

		if _, _, err := d9Client.cloudaccountK8S.UpdateOrganizationalID(d.Id(), k8s.CloudAccountUpdateOrganizationalIDRequest{
			OrganizationalUnitId: d.Get("organizational_unit_id").(string),
		}); err != nil {
			return err
		}
	}

	return resourceCloudAccountK8SRead(d, meta)
}

func createK8SCloudAccountRequest(d *schema.ResourceData) k8s.CloudAccountRequest {
	return k8s.CloudAccountRequest{
		Name:                 d.Get("name").(string),
		OrganizationalUnitID: d.Get("organizational_unit_id").(string),
	}
}
