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
  runtime_protection {
    enabled = true
  }
  admission_control {
    enabled = true
  }
  image_assurance {
    enabled = true
  }
  threat_intelligence {
    enabled = true
  }
}
```

## Argument Reference

The following arguments supported:

* `name` - (Required) The name of the kubernetes cluster as it will appear in Dome9 kubernetes cloud account.
* `organizational_unit_id` - (Optional) The Organizational Unit this cloud account will be attached to
* `runtime_protection` - (Optional) Runtime Protection which has the following configuration:
   * `enabled` - (Required) Is Runtime Protection enabled
* `admission_control` - (Optional) Admission Control which has the following configuration:
   * `enabled` - (Required) Is Admission Control enabled
* `image_assurance` - (Optional) Image Assurance which has the following configuration:
   * `enabled` - (Required) Is Image Assurance enabled
* `threat_intelligence` - (Optional) Threat Intelligence which has the following configuration:
   * `enabled` - (Required) Is Threat intelligence enabled
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
