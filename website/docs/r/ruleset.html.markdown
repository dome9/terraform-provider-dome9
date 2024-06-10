---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_ruleset"
sidebar_current: "docs-resource-dome9-ruleset"
description: |-
  Create ruleset in Dome9
---

# dome9_ruleset

This resource is used to create and manage rulesets in Dome9. Rulesets are sets of compliance rules.

## Example Usage

Basic usage:

```hcl
resource "dome9_ruleset" "ruleset" {
  name        = "some_ruleset"
  description = "this is the description of my ruleset"
  cloud_vendor = "aws"
  language = "en"
  hide_in_compliance = false
  rules {
    name = "some_rule2"
    logic = "EC2 should x"
    severity = "High"
    description = "rule description here"
    compliance_tag = "ct"
  
  }
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the ruleset in Dome9.
* `description` - (Optional) A description of the ruleset (what it represents); defaults to empty string.
* `cloud_vendor` - (Required) Cloud vendor that the ruleset is associated with, can be one of the following: `aws`, `azure`, `google`, or `imageassurance` (for Image Assurance rulesets).
* `language` - (Required) Language of the rules; defaults to 'en' (English).
* `hide_in_compliance` - (Required) hide in compliance - true/false.
*  [`rules`](#rules) - (Optional) List of rules in the ruleset.


### Rules

The `rules` supports the following arguments:
    
* `name` - (Required) Rule name.
* `logic` - (Required) Rule GSL logic. This is the text of the rule, using Dome9 GSL syntax.
* `severity` - (Optional) Rule severity (Default: "Low").
* `description` - (Optional) Rule description.
* `remediation` - (Optional) Rule remediation.
* `compliance_tag` - (Optional) A reference to a compliance standard.
* `domain` - (Optional) Rule domain.
* `priority` - (Optional) Rule priority.
* `control_title` - (Optional) Rule control title.
* `rule_id` - (Optional) Rule id.
* `category` - (Optional) Rule category.
* `is_default` - (Optional) is a default rule (Default: "false").


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Ruleset Id.
* `min_feature_tier` - Min feature tier.
* `created_time` - Rule set creation time.
* `updated_time` - Rule set last update time.
* `account_id` - The account id of the ruleset in Dome9.
* `system_bundle` - Is a system bundle or not.
* `rules_count` - The rules count.
* `is_template` - is a template rule.


## Import

Ruleset can be imported; use `<RULE SET ID>` as the import ID. 

For example:

```shell
terraform import dome9_rule_set.test 00000
```
