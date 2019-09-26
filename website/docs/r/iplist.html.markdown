---
layout: "dome9"
page_title: "Check Point Cloud Guard Dome9: dome9_iplist"
sidebar_current: "docs-resource-dome9-iplist"
description: |-
  Manages IP list.
---

# dome9_iplist

The IpList resource has methods to create and manage IP lists in Dome9. IP lists are groups of IP addresses (typically in customer cloud environments), on which common actions are applied. For example, a Security Group could be applied to a list, instead of applying it to each IP address in the list individually.

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

* `name` - (Required) The IP list name.
* `description` - (Optional) IP list description, defaults to empty string.
* `items` - (Optional) Items for IP list, defaults to empty list.

The `items` block supports:
    
* `ip` - (Required) IP.
* `comment` - (Required) Comment.

## Attributes Reference

* `id` - IP list Id.

## Import

IP list can be imported; use `<IP LIST ID>` as the import ID. For example:

```shell
terraform import dome9_iplist.test 00000
```
