package dome9

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func ExpandNotificationIDs(d *schema.ResourceData, key string) []string {
	notificationsIDsData := d.Get(key).([]interface{})
	notificationIDsList := make([]string, len(notificationsIDsData))
	for i, notificationID := range notificationsIDsData {
		notificationIDsList[i] = notificationID.(string)
	}
	return notificationIDsList
}
