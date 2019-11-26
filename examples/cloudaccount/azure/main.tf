resource "dome9_cloudaccount_azure" "azure_ca" {
  credentials = {
    client_id       = "CLIENT_ID"
    client_password = "CLIENT_PASSWORD"
  }

  name            = "sandbox"
  operation_mode  = "Read"
  subscription_id = "SUBSCRIPTION_ID"
  tenant_id       = "TENANT_ID"
}

data "dome9_cloudaccount_azure" "azure_ds" {
  id = "D9_CLOUD_ACCOUNT_ID"
}

output "get_name" {
  value = "${data.dome9_cloudaccount_azure.azure_ds.name}"
}

output "get_subscription_id" {
  value = "${data.dome9_cloudaccount_azure.azure_ds.subscription_id}"
}

output "get_tenant_id" {
  value = "${data.dome9_cloudaccount_azure.azure_ds.tenant_id}"
}

output "get_operation_mode" {
  value = "${data.dome9_cloudaccount_azure.azure_ds.operation_mode}"
}

output "get_creation_date" {
  value = "${data.dome9_cloudaccount_azure.azure_ds.creation_date}"
}

output "get_organizational_unit_id" {
  value = "${data.dome9_cloudaccount_azure.azure_ds.organizational_unit_id}"
}

output "get_organizational_unit_path" {
  value = "${data.dome9_cloudaccount_azure.azure_ds.organizational_unit_path}"
}

output "get_organizational_unit_name" {
  value = "${data.dome9_cloudaccount_azure.azure_ds.organizational_unit_name}"
}

output "get_vendor" {
  value = "${data.dome9_cloudaccount_azure.azure_ds.vendor}"
}
