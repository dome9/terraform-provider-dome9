package dome9

import (
	"strconv"

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
	d9Client := meta.(*Client)

	id, err := strconv.ParseInt(d.Get("id").(string), 10, 64)
	if err != nil {
		return err
	}

	ipList, _, err := d9Client.iplist.Get(id)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(id, 10))
	_ = d.Set("name", ipList.Name)
	_ = d.Set("description", ipList.Description)
	if err := d.Set("items", flattenIpListItems(ipList)); err != nil {
		return err
	}

	return nil
}
