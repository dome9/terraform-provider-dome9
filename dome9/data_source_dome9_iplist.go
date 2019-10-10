package dome9

import (
	"strconv"

	"github.com/dome9/terraform-provider-dome9/dome9/common/structservers"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIpList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIpListRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},

			// Complex computed value
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"comment": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIpListRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	id, err := strconv.ParseInt(d.Get("id").(string), 10, 64)
	if err != nil {
		return err
	}

	ipList, _, err := client.iplist.Get(id)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(id, 10))
	_ = d.Set("name", ipList.Name)
	_ = d.Set("description", ipList.Description)
	_ = d.Set("items", structservers.FlattenIpListItems(ipList))

	return nil
}
