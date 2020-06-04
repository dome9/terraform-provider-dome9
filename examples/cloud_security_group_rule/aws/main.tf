resource "dome9_cloud_security_group_rule" "aws_sg_rule" {
  dome9_security_group_id = "dome9_security_group_id"

  services {
    inbound {
      name = "FIRST_INBOUND_SERVICE_NAME"
      description = "DESCRIPTION"
      protocol_type = "PROTOCOL_TYPE"
      port = "PORT"
      open_for_all = false
      scope {
        type = "TYPE"
        data = {
          cidr = "CIDR"
          note = "NOTE"
        }
      }
    }

    inbound {
      name = "SECOND_INBOUND_SERVICE_NAME"
      description = "DESCRIPTION"
      protocol_type = "PROTOCOL_TYPE"
      port = "PORT"
      open_for_all = false
      scope {
        type = "TYPE"
        data = {
          cidr = "CIDR"
          note = "NOTE"
        }
      }
    }
  }
}

data "dome9_cloud_security_group_rule" "aws_sg_rule" {
  id = "dome9_security_group_id"
}
