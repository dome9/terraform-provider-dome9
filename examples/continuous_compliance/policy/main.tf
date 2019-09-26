provider "dome9" {
  dome9_access_id  = "--"
  dome9_secret_key = "--"
}

resource "dome9_continuouscompliance_policy" "test_policy" {
  cloud_account_id    = "CLOUD_ACCOUNT_ID"
  external_account_id = "EXTERNAL_ACCOUNT_ID"
  bundle_id           = 00000
  cloud_account_type  = "Azure" // options: ["Azure", "Aws", "Google", "Kubernetes"]
  notification_ids    = ["NOTIFICATION_IDS"]
}
