package structservers

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Used for strings list from schema to struct
func FlattenStringList(d *schema.ResourceData, key string) []string {
	sl := make([]string, 0)
	interfaceList := d.Get(key).([]interface{})
	for _, v := range interfaceList {
		sl = append(sl, v.(string))
	}

	return sl
}
