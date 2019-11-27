---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_rule_set"
sidebar_current: "docs-datasource-dome9-rule-set"
description: |-
  Get information about a rule set in Dome9.
---

# Data Source: dome9_rule_set

Use this data source to get information about a rule set in Dome9.

## Example Usage

```hcl
data "dome9_rule_set" "test" {
  id        = "d9-rule-set-id"
}

```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The id of the rule set in Dome9.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name	` - The rule set name.
* `description` - The rule set description.
* `cloud_vendor` - Cloud vendor that the rule set is associated with.
* `hide_in_compliance` - Whether or not a rule set is hidden in compliance assessment.
* `is_template` - Template or costume rule set.
* `created_time` - Rule set creation time.
* `updated_time` - Rule set last update time.
* `rules` - Rules in the rule set.
