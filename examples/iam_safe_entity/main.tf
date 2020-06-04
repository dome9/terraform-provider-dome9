resource "dome9_iam_safe_entity" "dome9_iam_safe_entity_re" {
  protection_mode           = "ProtectWithElevation"
  entity_type               = "User"
  entity_name               = "ENTITY_NAME"
  aws_cloud_account_id      = "00000000-0000-0000-0000-000000000000"
  dome9_users_id_to_protect = ["000000", "111111"]
}
