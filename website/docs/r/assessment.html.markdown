---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_assessment"
sidebar_current: "docs-resource-dome9-assessment"
description: Run assessment in Dome9
---

# dome9_assessment

Begins an aws environment onboarding process. gets onboarding parameters and returns parameters 
that should be set for an aws_cloudformation_stack resource in order to complete the onboarding

## Example Usage

Basic usage:

```hcl
resource "dome9_assessment" "test" {
  bundle_id = BUNDLE_ID
  dome9_cloud_account_id = DOME9_CLOUD_ACCOUNT_ID
  cloud_account_id = CLOUD_ACCOUNT_ID
  cloud_account_type = CLOUD_ACCOUNT_TYPE
  should_minimize_result = SHOULD_MINIMIZE_RESULT
  request_id = REQUEST_ID
}
```

## Argument Reference

The following arguments are supported:

* `bundle_id` - (Required) The ruleset id to run.
* `dome9_cloud_account_id` - (Required) dome9 cloud account id to run the ruleset on.
* `cloud_account_id` - (Required) Cloud account id to run the ruleset on.
* `cloud_account_type` - (Required) Cloud account type. Can be: `Aws`, `Azure`, `GCP`, `Kubernetes`, `Terraform`, `Generic`, `KubernetesRuntimeAssurance`, `ShiftLeft`, `SourceCodeAssurance`, `ImageAssurance`, `Alibaba`, `Cft`, `ContainerRegistry`, `Ers`.
* `should_minimize_result` - (Required) Should minimize result size.
* `request_id` - (Required) Request id.
* `name` - (Optional) Name of the ruleset.
* `description` - (Optional) Description of the request.
* `external_cloud_account_id` - (Optional) External cloud account id.

  
## Attributes Reference

* `request` - Request content.
  * `is_template` - Is CloudGuard template ruleset.
  * `id` - Ruleset id.
  * `dome9_cloud_account_id` - dome9 cloud account id to run the ruleset on.
  * `cloud_account_id` - Cloud account id to run the ruleset on.
  * `cloud_account_type` - Cloud account type. Can be: `Aws`, `Azure`, `GCP`, `Kubernetes`, `Terraform`, `Generic`, `KubernetesRuntimeAssurance`, `ShiftLeft`, `SourceCodeAssurance`, `ImageAssurance`, `Alibaba`, `Cft`, `ContainerRegistry`, `Ers`.
  * `should_minimize_result` - Should minimize result size.
  * `name` - Name of the ruleset.
  * `description` - Description of the request.
  * `external_cloud_account_id` - External cloud account id.
  * `request_id` - Request id.
* `tests` - List of all the tests that have been run.
  * `error` - 
  * `tested_count` - 
  * `relevant_count` - 
  * `non_complying_count` - 
  * `exclusion_stats` - 
    * `tested_count` - 
    * `relevant_count` - 
    * `non_complying_count` - 
  * `entity_results` - 
    * `validation_status` - 
    * `is_relevant` - 
    * `is_valid` - 
    * `is_excluded` - 
    * `exclusion_id` - 
    * `remediation_id` - 
    * `error` - 
    * `test_obj` - 
      * `id` - 
      * `dome9_id` - 
      * `entity_type` - 
      * `entity_index` - 
      * `custom_entity_comparison_hash` - 
  * `rule` - 
    * `name` - Rule name.
    * `severity` - Rule severity.
    * `logic` - Rule logic.
    * `description` - Rule description.
    * `remediation` - Rule remediation.
    * `cloudbots` - Rule cloudbots.
    * `compliance_tag` - Compliance tag.
    * `domain` - Rule domain.
    * `priority` - Rule priority.
    * `control_title` - Control title.
    * `rule_id` - Rule id.
    * `category` - Rule
    * `labels` - Rule labels.
    * `logic_hash` - Rule logic_hash.
    * `is_default` - Is default rule.
  * `test_passed` - Is test passed: true/false.
* `test_entities` - Test entities map.
* `exclusions` - 
  * `platform` - 
  * `id` - 
  * `rules` - 
    * `logic_hash` - 
    * `id` - 
    * `name` - 
  * `logic_expressions` - 
  * `ruleset_id` - 
  * `cloud_account_ids` - 
  * `comment` - 
  * `organizational_unit_ids` - 
  * `date_range` - 
    * `from` - 
    * `to` - 
* `remediations` - 
  * `platform` -
  * `id` -
  * `rules` -
    * `logic_hash` -
    * `id` -
    * `name` -
  * `logic_expressions` -
  * `ruleset_id` -
  * `cloud_account_ids` -
  * `comment` -
  * `cloud_bots` -
  * `organizational_unit_ids` -
  * `date_range` -
    * `from` -
    * `to` -
* `data_sync_status` - 
  * `entity_type` - 
  * `recently_successful_sync` - 
  * `general_fetch_permission_issues` - 
  * `entities_with_permission_issues` - 
    * `external_id` - 
    * `name` - 
    * `cloud_vendor_identifier` - 
* `created_time` - 
* `assessment_id` - 
* `triggered_by` - 
* `assessment_passed` - 
* `has_errors` - 
* `stats` - 
  * `passed` - 
  * `passed_rules_by_severity` - 
    * `informational` - 
    * `low` - 
    * `medium` - 
    * `high` - 
    * `critical` - 
  * `failed` - 
  * `failed_rules_by_severity` - 
    * `informational` -
    * `low` -
    * `medium` -
    * `high` -
    * `critical` -
  * `error` - 
  * `failed_tests` - 
  * `logically_tested` - 
  * `failed_entities` - 
  * `excluded_tests` - 
  * `excluded_failed_tests` - 
  * `excluded_rules` - 
  * `excluded_rules_by_severity` - 
    * `informational` -
    * `low` -
    * `medium` -
    * `high` -
    * `critical` -
* `has_data_sync_status_issues` - 
* `comparison_custom_id` - 
* `additional_fields` - 































 