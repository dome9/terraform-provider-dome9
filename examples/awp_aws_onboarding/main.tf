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

# There is a need to use this terraform module [terraform-dome9-awp-aws] to create all the prerequisites for the onboarding process (All the needed AWS Resources)
# Example for the module use:
module "terraform-dome9-awp-aws" {
	source = "github.com/dome9/terraform-dome9-awp-aws"
	awp_cloud_account_id = "<CLOUDGUARD_ACCOUNT_ID>"
	awp_scan_mode = "<SCAN_MODE>" # Valid Values = "inAccount" or "saas"
	# Optional customizations:
	# awp_cross_account_role_name = "CheckPoint-AWP-CrossAccount-Role"
	# awp_cross_account_role_external_id = "AWP_Fake@ExternalID123"

	# Optional account Settings
	# e.g:
	#   awp_account_settings_aws = {
	#     scan_machine_interval_in_hours = 24
	#     disabled_regions = ["ap-northeast-1", "ap-northeast-2", ...]
	#     max_concurrence_scans_per_region = 20
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
# The max_concurrence_scans_per_region attribute is used to specify the max concurrence scans per region of the agentless account settings of the Dome9 AWP AWS Onboarding.
# The custom_tags attribute is used to specify the custom tags of the agentless account settings of the Dome9 AWP AWS Onboarding.
resource "dome9_awp_aws_onboarding" "awp_aws_onboarding_test" {
	cloudguard_account_id = "dome9_cloudaccount_aws.aws_onboarding_account_test.id | <CLOUDGUARD_ACCOUNT_ID> | <EXTERNAL_AWS_ACCOUNT_NUMBER>"
	cross_account_role_name = "<AWP Cross account role name>"
	cross_account_role_external_id = "<AWP Cross account role external id>"
	scan_mode = "<SCAN_MODE>" # Valid Values = "inAccount" or "saas"
	agentless_account_settings {
		disabled_regions = ["us-east-1", "us-west-1", "ap-northeast-1", "ap-southeast-2"]
		scan_machine_interval_in_hours = 24
		max_concurrence_scans_per_region = 20
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
	depends_on = [
		dome9_awp_aws_onboarding.awp_aws_onboarding_test
	]
}