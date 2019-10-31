---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_iplist"
sidebar_current: "docs-datasource-dome9-iplist"
description: |-
  Get information about an IP list in Dome9.
---

# Data Source: dome9_iplist

Use this data source to get information about an IP list in Dome9.

## Example Usage

```hcl
data "dome9_iplist" "test" {
  id        = "IP List Id"
}

```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The id of the IP list in Dome9.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name	` - IP list name.
* `description` - IP list description.
* `items` - Items (IP addresses) in the IP list.
