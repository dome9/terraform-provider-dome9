package dome9

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/dome9/dome9-sdk-go/services/iplist"
	"github.com/dome9/terraform-provider-dome9/dome9/common/structservers"
)

func resourceIpList() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpListCreate,
		Read:   resourceIpListRead,
		Update: resourceIpListUpdate,
		Delete: resourceIpListDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"items": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"comment": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func expandingIpList(d *schema.ResourceData) *iplist.IpList {
	// Mandatory field
	ipList := iplist.IpList{
		Name: d.Get("name").(string),
	}
	// Optional fields
	if r, ok := d.GetOk("description"); ok {
		description := r.(string)
		ipList.Description = description
	}
	if itemsInterface, ok := d.GetOk("items"); ok {
		items := itemsInterface.([]interface{})
		log.Printf("[INFO] ------items from schema: %+v ---------\n", items)

		for _, item := range items {
			ipItem := item.(map[string]interface{})
			ip := ipItem["ip"].(string)
			comment := ipItem["comment"].(string)

			ipList.Items = append(ipList.Items, struct {
				Ip      string
				Comment string
			}{
				Ip:      ip,
				Comment: comment,
			},
			)

		}
		log.Printf("[INFO] ------iip list: %+v ---------\n", ipList)
	}

	return &ipList
}

func resourceIpListCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	ipList := expandingIpList(d)

	log.Printf("[INFO] Creating dome9 IP with name %s\n", ipList.Name)

	ipList, _, err := client.iplist.Create(ipList)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(ipList.Id, 10))
	return resourceIpListRead(d, meta)
}

func resourceIpListRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	ipList, _, err := client.iplist.Get(id)
	if err != nil {
		return nil
	}

	_ = d.Set("name", ipList.Name)
	_ = d.Set("description", ipList.Description)
	_ = d.Set("items", structservers.FlattenIpListItems(ipList))

	return nil
}

func resourceIpListUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	ipList := expandingIpList(d)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	ipList.Id = id
	log.Printf("[INFO] Updating IP list with id %d\n", id)

	_, err = client.iplist.Update(id, ipList)
	if err != nil {
		return err
	}

	return resourceIpListRead(d, meta)
}

func resourceIpListDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting IP list with id %v\n", id)

	if _, err := client.iplist.Delete(id); err != nil {
		return err
	}

	return nil
}
