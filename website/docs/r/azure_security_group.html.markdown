---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_azure_security_group"
sidebar_current: "docs-resource-dome9-azure-security-group"
description: |-
  Creates azure security group in Dome9
---

# dome9_azure_security_group

The Azure Security Group resource has methods to add and manage Azure Security Group policies for Azure cloud accounts that are managed by Dome9.

## Example Usage

Basic usage:

```hcl
resource "dome9_azure_security_group" "azure_sg" {
  dome9_security_group_name = "dome9_security_group_name"
  region                    = "australiaeast"
  resource_group            = "resource_group"
  dome9_cloud_account_id    = "dome9_cloud_account_id"

  description         = "description"
  is_tamper_protected = false

  inbound {
    name               = "name"
    description        = "description"
    priority           = 1000
    access             = "Allow"
    protocol           = "TCP"
    source_port_ranges = ["*"]

    source_scopes {
      type = "Tag"
      data = {
        name = "VirtualNetwork"
      }
    }
    destination_port_ranges = ["20-90"]

    destination_scopes {
      type = "CIDR"
      data = {
        cidr = "0.0.0.0/0"
        note = "Any"
      }
    }
    is_default = false
  }
}

```

## Argument Reference

The following arguments are supported:

* `dome9_security_group_name` - (Required) Name of the security group.
* `region` - (Required) Region can be one of the following: `centralus`, `eastus`, `eastus2`, `usgovlowa`, `usgovvirginia`, `northcentralus`, `southcentralus`, `westus`, `westus2`, `westcentralus`, `northeurope`, `westeurope`, `eastasia`, `southeastasia`, `japaneast`, `japanwest`, `brazilsouth`, `australiaeast`, `australiasoutheast`, `centralindia`, `southindia`, `westindia`, `chinaeast`, `chinanorth`, `canadacentral`, `canadaeast`, `germanycentral`, `germanynortheast`, `koreacentral`, `uksouth`, `ukwest`, `koreasout`
* `resource_group` - (Required) Azure resource group name.
* `dome9_cloud_account_id` - (Required) Cloud account id in Dome9.
* `description` - (Optional) Security group description.
* `is_tamper_protected` - (Optional) Is security group tamper protected.
* `tags` - (Optional) Security group tags list of `key`, `value`:
    * `key` - (Required) Tag key. 
    * `value` - (Required) Tag value.
* `inbound` - (Optional) Security group services.
* `outbound` - (Optional) Security group services.

The configuration of inbound and outbound is:
   * `name` - (Required) Service name.
   * `description` - (Optional) Service description.
   * `priority` - (Required) Service priority (a number between 100 and 4096)
   * `access` - (Optional) 	Service access (Allow / Deny).
   * `protocol` - (Required) Service protocol (UDP / TCP / ANY).
   * `source_port_ranges` - (Required) Source port ranges.
   * `destination_port_ranges` - (Required) Destination port ranges.
   * `source_scopes` - (Required) List of source scopes for the service (CIDR / IP List / Tag):
      * `type` - (Required) scope type.
      * `data` - (Required) scope data.
   * `destination_scopes` - (Required) List of destination scopes for the service (CIDR / IP List / Tag).
      * `type` - (Required) scope type.
      * `data` - (Required) scope data.
   * `is_default` - Gets or sets the default security rules of network security group.
        
## Attributes Reference

* `external_security_group_id` - Azure external security group id.
* `cloud_account_name` - Azure cloud account name.
* `last_updated_by_dome9` - Last updated by dome9.
 
## Import

The security group can be imported; use `<SESCURITY GROUP ID>` as the import ID. 

For example:

```shell
terraform import dome9_azure_security_group.test 00000000-0000-0000-0000-000000000000
```
