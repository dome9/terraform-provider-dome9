# IAM Safe Entity Example

This example will show you how to use Terraform to use IAM safe entity to protect your aws cloud account that protected by dome9.
This example codifies [this API](https://sc1.checkpoint.com/documents/CloudGuard_Dome9/Documentation/IAM-Safety/IAMSafety.html?tocpath=IAM%20Safety%7C_____1).

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
