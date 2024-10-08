package dome9

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var OrganizationalUnitSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"parent_id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"account_id": {
		Type:     schema.TypeInt,
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
	"oci_cloud_accounts_count": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"google_cloud_accounts_count": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"k8s_cloud_accounts_count": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"shift_left_cloud_accounts_count": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"alibaba_cloud_accounts_count": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"container_registry_cloud_accounts_count": {
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
	"oci_aggregate_cloud_accounts_count": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"google_aggregate_cloud_accounts_count": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"k8s_aggregate_cloud_accounts_count": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"shift_left_aggregate_cloud_accounts_count": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"alibaba_aggregate_cloud_accounts_count": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"container_registry_aggregate_cloud_accounts_count": {
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
}

func dataSourceOrganizationalUnit() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOrganizationalUnitRead,

		Schema: OrganizationalUnitSchema,
	}
}

func dataSourceOrganizationalUnitRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for Organizational Unit ID %s\n", id)

	resp, _, err := d9Client.organizationalUnit.Get(id)
	if err != nil {
		return err
	}

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
	_ = d.Set("oci_cloud_accounts_count", resp.Item.OciCloudAccountsCount)
	_ = d.Set("oci_aggregate_cloud_accounts_count", resp.Item.OciAggregateCloudAccountsCount)
	_ = d.Set("k8s_cloud_accounts_count", resp.Item.K8sCloudAccountsCount)
	_ = d.Set("k8s_aggregate_cloud_accounts_count", resp.Item.K8sAggregateCloudAccountsCount)
	_ = d.Set("shift_left_cloud_accounts_count", resp.Item.ShiftLeftCloudAccountsCount)
	_ = d.Set("shift_left_aggregate_cloud_accounts_count", resp.Item.ShiftLeftAggregateCloudAccountsCount)
	_ = d.Set("alibaba_cloud_accounts_count", resp.Item.AlibabaCloudAccountsCount)
	_ = d.Set("alibaba_aggregate_cloud_accounts_count", resp.Item.AlibabaAggregateCloudAccountsCount)
	_ = d.Set("container_registry_cloud_accounts_count", resp.Item.ContainerRegistryAccountsCount)
	_ = d.Set("container_registry_aggregate_cloud_accounts_count", resp.Item.ContainerRegistryAggregateCloudAccountsCount)
	_ = d.Set("sub_organizational_units_count", resp.Item.SubOrganizationalUnitsCount)
	_ = d.Set("is_root", resp.Item.IsRoot)
	_ = d.Set("is_parent_root", resp.Item.IsParentRoot)
	_ = d.Set("path_str", resp.Item.PathStr)

	return nil
}
