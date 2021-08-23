data "dome9_role" "role" {
  id = "ROLE_ID"
}

resource "dome9_service_account" "service_account_rs" {
  name        = "SERVICE_ACCOUNT_NAME"
  role_ids    = ["${data.dome9_role.role.id}"]
}

data "dome9_service_account" "service_account_ds" {
  id = "${dome9_service_account.service_account_rs.id}"
}

output "service_account_api_key" {
  value = "${data.dome9_service_account.service_account_ds.api_key_id}"
}