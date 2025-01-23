package dome9

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func expandNotificationIDs(d *schema.ResourceData, key string) []string {
	// Get the set from the resource data
	notificationsIDsData := d.Get(key).(*schema.Set).List()

	// Convert the set to a slice of strings
	notificationIDsList := make([]string, len(notificationsIDsData))
	for i, notificationID := range notificationsIDsData {
		notificationIDsList[i] = notificationID.(string)
	}

	return notificationIDsList
}
