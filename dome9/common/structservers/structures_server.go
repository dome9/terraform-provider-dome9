package structservers

import (
	"github.com/dome9/dome9-sdk-go/services/iplist"
)

func FlattenIpListItems(ipList *iplist.IpList) []interface{} {
	ipListItems := make([]interface{}, len(ipList.Items))
	for i, ipListItem := range ipList.Items {
		ipListItems[i] = map[string]interface{}{
			"ip":      ipListItem.Ip,
			"comment": ipListItem.Comment,
		}
	}

	return ipListItems
}
