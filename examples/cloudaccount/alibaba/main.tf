resource "dome9_cloudaccount_alibaba" "alibaba_ca" {
  name        = "sandbox"
  credentials = {
    access_key    = "ACCESS_KEY"
    access_secret = "ACCESS_SECRET"
  }
}

data "dome9_cloudaccount_alibaba" "alibaba_ds" {
  id = "D9_CLOUD_ACCOUNT_ID"
}

output "get_name" {
  value = "${data.dome9_cloudaccount_alibaba.alibaba_ds.name}"
}

output "get_alibaba_account_id" {
  value = "${data.dome9_cloudaccount_alibaba.alibaba_ds.alibaba_account_id}"
}

output "get_creation_date" {
  value = "${data.dome9_cloudaccount_alibaba.alibaba_ds.creation_date}"
}

output "get_credentials" {
  value = "${data.dome9_cloudaccount_alibaba.alibaba_ds.credentials}"
}

output "get_organizational_unit_id" {
  value = "${data.dome9_cloudaccount_alibaba.alibaba_ds.organizational_unit_id}"
}

output "get_organizational_unit_path" {
  value = "${data.dome9_cloudaccount_alibaba.alibaba_ds.organizational_unit_path}"
}

output "get_organizational_unit_name" {
  value = "${data.dome9_cloudaccount_alibaba.alibaba_ds.organizational_unit_name}"
}

output "get_vendor" {
  value = "${data.dome9_cloudaccount_alibaba.alibaba_ds.vendor}"
}
