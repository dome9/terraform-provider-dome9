resource "dome9_cloudaccount_kubernetes" "kubernetes_ca" {
  name                 = "sandbox"
}

data "dome9_cloudaccount_kubernetes" "kubernetes_ds" {
  id = "${dome9_cloudaccount_kubernetes.kubernetes_ca.id}"
}

output "get_name" {
  value = "${data.dome9_cloudaccount_kubernetes.kubernetes_ds.name}"
}

output "get_creation_date" {
  value = "${data.dome9_cloudaccount_kubernetes.kubernetes_ds.creation_date}"
}

output "get_organizational_unit_name" {
  value = "${data.dome9_cloudaccount_kubernetes.kubernetes_ds.organizational_unit_name}"
}

output "get_image_assurance_state" {
  value = "${data.dome9_cloudaccount_kubernetes.kubernetes_ds.image_assurance.0.enabled}"
}

output "get_admission_control_state" {
  value = "${data.dome9_cloudaccount_kubernetes.kubernetes_ds.admission_control.0.enabled}"
}

output "get_runtime_protection_state" {
  value = "${data.dome9_cloudaccount_kubernetes.kubernetes_ds.runtime_protection.0.enabled}"
}

output "get_flow_logs_state" {
  value = "${data.dome9_cloudaccount_kubernetes.kubernetes_ds.flow_logs.0.enabled}"
}