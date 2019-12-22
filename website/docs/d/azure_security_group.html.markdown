---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_azure_security_group"
sidebar_current: "docs-datasource-dome9-azure-security-group"
description: |-
  Get information about Azure Security Groups in Dome9
---

# Data Source: dome9_azure_security_group

Use this data source to get information about an Azure Security Group onboarded to Dome9.

## Example Usage

Basic usage:

```hcl
data "dome9_azure_security_group" "azure_sg_ds" {
  id = "SECURITY_GROUP_ID"
}

```

## Argument Reference
In addition to all arguments above, the following attributes are exported:

* `dome9_security_group_name` - (Required) Name of the Security Group.
* `region` - (Required) Security group region.
* `resource_group` - (Required) Azure resource group name.
* `dome9_cloud_account_id` - (Required) Cloud account id in Dome9.
* `description` - (Optional) Security group description.
* `is_tamper_protected` - (Optional) Is Security Group tamper protected?
* `tags` - (Optional) Security Group tags.
* `inbound` - (Optional) Security Group services.
* `outbound` - (Optional) Security Group services.
