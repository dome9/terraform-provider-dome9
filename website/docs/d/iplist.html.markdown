---
layout: "dome9"
page_title: "Check Point Cloud Guard Dome9: dome9_iplist"
sidebar_current: "docs-datasource-dome9-iplist"
description: |-
  Get information on IP list.
---

# Data Source: dome9_iplist

Use this data source to get an IP list.

## Example Usage

```hcl
data "dome9_iplist" "test" {
  id        = "IP List Id"
}

```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The IP List Id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name	` - IP List Name.
* `description` - IP List Description.
* `items` - Items in the IP list.
