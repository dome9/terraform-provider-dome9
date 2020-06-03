---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_aws_security_group"
sidebar_current: "docs-resource-dome9-aws-security-group"
description: |-
  Creates AWS Security Group in Dome9
---

# dome9_aws_security_group

This resource has methods to add and manage Security Groups in a cloud account that is managed by Dome9.

## Example Usage

Basic usage:

```hcl
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

```

Example for security group circular dependencies

```hcl
resource "dome9_aws_security_group" "aws_sg1" {
  dome9_security_group_name = "dome9_security_group_name"
  description               = "description"
  aws_region_id             = "aws_region_id"
  dome9_cloud_account_id    = "dome9_cloud_account_id"
}

resource "dome9_cloud_security_group_rule" "aws_sg1" {
  dome9_security_group_id = "${dome9_aws_security_group.aws_sg1.id}"
  services {
    outbound {
      name          = "HTTPS"
      description   = "HTTPS (TCP)"
      protocol_type = "TCP"
      port          = "8443"
      open_for_all  = false
      scope {
        type = "AWS"
        data = {
          extid = "${dome9_aws_security_group.aws_sg2.external_id}"
          note = "${dome9_aws_security_group.aws_sg2.external_id}"
        }
      }
    }
  }
}

##################

resource "dome9_aws_security_group" "aws_sg2" {
  dome9_security_group_name = "dome9_security_group_name"
  description               = "description"
  aws_region_id             = "aws_region_id"
  dome9_cloud_account_id    = "dome9_cloud_account_id"
}

resource "dome9_cloud_security_group_rule" "aws_sg2" {
  dome9_security_group_id = "${dome9_aws_security_group.aws_sg2.id}"
  services {
    outbound {
      name          = "HTTPS"
      description   = "HTTPS (TCP)"
      protocol_type = "TCP"
      port          = "8443"
      open_for_all  = false
      scope {
        type = "AWS"
        data = {
          extid = "${dome9_aws_security_group.aws_sg1.external_id}"
          note = "${dome9_aws_security_group.aws_sg1.external_id}"
        }
      }
    }
  }
}

```
## Argument Reference

The following arguments are supported:

* `dome9_security_group_name` - (Required) Name of the Security Group.
* `dome9_cloud_account_id` - (Required) Cloud account id in Dome9.
* `description` - (Optional) Security Group description.
* `aws_region_id` - (Optional) AWS region, in AWS format (e.g., "us-east-1"); default is us_east_1.
* `is_protected` - (Optional) Indicates the Security Group is in Protected mode.
    * Note: to set the protection mode, first create the Security Group, then update it with the desired protection mode value ('true' for Protected).
* `vpc_id` - (Optional) VPC id for VPC containing the Security Group.
* `vpc_name` - (Optional) Security Group VPC name.
* `tags` - (Optional) Security Group tags.
* `services` - (Optional) Security Group services.

### Security Group services

`services` has the these arguments:

* `inbound` - (Required) inbound service.
* `outbound` - (Required) outbound service. 

The configuration of inbound and outbound is:
   * `name` - (Required) Service name.
   * `description` - (Optional) Service description.
   * `protocol_type` - (Required) Service protocol type. Select from "ALL", "HOPOPT", "ICMP", "IGMP", "GGP", "IPV4", "ST", "TCP", "CBT", "EGP", "IGP", "BBN_RCC_MON", "NVP2", "PUP", "ARGUS", "EMCON", "XNET", "CHAOS", "UDP", "MUX", "DCN_MEAS", "HMP", "PRM", "XNS_IDP", "TRUNK1", "TRUNK2", "LEAF1", "LEAF2", "RDP", "IRTP", "ISO_TP4", "NETBLT", "MFE_NSP", "MERIT_INP", "DCCP", "ThreePC", "IDPR", "XTP", "DDP", "IDPR_CMTP", "TPplusplus", "IL", "IPV6", "SDRP", "IPV6_ROUTE", "IPV6_FRAG", "IDRP", "RSVP", "GRE", "DSR", "BNA", "ESP", "AH", "I_NLSP", "SWIPE", "NARP", "MOBILE", "TLSP", "SKIP", "ICMPV6", "IPV6_NONXT", "IPV6_OPTS", "CFTP", "SAT_EXPAK", "KRYPTOLAN", "RVD", "IPPC", "SAT_MON", "VISA", "IPCV", "CPNX", "CPHB", "WSN", "PVP", "BR_SAT_MON", "SUN_ND", "WB_MON", "WB_EXPAK", "ISO_IP", "VMTP", "SECURE_VMTP", "VINES", "TTP", "NSFNET_IGP", "DGP", "TCF", "EIGRP", "OSPFIGP", "SPRITE_RPC", "LARP", "MTP", "AX25", "IPIP", "MICP", "SCC_SP", "ETHERIP", "ENCAP", "GMTP", "IFMP", "PNNI", "PIM", "ARIS", "SCPS", "QNX", "AN", "IPCOMP", "SNP", "COMPAQ_PEER", "IPX_IN_IP", "VRRP", "PGM", "L2TP", "DDX", "IATP", "STP", "SRP", "UTI", "SMP", "SM", "PTP", "ISIS", "FIRE", "CRTP", "CRUDP", "SSCOPMCE", "IPLT", "SPS", "PIPE", "SCTP", "FC", "RSVP_E2E_IGNORE", "MOBILITY_HEADER", "UDPLITE", "MPLS_IN_IP", "MANET", "HIP", "SHIM6", "WESP" or "ROHC".
   * `port` - (Optional) Service type (port).
   * `open_for_all` - (Optional) Is open for all.
   * `scope` - (Optional) Service scope which has the following configuration:
      * `type` - (Required) scope type.
      * `data` - (Required) scope data.
        
## Attributes Reference

* `cloud_account_name` - AWS cloud account name.
* `external_id` - Security Group external id.

* Note: Just the following fields can be updated: services (inbound / outbound), tags and protection mode. 
## Import

The security group can be imported; use `<SESCURITY GROUP ID>` as the import ID. 

For example:

```shell
terraform import dome9_aws_security_group.test 00000000-0000-0000-0000-000000000000
```
