## 1.36.1 (September 29, 2024)
* Add support for Azure Org AWP SSE CMK encrypted disks scan to Azure org ([#226](https://github.com/dome9/terraform-provider-dome9/pull/226))

## 1.36.0 (September 16, 2024)
* Added support for CloudGuardResources and CodeSecurityResources permission at dome9_role resource

## 1.35.9 (September 12, 2024)
* Add support to AWP SSE CMK scanning

## 1.35.8 (September 03, 2024)
* Add support to AWP custom VPC

## 1.35.7 (August 20, 2024)


## 1.35.6 (August 18, 2024)
* Add retry mechanism to API calls

## 1.35.5 (August 06, 2024)
* Add dome9_vulnerability_policy resource and data

## 1.35.4 (August 01, 2024)
* Add support for 4 new Azure regions - Spain Central, Italy North, Poland Central, Poland Central

## 1.35.3 (August 01, 2024)
* Add support for a new AWS region ca-west-1(Calgary)

## 1.35.2 (July 30, 2024)
* Added retry for Integration Delete

## 1.35.1 (July 29, 2024)
* Added retry for dome9_cloudaccount_aws delete

## 1.35.0 (July 15, 2024)
* Add CloudGuard Integration support:
- `dome9_integration` (resource + data-source)
* **Notice:
* `dome9_continuous_compliance_notification` functionality has been replaced and the new implementation for this resource is `dome9_notification`[Learn more](https://registry.terraform.io/providers/dome9/dome9/latest/docs/data-sources/notification).**

## 1.34.0 (July 14, 2024)
* Add AWP Azure Onboarding Support

## 1.33.0 (July 10, 2024)
* Add support for the new Notification model:
- dome9_notification (data+resource)

## 1.32.0 (July 07, 2024)
* Al-2382-support-aws-centralized-scan-mode ([#182](https://github.com/dome9/terraform-provider-dome9/pull/182))

## 1.31.0 (July 04, 2024)
* Add AWS Organization Onboarding Support:
- dome9_aws_organization_onboarding (Resource + Data source)
- dome9_aws_organization_onboarding_management_stack (Data source)
- dome9_aws_organization_onboarding_member_account_configuration  (Data source)
* Bug fixes

## 1.30.1 (June 05, 2024)
- Added new "tenant_administrator_email_address" in OCI save temp data.
- Fixed an issue that caused an error while destroying a faulty unified onboarding.

## 1.30.0 (June 03, 2024)
* Go: 1.13.0 -> 1.19.2
* Terraform Plugin SDK: 1.11.0 -> 1.17.2

## 1.29.8 (May 30, 2024)
* Add support for Aws Unified Onboarding DELETE API

## 1.29.7 (April 04, 2024)
* CON-8179 : Onbaording AWS AWP ([#176](https://github.com/dome9/terraform-provider-dome9/pull/176))
* This pull request introduces changes related to the AWP AWS onboarding process in the Dome9 Terraform provider.
* Implementing AWP Onboarding API (https://github.com/dome9/dome9-sdk-go/releases/tag/v1.18.4)
* resource - dome9_awp_aws_onboarding
* data - dome9_awp_aws_onboarding
* data - dome9_awp_aws_onboarding_data

## 1.29.6 (February 05, 2024)
* Add Israel region

## 1.29.5 (August 02, 2023)
* Use new Dome9 SDK version

## 1.29.4 (June 14, 2023)
* fixed issue with regions names

## 1.29.3 (June 13, 2023)
* Added Support for a new Aws regions ("me-central-1", "ap-south-2", "ap-southeast-3", "ap-southeast-4", "eu-central-2", "eu-south-2")

## 1.29.2 (June 05, 2023)
* Added Support for a new Aws region me-central-1 (UAE)

## 1.29.0 (March 30, 2023)
* Added support for OCI vendor

## 1.28.6 (March 06, 2023)
* CON-4885 - K8s | TF Support for creating image assurance rulesets ([#166](https://github.com/dome9/terraform-provider-dome9/pull/166))

## 1.28.5 (November 20, 2022)
* CON-3826 - Add Support for IA Policy ([#160](https://github.com/dome9/terraform-provider-dome9/pull/160))

## 1.28.4 (October 13, 2022)
* Change SDK Version ([#157](https://github.com/dome9/terraform-provider-dome9/pull/157))
* Con 3824 k8s tf support for enabling ti within onboarding module ([#158](https://github.com/dome9/terraform-provider-dome9/pull/158))

## 1.28.2 (October 13, 2022)
* CON-3824 K8s TF Support for enabling TI within onboarding module ([#155](https://github.com/dome9/terraform-provider-dome9/pull/155))

## 1.28.1 (October 11, 2022)
* K8s | TF Support for enabling TI within onboarding module

## 1.28.0 (August 09, 2022)
* Added new Assessment resource and data source for Continuous Compliance

## 1.27.4 (August 04, 2022)


## 1.27.3 (August 04, 2022)
* Fix Dome9_cloudaccount_aws resource documentation

## 1.27.1 (June 24, 2022)
* Remove the `is_template` property from the `dome9_ruleset` resource documentation

## 1.27.0 (June 22, 2022)


## 1.26.0 (June 12, 2022)
* Support Canada Region

## 1.25.4 (May 04, 2022)
* fix some links, and explanations in the documentation

## 1.25.3 (April 25, 2022)


## 1.25.2 (April 24, 2022)


## 1.25.1 (April 24, 2022)
* Aws Unified Onboarding Resource Support
- dome9_aws_unified_onboarding resource
- dome9_aws_unified_onboarding data source
- aws_unified_onboarding_update_version_stack_config data source

## 1.25.0 (April 24, 2022)
* Aws Unified Onboarding Resource Support
- dome9_aws_unified_onboarding resource
- dome9_aws_unified_onboarding data source
- aws_unified_onboarding_update_version_stack_config data source

## 1.24.5 (January 03, 2022)
* Improve continuous_compliance_notification resource
- Added integration of Webhook notification with QRadar, Sumo, Jira.

## 1.24.4 (December 16, 2021)
* Improve  continuous_compliance_notification resource
- Added support for Teams integration
- Added support for Slack integration
- Added support for Webhook integration

## 1.24.3 (December 06, 2021)
* Update Ruleset recourse to work with the latest API

## 1.24.2 (November 29, 2021)
* Documentation fixes

## 1.24.1 (November 25, 2021)
* Documentation bug fix

## 1.24.0 (November 25, 2021)
* New AlibabaCloudAccount Resource

## 1.23.2 (August 25, 2021)
* Extend the dome9_cloudaccount_aws resource to support AWS china CloudAccount

## 1.23.1 (August 24, 2021)
* Extend the dome9_cloudaccount_aws resource to support AWS gov CloudAccount

## 1.23.0 (August 23, 2021)
- ServiceAccount Resource
- Adding new AWS region support (ap_northeast_3)
- Adding support for all CloudGuard DataCenters

## 1.22.0 (July 5, 2021)
* IMPROVEMENTS:
* Schema structure change for `resource_dome9_cloudaccount_kubernetes`,`data_source_dome9_cloudaccount_kubernetes`, ([#109](https://github.com/dome9/terraform-provider-dome9/pull/109))

## 1.21.1 (April 6, 2021)
BUG FIX
* Compliance Notification fix ([#104](https://github.com/dome9/terraform-provider-dome9/pulls))

## 1.21.0 (February 17, 2021)
* IMPROVEMENTS:
* Schema structure change for `data_aws_security_group`,`data_aws_security_group_rule`,`resource_dome9_aws_security_group`,`resource_dome9_aws_security_group_rule` ([#100](https://github.com/dome9/terraform-provider-dome9/pull/100))

## 1.20.1 (September 03, 2020)
* IMPROVEMENTS:
    - Support new two aws regions: Cape Town (af-south-1) and Milan (eu-south-1) ([#90](https://github.com/terraform-providers/terraform-provider-dome9/pull/90))
    - Support continues complience policy v2 ([#89](https://github.com/terraform-providers/terraform-provider-dome9/pull/89))
## 1.20.0 (July 29, 2020)
* **New Resource:** `dome9_cloudaccount_kubernetes` ([#88](https://github.com/terraform-providers/terraform-provider-dome9/pull/88))
* **New Data Source:** `dome9_cloudaccount_kubernetes`([#88](https://github.com/terraform-providers/terraform-provider-dome9/pull/88))
## 1.19.0 (June 04, 2020)
* **New Resource:** `dome9_aws_security_group_rule` ([#84](https://github.com/terraform-providers/terraform-provider-dome9/pull/84))
* **New Data Source:** `dome9_aws_security_group_rule`([#84](https://github.com/terraform-providers/terraform-provider-dome9/pull/84))
## 1.18.1 (April 22, 2020)

* IMPROVEMENTS:
    - Support new two aws regions: Bahrain (me-south-1) and Hong Kong (ap-east-1)
    - Using go vet rather than go lint
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

