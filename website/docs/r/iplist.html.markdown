---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_iplist"
sidebar_current: "docs-resource-dome9-iplist"
description: |-
  Create IP lists in Dome9
---

# dome9_iplist

This resource is used  to create and manage IP lists in Dome9. IP lists are groups of IP addresses (typically in customer cloud environments), on which common actions are applied. For example, a Security Group could be applied to a list, instead of applying it to each IP address in the list individually.

## Example Usage

Basic usage:

```hcl
resource "dome9_iplist" "iplist" {
  name        = "NAME"
  description = "DESCRIPTION"

  items = [
    {
      ip      = "1.1.1.1"
      comment = "COMMENT1"
    },
    {
      ip      = "2.2.2.2"
      comment = "COMMENT2"
    },
  ]
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the IP list in Dome9
* `description` - (Optional) A description of the list (what it represents); defaults to empty string
* `items` - (Optional) the individual IP addresses for the list; defaults to empty list

### Items 

The `items` supports the following arguments:
    
* `ip` - (Optional) IP address
* `comment` - (Optional) Comment

## Attributes Reference

* `id` - IP list Id

## Import

IP list can be imported; use `<IP LIST ID>` as the import ID. 

For example:

```shell
terraform import dome9_iplist.test 00000
```
