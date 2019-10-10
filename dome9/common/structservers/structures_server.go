package structservers

import (
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"
	"github.com/dome9/dome9-sdk-go/services/iplist"
)

// cloud account AWS

func FlattenNetSec(resp *aws.CloudAccountResponse) []interface{} {
	m := map[string]interface{}{
		"regions": FlattenRegions(resp),
	}

	return []interface{}{m}
}

func FlattenRegions(resp *aws.CloudAccountResponse) []interface{} {
	netSecRegions := make([]interface{}, len(resp.NetSec.Regions))
	for i, val := range resp.NetSec.Regions {
		netSecRegions[i] = map[string]interface{}{
			"region":             val.Region,
			"name":               val.Name,
			"hidden":             val.Hidden,
			"new_group_behavior": val.NewGroupBehavior,
		}
	}

	return netSecRegions
}

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
