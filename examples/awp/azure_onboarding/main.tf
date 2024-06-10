terraform {
	required_providers {
		dome9 = {
		source = "dome9/dome9"
		version = "1.29.7"
		}
		azurerm = {
		source = "hashicorp/azurerm"
		version = "3.99.0"
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


# Onboarding Azure Account to CloudGuard Dome9 Account
# This resource is optional and can be ignored and you need to pass CloudGuard account id to the module directly at the parameter awp_cloud_account_id.
# to know how to get the credentials for the onboarding process, please refer to the following link:
## https://sc1.checkpoint.com/documents/CloudGuard_Dome9/Documentation/Assets/Azure/Azure.htm
resource "dome9_cloudaccount_azure" "azure_ca" {
	client_id       = "CLIENT_ID"
	client_password = "CLIENT_PASSWORD"
	name            = "sandbox"
	operation_mode  = "Read"
	subscription_id = "SUBSCRIPTION_ID"
	tenant_id       = "TENANT_ID"
}

# There is a need to use this terraform module [terraform-dome9-awp-azure] to create all the prerequisites for the onboarding process (All the needed Azure Resources)
# Example for the module use:
module "terraform-dome9-awp-azure" {
	source = "github.com/dome9/terraform-dome9-awp-azure"
	awp_cloud_account_id = "<CLOUDGUARD_ACCOUNT_ID> or <AZURE_SUBSCRIPTION_ID>"
	awp_scan_mode = "<SCAN_MODE>" # Valid Values = "inAccount", "saas", "inAccountHub" or "inAccountSub"
    # awp_centralized_cloud_account_id = "<CENTRALIZED_CLOUD_ACCOUNT_ID> OR <CENTRALIZED_SUBSCRIPTION_ID>"

	# Optional customizations:
	# e.g:
	# awp_is_scanned_hub        = false
    # management_group_id       = "management group id"


	# Optional account Settings
	# e.g:
	#   awp_account_settings_azure = {
	#     scan_machine_interval_in_hours = 24
    #     skip_function_apps_scan = false
	#     disabled_regions = ["eastus", "westus", ...] # List of regions to disable
	#     max_concurrent_scans_per_region = 20
	#     custom_tags = {
	#       tag1 = "value1"
	#       tag2 = "value2"
	#       tag3 = "value3"
	#       ...
	#     }
	# }
}