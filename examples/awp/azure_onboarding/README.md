# AWP Azure Onboarding Example

This example will show you how to use Terraform to onboarding AWP Azure cloud account that protected by dome9.
This example codifies [this API](https://docs.cgn.portal.checkpoint.com/reference/agentless).

To run, configure your Dome9 provider as described in https://www.terraform.io/docs/providers/dome9/index.html

## Run the example

From inside of this directory:

```bash
terraform init
terraform plan -out theplan
terraform apply theplan
```

## Destroy 💥

```bash
terraform destroy
```
