package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-dome9/dome9/common/providerconst"
)

func srlDescriptorSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringInSlice(providerconst.SRLTypes, true),
				},
				"main_id": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"region": {
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringInSlice(providerconst.AWSRegions, true),
				},
				"security_group_id": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"traffic": {
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringInSlice(providerconst.PermissionTrafficOptions, true),
				},
			},
		},
	}
}

func srlDescriptorDataSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"main_id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"region": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"security_group_id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"traffic": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}
