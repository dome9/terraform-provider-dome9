# Configure the Dome9 Provider, works with US region by default
provider "dome9" {
  dome9_access_id     = "DOME9_API_ACCESS_ID"
  dome9_secret_key    = "DOME9_API_SECRET"
}

resource "dome9_admission_control_policy" "admission_control_policy_test" {
  target_id = "1212c5a5-ccc7-451a-8b3f-f19b5fc041f4"
  ruleset_id = -2001
  // options: ["Environment", "OrganizationalUnit"]
  target_type = "Environment"
  notification_ids = ["52008dbc-6fdc-45d5-b588-de2afe13ac5d"]
  admission_control_action = "Detection"
  admission_control_unscanned_action = "Detection"
}

data "dome9_admission_control_policy" "admission_control_policy_test" {
  id = dome9_admission_control_policy.admission_control_policy_test.id
}

output "getPolicyId" {
  value = dome9_admission_control_policy.admission_control_policy_test.id
}

output "getTargetId" {
  value = dome9_admission_control_policy.admission_control_policy_test.target_id
}

output "getAction" {
  value = dome9_admission_control_policy.admission_control_policy_test.action
}
