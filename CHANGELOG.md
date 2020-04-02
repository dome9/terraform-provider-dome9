## 1.18.0 (April 02, 2020)

* Cloud vendors in ruleset are sensitive ([#73](https://github.com/terraform-providers/terraform-provider-dome9/pull/73))

BUG FIXES:
* The dome9_aws_security_group.external_id property does not return the AWS security group ID, it does return the VPC ID instead ([#75](https://github.com/terraform-providers/terraform-provider-dome9/issues/75))

## 1.17.0 (February 24, 2020)

* Google cloud vendor fix ([#69](https://github.com/terraform-providers/terraform-provider-dome9/pull/69))

## 1.16.0 (December 29, 2019)

IMPROVEMENTS:
* `dome9_user` and `dome9_aws_role` supports additional permission fields ([#57](https://github.com/terraform-providers/terraform-provider-dome9/pull/57)) ([#62](https://github.com/terraform-providers/terraform-provider-dome9/pull/62))
* `dome9_cloudaccount_aws` supports `restricted_iam_entities` fields ([#63](https://github.com/terraform-providers/terraform-provider-dome9/pull/63))


## 1.15.0 (December 19, 2019)

FEATURES: 

* **New Resource:** `dome9_azure_security_group` ([#41](https://github.com/terraform-providers/terraform-provider-dome9/pull/41))
* **New Data Source:** `dome9_azure_security_group`([#41](https://github.com/terraform-providers/terraform-provider-dome9/pull/41))
* **New Resource:** `dome9_organizational_unit` ([#44](https://github.com/terraform-providers/terraform-provider-dome9/pull/44))
* **New Data Source:** `dome9_organizational_unit`([#44](https://github.com/terraform-providers/terraform-provider-dome9/pull/44))
* **New Resource:** `dome9_attach_iam_safe` ([#49](https://github.com/terraform-providers/terraform-provider-dome9/pull/49))
* **New Resource:** `dome9_user` ([#54](https://github.com/terraform-providers/terraform-provider-dome9/pull/54))
* **New Data Source:** `dome9_user`([#54](https://github.com/terraform-providers/terraform-provider-dome9/pull/54))

IMPROVEMENTS:
* Schema structure change for `dome9_cloudaccount_azure` and `dome9_cloudaccount_gcp`([#47](https://github.com/terraform-providers/terraform-provider-dome9/pull/47)) ([#48](https://github.com/terraform-providers/terraform-provider-dome9/pull/48))
* New field `iam_safe` in `dome9_cloudaccount_aws` resource ([#51](https://github.com/terraform-providers/terraform-provider-dome9/pull/51))


BUG FIXES:
* Sensitive strings will be hidden in Terraform plan for Azure and GCP cloud account onboarding ([#47](https://github.com/terraform-providers/terraform-provider-dome9/pull/47)) ([#48](https://github.com/terraform-providers/terraform-provider-dome9/pull/48))



## 1.14.0 (December 11, 2019)

FEATURES:

* **New Resource:** `dome9_aws_security_group` ([#33](https://github.com/terraform-providers/terraform-provider-dome9/pull/33))
* **New Data Source:** `dome9_aws_security_group`([#33](https://github.com/terraform-providers/terraform-provider-dome9/pull/33))
* **New Resource:** `dome9_aws_role` ([#35](https://github.com/terraform-providers/terraform-provider-dome9/pull/35))
* **New Data Source:** `dome9_aws_role`([#35](https://github.com/terraform-providers/terraform-provider-dome9/pull/35))

IMPROVEMENTS:
* Resource `dome9_rule_set` renamed to `dome9_ruleset`([#30](https://github.com/terraform-providers/terraform-provider-dome9/pull/30))

BUG FIXES:
* Support attachment of cloud account resources to organizational unit on creation ([#29](https://github.com/terraform-providers/terraform-provider-dome9/pull/29))
* Sensitive strings will be hidden in Terraform plan for Azure and GCP cloud account onboarding ([#40](https://github.com/terraform-providers/terraform-provider-dome9/pull/40))

## 1.13.0 (December 02, 2019)

FEATURES:

* **New Resource:** `dome9_rule_set` ([#28](https://github.com/terraform-providers/terraform-provider-dome9/pull/28))
* **New Data Source:** `dome9_rule_set`([#28](https://github.com/terraform-providers/terraform-provider-dome9/pull/28))

IMPROVEMENTS:
* .travis.yml: Add tflint process ([#22](https://github.com/terraform-providers/terraform-provider-dome9/issues/22))
* Align documentation code in GCP and AWS data resources and fixed continuous compliance documentation typo ([#23](https://github.com/terraform-providers/terraform-provider-dome9/issues/23), [#26](https://github.com/terraform-providers/terraform-provider-dome9/issues/26))

BUG FIXES:
* Don't print API secret key when provider is configured. API secret key now marked as sensitive in schema ([#18](https://github.com/terraform-providers/terraform-provider-dome9/issues/18))
* Removed linting issues ([#21](https://github.com/terraform-providers/terraform-provider-dome9/issues/21))
* .travis.yml now runs with go 1.13.x ([#22](https://github.com/terraform-providers/terraform-provider-dome9/issues/22))

## 1.12.1 (November 04, 2019)

IMPROVEMENTS:
* Documentation addition for `dome9_continuous_compliance_notification` resource and data source ([#14](https://github.com/terraform-providers/terraform-provider-dome9/issues/7))
* Remove double timestamp from logger ([#11](https://github.com/terraform-providers/terraform-provider-dome9/issues/11))
* Update examples to be more understandable and follow same convention ([#12](https://github.com/terraform-providers/terraform-provider-dome9/issues/12))

## 1.12.0 (November 01, 2019)

FEATURES:

* **New Resource:** `dome9_continuous_compliance_notification` ([#2](https://github.com/terraform-providers/terraform-provider-dome9/issues/2))
* **New Data Source:** `dome9_continuous_compliance_notification`([#2](https://github.com/terraform-providers/terraform-provider-dome9/issues/2))

BUG FIXES:

* documention fixes: links and phrasing ([#3](https://github.com/terraform-providers/terraform-provider-github/issues/3))

## 1.11.1 (October 28, 2019)

