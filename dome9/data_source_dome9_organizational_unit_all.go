package dome9

import (
	"github.com/dome9/dome9-sdk-go/services/organizationalunits"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceOrganizationalUnitAll() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOrganizationalUnitAllRead,

		Schema: map[string]*schema.Schema{
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: OrganizationalUnitSchema,
				},
			},
		},
	}
}

func dataSourceOrganizationalUnitAllRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	log.Printf("[INFO] Getting all data for Organizational Units \n")

	resp, _, err := d9Client.organizationalUnit.GetAll()
	d.SetId("all_organizational_units")
	if err != nil {
		return err
	}

	if err := d.Set("items", flattenOrganizationalUnitItems(*resp)); err != nil {
		return err
	}
	return nil
}

func flattenOrganizationalUnitItems(resp []organizationalunits.OUResponse) []interface{} {
	var ouListItems []interface{}

	for _, ou := range resp {
		flatItem := map[string]interface{}{
			"id":                                                ou.Item.ID,
			"name":                                              ou.Item.Name,
			"parent_id":                                         ou.ParentID,
			"account_id":                                        ou.Item.AccountID,
			"path":                                              ou.Item.Path,
			"created":                                           ou.Item.Created.Format("2006-01-02 15:04:05"),
			"updated":                                           ou.Item.Updated.Format("2006-01-02 15:04:05"),
			"aws_cloud_accounts_count":                          ou.Item.AwsCloudAcountsCount,
			"aws_aggregate_cloud_accounts_count":                ou.Item.AwsAggregatedCloudAcountsCount,
			"azure_cloud_accounts_count":                        ou.Item.AzureCloudAccountsCount,
			"azure_aggregate_cloud_accounts_count":              ou.Item.AzureAggregateCloudAccountsCount,
			"google_cloud_accounts_count":                       ou.Item.GoogleCloudAccountsCount,
			"google_aggregate_cloud_accounts_count":             ou.Item.GoogleAggregateCloudAccountsCount,
			"oci_cloud_accounts_count":                          ou.Item.OciCloudAccountsCount,
			"oci_aggregate_cloud_accounts_count":                ou.Item.OciAggregateCloudAccountsCount,
			"k8s_cloud_accounts_count":                          ou.Item.K8sCloudAccountsCount,
			"k8s_aggregate_cloud_accounts_count":                ou.Item.K8sAggregateCloudAccountsCount,
			"shift_left_cloud_accounts_count":                   ou.Item.ShiftLeftCloudAccountsCount,
			"shift_left_aggregate_cloud_accounts_count":         ou.Item.ShiftLeftAggregateCloudAccountsCount,
			"alibaba_cloud_accounts_count":                      ou.Item.AlibabaCloudAccountsCount,
			"alibaba_aggregate_cloud_accounts_count":            ou.Item.AlibabaAggregateCloudAccountsCount,
			"container_registry_cloud_accounts_count":           ou.Item.ContainerRegistryAccountsCount,
			"container_registry_aggregate_cloud_accounts_count": ou.Item.ContainerRegistryAggregateCloudAccountsCount,
			"sub_organizational_units_count":                    ou.Item.SubOrganizationalUnitsCount,
			"is_root":                                           ou.Item.IsRoot,
			"is_parent_root":                                    ou.Item.IsParentRoot,
			"path_str":                                          ou.Item.PathStr,
		}
		ouListItems = append(ouListItems, flatItem)

		if len(ou.Children) > 0 {
			childItems := flattenOrganizationalUnitItems(ou.Children)
			ouListItems = append(ouListItems, childItems...)
		}
	}
	return ouListItems
}
