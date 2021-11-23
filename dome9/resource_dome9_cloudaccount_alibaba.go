package dome9

import (
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/alibaba"
)

func resourceCloudAccountAlibaba() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudAccountAlibabaCreate,
		Read:   resourceCloudAccountAlibabaRead,
		Update: resourceCloudAccountAlibabaUpdate,
		Delete: resourceCloudAccountAlibabaDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"access_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"access_secret": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"vendor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_id": {
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

func resourceCloudAccountAlibabaCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandCloudAccountAlibabaRequest(d)
	log.Printf("[INFO] Creating Alibaba Cloud Account with request %+v\n", req)

	resp, _, err := d9Client.cloudaccountAlibaba.Create(req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created Alibaba CloudAccount. ID: %v\n", resp.ID)
	d.SetId(resp.ID)

	return resourceCloudAccountAlibabaRead(d, meta)
}

func resourceCloudAccountAlibabaRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.cloudaccountAlibaba.Get(d.Id())

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
	_ = d.Set("account_id", resp.AccountId)
	_ = d.Set("vendor", resp.Vendor)
	_ = d.Set("creation_date", resp.CreationDate)
	_ = d.Set("organizational_unit_id", resp.OrganizationalUnitID)
	_ = d.Set("organizational_unit_path", resp.OrganizationalUnitPath)
	_ = d.Set("organizational_unit_name", resp.OrganizationalUnitName)

	return nil
}

func resourceCloudAccountAlibabaDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting Alibaba Cloud Account ID: %v\n", d.Id())
	if _, err := d9Client.cloudaccountAlibaba.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceCloudAccountAlibabaUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Println("An updated occurred")

	if d.HasChange("name") {
		log.Println("The name has been changed")

		if resp, _, err := d9Client.cloudaccountAlibaba.UpdateName(d.Id(), alibaba.CloudAccountUpdateNameRequest{
			Name: d.Get("name").(string),
		}); err != nil {
			return err
		} else {
			log.Printf("resourceCloudAccountAlibabaUpdate response is: %+v\n", resp)
		}
	}

	if d.HasChange("access_key") || d.HasChange("access_secret") {
		log.Println("The credentials has been changed")

		if resp, _, err := d9Client.cloudaccountAlibaba.UpdateCredentials(d.Id(), alibaba.CloudAccountUpdateCredentialsRequest{
			ApplicationID:  d.Get("access_key").(string),
			ApplicationKey: d.Get("access_secret").(string),
		}); err != nil {
			return err
		} else {
			log.Printf("resourceCloudAccountAlibabaUpdate response is: %+v\n", resp)
		}
	}

	if d.HasChange("organizational_unit_id") {
		log.Println("The organizational unit id has been changed")

		if resp, _, err := d9Client.cloudaccountAlibaba.UpdateOrganizationalID(d.Id(), alibaba.CloudAccountUpdateOrganizationalIDRequest{
			OrganizationalUnitID: d.Get("organizational_unit_id").(string),
		}); err != nil {
			return err
		} else {
			log.Printf("resourceCloudAccountAlibabaUpdate response is: %+v\n", resp)
		}
	}

	return nil
}

func expandCloudAccountAlibabaRequest(d *schema.ResourceData) alibaba.CloudAccountRequest {
	req := alibaba.CloudAccountRequest{
		Name:           d.Get("name").(string),
		AccountId:       d.Get("account_id").(string),
		Credentials: alibaba.CloudAccountCredentials{
			AccessKey:       d.Get("access_key").(string),
			AccessSecret: d.Get("access_secret").(string),
		},
		OrganizationalUnitID:   d.Get("organizational_unit_id").(string),
		OrganizationalUnitPath: d.Get("organizational_unit_path").(string),
		OrganizationalUnitName: d.Get("organizational_unit_name").(string),
	}

	if r, ok := d.GetOk("creation_date"); ok {
		formatTemplate := "2006-01-02 15:04:05"
		creationDateTime, _ := time.Parse(formatTemplate, r.(string))
		req.CreationDate = creationDateTime
	}
	return req
}
