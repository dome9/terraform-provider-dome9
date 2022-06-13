# Admission Control Policy Example

This example will show you how to use Terraform to implement the Dome9 Admission Control Policy resource.
This example codifies [this API](https://api-v2-docs.dome9.com/#dome9-api-KubernetesAdmissionControl).

To run, configure your Dome9 provider as described in https://www.terraform.io/docs/providers/dome9/index.html

## Run the example

From inside this directory:

```bash
terraform init
terraform plan -out theplan
terraform apply theplan
```

## Destroy 💥

```bash
terraform destroy
```
