---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_kubernetes"
sidebar_current: "docs-datasource-dome9-cloudaccount-kubernetes"
description: |-
  Get information about kubernetes cloud account onboarded to Dome9.
---

# Data Source: dome9_cloudaccount_kubernetes

Use this data source to get information about a kubernetes cloud account onboarded to Dome9.

## Example Usage

```hcl
data "dome9_cloudaccount_kubernetes" "test" {
  id = "d9-kubernetes-cloud-account-id"
}

```

## Argument Reference

The following arguments supported:

* `id` - (Required) The Dome9 id for the kubernetes account

## Attributes Reference

In addition to all arguments above, the following attributes exported:

* `name` - The name of the kubernetes cluster as it appears in Dome9 kubernetes cloud account.
* `creation_date` - Date account was onboarded to Dome9.
* `vendor` - The cloud provider ("kubernetes").
* `organizational_unit_id` - Organizational unit id.
* `organizational_unit_path` - Organizational unit path.
* `organizational_unit_name` - Organizational unit name.
* `cluster_version` - The onboarded cluster version.
* `runtime_protection` - Runtime Protection details
    * `enabled` - Is Runtime Protection enabled
* `admission_control` - Admission Control details
    * `enabled` - Is Admission Control enabled
* `image_assurance` - Image Assurance details
    * `enabled` - Is Image Assurance enabled
    