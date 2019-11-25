## 1.12.2 (Unreleased)

IMPROVEMENTS:
* .travis.yml: Add tflint process ([#22](https://github.com/terraform-providers/terraform-provider-dome9/issues/22))

BUG FIXES:
* Don't print API secret key when provider is configured. API secret key now marked as sensitive in schema ([#18](https://github.com/terraform-providers/terraform-provider-dome9/issues/18))
* Removed linting issues ([#21](https://github.com/terraform-providers/terraform-provider-dome9/issues/21))
* .travis.yml now runs with go 1.13.x ([#22](https://github.com/terraform-providers/terraform-provider-dome9/issues/22))

## 1.12.1 (November 04, 2019)

IMPROVEMENTS:
* Documentation addition for `dome9_continuous_compliance_notification` resource and data source ([#14](https://github.com/terraform-providers/terraform-provider-dome9/issues/7))
* Remove double timestamp from logger ([#11](https://github.com/terraform-providers/terraform-provider-dome9/issues/11))
* Update examples to be more understandable and follow same convention ([#12](https://github.com/terraform-providers/terraform-provider-dome9/issues/12))
* Align documentation code in GCP and AWS data resources

## 1.12.0 (November 01, 2019)

FEATURES:

* **New Resource:** `dome9_continuous_compliance_notification` ([#2](https://github.com/terraform-providers/terraform-provider-dome9/issues/2))
* **New Data Source:** `dome9_continuous_compliance_notification`([#2](https://github.com/terraform-providers/terraform-provider-dome9/issues/2))

BUG FIXES:

* documention fixes: links and phrasing ([#3](https://github.com/terraform-providers/terraform-provider-github/issues/3))

## 1.11.1 (October 28, 2019)

