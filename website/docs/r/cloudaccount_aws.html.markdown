---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_AWS"
sidebar_current: "docs-resource-dome9-cloudaccount-AWS"
description: |-
  Onboard AWS cloud account
---

# dome9_cloudaccount_AWS

This resource is used to onboard AWS cloud accounts to Dome9. This is the first and pre-requisite step in order to apply  Dome9 features, such as compliance testing, on the account.

## Example Usage

Basic usage:

```hcl
resource "dome9_cloudaccount_AWS" "test" {
  name  = "ACCOUNT NAME"
 
  credentials  {
    ARN    = "ARN"
    secret = "SECRET"
    type   = "RoleBased"
  }

  organizational_unit_id = "ORGANIZATIONAL UNIT ID"

  net_sec {
    regions {
      new_group_behavior = "ReadOnly"
      region             = "us_east_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "us_west_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "eu_west_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "ap_southeast_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "ap_northeast_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "us_west_2"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "sa_east_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "ap_southeast_2"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "eu_central_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "ap_northeast_2"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "ap_south_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "us_east_2"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "ca_central_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "eu_west_2"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "eu_west_3"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "eu_north_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "ap_east_1"
    }
    regions {
      new_group_behavior = "ReadOnly"
      region             = "me_south_1"
    }
	regions {
      new_group_behavior = "ReadOnly"
      region             = "af_south_1"
    }
	regions {
      new_group_behavior = "ReadOnly"
      region             = "eu_south_1"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of AWS account in Dome9
* `credentials` - (Required) The information needed for Dome9 System in order to connect to the AWS cloud account
* `organizational_unit_id` - (Optional) The Organizational Unit that this cloud account will be attached to

### Credentials

`credentials` has the following arguments:

* `arn` - (Required) AWS Role ARN (to be assumed by Dome9)
* `secret` - (Required) The AWS role External ID (Dome9  will have to use this secret in order to assume the role)
* `type` - (Required) The cloud account onboarding method. Set to "RoleBased".

### Network security configuration

`net_sec` has the these arguments:

* `Regions` - (Required) list of the supported regions, and their configuration:
    * `new_group_behavior` - (Required) The network security configuration. Select "ReadOnly", "FullManage", or "Reset".
    * `region` - (Required) AWS region, in AWS format (e.g., "us-east-1")

## Attributes Reference

* `id` - The id of the account in Dome9.
* `vendor` - The cloud provider ("AWS").
* `external_account_number` - The AWS account number.
* `is_fetching_suspended` - Fetching suspending status.
* `creation_date` - Date the account was onboarded to Dome9.
* `full_protection` - The protection mode for existing security groups in the account.
* `allow_read_only` - The AWS cloud account operation mode. true for "Full-Manage", false for "Readonly".
* `net_sec` - The network security configuration for the AWS cloud account. If not given, sets to default value.
* `IAM_safe` - IAM safe entity details
    * `AWS_group_ARN` - AWS group ARN  
    * `AWS_policy_ARN` - AWS policy ARN  
    * `mode` - Mode  
    * `restricted_IAM_entities` - Restricted IAM safe entities, which have the following fields:  
		* `roles_ARNs` - Restricted IAM safe entities roles ARNs
		* `users_ARNs` - Restricted IAM safe entities users ARNs

## Import

AWS cloud account can be imported; use `<AWS CLOUD ACCOUNT ID>` as the import ID. 

For example:

```shell
terraform import dome9_cloudaccount_AWS.test 00000000-0000-0000-0000-000000000000
```
