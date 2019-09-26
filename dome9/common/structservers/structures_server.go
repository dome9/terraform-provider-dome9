package structservers

import (
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"
)

// cloud account AWS

func FlattenNetSec(resp *aws.CloudAccountResponse) []interface{} {
	m := map[string]interface{}{
		"regions": FlattenRegions(resp),
	}

	return []interface{}{m}
}

func FlattenRegions(resp *aws.CloudAccountResponse) []interface{} {
	// fixed size interface array
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
