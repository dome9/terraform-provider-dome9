resource "dome9_attach_iam_safe" "dome9_attach_iam_safe_re" {
  aws_cloud_account_id = "00000000-0000-0000-0000-000000000000"
  aws_group_arn        = "AWS_GROUP_ARN"
  aws_policy_arn       = "AWS_POLICY_ARN"
}

data "dome9_attach_iam_safe" "dome9_attach_iam_safe_ds" {
  id = "00000000-0000-0000-0000-000000000000"
}

output "aws_cloud_account_id" {
  value = "${data.dome9_attach_iam_safe.dome9_attach_iam_safe_ds.aws_cloud_account_id}"
}

output "aws_group_arn" {
  value = "${data.dome9_attach_iam_safe.dome9_attach_iam_safe_ds.aws_group_arn}"
}

output "aws_policy_arn" {
  value = "${data.dome9_attach_iam_safe.dome9_attach_iam_safe_ds.aws_policy_arn}"
}

output "mode" {
  value = "${data.dome9_attach_iam_safe.dome9_attach_iam_safe_ds.mode}"
}
