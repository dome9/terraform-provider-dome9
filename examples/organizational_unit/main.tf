resource "dome9_organizational_unit" "test_ou" {
  name               = "some_organizational_unit"
  parent_id          = "00000000-0000-0000-0000-000000000000"
}

data "dome9_organizational_unit" "test_ou_ds" {
  id = "OU_ID"
}

output "get_organizational_unit_name" {
  value = "${data.dome9_organizational_unit.test_ou_ds.name}"
}

output "get_organizational_unit_parent_id" {
  value = "${data.dome9_organizational_unit.test_ou_ds.parent_id}"
}
