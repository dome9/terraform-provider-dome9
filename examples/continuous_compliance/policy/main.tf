resource "dome9_continuous_compliance_policy" "test_policy" {
  cloud_account_id    = "CLOUD_ACCOUNT_ID"
  external_account_id = "EXTERNAL_ACCOUNT_ID"
  bundle_id           = 00000
  cloud_account_type  = "Azure" // options: ["Azure", "Aws", "Google", "Kubernetes"]
  notification_ids    = ["NOTIFICATION_IDS"]
}
