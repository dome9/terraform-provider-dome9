---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_organizational_unit"
sidebar_current: "docs-resource-dome9-organizational-unit"
description: |-
  Create organizational unit in Dome9
---

# dome9_organizational_unit

This resource is used to create and manage Organizational Unit in Dome9. An Organizational Unit is a group of cloud accounts representing, for example, a business unit or geographical region.

## Example Usage

Basic usage:

```hcl
resource "dome9_organizational_unit" "test_ou" {
  name      = "some_organizational_unit"
  parent_id = "00000000-0000-0000-0000-000000000000"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the organizational unit in Dome9.
* `parent_id` - (Optional) The organizational unit parent ID.

## Attributes Reference

* `id` - Organizational unit Id
* `account_id` - Dome9 internal account ID.
* `path` - Organizational Unit full path (IDs).
* `path_str` - Organizational Unit full path (names).
* `created` - Organizational Unit creation time.
* `updated` - Organizational Unit update time.
* `aws_cloud_accounts_count` - Number of AWS cloud accounts in the Organizational Unit.
* `azure_cloud_accounts_count` - Number of Azure cloud accounts in the Organizational Unit.
* `oci_cloud_accounts_count` - Number of OCI cloud accounts in the Organizational Unit.
* `google_cloud_accounts_count` - Number of GCP cloud accounts in the Organizational Unit.
* `k8s_cloud_accounts_count` - Number of K8S cloud accounts in the Organizational Unit.
* `shift_left_cloud_accounts_count` - Number of Shift Left cloud accounts in the Organizational Unit.
* `alibaba_cloud_accounts_count` - Number of Alibaba cloud accounts in the Organizational Unit.
* `container_registry_cloud_accounts_count` - Number of Container Registry cloud accounts in the Organizational Unit.
* `aws_aggregated_cloud_accounts_count` - Number of AWS cloud accounts in the Organizational Unit and its children.
* `azure_aggregate_cloud_accounts_count` - Number of Azure cloud accounts in the Organizational Unit and its children.
* `oci_aggregate_cloud_accounts_count` - Number of OCI cloud accounts in the Organizational Unit and its children.
* `google_aggregate_cloud_accounts_count` - Number of GCP cloud accounts in the Organizational Unit and its children.
* `k8s_aggregate_cloud_accounts_count` - Number of K8S cloud accounts in the Organizational Unit and its children.
* `shift_left_aggregate_cloud_accounts_count` - Number of Shift Left cloud accounts in the Organizational Unit and its children.
* `alibaba_aggregate_cloud_accounts_count` - Number of Alibaba cloud accounts in the Organizational Unit and its children.
* `container_registry_aggregate_cloud_accounts_count` - Number of Container Registry cloud accounts in the Organizational Unit and its children.
* `sub_organizational_units_count` - Number of sub Organizational Units.
* `is_root` - Is Organizational Unit root.
* `is_parent_root` - Is the parent of Organizational Unit root.


## Import

Organizational unit can be imported; use `<ORGANIZATIONAL UNIT ID>` as the import ID. 

For example:

```shell
terraform import dome9_organizational_unit.test 00000
```
