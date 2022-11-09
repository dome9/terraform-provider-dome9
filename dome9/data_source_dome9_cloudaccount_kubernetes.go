package dome9

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/testing/variable"
)

func dataSourceCloudAccountKubernetes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKubernetesRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
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
				Computed: true,
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
			"runtime_protection": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"admission_control": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"create_default_policy": {
							Type:     schema.TypeBool,
							Required: false,
							Default:  true,
						},
					},
				},
			},
			"image_assurance": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"threat_intelligence": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceKubernetesRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for cloud account %s with id %s\n", variable.CloudAccountKubernetesVendor, id)

	resp, _, err := d9Client.cloudaccountKubernetes.Get(id)
	if err != nil {
		return err
	}

	d.SetId(resp.ID)
	_ = d.Set("name", resp.Name)
	_ = d.Set("creation_date", resp.CreationDate.Format("2006-01-02 15:04:05"))
	_ = d.Set("vendor", resp.Vendor)
	_ = d.Set("organizational_unit_id", resp.OrganizationalUnitID)
	_ = d.Set("organizational_unit_path", resp.OrganizationalUnitPath)
	_ = d.Set("organizational_unit_name", resp.OrganizationalUnitName)
	_ = d.Set("cluster_version", resp.ClusterVersion)
	_ = d.Set("runtime_protection", expandRuntimeProtectionConfig(resp))
	_ = d.Set("admission_control", expandAdmissionControlConfig(resp))
	_ = d.Set("image_assurance", expandImageAssuranceConfig(resp))
	_ = d.Set("threat_intelligence", expandThreatIntelligenceConfig(resp))
	return nil
}
