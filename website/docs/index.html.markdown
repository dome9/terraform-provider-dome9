---
layout: "dome9"
page_title: "Provider: Check Point Cloud Guard Dome9"
sidebar_current: "docs-dome9-index"
description: |-
  The Check Point Cloud Guard Dome9 provider is used to interact with Dome9 resources. The provider needs to be configured with the proper credentials before it can be used.
---

# Check Point CloudGuard Dome9 Provider

The Check Point Cloud Guard Dome9 provider is used to interact with the many resources
supported by [Check Point Cloud Guard Dome9](https://www.dome9.com).

Use the navigation to the left to read about the available resources.

## Authentication

This provider requires a Dome9 API access ID and secret key in order to manage the resources.

To manage the full selection of resources, provide a
[Dome9 access id & secret key](https://secure.dome9.com/v2/settings/credentials) 
from an account with admin access permissions.

[Dome9 API documentation](/docs/cloud/api/index.html)
for more details about access to specific resources.

There are three ways to provide the required access key and secret:

- On the CLI, omit the `provider` block from your tf file, the CLI will ask for proper credentials.
  [CLI config file](/docs/commands/cli-config.html#credentials).
- Set the `DOME9_ACCESS_ID` and `DOME9_SECRET_KEY` environment variable.
- Fill the provider block with the appropriate arguments:    

## Example Usage

```hcl
# Configure the Dome9 Provider
provider "dome9" {
  dome9_access_id     = "${var.access_id}"
  dome9_secret_key    = "${var.secret_key}"
}

# Create an organization
resource "dome9_cloudaccount_aws" "account" {
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `dome9_access_id` - (Required) Dome9 access ID.
* `dome9_secret_key` - (Required) Dome9 access key.