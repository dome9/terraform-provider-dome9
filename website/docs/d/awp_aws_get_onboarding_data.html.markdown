---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_awp_aws_get_onboarding_data"
sidebar_current: "docs-datasource-dome9-awp-aws-get-onboarding-data"
description: |-
  Get information about AWS AWP onboarding data in Dome9.
---

# Data Source: dome9_awp_aws_get_onboarding_data

Use this data source to get information about AWS AWP onboarding data in Dome9.

## Example Usage

```hcl
data "dome9_awp_aws_get_onboarding_data" "test" {
  cloud_account_id = "d9-aws-cloud-account-id"
}

```

## Argument Reference

The following arguments supported:

* `cloud_account_id` - (Required) The Dome9 id for the onboarded AWS account, 
  * it can be the dome9 cloudguard account id or the external aws account id.

## Attributes Reference

In addition to all arguments above, the following attributes exported:

* `stage` - The stage of the AWP AWS onboarding process(i.e "prod-us").
* `region` - The region of the AWP AWS onboarding process.
* `cloud_guard_backend_account_id` - The CloudGuard AWS backend account ID.
* `agentless_bucket_name` - The name of the agentless s3 bucket.
* `remote_functions_prefix_key` - The prefix key for remote functions. 
* `remote_snapshots_utils_function_name` - The name of the remote snapshots utility function. 
* `remote_snapshots_utils_function_run_time` - The runtime of the remote snapshots utility function.
* `remote_snapshots_utils_function_time_out` - The timeout for the remote snapshots utility function.
* `awp_client_side_security_group_name` - The name of the AWP client-side security group.
* `cross_account_role_external_id` - The external ID for the cross-account role.
* `remote_snapshots_utils_function_s3_pre_signed_url` - The pre-signed URL for the remote snapshots utility function.