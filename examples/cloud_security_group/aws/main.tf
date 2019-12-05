resource "dome9_aws_security_group" "aws_sg" {
  dome9_security_group_name = "dome9_security_group_name"
  description               = "description"
  aws_region_id             = "aws_region_id"
  dome9_cloud_account_id    = "dome9_cloud_account_id"

  services {
    inbound {
      name          = "FIRST_INBOUND_SERVICE_NAME"
      description   = "DESCRIPTION"
      protocol_type = "PROTOCOL_TYPE"
      port          = "PORT"
      open_for_all  = false
      scope {
        type = "TYPE"
        data = {
          cidr = "CIDR"
          note = "NOTE"
        }
      }
    }
    inbound {
      name          = "SECOND_INBOUND_SERVICE_NAME"
      description   = "DESCRIPTION"
      protocol_type = "PROTOCOL_TYPE"
      port          = "PORT"
      open_for_all  = false
      scope {
        type = "TYPE"
        data = {
          cidr = "CIDR"
          note = "NOTE"
        }
      }
    }
    outbound {
      name          = "NAME"
      description   = "DESCRIPTION"
      protocol_type = "PROTOCOL_TYPE"
      port          = ""
      open_for_all  = true
    }
  }

  tags = {
    tag-key = "TAG-VALUE"
  }
}

data "dome9_aws_security_group" "aws_sg_ds" {
  id = "SECURITY_GROUP_ID"
}

output "get_security_group_name" {
  value = "${data.dome9_aws_security_group.aws_sg_ds.dome9_security_group_name}"
}

output "get_aws_region_id" {
  value = "${data.dome9_aws_security_group.aws_sg_ds.aws_region_id}"
}

output "description" {
  value = "${data.dome9_aws_security_group.aws_sg_ds.description}"
}

output "aws_region_id" {
  value = "${data.dome9_aws_security_group.aws_sg_ds.aws_region_id}"
}

output "services" {
  value = "${data.dome9_aws_security_group.aws_sg_ds.services}"
}
