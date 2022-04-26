---
layout: "dome9"
page_title: "Provider: Check Point CloudGuard Dome9"
sidebar_current: "docs-dome9-index"
description: |-
   The Check Point CloudGuard Dome9 provider is used to interact with Dome9 security posture platform, to onboard cloud accounts and configure compliance policies. To use this  provider, you must create Dome9 API credentials.
---

# Check Point CloudGuard Dome9 Provider

The Check Point CloudGuard Dome9 provider is to interact with [CloudGuard](https://www.checkpoint.com/dome9/) security posture platform to onboard cloud environments and configure compliance policies.
To use this  provider, you must create CloudGuard API credentials.


Use the navigation menu on the left to read about the available resources.

## Authentication

This provider requires a CloudGuard API Key and Key secret in order to manage the resources.
You can generate these credentials in the CloudGuard portal or in the Infinity portal.
- In the CloudGuard portal, use [https://secure.dome9.com/v2/settings/credentials](https://secure.dome9.com/v2/settings/credentials) and, for more details, see [CloudGuard documentation](https://sc1.checkpoint.com/documents/CloudGuard_Dome9/Documentation/Settings/Credentials.htm?cshid=API_V2)
- In the Infinity portal,  use [https://portal.checkpoint.com/dashboard/cloudguard#/settings/credentials](https://portal.checkpoint.com/dashboard/cloudguard#/settings/credentials) and see [CloudGuard documentation](https://sc1.checkpoint.com/documents/Infinity_Portal/WebAdminGuides/EN/CloudGuard-PM-Admin-Guide/Documentation/Settings/Credentials.htm?cshid=API_V2)

To manage the full selection of resources, provide the credentials from an account with administrative access permissions.


You can use the Key and Secret in the following ways:

- On the CLI, omit the `provider` block from your tf file, and the CLI will prompt for proper credentials.
  [CLI config file](/docs/commands/cli-config.html#credentials).
- Set the `DOME9_ACCESS_ID` and `DOME9_SECRET_KEY` environment variables.
- [Optional] The provider works by default with US region. Set 'base_url' with one of the following 
  URLs for working with other supported regions.
  Support regions URLs list:
    - N.Virginia [DEFAULT]: 'https://api.dome9.com/v2/' 
    - Ireland : 'https://api.eu1.dome9.com/v2/'
    - Singapore : 'https://api.ap1.dome9.com/v2/'
    - Sydney : 'https://api.ap2.dome9.com/v2/'
    - Mumbai : 'https://api.ap3.dome9.com/v2/'
- Fill the provider block with the appropriate arguments:    


```hcl
# Configure the Dome9 Provider, works with US region by default
provider "dome9" {
  dome9_access_id     = "${var.access_id}"
  dome9_secret_key    = "${var.secret_key}"
}

# Create an organization
resource "dome9_cloudaccount_aws" "account" {
  # ...
}
```

```hcl
# Configure the Dome9 Provider None-US regions
provider "dome9" {
  dome9_access_id     = "${var.access_id}"
  dome9_secret_key    = "${var.secret_key}"
  base_url            = "${var.base_url}"
}

# Create an organization
resource "dome9_cloudaccount_aws" "account" {
  # ...
}
```

### Argument Reference

* `dome9_access_id` - (Required) the Dome9 API Key
* `dome9_secret_key` - (Required) the Dome9  key secret