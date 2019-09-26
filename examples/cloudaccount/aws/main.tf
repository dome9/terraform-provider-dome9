provider "dome9" {
  dome9_access_id  = "--"
  dome9_secret_key = "--"
}

resource "dome9_cloudaccount_aws" "aws_ca" {
  name = "account name"

  credentials = {
    arn    = "ARN"
    secret = "SECRET"
    type   = "RoleBased"
  }
}
