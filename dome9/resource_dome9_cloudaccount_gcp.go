package dome9

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/gcp"
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
			"service_account_credentials": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"project_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"private_key_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"private_key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"client_email": {
							Type:     schema.TypeString,
							Required: true,
						},
						"client_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"auth_uri": {
							Type:     schema.TypeString,
							Required: true,
						},
						"token_uri": {
							Type:     schema.TypeString,
							Required: true,
						},
						"auth_provider_x509_cert_url": {
							Type:     schema.TypeString,
							Required: true,
						},
						"client_x509_cert_url": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
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

func constructCloudAccountRequestGCP(d *schema.ResourceData) *gcp.CloudAccountRequest {
	req := gcp.CloudAccountRequest{
		Name: d.Get("name").(string),
	}

	if r, ok := d.GetOk("gsuite_user"); ok {
		req.GsuiteUser = r.(string)
	}
	if r, ok := d.GetOk("domain_name"); ok {
		req.DomainName = r.(string)
	}
	if r, ok := d.GetOk("service_account_credentials.type"); ok {
		req.ServiceAccountCredentials.Type = r.(string)
	}
	if r, ok := d.GetOk("service_account_credentials.project_id"); ok {
		req.ServiceAccountCredentials.ProjectID = r.(string)
	}
	if r, ok := d.GetOk("service_account_credentials.private_key_id"); ok {
		req.ServiceAccountCredentials.PrivateKeyID = r.(string)
	}
	if r, ok := d.GetOk("service_account_credentials.private_key"); ok {
		req.ServiceAccountCredentials.PrivateKey = r.(string)
	}
	if r, ok := d.GetOk("service_account_credentials.client_email"); ok {
		req.ServiceAccountCredentials.ClientEmail = r.(string)
	}
	if r, ok := d.GetOk("service_account_credentials.client_id"); ok {
		req.ServiceAccountCredentials.ClientID = r.(string)
	}
	if r, ok := d.GetOk("service_account_credentials.client_id"); ok {
		req.ServiceAccountCredentials.ClientID = r.(string)
	}
	if r, ok := d.GetOk("service_account_credentials.auth_uri"); ok {
		req.ServiceAccountCredentials.AuthURI = r.(string)
	}
	if r, ok := d.GetOk("service_account_credentials.auth_provider_x509_cert_url"); ok {
		req.ServiceAccountCredentials.AuthProviderX509CertURL = r.(string)
	}
	if r, ok := d.GetOk("service_account_credentials.client_x509_cert_url"); ok {
		req.ServiceAccountCredentials.ClientX509CertURL = r.(string)
	}

	log.Printf("[INFO] Creating GCP Cloud Account request: %+v\n", req)

	return &req
}

func resourceCloudAccountGCPCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	req := constructCloudAccountRequestGCP(d)
	log.Printf("[INFO] Creating GCP Cloud Account with request %+v\n", req)
	resp, _, err := client.cloudaccountGCP.Create(*req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created GCP CloudAccount. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceCloudAccountGCPRead(d, meta)
}

func resourceCloudAccountGCPRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	getCloudAccountQueryParams := cloudaccounts.QueryParameters{ID: d.Id()}
	resp, _, err := client.cloudaccountGCP.Get(&getCloudAccountQueryParams)
	if err != nil {
		return nil
	}

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
	client := meta.(*Client)
	log.Printf("[INFO] Deleting GCP Cloud Account ID: %v\n", d.Id())

	if _, err := client.cloudaccountGCP.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceCloudAccountGCPUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	log.Println("An updated occurred")

	if d.HasChange("name") {
		log.Println("The name has been changed")

		if resp, _, err := client.cloudaccountGCP.UpdateName(d.Id(), gcp.CloudAccountUpdateNameRequest{
			Name: d.Get("name").(string),
		}); err != nil {
			return err

		} else {
			log.Printf("resourceCloudAccountGCPUpdate response is: %+v\n", resp)
		}
	}

	if d.HasChange("organizational_unit_id") {
		log.Println("The organizational unit id has been changed")

		if resp, _, err := client.cloudaccountGCP.UpdateOrganizationalID(d.Id(), gcp.CloudAccountUpdateOrganizationalIDRequest{
			OrganizationalUnitID: d.Get("organizational_unit_id").(string),
		}); err != nil {
			return err

		} else {
			log.Printf("resourceCloudAccountGCPUpdate response is: %+v\n", resp)
		}
	}

	if d.HasChange("gsuite_user") || d.HasChange("domain_name") {
		log.Println("The gsuite user or domain name has been changed")

		if resp, _, err := client.cloudaccountGCP.UpdateAccountGSuite(d.Id(), gcp.CloudAccountUpdateGSuite{
			GsuiteUser: d.Get("gsuite_user").(string),
			DomainName: d.Get("domain_name").(string),
		}); err != nil {
			return err

		} else {
			log.Printf("resourceCloudAccountGCPUpdate response is: %+v\n", resp)
		}
	}

	if d.HasChange("service_account_credentials") {
		log.Println("The service account credentials user or domain name has been changed")

		if resp, _, err := client.cloudaccountGCP.UpdateCredentials(d.Id(), gcp.CloudAccountUpdateCredentialsRequest{
			Name: d.Get("name").(string),
			ServiceAccountCredentials: struct {
				Type                    string `json:"type,omitempty"`
				ProjectID               string `json:"project_id,omitempty"`
				PrivateKeyID            string `json:"private_key_id,omitempty"`
				PrivateKey              string `json:"private_key,omitempty"`
				ClientEmail             string `json:"client_email,omitempty"`
				ClientID                string `json:"client_id,omitempty"`
				AuthURI                 string `json:"auth_uri,omitempty"`
				TokenURI                string `json:"token_uri,omitempty"`
				AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url,omitempty"`
				ClientX509CertURL       string `json:"client_x509_cert_url,omitempty"`
			}{
				Type:                    d.Get("service_account_credentials.type").(string),
				ProjectID:               d.Get("service_account_credentials.project_id").(string),
				PrivateKeyID:            d.Get("service_account_credentials.private_key_id").(string),
				PrivateKey:              d.Get("service_account_credentials.private_key").(string),
				ClientEmail:             d.Get("service_account_credentials.client_email").(string),
				ClientID:                d.Get("service_account_credentials.client_id").(string),
				AuthURI:                 d.Get("service_account_credentials.auth_uri").(string),
				TokenURI:                d.Get("service_account_credentials.token_uri").(string),
				AuthProviderX509CertURL: d.Get("service_account_credentials.auth_provider_x509_cert_url").(string),
				ClientX509CertURL:       d.Get("service_account_credentials.client_x509_cert_url").(string),
			},
		}); err != nil {
			return err

		} else {
			log.Printf("resourceCloudAccountGCPUpdate response is: %+v\n", resp)
		}
	}

	return nil
}
