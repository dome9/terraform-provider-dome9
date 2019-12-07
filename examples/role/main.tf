resource "dome9_role" "role_rs" {
  name        = "ROLE_NAME"
  description = "ROLE_DESC"

  permissions {
    access               = ["string"]
    manage               = ["string"]
    rulesets             = ["string"]
    notifications        = ["string"]
    policies             = ["string"]
    alert_actions        = ["string"]
    create               = ["string"]
    view                 = ["string"]
    on_boarding          = ["string"]
    cross_account_access = ["string"]
  }
}

data "dome9_role" "role_ds" {
  id = 00000
}

output "getId" {
  value = "${data.dome9_role.data.id}"
}

output "getName" {
  value = "${data.dome9_role.data.name}"
}

output "getDescriptione" {
  value = "${data.dome9_role.data.description}"
}
