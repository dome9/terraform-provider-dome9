package dome9

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/dome9/dome9-sdk-go/dome9/client"
	"github.com/dome9/dome9-sdk-go/services/iplist"
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

func resourceIpListCreate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	ipListRequest := expandIpList(d)
	log.Printf("[INFO] Creating dome9 IpList with request\n%+v\n", ipListRequest)

	ipList, _, err := d9Client.iplist.Create(&ipListRequest)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(ipList.Id, 10))
	return resourceIpListRead(d, meta)
}

func resourceIpListRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	ipList, _, err := d9Client.iplist.Get(id)
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing ip list %s from state because it no longer exists in Dome9", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	_ = d.Set("name", ipList.Name)
	_ = d.Set("description", ipList.Description)
	if err := d.Set("items", flattenIpListItems(ipList)); err != nil {
		return err
	}

	return nil
}

func resourceIpListUpdate(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	ipListRequest := expandIpList(d)
	ipListRequest.Id = id
	log.Printf("[INFO] Updating IpList with name %s\n", ipListRequest.Name)

	if _, err := d9Client.iplist.Update(id, &ipListRequest); err != nil {
		return err
	}

	return resourceIpListRead(d, meta)
}

func resourceIpListDelete(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting IP list with id %v\n", id)

	if _, err := d9Client.iplist.Delete(id); err != nil {
		return err
	}

	return nil
}

func expandIpList(d *schema.ResourceData) iplist.IpList {
	ipList := iplist.IpList{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Items:       expandIpListItems(d),
	}

	return ipList
}

func expandIpListItems(d *schema.ResourceData) []iplist.Item {
	var ipListItems []iplist.Item
	if itemsInterface, ok := d.GetOk("items"); ok {
		items := itemsInterface.([]interface{})
		ipListItems = make([]iplist.Item, len(items))
		for i, item := range items {
			ipItem := item.(map[string]interface{})
			ipListItems[i] = iplist.Item{
				Ip:      ipItem["ip"].(string),
				Comment: ipItem["comment"].(string),
			}
		}
	}

	return ipListItems
}

func flattenIpListItems(ipList *iplist.IpList) []interface{} {
	ipListItems := make([]interface{}, len(ipList.Items))
	for i, ipListItem := range ipList.Items {
		ipListItems[i] = map[string]interface{}{
			"ip":      ipListItem.Ip,
			"comment": ipListItem.Comment,
		}
	}

	return ipListItems
}
