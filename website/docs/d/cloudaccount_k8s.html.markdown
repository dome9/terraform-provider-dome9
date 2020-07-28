---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_k8s"
sidebar_current: "docs-datasource-dome9-cloudaccount-k8s"
description: |-
  Get information about k8s cloud account onboarded to Dome9.
---

# Data Source: dome9_cloudaccount_k8s

Use this data source to get information about a k8s cloud account onboarded to Dome9.

## Example Usage

```hcl
data "dome9_cloudaccount_k8s" "test" {
  id = "d9-K8S-cloud-account-id"
}

```

## Argument Reference

The following arguments supported:

* `id` - (Required) The Dome9 id for the k8s account

## Attributes Reference

In addition to all arguments above, the following attributes exported:

* `name` - The cloud account name in Dome9.
* `creation_date` - Date account was onboarded to Dome9.
* `vendor` - The cloud provider ("kubernetes").
* `organizational_unit_id` - Organizational unit id.
* `organizational_unit_path` - Organizational unit path.
* `organizational_unit_name` - Organizational unit name.
* `cluster_version` - The onboarded cluster version.
