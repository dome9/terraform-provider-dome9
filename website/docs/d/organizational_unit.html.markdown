---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_organizational_unit"
sidebar_current: "docs-datasource-dome9-organizational-unit"
description: |-
  Get information about an Organizational Unit in Dome9.
---

# Data Source: dome9_organizational_unit

Use this data source to get information about a Organizational Unit in Dome9.

## Example Usage

```hcl
data "dome9_organizational_unit" "test" {
  id = "d9-organizational-unit-id"
}

```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The ID of the organizational unit in Dome9.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - The name of the Organizational Unit in Dome9.
* `parent_id` - The Organizational Unit parent ID.
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
