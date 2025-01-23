package dome9

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func expandNotificationIDs(d *schema.ResourceData, key string) []string {
	rawData := d.Get(key).([]interface{}) // Take the list as-is
	uniqueSet := make(map[string]struct{})

	// Convert the list to a set (map) to ensure uniqueness
	for _, id := range rawData {
		uniqueSet[id.(string)] = struct{}{}
	}

	// Convert the set back to a slice
	notificationIDsList := make([]string, 0, len(uniqueSet))
	for id := range uniqueSet {
		notificationIDsList = append(notificationIDsList, id)
	}

	return notificationIDsList
}
