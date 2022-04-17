# Attach IAM safe Example

This example will show you how to use Terraform to create Unified Onboarding for aws cloud account and protect him with CloudGuard.
This example codifies [this API](https://api-v2-docs.dome9.com/index.html#dome9-api-AwsUnifiedOnboarding).

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
