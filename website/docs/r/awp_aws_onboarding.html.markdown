---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_awp_aws_onboarding"
sidebar_current: "docs-resource-dome9-awp-aws-onboarding"
description: |-
  Creates an AWP AWS Onboarding in Dome9
---

# dome9_awp_aws_onboarding

This resource is used to create and modify AWP AWS Onboarding in CloudGuard Dome9.

## Example Usage

Basic usage:

```hcl
resource "dome9_awp_aws_onboarding" "test_awp_aws_onboarding" {
  cloudguard_account_id = "CloudGuard Account ID or External AWS Account ID"
  cross_account_role_name = "Cross Account Role Name"
  cross_account_role_external_id = "Cross Account Role External ID"
  scan_mode = "inAccount"
  agentless_account_settings {
    disabled_regions = ["us-east-1", "us-west-1"]
    scan_machine_interval_in_hours = 24
    max_concurrence_scans_per_region = 6
    skip_function_apps_scan = true
    custom_tags = {
      tag1 = "value1"
      tag2 = "value2"
    }
  }
}

```

## Argument Reference

The following arguments are supported:

* `cloudguard_account_id` - (Required) The CloudGuard account id.
* `centralized_cloud_account_id` - (Optional) The centralized cloud account id.
* `cross_account_role_name` - (Required) The name of the cross account role.
* `cross_account_role_external_id` - (Required) The external id of the cross account role.
* `scan_mode` - (Required) The scan mode. Valid values are "inAccount", "saas", "inAccountHub", "inAccountSub".
* `agentless_account_settings` - (Optional) The agentless account settings.
  * `disabled_regions` - (Optional) The disabled regions. valid values are "us-east-1", "us-west-1", "us-west-2", "eu-west-1", "eu-central-1", "ap-northeast-1", "ap-southeast-1", "ap-southeast-2", "ap-northeast-2", "ap-south-1", "sa-east-1".
  * `scan_machine_interval_in_hours` - (Optional) The scan machine interval in hours
  * `max_concurrence_scans_per_region` - (Optional) The max concurrence scans per region
  * `skip_function_apps_scan` - (Optional) Whether to skip function apps scan. Default is false.
  * `custom_tags` - (Optional) The custom tags.
* `should_create_policy` - (Optional) Whether to create a policy. Default is true.
    
## Attributes Reference

* `missing_awp_private_network_regions` - The missing AWP private network regions.
* `account_issues` - The account issues.
* `cloud_account_id` - The cloud guard account id.
* `agentless_protection_enabled` - Whether agentless protection is enabled.
* `cloud_provider` - The cloud provider.
* `should_update` - Whether to update.
* `is_org_onboarding` - Whether is org onboarding.

## Import

The AWP AWS Onboarding can be imported; use <ONBOARDING ID> as the import ID.

For example:

```shell
terraform import dome9_awp_aws_onboarding.test_awp_aws_onboarding 00000000-0000-0000-0000-000000000000
```
