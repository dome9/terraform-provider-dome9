---
layout: "dome9"
page_title: "Check Point CloudGuard: dome9_integration"
sidebar_current: "docs-resource-dome9-integration"
description: |- 
  Creates and manages CloudGuard integrations.
---

# dome9\_integration

This resource is used to create and manage CloudGuard integrations. Integrations in CloudGuard allow you to connect and configure supported third-party services and tools.

## Example Usage

Basic usage:

```hcl
resource "dome9_integration" "example_integration" {
  name           = "Example Integration"
  type           = "webhook"
  configuration  = jsonencode({
    Url                = "https://example.com/webhook"
    MethodType         = "Post"
    AuthType           = "BasicAuth"
    Username           = "example-username"
    Password           = "example-password"
    IgnoreCertificate  = true
  })
}
```


## Argument Reference

The following arguments are supported:

- `name` - (Required) The name of the integration.
- `type` - (Required) The type of the integration. Possible values are:
    - `SNS`
    - `Email`
    - `PagerDuty`
    - `AwsSecurityHub`
    - `AzureDefender`
    - `GcpSecurityCommandCenter`
    - `Webhook`
    - `ServiceNow`
    - `Splunk`
    - `Jira`
    - `SumoLogic`
    - `QRadar`
    - `Slack`
    - `Teams`
- `configuration` - (Required) The configuration of the integration in JSON format. The configuration should contain all required details for the integration configuration.
    - Configuration details for each integration type can be found in the [CloudGuard API documentation](https://docs.cgn.portal.checkpoint.com/reference/integration_createintegration_post_v2integration).


## Import

This resource can be imported using the integration ID, which can be found in the CloudGuard console.

```shell
terraform import dome9_integration.example <INTEGRATION_ID>
```
