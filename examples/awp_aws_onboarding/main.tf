
# Dome9 Provider Configurations
terraform {
	required_providers {
		dome9 = {
			source = "dome9/dome9"
			version = "1.29.6"
		}
	}
}

provider "dome9" {
	dome9_access_id     = "DOME9_CLOUDGUARD_API_ACCESS_ID"
	dome9_secret_key    = "DOME9_CLOUDGUARD_API_SECRET_KEY"
	base_url            = "DOME9_API_BASE_URL"
}

terraform {
	required_providers {
		aws = {
			source = "hashicorp/aws"
			version = "5.37.0"
		}
	}
}

provider "aws" {
	region     = "us-east-1"
	profile = "custom"
}

data "dome9_awp_aws_get_onboarding_data" "dome9_awp_aws_onboarding_data_source" {
	cloudguard_account_id = "ae481d4a-603b-4fa6-8f31-6c6d57920e96"
	scan_mode = "inAccount"
}

#onboarding Enable/Disable AWP on AWS Account

resource "dome9_awp_aws_onboarding" "awp_onboarding_on_aws" {
	cloudguard_account_id = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode
	cross_account_role_name = "CloudGuardAWPCrossAccountRole" # default value can be applicable
	# should be similar to "NjM0NzI5NTk3NjIzLWFlNDgxZDRhLTYwM2ItNGZhNi04ZjMxLTZjNmQ1NzkyMGU5Ng=="
	cross_account_external_id = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.cross_account_external_id
	scan_mode = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode
	should_create_policy = false
	force_delete = false
	account_settings = {

	}
	version = ""
	# Add depends_on to ensure this resource is created last
	depends_on = [
		aws_iam_role_policy_attachment.CloudGuardAWPCrossAccountRoleAttachment
	]
}

data "aws_partition" "current" {}

data "aws_region" "current" {}

data "aws_caller_identity" "current" {}

# Cross account role to allow CloudGuard access
resource "aws_iam_role" "CloudGuardAWPCrossAccountRole" {
	name               = "CloudGuardAWPCrossAccountRole"
	description        = "CloudGuard AWP Cross Account Role"
	assume_role_policy = jsonencode({
		Version   = "2012-10-17"
		Statement = [{
			Effect    = "Allow"
			Principal = {
				AWS = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.d9_aws_account_id
			}
			Action    = "sts:AssumeRole"
			Condition = {
				StringEquals = {
					"sts:ExternalId" = "${data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.cross_account_external_id}"
				}
			}
		}]
	})

	depends_on = [aws_lambda_function.CloudGuardAWPSnapshotsUtilsFunction]
}

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
				Resource = "arn:${data.aws_partition.current.partition}:cloudformation:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:stack/stackName/*"
			},
			{
				Effect   = "Allow"
				Action   = "s3:GetObject"
				Resource = "arn:${data.aws_partition.current.partition}:s3:::agentless-prod-us/remote_functions*"
			}
		]
	})
}

resource "aws_iam_role_policy_attachment" "CloudGuardAWPCrossAccountRoleAttachment" {
	role       = aws_iam_role.CloudGuardAWPCrossAccountRole.name
	policy_arn = aws_iam_policy.CloudGuardAWP.arn
}
# end resources for CloudGuardAWPCrossAccountRole

# Cross account role policy
resource "aws_iam_policy" "CloudGuardAWPCrossAccountRolePolicy" {
	count = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode == "inAccount" ? 1 : 0
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

resource "aws_iam_policy" "CloudGuardAWPCrossAccountRolePolicy_SaaS" {
	count = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode == "saas" ? 1 : 0
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
				Resource = [aws_kms_key.CloudGuardAWPKey.arn]
			},
			{
				Effect = "Allow"
				Action = [
					"kms:PutKeyPolicy",
					"kms:ScheduleKeyDeletion",
					"kms:CancelKeyDeletion",
					"kms:TagResource",
				]
				Resource = aws_kms_key.CloudGuardAWPKey.arn
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

resource "aws_iam_policy_attachment" "CloudGuardAWPCrossAccountRolePolicyAttachment" {
	count       = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode == "inAccount" ? 1 : 0
	name       = "CloudGuardAWPCrossAccountRolePolicyAttachment"
	policy_arn  = aws_iam_policy.CloudGuardAWPCrossAccountRolePolicy[count.index].arn
	roles       = [aws_iam_role.CloudGuardAWPCrossAccountRole.name]
}

resource "aws_iam_policy_attachment" "CloudGuardAWPCrossAccountRolePolicyAttachment_SaaS" {
	count = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode == "saas" ? 1 : 0
	name       = "CloudGuardAWPCrossAccountRolePolicyAttachment_SaaS"
	policy_arn  = aws_iam_policy.CloudGuardAWPCrossAccountRolePolicy_SaaS[count.index].arn
	roles       = [aws_iam_role.CloudGuardAWPCrossAccountRole.name]
}
# END Cross account role policy

# AWP proxy lambda function
resource "aws_lambda_function" "CloudGuardAWPSnapshotsUtilsFunction" {
	function_name    = "CloudGuardAWPSnapshotsUtils"
	handler          = "snapshots_utils.lambda_handler"
	description      = "CloudGuard AWP Proxy for managing remote actions and resources"
	role             = aws_iam_role.CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.arn
	runtime          = "python3.9"
	memory_size      = 256
	timeout          = 900
	s3_bucket        = "agentless-prod-us"
	s3_key           = "remote_functions/CloudGuardAWPSnapshotsUtils7.zip"

	environment {
		variables = {
			CP_AWP_AWS_ACCOUNT        = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.d9_aws_account_id
			CP_AWP_MR_KMS_KEY_ID      = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode == "saas" ? aws_kms_key.CloudGuardAWPKey.arn : ""
			CP_AWP_SCAN_MODE          = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode
			CP_AWP_SECURITY_GROUP_NAME = "CloudGuardAWPSecurityGroup"
			AWS_PARTITION             = data.aws_partition.current.partition
		}
	}

	tags = {
		Owner = "CG.AWP"
	}

	# Use provisioners to invoke the Lambda function after creation and destruction
	# Define Provisioners to do some equivalent to AWS Custom Resource
	# Create Provisioner is not relevant
	provisioner "local-exec" {
		when    = "create"
		command = <<EOF
aws lambda invoke \
    --function-name ${aws_lambda_function.CloudGuardAWPSnapshotsUtilsFunction.function_name} \
    --invocation-type RequestResponse \
    --payload '{"action": "create"}' \
    /dev/null
EOF
	}

	provisioner "local-exec" {
		when    = "destroy"
		command = <<EOF
aws lambda invoke \
    --function-name ${aws_lambda_function.CloudGuardAWPSnapshotsUtilsFunction.function_name} \
    --invocation-type RequestResponse \
    --payload '{"action": "destroy"}' \
    /dev/null
EOF
	}
}

resource "aws_lambda_permission" "allow_cloudguard" {
	statement_id  = "AllowExecutionFromCloudGuard"
	action        = "lambda:InvokeFunction"
	function_name = aws_lambda_function.CloudGuardAWPSnapshotsUtilsFunction.function_name
	principal     = "s3.amazonaws.com"
	source_arn    = "arn:${data.aws_partition.current.partition}:s3:::agentless-prod-us/*"
}
# END AWP proxy lambda function

resource "aws_cloudwatch_log_group" "CloudGuardAWPSnapshotsUtilsLogGroup" {
	name              = "/aws/lambda/CloudGuardAWPSnapshotsUtils"
	retention_in_days = 30
}

# AWP proxy lambda function role
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

resource "aws_iam_role_policy_attachment" "CloudGuardAWPSnapshotsUtilsLambdaExecutionRoleAttachment" {
	role       = aws_iam_role.CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.name
	policy_arn = aws_iam_policy.CloudGuardAWPSnapshotsPolicy.arn
}
# END AWP proxy lambda function role

# AWP proxy lambda function role policy
resource "aws_iam_policy" "CloudGuardAWPLambdaExecutionRolePolicy" {
	count       = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode == "inAccount" ? 1 : 0
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
				Condition = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode == "inAccount" ? {
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
				Condition = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode == "inAccount" ? {
					StringEquals = {
						"aws:ResourceTag/Owner" = "CG.AWP"
					}
				} : null
			},
		]
	})
}

resource "aws_iam_policy" "CloudGuardAWPLambdaExecutionRolePolicy_SaaS" {
	count       = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode == "saas" ? 1 : 0
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

resource "aws_iam_policy_attachment" "CloudGuardAWPLambdaExecutionRolePolicyAttachment" {
	count       = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode == "inAccount" ? 1 : 0
	name = "CloudGuardAWPLambdaExecutionRolePolicyAttachment"
	policy_arn  = aws_iam_policy.CloudGuardAWPLambdaExecutionRolePolicy[count.index].arn
	roles       = [aws_iam_role.CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.name]
}

resource "aws_iam_policy_attachment" "CloudGuardAWPLambdaExecutionRolePolicyAttachment_SaaS" {
	count       = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode == "saas" ? 1 : 0
	name = "CloudGuardAWPLambdaExecutionRolePolicyAttachment"
	policy_arn  = aws_iam_policy.CloudGuardAWPLambdaExecutionRolePolicy_SaaS[count.index].arn
	roles       = [aws_iam_role.CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.name]
}
# END AWP proxy lambda function role policy

# AWP MR key for snapshot re-encryption
resource "aws_kms_key" "CloudGuardAWPKey" {
	count       = data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.scan_mode == "saas" ? 1 : 0

	description          = "CloudGuard AWP Multi-Region primary key for snapshots re-encryption (for Saas mode only)"
	enable_key_rotation  = true
	pending_window_in_days = 7

	# Conditionally set multi-region based on IsChinaPartition
	multi_region         = data.aws_partition.current.partition == "aws-cn" ? false : true

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
					AWS = "arn:${data.aws_partition.current.partition}:iam::${data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.d9_aws_account_id}:root"
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
					AWS = "arn:${data.aws_partition.current.partition}:iam::${data.dome9_awp_aws_get_onboarding_data.dome9_awp_aws_onboarding_data_source.d9_aws_account_id}:root"
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