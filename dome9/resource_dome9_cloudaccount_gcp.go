package dome9

import (
	"log"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/gcp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
)

func resourceCloudAccountGCP() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudAccountGCPCreate,
		Read:   resourceCloudAccountGCPRead,
		Update: resourceCloudAccountGCPUpdate,
		Delete: resourceCloudAccountGCPDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private_key_id": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"private_key": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"client_email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_x509_cert_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vendor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gsuite_user": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
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

func resourceCloudAccountGCPCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandCloudAccountGCPRequest(d)
	log.Printf("[INFO] Creating GCP Cloud Account with request %+v\n", req)
	resp, _, err := d9Client.cloudaccountGCP.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created GCP CloudAccount. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceCloudAccountGCPRead(d, meta)
}

func resourceCloudAccountGCPRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	getCloudAccountQueryParams := cloudaccounts.QueryParameters{ID: d.Id()}
	resp, _, err := d9Client.cloudaccountGCP.Get(&getCloudAccountQueryParams)

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing GCP cloud account %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Reading account response and settings states: %+v\n", resp)
	d.SetId(resp.ID)
	_ = d.Set("name", resp.Name)
	_ = d.Set("project_id", resp.ProjectID)
	// Converting the timestamp to string in the format yyyy-MM-dd HH:mm:ss
	_ = d.Set("creation_date", resp.CreationDate.Format("2006-01-02 15:04:05"))
	_ = d.Set("organizational_unit_id", resp.OrganizationalUnitID)
	_ = d.Set("organizational_unit_path", resp.OrganizationalUnitPath)
	_ = d.Set("organizational_unit_name", resp.OrganizationalUnitName)
	_ = d.Set("gsuite_user", resp.GSuite.GSuiteUser)
	_ = d.Set("domain_name", resp.GSuite.DomainName)
	_ = d.Set("vendor", resp.Vendor)

	return nil
}

func resourceCloudAccountGCPDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting GCP Cloud Account ID: %v\n", d.Id())

	if _, err := d9Client.cloudaccountGCP.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceCloudAccountGCPUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Println("An updated occurred")

	if d.HasChange("name") {
		log.Println("The name has been changed")

		if resp, _, err := d9Client.cloudaccountGCP.UpdateName(d.Id(), gcp.CloudAccountUpdateNameRequest{
			Name: d.Get("name").(string),
		}); err != nil {
			return err
		} else {
			log.Printf("resourceCloudAccountGCPUpdate response is: %+v\n", resp)
		}
	}

	if d.HasChange("organizational_unit_id") {
		log.Println("The organizational unit id has been changed")

		if resp, _, err := d9Client.cloudaccountGCP.UpdateOrganizationalID(d.Id(), gcp.CloudAccountUpdateOrganizationalIDRequest{
			OrganizationalUnitID: d.Get("organizational_unit_id").(string),
		}); err != nil {
			return err
		} else {
			log.Printf("resourceCloudAccountGCPUpdate response is: %+v\n", resp)
		}
	}

	if d.HasChange("gsuite_user") || d.HasChange("domain_name") {
		log.Println("The gsuite user or domain name has been changed")

		if resp, _, err := d9Client.cloudaccountGCP.UpdateAccountGSuite(d.Id(), gcp.GSuite{
			GSuiteUser: d.Get("gsuite_user").(string),
			DomainName: d.Get("domain_name").(string),
		}); err != nil {
			return err
		} else {
			log.Printf("resourceCloudAccountGCPUpdate response is: %+v\n", resp)
		}
	}

	if credentialsHasChange(d) {
		log.Println("The service account credentials user or domain name has been changed")

		if resp, _, err := d9Client.cloudaccountGCP.UpdateCredentials(d.Id(), gcp.CloudAccountUpdateCredentialsRequest{
			Name:                      d.Get("name").(string),
			ServiceAccountCredentials: expandServiceAccountCredentials(d),
		}); err != nil {
			return err
		} else {
			log.Printf("resourceCloudAccountGCPUpdate response is: %+v\n", resp)
		}
	}

	return nil
}

func credentialsHasChange(d *schema.ResourceData) bool {
	return d.HasChange("project_id") || d.HasChange("private_key_id") || d.HasChange("private_key") || d.HasChange("client_email") || d.HasChange("client_id") || d.HasChange("client_x509_cert_url")
}

func expandCloudAccountGCPRequest(d *schema.ResourceData) gcp.CloudAccountRequest {
	req := gcp.CloudAccountRequest{
		Name:                      d.Get("name").(string),
		ServiceAccountCredentials: expandServiceAccountCredentials(d),
		GsuiteUser:                d.Get("gsuite_user").(string),
		DomainName:                d.Get("domain_name").(string),
		OrganizationalUnitID:      d.Get("organizational_unit_id").(string),
	}

	return req
}

func expandServiceAccountCredentials(d *schema.ResourceData) gcp.ServiceAccountCredentials {
	return gcp.ServiceAccountCredentials{
		Type:                    providerconst.GCPCloudAccountType,
		ProjectID:               d.Get("project_id").(string),
		PrivateKeyID:            d.Get("private_key_id").(string),
		PrivateKey:              d.Get("private_key").(string),
		ClientEmail:             d.Get("client_email").(string),
		ClientID:                d.Get("client_id").(string),
		AuthURI:                 providerconst.GCPCloudAccountAuthUri,
		TokenURI:                providerconst.GCPCloudAccountTokenUri,
		AuthProviderX509CertURL: providerconst.GCPCloudAccountAuthProviderX509CertUrl,
		ClientX509CertURL:       d.Get("client_x509_cert_url").(string),
	}
}
