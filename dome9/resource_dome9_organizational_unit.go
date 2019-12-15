package dome9

import (
	"log"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/organizationalunits"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceOrganizationalUnit() *schema.Resource {
	return &schema.Resource{
		Create: resourceOrganizationalUnitCreate,
		Read:   resourceOrganizationalUnitRead,
		Update: resourceOrganizationalUnitUpdate,
		Delete: resourceOrganizationalUnitDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_cloud_accounts_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"azure_cloud_accounts_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"google_cloud_accounts_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"aws_aggregate_cloud_accounts_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"azure_aggregate_cloud_accounts_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"google_aggregate_cloud_accounts_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sub_organizational_units_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_root": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_parent_root": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"path_str": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceOrganizationalUnitCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	req := expandOrganizationalUnitRequest(d)
	log.Printf("[INFO] Creating Organizational Unit with request\n%+v\n", req)
	resp, _, err := d9Client.organizationalUnit.Create(&req)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Created Organizational Unit. ID: %v\n", resp.Item.ID)
	d.SetId(resp.Item.ID)

	return resourceOrganizationalUnitRead(d, meta)
}

func resourceOrganizationalUnitRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	resp, _, err := d9Client.organizationalUnit.Get(d.Id())

	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing Organizational Unit %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Reading Organizational Unit response and settings states: %+v\n", resp)
	d.SetId(resp.Item.ID)
	_ = d.Set("name", resp.Item.Name)
	_ = d.Set("parent_id", resp.Item.ParentID)
	_ = d.Set("account_id", resp.Item.AccountID)
	_ = d.Set("path", resp.Item.Path)
	_ = d.Set("created", resp.Item.Created.Format("2006-01-02 15:04:05"))
	_ = d.Set("updated", resp.Item.Updated.Format("2006-01-02 15:04:05"))
	_ = d.Set("aws_cloud_accounts_count", resp.Item.AwsCloudAcountsCount)
	_ = d.Set("aws_aggregate_cloud_accounts_count", resp.Item.AwsAggregatedCloudAcountsCount)
	_ = d.Set("azure_cloud_accounts_count", resp.Item.AzureCloudAccountsCount)
	_ = d.Set("azure_aggregate_cloud_accounts_count", resp.Item.AzureAggregateCloudAccountsCount)
	_ = d.Set("google_cloud_accounts_count", resp.Item.GoogleCloudAccountsCount)
	_ = d.Set("google_aggregate_cloud_accounts_count", resp.Item.GoogleAggregateCloudAccountsCount)
	_ = d.Set("sub_organizational_units_count", resp.Item.SubOrganizationalUnitsCount)
	_ = d.Set("is_root", resp.Item.IsRoot)
	_ = d.Set("is_parent_root", resp.Item.IsParentRoot)
	_ = d.Set("path_str", resp.Item.PathStr)

	return nil
}

func resourceOrganizationalUnitDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Printf("[INFO] Deleting Organizational Unit ID: %v\n", d.Id())

	if _, err := d9Client.organizationalUnit.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceOrganizationalUnitUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	log.Println("An updated occurred")

	if d.HasChange("name") || d.HasChange("parent_id") {
		log.Println("The name or parent ID has been changed")

		if _, err := d9Client.organizationalUnit.Update(d.Id(), &organizationalunits.OURequest{
			Name:     d.Get("name").(string),
			ParentID: d.Get("parent_id").(string),
		}); err != nil {
			return err
		}
	}

	return nil
}

func expandOrganizationalUnitRequest(d *schema.ResourceData) organizationalunits.OURequest {
	return organizationalunits.OURequest{
		Name:     d.Get("name").(string),
		ParentID: d.Get("parent_id").(string),
	}
}
