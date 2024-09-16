resource "dome9_role" "role_rs" {
  name        = "ROLE_NAME"
  description = "ROLE_DESC"
  access {
    type              = "AWS"
    main_id           = "MAIN_ID"
    region            = "us_east_1"
    security_group_id = "SECURITY_GROUP_ID"
    traffic           = "All Traffic"
  }
  access {
    type    = "OrganizationalUnit"
    main_id = "00000000-0000-0000-0000-000000000000"
  }

  manage {
      type = "CodeSecurityResources"
      main_id = "Member"
    }

  view {
      type = "CloudGuardResources"
  }

  permit_notifications = false
  permit_rulesets      = false
  permit_policies      = false
  permit_alert_actions = false
  permit_on_boarding   = false
  create               = []
  cross_account_access = []
}


data "dome9_role" "data" {
  id = "${dome9_role.role_rs.id}"
}

output "getId" {
  value = "${data.dome9_role.data.id}"
}

output "getDescription" {
  value = "${data.dome9_role.data.description}"
}

output "getItems" {
  value = "${data.dome9_role.data.access}"
}
