# Aws security group Example

This example will show you how to use Terraform to implement the CloudGuard Dome9 Aws security group rule resource.
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
