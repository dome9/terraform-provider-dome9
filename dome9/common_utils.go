package dome9

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func expandNotificationIDs(d *schema.ResourceData, key string) []string {
	notificationsIDsData := schema.NewSet(schema.HashString, d.Get(key).([]interface{})).List()
	notificationIDsList := make([]string, len(notificationsIDsData))
	for i, notificationID := range notificationsIDsData {
		notificationIDsList[i] = notificationID.(string)
	}
	return notificationIDsList
}
