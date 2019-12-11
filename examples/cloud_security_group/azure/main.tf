resource "dome9_azure_security_group" "azure_sg" {
  dome9_security_group_name = "DOME9_SECURITY_GROUP_NAME"
  region                    = "REGION"
  resource_group            = "RESOURCE_GROUP"
  dome9_cloud_account_id    = "DOME9_CLOUD_ACCOUNT_ID"
  description               = "DESCRIPTION"
  is_tamper_protected       = false

  inbound {
    name               = "NAME"
    description        = "DESCRIPTION"
    priority           = 0
    access             = "ACCESS"
    protocol           = "PRIORITY"
    source_port_ranges = ["SOURCE_PORT_RANGES"]

    source_scopes {
      type = "TYPE"
      data = {
        //        data fields
      }
    }
    destination_port_ranges = ["20-90"]

    destination_scopes {
      type = "TYPE"
      data = {
        //          data fields
      }
    }
    is_default = false
  }
}

data "dome9_azure_security_group" "azure_sg_ds" {
  id = "ID"
}

output "security_group_name" {
  value = "${data.dome9_azure_security_group.azure_sg_ds.dome9_security_group_name}"
}

output "region" {
  value = "${data.dome9_azure_security_group.azure_sg_ds.region}"
}

output "resource_group" {
  value = "${data.dome9_azure_security_group.azure_sg_ds.resource_group}"
}

output "cloud_account_id" {
  value = "${data.dome9_azure_security_group.azure_sg_ds.dome9_cloud_account_id}"
}

output "description" {
  value = "${data.dome9_azure_security_group.azure_sg_ds.description}"
}

output "is_tamper_protected" {
  value = "${data.dome9_azure_security_group.azure_sg_ds.is_tamper_protected}"
}

output "inbound" {
  value = "${data.dome9_azure_security_group.azure_sg_ds.inbound}"
}
