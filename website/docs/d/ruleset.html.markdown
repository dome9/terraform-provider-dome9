---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_ruleset"
sidebar_current: "docs-datasource-dome9-ruleset"
description: |-
  Get information about a ruleset in Dome9.
---

# Data Source: dome9_ruleset

Use this data source to get information about a ruleset in Dome9.

## Example Usage

```hcl
data "dome9_ruleset" "test" {
  id        = "d9-rule-set-id"
}

```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The id of the ruleset in Dome9.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - The name of the ruleset in Dome9.
* `description` - A description of the ruleset (what it represents); defaults to empty string.
* `cloud_vendor` - Cloud vendor that the ruleset is associated with, can be one of the following: `aws`, `azure` or `google`.
* `language` - Language of the rules; defaults to 'en' (English).
* `hide_in_compliance` - hide in compliance - true/false.
* `is_template` - is a template rule.
* `created_time` - Rule set creation time.
* `updated_time` - Rule set last update time.
* `rules` - List of rules in the ruleset.

### Rules

The `rules` supports the following attributes:

* `name` - Rule name.
* `logic` - Rule GSL logic. This is the text of the rule, using Dome9 GSL syntax.
* `severity` - Rule severity (Default: "Low").
* `description` - Rule description.
* `remediation` - Rule remediation.
* `compliance_tag` - A reference to a compliance standard.
* `domain` - Rule domain.
* `priority` - Rule priority.
* `control_title` - Rule control title.
* `rule_id` - Rule id.
* `Category` - Rule category.
* `logic_hash` - Rule logic hash.
* `is_default` - is a default rule (Default: "false").