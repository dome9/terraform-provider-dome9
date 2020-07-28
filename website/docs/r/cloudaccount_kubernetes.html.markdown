---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_cloudaccount_kubernetes"
sidebar_current: "docs-resource-dome9-cloudaccount-kubernetes"
description: |-
  Onboard kubernetes cloud account
---

# dome9_cloudaccount_kubernetes

This resource used to onboard kubernetes cloud accounts to Dome9. This is the first and pre-requisite step in order to apply Dome9 kubernetes features on the account.

## Example Usage

Basic usage:

```hcl
resource "dome9_cloudaccount_kubernetes" "test" {
  name  = "CLUSTER NAME"
  organizational_unit_id = "ORGANIZATIONAL UNIT ID"
}
```

## Argument Reference

The following arguments supported:

* `name` - (Required) The name of the kubernetes cluster as it will appear in Dome9 kubernetes cloud account.
* `organizational_unit_id` - (Optional) The Organizational Unit this cloud account will be attached to

## Attributes Reference

* `id` - The id of the account in Dome9.
* `creation_date` - Date account was onboarded to Dome9.
* `vendor` - The cloud provider ("kubernetes").
* `organizational_unit_path` - Organizational unit path.
* `organizational_unit_name` - Organizational unit name.
* `cluster_version` - The onboarded cluster version.

## Import

Kubernetes cloud account can be imported; use `<KUBERNETES CLOUD ACCOUNT ID>` as the import ID. 

For example:

```shell
terraform import dome9_cloudaccount_kubernetes.test 00000000-0000-0000-0000-000000000000
```
