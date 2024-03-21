---
layout: "dome9"
page_title: "Check Point CloudGuard Dome9: dome9_awp_aws_onboarding"
sidebar_current: "docs-resource-dome9-awp-aws-onboarding"
description: |-
  Creates an AWP AWS Onboarding in Dome9
---

# dome9_awp_aws_onboarding

This resource is used to create and modify AWP AWS Onboarding in CloudGuard Dome9.

## Example Usage

Basic usage:

```hcl
# Required Providers Configuration Block for Dome9, AWS, HTTP, and Local
terraform {
  required_providers {
    dome9 = {
      source = "dome9/dome9"
      version = "1.29.6"
    }
    aws = {
      source = "hashicorp/aws"
      version = "5.39.1"
    }
    http = {
      source  = "hashicorp/http"
      version = "3.4.2"
    }
    local = {
      source  = "hashicorp/local"
      version = "2.5.1"
    }
  }
}
# The Dome9 provider is used to interact with the resources supported by Dome9.
# The provider needs to be configured with the proper credentials before it can be used.
# Use the dome9_access_id and dome9_secret_key attributes of the provider to provide the Dome9 access key and secret key.
# The base_url attribute is used to specify the base URL of the Dome9 API.
# The Dome9 provider supports several options for providing these credentials. The following example demonstrates the use of static credentials:
#you can read the Dome9 provider documentation to understand the full set of options available for providing credentials.
#https://registry.terraform.io/providers/dome9/dome9/latest/docs#authentication
provider "dome9" {
  dome9_access_id     = "DOME9_ACCESS_ID"
  dome9_secret_key    = "DOME9_SECRET_KEY"
  base_url            = "https://api.dome9.com/v2/"
}

# AWS Provider Configurations
# The AWS provider is used to interact with the resources supported by AWS.
# The provider needs to be configured with the proper credentials before it can be used.
# Use the access_key, secret_key, and token attributes of the provider to provide the credentials.
# also you can use the shared_credentials_file attribute to provide the path to the shared credentials file.
# The AWS provider supports several options for providing these credentials. The following example demonstrates the use of static credentials:
#you can read the AWS provider documentation to understand the full set of options available for providing credentials.
#https://registry.terraform.io/providers/hashicorp/aws/latest/docs#authentication-and-configuration
provider "aws" {
  region     = "AWS_REGION"
  access_key = "AWS_ACCESS_KEY"
  secret_key = "AWS_SECRET_KEY"
  token      = "AWS_SESSION_TOKEN"
}

# The resource block defines a Dome9 AWS Cloud Account onboarding.
# The Dome9 AWS Cloud Account onboarding resource allows you to onboard an AWS account to Dome9.
# this resource is optional and can be ignored and you need to pass CloudGuard account id Dome9 AWP AWS Onboarding resource and "dome9_awp_aws_get_onboarding_data" data source.
/*
resource "dome9_cloudaccount_aws" "aws_onboarding_account_test" {
	name  = "aws_onboarding_account_test"
	credentials  {
		arn    = "arn:aws:iam::478980137264:role/CloudGuard-Connect"
		secret = "@R2PUjk0up42HHDtD9CByVF8"
		type   = "RoleBased"
	}
	net_sec {
		regions {
			new_group_behavior = "ReadOnly"
			region             = "us_west_2"
		}
	}
}
*/

# The dome9_awp_aws_get_onboarding_data data source allows you to get the onboarding data of an AWS account.
# you can pass the CloudGuard account id to get the onboarding data of the AWS account or the external account number for the AWS account.
data "dome9_awp_aws_get_onboarding_data" "dome9_awp_aws_onboarding_data_source" {
  cloud_account_id = "CLOUDGUARD_ACCOUNT_ID or EXTERNAL_AWS_ACCOUNT_NUMBER"
}

# The local block defines a local value that can be used to store the data that is used in multiple places in the configuration.
# the scan_mode is used to define the scan mode of the Dome9 AWP AWS Onboarding.
# the valid values are "inAccount" and "saas". you need to select one of them based on the scan mode of the Dome9 AWP AWS Onboarding.
locals {
  scan_mode = "inAccount or saas" # the valid values are "inAccount" and "saas" when onboarding the AWS account to Dome9 AWP.
  stage = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.stage
  region = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.region
  cloud_guard_backend_account_id = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.cloud_guard_backend_account_id
  agentless_bucket_name = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.agentless_bucket_name
  remote_functions_prefix_key = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.remote_functions_prefix_key
  remote_snapshots_utils_function_name = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.remote_snapshots_utils_function_name
  remote_snapshots_utils_function_run_time = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.remote_snapshots_utils_function_run_time
  remote_snapshots_utils_function_time_out = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.remote_snapshots_utils_function_time_out
  awp_client_side_security_group_name = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.awp_client_side_security_group_name
  cross_account_role_external_id = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.cross_account_role_external_id
  remote_snapshots_utils_function_s3_pre_signed_url = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.remote_snapshots_utils_function_s3_pre_signed_url
}

# <All these AWS Resources SHOULD NOT be changed in its configuration, otherwise, the awp onboarding process will not work properly>
# CloudGuardAWPCrossAccountRole : The IAM role that is used to allow AWP to access the AWS account.
# CloudGuardAWPCrossAccountRolePolicy : The IAM policy that is used to define the permissions for the CloudGuardAWPCrossAccountRole.
# CloudGuardAWPSnapshotsUtilsFunction : The Lambda function that is used to manage remote actions and resources.
# CloudGuardAWPSnapshotsUtilsFunctionZip : The local file that is used to store the remote function file to be used in the lambda function.
# CloudGuardAWPSnapshotsUtilsLogGroup : The CloudWatch log group that is used to store the logs of the CloudGuardAWPSnapshotsUtilsFunction.
# CloudGuardAWPSnapshotsUtilsLambdaExecutionRole : The IAM role that is used to allow the CloudGuardAWPSnapshotsUtilsFunction to execute.
# CloudGuardAWPSnapshotsPolicy : The IAM policy that is used to define the permissions for the CloudGuardAWPSnapshotsUtilsFunction.
# CloudGuardAWPLambdaExecutionRolePolicy : The IAM policy that is used to define the permissions for the CloudGuardAWPSnapshotsUtilsFunction.
# CloudGuardAWPLambdaExecutionRolePolicy_SaaS : The IAM policy that is used to define the permissions for the CloudGuardAWPSnapshotsUtilsFunction in SaaS mode.
# CloudGuardAWPKey : The KMS key that is used to re-encrypt the snapshots in SaaS mode.
# CloudGuardAWPKeyAlias : The KMS key alias that is used to reference the KMS key in SaaS mode.
# CloudGuardAWPSnapshotsUtilsCleanupFunctionInvocation : The Lambda invocation that is used to clean up the resources after the onboarding process.
# The data block defines a data source that can be used to get the current AWS partition.
data "aws_partition" "current" {}
# The data block defines a data source that can be used to get the current AWS region.
data "aws_region" "current" {}
# The data block defines a data source that can be used to get the current AWS caller identity.
data "aws_caller_identity" "current" {}

# Cross account role to allow CloudGuard access
# The CloudGuardAWPCrossAccountRole resource defines an IAM role that is used to allow AWP to access the AWS account.
resource "aws_iam_role" "CloudGuardAWPCrossAccountRole" {
  name               = "CloudGuardAWPCrossAccountRole"
  description        = "CloudGuard AWP Cross Account Role"
  assume_role_policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [{
      Effect    = "Allow"
      Principal = {
        AWS = local.cloud_guard_backend_account_id
      }
      Action    = "sts:AssumeRole"
      Condition = {
        StringEquals = {
          "sts:ExternalId" = local.cross_account_role_external_id
        }
      }
    }]
  })

  depends_on = [aws_lambda_function.CloudGuardAWPSnapshotsUtilsFunction]
}

# The CloudGuardAWPCrossAccountRolePolicy resource defines an IAM policy that is used to define the permissions for the CloudGuardAWPCrossAccountRole.
resource "aws_iam_policy" "CloudGuardAWP" {
  name        = "CloudGuardAWP"
  description = "Policy for CloudGuard AWP"

  policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = [
          "ec2:DescribeInstances",
          "ec2:DescribeSnapshots",
          "ec2:DescribeRegions",
          "ec2:DescribeVolumes"
        ]
        Resource = "*"
      },
      {
        Effect   = "Allow"
        Action   = [
          "lambda:InvokeFunction",
          "lambda:GetFunction",
          "lambda:GetLayerVersion",
          "lambda:TagResource",
          "lambda:ListTags",
          "lambda:UntagResource",
          "lambda:UpdateFunctionCode",
          "lambda:UpdateFunctionConfiguration",
          "lambda:GetFunctionConfiguration"
        ]
        Resource = aws_lambda_function.CloudGuardAWPSnapshotsUtilsFunction.arn
      },
      {
        Effect   = "Allow"
        Action   = "cloudformation:DescribeStacks"
        Resource = "arn:${data.aws_partition.current.partition}:cloudformation:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:stack/*"
      },
      {
        Effect   = "Allow"
        Action   = "s3:GetObject"
        Resource = "arn:${data.aws_partition.current.partition}:s3:::${local.agentless_bucket_name}/${local.remote_functions_prefix_key}*"
      }
    ]
  })
}

# The CloudGuardAWPCrossAccountRoleAttachment resource attaches the CloudGuardAWPCrossAccountRolePolicy to the CloudGuardAWPCrossAccountRole.
resource "aws_iam_role_policy_attachment" "CloudGuardAWPCrossAccountRoleAttachment" {
  role       = aws_iam_role.CloudGuardAWPCrossAccountRole.name
  policy_arn = aws_iam_policy.CloudGuardAWP.arn
}
# end resources for CloudGuardAWPCrossAccountRole

# Cross account role policy
# The CloudGuardAWPCrossAccountRolePolicy resource defines an IAM policy that is used to define the permissions for the CloudGuardAWPCrossAccountRole.
resource "aws_iam_policy" "CloudGuardAWPCrossAccountRolePolicy" {
  count = local.scan_mode == "inAccount" ? 1 : 0
  name        = "CloudGuardAWPCrossAccountRolePolicy"
  description = "Policy for CloudGuard AWP Cross Account Role"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "ec2:CreateSecurityGroup",
          "ec2:DescribeManagedPrefixLists",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeSecurityGroupRules",
          "ec2:RevokeSecurityGroupEgress",
          "ec2:AuthorizeSecurityGroupEgress",
          "ec2:CreateTags",
        ]
        Resource = "*"
      },
      {
        Effect = "Allow"
        Action = [
          "ec2:DeleteSecurityGroup",
        ]
        Resource = "*"
        Condition = {
          StringEquals = {
            "aws:ResourceTag/Owner" = "CG.AWP"
          }
        }
      },
    ]
  })
}

# The CloudGuardAWPCrossAccountRolePolicy_SaaS resource defines an IAM policy that is used to define the permissions for the CloudGuardAWPCrossAccountRole in SaaS mode.
resource "aws_iam_policy" "CloudGuardAWPCrossAccountRolePolicy_SaaS" {
  count = local.scan_mode == "saas" ? 1 : 0
  name        = "CloudGuardAWPCrossAccountRolePolicy_SaaS"
  description = "Policy for CloudGuard AWP Cross Account Role - SaaS Mode"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "kms:DescribeKey",
          "kms:ReplicateKey",
        ]
        Resource = [aws_kms_key.CloudGuardAWPKey[count.index].arn]
      },
      {
        Effect = "Allow"
        Action = [
          "kms:PutKeyPolicy",
          "kms:ScheduleKeyDeletion",
          "kms:CancelKeyDeletion",
          "kms:TagResource",
        ]
        Resource = aws_kms_key.CloudGuardAWPKey[count.index].arn
      },
      {
        Effect = "Allow"
        Action = [
          "kms:CreateKey",
        ]
        Resource = "*"
      },
    ]
  })
}

# The CloudGuardAWPCrossAccountRolePolicyAttachment resource attaches the CloudGuardAWPCrossAccountRolePolicy to the CloudGuardAWPCrossAccountRole.
resource "aws_iam_policy_attachment" "CloudGuardAWPCrossAccountRolePolicyAttachment" {
  count       = local.scan_mode == "inAccount" ? 1 : 0
  name       = "CloudGuardAWPCrossAccountRolePolicyAttachment"
  policy_arn  = aws_iam_policy.CloudGuardAWPCrossAccountRolePolicy[count.index].arn
  roles       = [aws_iam_role.CloudGuardAWPCrossAccountRole.name]
}

# The CloudGuardAWPCrossAccountRolePolicyAttachment_SaaS resource attaches the CloudGuardAWPCrossAccountRolePolicy_SaaS to the CloudGuardAWPCrossAccountRole.
resource "aws_iam_policy_attachment" "CloudGuardAWPCrossAccountRolePolicyAttachment_SaaS" {
  count = local.scan_mode == "saas" ? 1 : 0
  name       = "CloudGuardAWPCrossAccountRolePolicyAttachment_SaaS"
  policy_arn  = aws_iam_policy.CloudGuardAWPCrossAccountRolePolicy_SaaS[count.index].arn
  roles       = [aws_iam_role.CloudGuardAWPCrossAccountRole.name]
}
# END Cross account role policy

# The CloudGuardAWPSnapshotsUtilsFunctionZip resource defines http data source to download the remote function file from S3 pre-signed URL.
data "http" "CloudGuardAWPSnapshotsUtilsFunctionZip" {
  url = local.remote_snapshots_utils_function_s3_pre_signed_url
  method = "GET"
  request_headers = {
    Accept = "application/zip"
  }
}

# The CloudGuardAWPSnapshotsUtilsFunctionZip resource defines a local file that is used to store the remote function file to be used in the lambda function.
resource "local_file" "CloudGuardAWPSnapshotsUtilsFunctionZip" {
  filename = "${local.remote_snapshots_utils_function_name}7.zip"
  content_base64 = data.http.CloudGuardAWPSnapshotsUtilsFunctionZip.response_body_base64
}

# AWP proxy lambda function
# The CloudGuardAWPSnapshotsUtilsFunction resource defines a lambda function that is used to manage remote actions and resources.
resource "aws_lambda_function" "CloudGuardAWPSnapshotsUtilsFunction" {
  function_name    = local.remote_snapshots_utils_function_name
  handler          = "snapshots_utils.lambda_handler"
  description      = "CloudGuard AWP Proxy for managing remote actions and resources"
  role             = aws_iam_role.CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.arn
  runtime          = "python3.9"
  memory_size      = 256
  timeout          = local.remote_snapshots_utils_function_time_out
  filename         = local_file.CloudGuardAWPSnapshotsUtilsFunctionZip.filename

  environment {
    variables = {
      CP_AWP_AWS_ACCOUNT        = local.cloud_guard_backend_account_id
      CP_AWP_MR_KMS_KEY_ID      = local.scan_mode == "saas" ? aws_kms_key.CloudGuardAWPKey[0].arn : ""
      CP_AWP_SCAN_MODE          = local.scan_mode
      CP_AWP_SECURITY_GROUP_NAME = local.awp_client_side_security_group_name
      AWS_PARTITION             = data.aws_partition.current.partition
      CP_AWP_LOG_LEVEL		  = "DEBUG"
    }
  }

  tags = {
    Owner = "CG.AWP"
  }
}

resource "aws_lambda_permission" "allow_cloudguard" {
  statement_id  = "AllowExecutionFromCloudGuard"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.CloudGuardAWPSnapshotsUtilsFunction.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = "arn:${data.aws_partition.current.partition}:s3:::${local.agentless_bucket_name}/*"
}
# END AWP proxy lambda function

# CloudGuardAWPSnapshotsUtilsLogGroup : The CloudWatch log group that is used to store the logs of the CloudGuardAWPSnapshotsUtilsFunction.
resource "aws_cloudwatch_log_group" "CloudGuardAWPSnapshotsUtilsLogGroup" {
  name              = "/aws/lambda/CloudGuardAWPSnapshotsUtils"
  retention_in_days = 30
  depends_on = [
    aws_lambda_function.CloudGuardAWPSnapshotsUtilsFunction
  ]
}

# AWP proxy lambda function role
# The CloudGuardAWPSnapshotsUtilsLambdaExecutionRole resource defines an IAM role that is used to allow the CloudGuardAWPSnapshotsUtilsFunction to execute.
resource "aws_iam_role" "CloudGuardAWPSnapshotsUtilsLambdaExecutionRole" {
  name               = "CloudGuardAWPLambdaExecutionRole"
  description        = "CloudGuard AWP proxy lambda function execution role"
  assume_role_policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Effect    = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
        Action    = "sts:AssumeRole"
      }
    ]
  })

  tags = {
    Owner = "CG.AWP"
  }
}

# The CloudGuardAWPSnapshotsPolicy resource defines an IAM policy that is used to define the permissions for the CloudGuardAWPSnapshotsUtilsFunction.
resource "aws_iam_policy" "CloudGuardAWPSnapshotsPolicy" {
  name        = "CloudGuardAWPSnapshotsPolicy"
  description = "Policy for managing snapshots at client side and delete AWP keys"

  policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = [
          "ec2:CreateTags",
          "ec2:CopySnapshot",
          "ec2:CreateSnapshot",
          "ec2:CreateSnapshots",
          "ec2:DescribeSnapshots",
          "ec2:DescribeRegions"
        ]
        Resource = "*"
      },
      {
        Effect   = "Allow"
        Action   = [
          "ec2:DeleteSnapshot"
        ]
        Resource = "*"
        Condition = {
          StringEquals = {
            "aws:ResourceTag/Owner" = "CG.AWP"
          }
        }
      },
      {
        Effect   = "Allow"
        Action   = [
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Resource = [aws_cloudwatch_log_group.CloudGuardAWPSnapshotsUtilsLogGroup.arn]
      }
    ]
  })
}

# The CloudGuardAWPSnapshotsUtilsLambdaExecutionRoleAttachment resource attaches the CloudGuardAWPSnapshotsPolicy to the CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.
resource "aws_iam_role_policy_attachment" "CloudGuardAWPSnapshotsUtilsLambdaExecutionRoleAttachment" {
  role       = aws_iam_role.CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.name
  policy_arn = aws_iam_policy.CloudGuardAWPSnapshotsPolicy.arn
}
# END AWP proxy lambda function role

# AWP proxy lambda function role policy
# The CloudGuardAWPLambdaExecutionRolePolicy resource defines an IAM policy that is used to define the permissions for the CloudGuardAWPSnapshotsUtilsFunction.
resource "aws_iam_policy" "CloudGuardAWPLambdaExecutionRolePolicy" {
  count       = local.scan_mode == "inAccount" ? 1 : 0
  name        = "CloudGuardAWPLambdaExecutionRolePolicy"
  description = "Policy for CloudGuard AWP Lambda Execution Role"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "ec2:RunInstances",
        ]
        Resource = "*"
      },
      {
        Effect = "Allow"
        Action = [
          "ec2:TerminateInstances",
          "ec2:DeleteVolume",
        ]
        Resource = "*"
        Condition = local.scan_mode == "inAccount" ? {
          StringEquals = {
            "aws:ResourceTag/Owner" = "CG.AWP"
          }
        } : null
      },
      {
        Effect = "Allow"
        Action = [
          "iam:CreateServiceLinkedRole",
        ]
        Resource = ["arn:${data.aws_partition.current.partition}:iam::${data.aws_caller_identity.current.account_id}:role/aws-service-role/spot.amazonaws.com/AWSServiceRoleForEC2Spot"]
      },
      {
        Effect = "Allow"
        Action = [
          "kms:Decrypt",
          "kms:DescribeKey",
          "kms:GenerateDataKey*",
          "kms:CreateGrant",
          "kms:Encrypt",
          "kms:ReEncrypt*",
        ]
        Resource = "*"
      },
      {
        Effect = "Allow"
        Action = [
          "ec2:CreateVpc",
          "ec2:CreateSecurityGroup",
          "ec2:CreateSubnet",
          "ec2:DescribeInstances",
          "ec2:DescribeVolumes",
          "ec2:DescribeVpcs",
          "ec2:DescribeSubnets",
          "ec2:DescribeRouteTables",
          "ec2:DescribeNetworkAcls",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeInternetGateways",
          "ec2:DescribeSecurityGroupRules",
          "ec2:ModifySubnetAttribute",
          "ec2:CreateVpcEndpoint",
          "ec2:DescribeVpcEndpoints",
        ]
        Resource = "*"
      },
      {
        Effect = "Allow"
        Action = [
          "ec2:AssociateRouteTable",
          "ec2:DeleteVpc",
          "ec2:DeleteSubnet",
          "ec2:DeleteVolume",
          "ec2:DeleteInternetGateway",
          "ec2:RevokeSecurityGroupEgress",
          "ec2:RevokeSecurityGroupIngress",
          "ec2:AuthorizeSecurityGroupEgress",
          "ec2:DeleteSecurityGroup",
          "ec2:DeleteVpcEndpoints",
          "ec2:CreateNetworkAclEntry",
        ]
        Resource = "*"
        Condition = local.scan_mode == "inAccount" ? {
          StringEquals = {
            "aws:ResourceTag/Owner" = "CG.AWP"
          }
        } : null
      },
    ]
  })
}

# The CloudGuardAWPLambdaExecutionRolePolicyAttachment resource attaches the CloudGuardAWPLambdaExecutionRolePolicy to the CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.
resource "aws_iam_policy" "CloudGuardAWPLambdaExecutionRolePolicy_SaaS" {
  count       = local.scan_mode == "saas" ? 1 : 0
  name        = "CloudGuardAWPLambdaExecutionRolePolicy_SaaS"
  description = "Policy for CloudGuard AWP Lambda Execution Role - SaaS Mode"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "ec2:ModifySnapshotAttribute",
        ]
        Resource = "*"
      },
      {
        Effect = "Allow"
        Action = [
          "kms:ReEncrypt*",
          "kms:Encrypt",
        ]
        Resource = "*"
      },
      {
        Effect = "Allow"
        Action = [
          "kms:Decrypt",
          "kms:DescribeKey",
          "kms:GenerateDataKey*",
          "kms:CreateGrant",
        ]
        Resource = "*"
      },
      {
        Effect = "Allow"
        Action = [
          "kms:ScheduleKeyDeletion",
        ]
        Resource = "*"
        Condition = {
          StringEquals = {
            "aws:ResourceTag/Owner" = "CG.AWP"
          }
        }
      },
    ]
  })
}

# The CloudGuardAWPLambdaExecutionRolePolicyAttachment resource attaches the CloudGuardAWPLambdaExecutionRolePolicy to the CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.
resource "aws_iam_policy_attachment" "CloudGuardAWPLambdaExecutionRolePolicyAttachment" {
  count       = local.scan_mode == "inAccount" ? 1 : 0
  name = "CloudGuardAWPLambdaExecutionRolePolicyAttachment"
  policy_arn  = aws_iam_policy.CloudGuardAWPLambdaExecutionRolePolicy[count.index].arn
  roles       = [aws_iam_role.CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.name]
}

# The CloudGuardAWPLambdaExecutionRolePolicyAttachment_SaaS resource attaches the CloudGuardAWPLambdaExecutionRolePolicy_SaaS to the CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.
resource "aws_iam_policy_attachment" "CloudGuardAWPLambdaExecutionRolePolicyAttachment_SaaS" {
  count       = local.scan_mode == "saas" ? 1 : 0
  name = "CloudGuardAWPLambdaExecutionRolePolicyAttachment"
  policy_arn  = aws_iam_policy.CloudGuardAWPLambdaExecutionRolePolicy_SaaS[count.index].arn
  roles       = [aws_iam_role.CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.name]
}
# END AWP proxy lambda function role policy

# aws_lambda_invocation : The Lambda invocation that is used to clean up the resources after the onboarding process.
resource "aws_lambda_invocation" "CloudGuardAWPSnapshotsUtilsCleanupFunctionInvocation" {
  function_name = aws_lambda_function.CloudGuardAWPSnapshotsUtilsFunction.function_name
  input = jsonencode({
    "target_account_id" : data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.cloud_account_id
  })
  lifecycle_scope = "CRUD"
  depends_on      = [
    aws_iam_policy_attachment.CloudGuardAWPLambdaExecutionRolePolicyAttachment,
    aws_iam_policy_attachment.CloudGuardAWPLambdaExecutionRolePolicyAttachment_SaaS
  ]
}

# AWP MR key for snapshot re-encryption
# The CloudGuardAWPKey resource defines a KMS key that is used to re-encrypt the snapshots in SaaS mode.
resource "aws_kms_key" "CloudGuardAWPKey" {
  count       = local.scan_mode == "saas" ? 1 : 0
  description          = "CloudGuard AWP Multi-Region primary key for snapshots re-encryption (for Saas mode only)"
  enable_key_rotation  = true
  deletion_window_in_days = 7

  # Conditionally set multi-region based on IsChinaPartition
  multi_region = data.aws_partition.current.partition == "aws-cn" ? false : true

  policy = jsonencode({
    Version = "2012-10-17"
    Id      = "cloud-guard-awp-key"
    Statement = [
      {
        Sid       = "Enable IAM User Permissions"
        Effect    = "Allow"
        Principal = {
          AWS = "arn:${data.aws_partition.current.partition}:iam::${data.aws_caller_identity.current.account_id}:root"
        }
        Action    = "kms:*"
        Resource  = "*"
      },
      {
        Sid       = "Allow usage of the key"
        Effect    = "Allow"
        Principal = {
          AWS = "arn:${data.aws_partition.current.partition}:iam::${local.cloud_guard_backend_account_id}:root"
        }
        Action = [
          "kms:DescribeKey",
          "kms:Encrypt",
          "kms:Decrypt",
          "kms:ReEncrypt*",
          "kms:GenerateDataKey*",
          "kms:PutKeyPolicy",
          "kms:ScheduleKeyDeletion",
          "kms:CancelKeyDeletion",
        ]
        Resource = "*"
      },
      {
        Sid       = "Allow attachment of persistent resources"
        Effect    = "Allow"
        Principal = {
          AWS = "arn:${data.aws_partition.current.partition}:iam::${local.cloud_guard_backend_account_id}:root"
        }
        Action = [
          "kms:CreateGrant",
          "kms:ListGrants",
          "kms:RevokeGrant",
        ]
        Resource = "*"
        Condition = {
          Bool = {
            "kms:GrantIsForAWSResource" = true
          }
        }
      },
    ]
  })
}
#END AWP MR key for snapshot re-encryption

# The CloudGuardAWPKeyAlias resource defines a KMS key alias that is used to reference the KMS key in SaaS mode.
resource "aws_kms_alias" "CloudGuardAWPKeyAlias" {
  count      = local.scan_mode == "saas" ? 1 : 0
  name       = "alias/CloudGuardAWPKey"
  target_key_id = aws_kms_key.CloudGuardAWPKey[count.index].arn
  depends_on = [
    aws_kms_key.CloudGuardAWPKey
  ]
}

# The dome9_awp_aws_onboarding resource defines a Dome9 AWP AWS Onboarding.
# The Dome9 AWP AWS Onboarding resource allows you to onboard an AWS account to Dome9 AWP.
# The cloudguard_account_id attribute is used to specify the CloudGuard account id of the AWS account.
# The cross_account_role_name attribute is used to specify the name of the cross account role that is used to allow AWP to access the AWS account.
# The cross_account_role_external_id attribute is used to specify the external id of the cross account role that is used to allow AWP to access the AWS account.
# The scan_mode attribute is used to specify the scan mode of the Dome9 AWP AWS Onboarding. The valid values are "inAccount" and "saas".
# The agentless_account_settings attribute is used to specify the agentless account settings of the Dome9 AWP AWS Onboarding.
# The disabled_regions attribute is used to specify the disabled regions of the agentless account settings of the Dome9 AWP AWS Onboarding.
# The scan_machine_interval_in_hours attribute is used to specify the scan machine interval in hours of the agentless account settings of the Dome9 AWP AWS Onboarding.
# The max_concurrence_scans_per_region attribute is used to specify the max concurrence scans per region of the agentless account settings of the Dome9 AWP AWS Onboarding.
# The skip_function_apps_scan attribute is used to specify whether to skip the function apps scan of the agentless account settings of the Dome9 AWP AWS Onboarding.
# The custom_tags attribute is used to specify the custom tags of the agentless account settings of the Dome9 AWP AWS Onboarding.
resource "dome9_awp_aws_onboarding" "awp_aws_onboarding_test" {
  cloudguard_account_id = "CLOUDGUARD_ACCOUNT_ID or EXTERNAL_AWS_ACCOUNT_NUMBER"
  cross_account_role_name = aws_iam_role.CloudGuardAWPCrossAccountRole.name
  cross_account_role_external_id = local.cross_account_role_external_id
  scan_mode = local.scan_mode
  agentless_account_settings {
    disabled_regions = ["us-east-1", "us-west-1", "ap-northeast-1", "ap-southeast-2"]
    scan_machine_interval_in_hours = 24
    max_concurrence_scans_per_region = 6
    skip_function_apps_scan = true
    custom_tags = {
      tag1 = "value1"
      tag2 = "value2"
      tag3 = "value3"
    }
  }
  depends_on = [
    aws_iam_policy_attachment.CloudGuardAWPLambdaExecutionRolePolicyAttachment,
    aws_iam_policy_attachment.CloudGuardAWPLambdaExecutionRolePolicyAttachment_SaaS,
    aws_iam_role.CloudGuardAWPCrossAccountRole,
    aws_iam_role_policy_attachment.CloudGuardAWPCrossAccountRoleAttachment
  ]
}

# The dome9_awp_aws_onboarding data source allows you to get the onboarding data of an AWS account.
data "dome9_awp_aws_onboarding" "awp_aws_onboarding_test" {
  id = dome9_awp_aws_onboarding.awp_aws_onboarding_test.cloudguard_account_id
  depends_on = [
    dome9_awp_aws_onboarding.awp_aws_onboarding_test
  ]
}
```

## Argument Reference

The following arguments are supported:

* `cloudguard_account_id` - (Required) The CloudGuard account id.
* `centralized_cloud_account_id` - (Optional) The centralized cloud account id.
* `cross_account_role_name` - (Required) The name of the cross account role.
* `cross_account_role_external_id` - (Required) The external id of the cross account role.
* `scan_mode` - (Required) The scan mode. Valid values are "inAccount", "saas", "inAccountHub", "inAccountSub".
* `agentless_account_settings` - (Optional) The agentless account settings.
  * `disabled_regions` - (Optional) The disabled regions. valid values are "af-south-1", "ap-south-1", "eu-north-1", "eu-west-3", "eu-south-1", "eu-west-2", "eu-west-1", "ap-northeast-3", "ap-northeast-2", "me-south-1", "ap-northeast-1", "me-central-1", "ca-central-1", "sa-east-1", "ap-east-1", "ap-southeast-1", "ap-southeast-2", "eu-central-1", "ap-southeast-3", "us-east-1", "us-east-2", "us-west-1", "us-west-2"
  * `scan_machine_interval_in_hours` - (Optional) The scan machine interval in hours
  * `max_concurrence_scans_per_region` - (Optional) The max concurrence scans per region
  * `skip_function_apps_scan` - (Optional) Whether to skip function apps scan. Default is false.
  * `custom_tags` - (Optional) The custom tags.
* `should_create_policy` - (Optional) Whether to create a policy. Default is true.
    
## Attributes Reference

* `missing_awp_private_network_regions` - The missing AWP private network regions.
* `account_issues` - The account issues.
* `cloud_account_id` - The cloud guard account id.
* `agentless_protection_enabled` - Whether agentless protection is enabled.
* `cloud_provider` - The cloud provider.
* `should_update` - Whether to update.
* `is_org_onboarding` - Whether is org onboarding.

## Import

The AWP AWS Onboarding can be imported; use <ONBOARDING ID> as the import ID.

For example:

```shell
terraform import dome9_awp_aws_onboarding.test_awp_aws_onboarding 00000000-0000-0000-0000-000000000000
```
