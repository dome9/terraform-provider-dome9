resource "dome9_continuous_compliance_policy" "test_policy" {
  target_id = "TARGET_ID"
  ruleset_id = 00000
  target_type = "Azure"
  // options: ["Aws", "Azure", "Gcp", "Kubernetes", "OrganizationalUnit"]
  notification_ids = [
    "NOTIFICATION_IDS"]
}
