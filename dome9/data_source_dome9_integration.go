package dome9

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIntegration() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIntegrationRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configuration": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIntegrationRead(d *schema.ResourceData, m interface{}) error {
	d.SetId(d.Get("id").(string))
	return resourceIntegrationRead(d, m)
}
