resource "dome9_cloudaccount_gcp" "gcp_ca" {
  name = "sandbox"

  service_account_credentials = {
    auth_provider_x509_cert_url = "https://www.googleapis.com/oauth2/v1/certs"
    auth_uri                    = "https://accounts.google.com/o/oauth2/auth"
    client_email                = "EMAIL@ADDRESS.COM"
    client_id                   = "CID"
    client_x509_cert_url        = "CERT URL"
    private_key                 = "KEY"
    private_key_id              = "PRIVATE"
    project_id                  = "ID"
    token_uri                   = "https://oauth2.googleapis.com/token"
    type                        = "service_account"
  }
}

data "dome9_cloudaccount_gcp" "gcp_ds" {
  id = "D9_CLOUD_ACOUNT_ID"
}

output "get_name" {
  value = "${data.dome9_cloudaccount_gcp.gcp_ds.name}"
}

output "get_project_id" {
  value = "${data.dome9_cloudaccount_gcp.gcp_ds.project_id}"
}

output "get_creation_date" {
  value = "${data.dome9_cloudaccount_gcp.gcp_ds.creation_date}"
}

output "get_organizational_unit_id" {
  value = "${data.dome9_cloudaccount_gcp.gcp_ds.organizational_unit_id}"
}

output "get_organizational_unit_path" {
  value = "${data.dome9_cloudaccount_gcp.gcp_ds.organizational_unit_path}"
}

output "get_organizational_unit_name" {
  value = "${data.dome9_cloudaccount_gcp.gcp_ds.organizational_unit_name}"
}

output "get_gsuite" {
  value = "${data.dome9_cloudaccount_gcp.gcp_ds.gsuite_user}"
}

output "get_domain_name" {
  value = "${data.dome9_cloudaccount_gcp.gcp_ds.domain_name}"
}

output "get_vendor" {
  value = "${data.dome9_cloudaccount_gcp.gcp_ds.vendor}"
}
