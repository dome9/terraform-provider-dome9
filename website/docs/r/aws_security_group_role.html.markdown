---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloud_security_group_role"
sidebar_current: "docs-resource-dome9-aws-security-group-role"
description: |-
  Bound input and output services to AWS Security Group in Dome9
---

# dome9_cloud_security_group_role

This resource has methods to add and manage input and output services to Security Groups in a cloud account that is managed by Dome9.

## Example Usage

Basic usage:

```hcl
resource "dome9_cloud_security_group_role" "aws_sg_role" {
  dome9_security_group_id    = "dome9_security_group_id"
  
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
}

```

## Argument Reference

The following arguments are supported:

* `dome9_security_group_id` - (Required) Dome9 security group id.
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
