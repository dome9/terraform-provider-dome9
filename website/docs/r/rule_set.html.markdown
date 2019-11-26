---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_rule_set"
sidebar_current: "docs-resource-dome9-rule-set"
description: |-
  Create rule set in Dome9
---

# dome9_rule_set

This resource is used to create and manage rule sets in Dome9. Rule sets are a bundle of compliance rules.

## Example Usage

Basic usage:

```hcl
resource "dome9_rule_set" "ruleset" {
  name        = "some_ruleset"
  description = "this is the descrption of my ruleset"
  cloud_vendor = "aws"
  language = "en"
  hide_in_compliance = false
  is_template = false
  rules {
    name = "some_rule2"
    logic = "EC2 should x"
    severity = "High"
    description = "rule description here"
    compliance_tag = "ct"
    domain = "test"
    priority = "high"
    control_title = "ct"
    rule_id = ""
    is_default = false
  }
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the rule set in Dome9.
* `description` - (Optional) A description of the rule set (what it represents); defaults to empty string.
* `cloud_vendor` - (Required) Cloud vendor that the rule set is associated with.
* `language` - (Optional) Language of the rules; defaults to english.
* `hide_in_compliance` - (Optional) Whether or not a rule set is hidden in compliance assessment; defaults to false.
* `is_template` - (Optional) Template or costume rule set; defaults to false.

### Rules 

The `rules` supports the following arguments:
    
* `name` - (Optional) Rule name.
* `logic` - (Optional) Rule logic.
* `severity` - (Optional) Rule severity.
* `description` - (Optional) Rule description.
* `compliance_tag` - (Optional) Compliance tag.
* `domain` - (Optional) Rule domain.
* `priority` - (Optional) Rule priority.
* `control_title` - (Optional) Rule control title.
* `rule_id` - (Optional) Rule ID.
* `is_default` - (Optional) Is rule default.

## Attributes Reference

* `id` - Rule set Id

## Import

Rule set can be imported; use `<RULE SET ID>` as the import ID. 

For example:

```shell
terraform import dome9_rule_set.test 00000
```
