## Add the account for IAM User access for GovCloud
resource "dome9_cloudaccount_aws" "aws_dome9_account_gov" {
  name = "account name"

  vendor = "awsGov"

  credentials {
    arn     = ""
    api_key = "AWS_ACCES_KEY_ID"
    secret  = "AWS_SECRET_ACCESS_KEY"
    type    = "UserBased"
  }

}
