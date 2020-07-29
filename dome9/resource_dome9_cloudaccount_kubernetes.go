package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/k8s"
	"log"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudAccountKubernetes() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudAccountKubernetesCreate,
		Read:   resourceCloudAccountKubernetesRead,
		Update: resourceCloudAccountKubernetesUpdate,
		Delete: resourceCloudAccountKubernetesDelete,
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

func resourceCloudAccountKubernetesCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := createKubernetesCloudAccountRequest(d)
	log.Printf("[INFO] Creating Kubernetes Cloud Account with request\n%+v\n", req)
	resp, _, err := d9Client.cloudaccountKubernetes.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created Kubernetes CloudAccount. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceCloudAccountKubernetesRead(d, meta)
}

func resourceCloudAccountKubernetesRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	resp, _, err := d9Client.cloudaccountKubernetes.Get(d.Id())

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() { // 404 response code
			log.Printf("[WARN] Removing Kubernetes cloud account %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Reading Kubernetes account response and settings states: %+v\n", resp)
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

func resourceCloudAccountKubernetesDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting Kubernetes Cloud Account ID: %v\n", d.Id())

	if _, err := d9Client.cloudaccountKubernetes.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceCloudAccountKubernetesUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Println("An update occurred for Kubernetes account")

	if d.HasChange("name") {
		log.Println("The name has been changed")

		if _, _, err := d9Client.cloudaccountKubernetes.UpdateName(d.Id(), k8s.CloudAccountUpdateNameRequest{
			Name: d.Get("name").(string),
		}); err != nil {
			return err
		}
	}

	if d.HasChange("organizational_unit_id") {
		log.Println("The Organizational Unit ID has been changed")

		if _, _, err := d9Client.cloudaccountKubernetes.UpdateOrganizationalID(d.Id(), k8s.CloudAccountUpdateOrganizationalIDRequest{
			OrganizationalUnitId: d.Get("organizational_unit_id").(string),
		}); err != nil {
			return err
		}
	}

	return resourceCloudAccountKubernetesRead(d, meta)
}

func createKubernetesCloudAccountRequest(d *schema.ResourceData) k8s.CloudAccountRequest {
	return k8s.CloudAccountRequest{
		Name:                 d.Get("name").(string),
		OrganizationalUnitID: d.Get("organizational_unit_id").(string),
	}
}
