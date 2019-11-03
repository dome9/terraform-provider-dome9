---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_continuous_compliance_notification"
sidebar_current: "docs-datasource-dome9-continuous-compliance-notification"
description: |-
  Get information about a Dome9 continuous compliance notification policy.
---

# Data Source: dome9_continuous_compliance_notification

Use this data source to get information about a Dome9 continuous compliance notification policy.

## Example Usage

```hcl
data "dome9_continuous_compliance_notification" "test" {
  id = "${%s.id}"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The id for the continuous compliance notification policy in Dome9. 

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Notification policy name.
* `description` - Description of the notification.
* `alerts_console` - Include in the alerts console.
* `scheduled_report` - Scheduled report details.
* `change_detection` - Change detection options.
* `gcp_security_command_center_integration` - GCP security command center details
