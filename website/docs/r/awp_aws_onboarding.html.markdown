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
provider "dome9" {
  dome9_access_id     = "DOME9_ACCESS_ID"
  dome9_secret_key    = "DOME9_SECRET_KEY"
  base_url            = "https://api.dome9.com/v2/"
}
provider "aws" {
  region     = "AWS_REGION"
  access_key = "AWS_ACCESS_KEY"
  secret_key = "AWS_SECRET_KEY"
  token      = "AWS_SESSION_TOKEN"
}


# There is a need to use this terraform module [terraform-dome9-awp-aws] to create all the prerequisites for the onboarding process (All the needed AWS Resources)
# for further information please refer to the module documentation [terraform-dome9-awp-aws](https://registry.terraform.io/modules/dome9/awp-aws/dome9/latest)
# for more examples (simple and complete), you can visit this github examples [terraform-dome9-awp-aws](https://github.com/dome9/terraform-dome9-awp-aws/blob/master/examples)
# Example for the module use:
module "terraform-dome9-awp-aws" {
	source = "github.com/dome9/terraform-dome9-awp-aws"
	awp_cloud_account_id = "<CLOUDGUARD_ACCOUNT_ID> or <AWS_ACCOUNT_ID>"
	awp_scan_mode = "<SCAN_MODE>" # Possible values: "inAccount", "saas", "inAccountHub", "inAccountSub"

	# Optional customizations:
	# e.g:
	# awp_cross_account_role_name = "<CROSS_ACCOUNT_ROLE_NAME>"
	# awp_cross_account_role_external_id = "<CROSS_ACCOUNT_ROLE_EXTERNAL_ID>"
  # the following parameter is required for "InAccountSub" scan mode
  # awp_centralized_account_id = "<CENTRALIZED_ACCOUNT_ID> or <AWS_ACCOUNT_ID>" # centralized account-id where AWP scanner runs

	# Optional account Settings
	# e.g:
	#   awp_account_settings_aws = {
	#     scan_machine_interval_in_hours = 24
	#     disabled_regions = ["ap-northeast-1", "ap-northeast-2", ...] # List of regions to disable
	#     max_concurrent_scans_per_region = 20
	#     custom_tags = {
	#       tag1 = "value1"
	#       tag2 = "value2"
	#       tag3 = "value3"
	#       ...
	#     }
	# }
}

# The dome9_awp_aws_onboarding resource defines a Dome9 AWP AWS Onboarding.
# The Dome9 AWP AWS Onboarding resource allows you to onboard an AWS account to Dome9 AWP.
# The cloudguard_account_id attribute is used to specify the CloudGuard account id of the AWS account.
# The cross_account_role_name attribute is used to specify the name of the cross account role that is used to allow AWP to access the AWS account.
# The cross_account_role_external_id attribute is used to specify the external id of the cross account role that is used to allow AWP to access the AWS account.
# The scan_mode attribute is used to specify the scan mode of the Dome9 AWP AWS Onboarding. The valid values are "inAccount" and "saas".
# The agentless_account_settings attribute is used to specify the agentless account settings of the Dome9 AWP AWS Onboarding.
# The disabled_regions attribute is used to specify the disabled regions of the agentless account settings of the Dome9 AWP AWS Onboarding.
# The scan_machine_interval_in_hours attribute is used to specify the scan machine interval in hours of the agentless account settings of the Dome9 AWP AWS Onboarding.
# The max_concurrent_scans_per_region attribute is used to specify the max concurrent scans per region of the agentless account settings of the Dome9 AWP AWS Onboarding.
# The custom_tags attribute is used to specify the custom tags of the agentless account settings of the Dome9 AWP AWS Onboarding.
resource "dome9_awp_aws_onboarding" "awp_aws_onboarding_test" {
  cloudguard_account_id = "dome9_cloudaccount_aws.aws_onboarding_account_test.id | <CLOUDGUARD_ACCOUNT_ID> | <EXTERNAL_AWS_ACCOUNT_NUMBER>"
  cross_account_role_name = "<AWP Cross account role name>"
  cross_account_role_external_id = "<AWP Cross account role external id>"
  scan_mode = "<SCAN_MODE>" # Possible values: "inAccount", "saas", "inAccountHub", "inAccountSub"
	awp_centralized_account_id = "<CENTRALIZED_ACCOUNT_ID> or <AWS_ACCOUNT_ID>" # required for "InAccountSub" scan mode, it is the centralized account-id where AWP scanner runs

  # Optional account Settings (supported for 'inAccount' and 'saas' scan modes)
  # e.g:
  agentless_account_settings {
    disabled_regions = ["us-east-1", "us-west-1", "ap-northeast-1", "ap-southeast-2"]
    scan_machine_interval_in_hours = 24
    max_concurrent_scans_per_region = 20
    custom_tags = {
      tag1 = "value1"
      tag2 = "value2"
      tag3 = "value3"
    }
  }
}

# The dome9_awp_aws_onboarding data source allows you to get the onboarding data of an AWS account (Optional).
data "dome9_awp_aws_onboarding" "awp_aws_onboarding_test" {
  id = dome9_awp_aws_onboarding.awp_aws_onboarding_test.cloudguard_account_id
}
```

## Argument Reference

The following arguments are supported:

* `cloudguard_account_id` - (Required) The CloudGuard account id.
* `cross_account_role_name` - (Required) The name of the cross account role.
* `cross_account_role_external_id` - (Required) The external id of the cross account role.
* `scan_mode` - (Required) The scan mode. Valid values are "inAccount", "saas", "inAccountHub", "inAccountSub".
* `awp_centralized_account_id` - (Optional) The centralized cloud account id, required (and only relevant) for "inAccountSub" scan mode
* `agentless_account_settings` - (Optional) The agentless account settings.
  * `disabled_regions` - (Optional) The disabled regions. valid values are "af-south-1", "ap-south-1", "eu-north-1", "eu-west-3", "eu-south-1", "eu-west-2", "eu-west-1", "ap-northeast-3", "ap-northeast-2", "me-south-1", "ap-northeast-1", "me-central-1", "ca-central-1", "sa-east-1", "ap-east-1", "ap-southeast-1", "ap-southeast-2", "eu-central-1", "ap-southeast-3", "us-east-1", "us-east-2", "us-west-1", "us-west-2"
  * `scan_machine_interval_in_hours` - (Optional) The scan machine interval in hours
  * `max_concurrent_scans_per_region` - (Optional) The max concurrent scans per region
  * `in_account_scanner_vpc` - (Optional) The VPC mode. Valid values are "ManagedByAWP" or "ManagedByCustomer".
  * `custom_tags` - (Optional) The custom tags.
* `should_create_policy` - (Optional) Whether to create a policy. Default is true.
    
## Attributes Reference

* `missing_awp_private_network_regions` - The missing AWP private network regions.
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
