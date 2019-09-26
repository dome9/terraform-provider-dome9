# CloudAccount Example

This example will show you how to use Terraform to implement the Dome9 GCP cloud account resource.
This example codifies [this api](https://api-v2-docs.dome9.com/#Dome9-API-GoogleCloudAccount)

## Run the example

From inside of this directory:

```bash
export DOME9_ACCESS_ID=(this is a secret)
export DOME9_SECRET_KEY=(this is a secret)
terraform init
terraform plan -out theplan
terraform apply theplan
```

## Remove the example

```bash
terraform destroy
```
