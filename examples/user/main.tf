resource "dome9_users" "users_sg" {
  email = "EMAIL"
  first_name = "FIRST_NAME"
  last_name = "LAST_NAME"
  is_sso_enabled = false
}

data "dome9_users" "users_ds" {
  id = "ID"
}

output "user_email" {
  value = "${data.dome9_users.users_ds.email}"
}

output "user_is_sso_enabled" {
  value = "${data.dome9_users.users_ds.is_sso_enabled}"
}