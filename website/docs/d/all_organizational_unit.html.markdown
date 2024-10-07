---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_organizational_unit"
sidebar_current: "docs-datasource-dome9-organizational-unit"
description: |-
  Get information about an Organizational Unit in Dome9.
---

# Data Source: dome9_all_organizational_unit

Use this data source to get information about all Organizational Units in Dome9.

## Example Usage

```hcl
data "dome9_all_organizational_unit" "test" {}
```

## Argument Reference

No arguments are needed.

## Attributes Reference

Returns a list of `dome9_organizational_unit`.

For more details, see the [dome9_organizational_unit documentation](./organizational_unit.html.markdown).