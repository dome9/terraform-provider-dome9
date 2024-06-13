terraform {
	required_providers {
		dome9 = {
			source = "dome9/dome9"
			version = ">=1.29.7"
		}
		aws = {
			source  = "hashicorp/aws"
			version = ">= 3.0"
		}
	}
}

# The Dome9 provider is used to interact with the resources supported by Dome9.
# The provider needs to be configured with the proper credentials before it can be used.
# Use the dome9_access_id and dome9_secret_key attributes of the provider to provide the Dome9 access key and secret key.
# The base_url attribute is used to specify the base URL of the Dome9 API.
# The Dome9 provider supports several options for providing these credentials. The following example demonstrates the use of static credentials:
#you can read the Dome9 provider documentation to understand the full set of options available for providing credentials.
#https://registry.terraform.io/providers/dome9/dome9/latest/docs#authentication
provider "dome9" {
	dome9_access_id     = "DOME9_ACCESS_ID"
	dome9_secret_key    = "DOME9_SECRET_KEY"
	base_url            = "https://api.dome9.com/v2/"
}

# AWS Provider Configurations
# The AWS provider is used to interact with the resources supported by AWS.
# The provider needs to be configured with the proper credentials before it can be used.
# Use the access_key, secret_key, and token attributes of the provider to provide the credentials.
# also you can use the shared_credentials_file attribute to provide the path to the shared credentials file.
# The AWS provider supports several options for providing these credentials. The following example demonstrates the use of static credentials:
#you can read the AWS provider documentation to understand the full set of options available for providing credentials.
#https://registry.terraform.io/providers/hashicorp/aws/latest/docs#authentication-and-configuration
provider "aws" {
	region     = "AWS_REGION"
	access_key = "AWS_ACCESS_KEY"
	secret_key = "AWS_SECRET_KEY"
	token      = "AWS_SESSION_TOKEN"
}

# Onboarding AWS Account to CloudGuard Dome9 Account
# This resource is optional and can be ignored and you need to pass CloudGuard account id to the module directly at the parameter awp_cloud_account_id.
# to know how to get the credentials for the onboarding process, please refer to the following link:
## https://sc1.checkpoint.com/documents/CloudGuard_Dome9/Documentation/Assets/AWS/OnboardAWS.htm
resource "dome9_cloudaccount_aws" "aws_onboarding_account_test" {
	name  = "aws_onboarding_account_test"
	credentials  {
		arn    = "CloudGuard Connect Role ARN"
		secret = "CloudGuard Connect Role Secret"
		type   = "RoleBased"
	}
	net_sec {
		regions {
			new_group_behavior = "ReadOnly"
			region             = "us_west_2"
		}
	}
}

# There is a need to use this terraform module [terraform-dome9-awp-aws] to create all the prerequisites for the onboarding process (All the needed AWS Resources)
# Example for the module use:
module "terraform-dome9-awp-aws" {
	source = "github.com/dome9/terraform-dome9-awp-aws"
	awp_cloud_account_id = "<CLOUDGUARD_ACCOUNT_ID> or <AWS_ACCOUNT_ID>"
	awp_scan_mode = "<SCAN_MODE>" # Valid Values = "inAccount" or "saas"

	# Optional customizations:
	# e.g:
	# awp_cross_account_role_name = "<CROSS_ACCOUNT_ROLE_NAME>"
	# awp_cross_account_role_external_id = "<CROSS_ACCOUNT_ROLE_EXTERNAL_ID>"

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