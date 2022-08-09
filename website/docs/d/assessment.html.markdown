---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_assessment"
sidebar_current: "docs-resource-dome9-assessment"
description: Run assessment in Dome9
---

# dome9_assessment

Use this data source to get information about an assessment.

## Example Usage

Basic usage:

```hcl
data "dome9_assessment" "test" {
  id = ASSESSMENT_ID
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) Assessment id.


  
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
  * `error` - Test error.
  * `tested_count` - Number of assets that tested.
  * `relevant_count` - Number of assets that relevant to the test.
  * `non_complying_count` - Number of assets that non-complying to the test.
  * `exclusion_stats` - Exclusion stats.
    * `tested_count` - Number of assets that has been excluded from the test.
    * `relevant_count` - Number of assets that has been relevant to the test.
    * `non_complying_count` - Number of assets that has been non-complying to the test.
  * `entity_results` - Entity results.
    * `validation_status` - Can be: `Relevant`, `Valid`, `Excluded`.
    * `is_relevant` - Means if entity is relevant for this rule. for example rule = "Instance where name like '%db%' should have...", returns false in instance name = 'scheduler1'.
    * `is_valid` - Means if entity is compliant. for example for rule="Instance should have vpc", return true if instance i-123 is under vpc-234, and false if not.
    * `is_excluded` - Means if entity is excluded. for example for rule="Instance should have vpc exclude name='test'", return true if instance name is test, and false if not.
    * `exclusion_id` - Guid, can be Null.
    * `remediation_id` - Guid, can be Null.
    * `error` - Entity result error.
    * `test_obj` - The object that has been tested.
      * `id` - Id of the object.
      * `dome9_id` - Dome9 id of the object.
      * `entity_type` - Entity type of the object.
      * `entity_index` - Entity index of the object.
      * `custom_entity_comparison_hash` - Custom entity comparison hash of the object.
  * `rule` - Rule.
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
    * `category` - Rule category.
    * `labels` - Rule labels.
    * `logic_hash` - Rule logic_hash.
    * `is_default` - Is default rule.
  * `test_passed` - Is test passed: true/false.
* `exclusions` - List of exclusions associated with this assessment.
  * `platform` - Exclusions platform, can be: `Aws`, `Azure`, `GCP`, `Kubernetes`, `Terraform`, `Generic`, `KubernetesRuntimeAssurance`, `ShiftLeft`, `SourceCodeAssurance`, `ImageAssurance`, `Alibaba`, `Cft`, `ContainerRegistry`, `Ers`.
  * `id` - Exclusion ID.
  * `rules` - List of rules to apply the exclusion on.
    * `logic_hash` - Rule logic hash.
    * `id` - Rule ID.
    * `name` - Rule name.
  * `logic_expressions` - The GSL logic expressions of the exclusion.
  * `ruleset_id` - Ruleset ID to apply exclusion on.
  * `cloud_account_ids` - List of cloud account IDs to apply exclusion on.
  * `comment` - Comment text (free text).
  * `organizational_unit_ids` - List of organizational unit IDs to apply exclusion on.
  * `date_range` - Effective date range for the exclusion.
    * `from` - From date time.
    * `to` - To date time.
* `remediations` - List of remediations associated with this assessment.
  * `platform` - Remediation platform, can be: `Aws`, `Azure`, `GCP`, `Kubernetes`, `Terraform`, `Generic`, `KubernetesRuntimeAssurance`, `ShiftLeft`, `SourceCodeAssurance`, `ImageAssurance`, `Alibaba`, `Cft`, `ContainerRegistry`, `Ers`.
  * `id` - Exclusion ID.
  * `rules` - List of rules to apply the exclusion on.
    * `logic_hash` - Rule logic hash.
    * `id` - Rule ID.
    * `name` - Rule name.
  * `logic_expressions` - The GSL logic expressions of the exclusion.
  * `ruleset_id` - Ruleset ID to apply exclusion on.
  * `cloud_account_ids` - List of cloud account IDs to apply exclusion on.
  * `comment` - Comment text (free text).
  * `cloud_bots` - Cloud bots execution expressions.
  * `organizational_unit_ids` - List of organizational unit IDs to apply exclusion on.
  * `date_range` - Effective date range for the exclusion.
    * `from` - From date time.
    * `to` - To date time.
* `data_sync_status` - Data sync status - list of entities.
  * `entity_type` - Entity type.
  * `recently_successful_sync` - Is recently successful sync. True/False.
  * `general_fetch_permission_issues` - Is general fetch permission issues. True/False.
  * `entities_with_permission_issues` - List entities with permission issues.
    * `external_id` - Entity external id.
    * `name` - Entity name.
    * `cloud_vendor_identifier` - Entity cloud vendor identifier.
* `created_time` - Date of assessment.
* `assessment_id` - Assessment id.
* `triggered_by` - Reason for assessment.
* `assessment_passed` - Is assessment_passed. True/False.
* `has_errors` - Is assessment has errors. True/False.
* `stats` - Summary statistics for assessment.
  * `passed` - Number of passed rules.
  * `passed_rules_by_severity` - Passed rules divided by severity.
    * `informational` - Informational.
    * `low` - Low.
    * `medium` - Medium.
    * `high` - High.
    * `critical` - Critical.
  * `failed` - Number of failed rules.
  * `failed_rules_by_severity` - Failed rules divided by severity.
    * `informational` - Informational.
    * `low` - Low.
    * `medium` - Medium.
    * `high` - High.
    * `critical` - Critical.
  * `error` - Number of errors
  * `failed_tests` - Number of failed tests.
  * `logically_tested` - Total number of tests performed.
  * `failed_entities` - Number of failed entities.
  * `excluded_tests` - Number of excluded tests.
  * `excluded_failed_tests` - Number of excluded tests that also failed.
  * `excluded_rules` - Number of rules that contains only excluded tests.
  * `excluded_rules_by_severity` - Excluded rules divided by severity.
    * `informational` - Informational.
    * `low` - Low.
    * `medium` - Medium.
    * `high` - High.
    * `critical` - Critical.
* `has_data_sync_status_issues` - Is has data sync status issues. True/False.
* `comparison_custom_id` - Comparison custom id.
* `additional_fields` - Additional fields.































 