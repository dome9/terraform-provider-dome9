# Dome9 Provider Configurations
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
	}
}

provider "dome9" {
	dome9_access_id     = "DOME9_ACCESS_ID"
	dome9_secret_key    = "DOME9_SECRET_KEY"
	base_url            = "https://api.us7.falconetix.com/v2/"
}

provider "aws" {
	region     = "us-west-2"
	access_key = ""
	secret_key = ""
	token      = ""
}

resource "dome9_cloudaccount_aws" "aws_onboarding_account_test" {
	name  = "aws_onboarding_account_test"
	credentials  {
		arn    = "ARN for IAM Role"
		secret = "Secret for IAM Role"
		type   = "RoleBased"
	}
	net_sec {
		regions {
			new_group_behavior = "ReadOnly"
			region             = "us_west_2"
		}
	}
}

data "dome9_awp_aws_get_onboarding_data" "dome9_awp_aws_onboarding_data_source" {
	cloud_account_id = dome9_cloudaccount_aws.aws_onboarding_account_test.external_account_number
	depends_on = [
		dome9_cloudaccount_aws.aws_onboarding_account_test
	]
}

locals {
	scan_mode = "saas"
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

resource "aws_iam_role_policy_attachment" "CloudGuardAWPCrossAccountRoleAttachment" {
	role       = aws_iam_role.CloudGuardAWPCrossAccountRole.name
	policy_arn = aws_iam_policy.CloudGuardAWP.arn
}
# end resources for CloudGuardAWPCrossAccountRole

# Cross account role policy
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

resource "aws_iam_policy_attachment" "CloudGuardAWPCrossAccountRolePolicyAttachment" {
	count       = local.scan_mode == "inAccount" ? 1 : 0
	name       = "CloudGuardAWPCrossAccountRolePolicyAttachment"
	policy_arn  = aws_iam_policy.CloudGuardAWPCrossAccountRolePolicy[count.index].arn
	roles       = [aws_iam_role.CloudGuardAWPCrossAccountRole.name]
}

resource "aws_iam_policy_attachment" "CloudGuardAWPCrossAccountRolePolicyAttachment_SaaS" {
	count = local.scan_mode == "saas" ? 1 : 0
	name       = "CloudGuardAWPCrossAccountRolePolicyAttachment_SaaS"
	policy_arn  = aws_iam_policy.CloudGuardAWPCrossAccountRolePolicy_SaaS[count.index].arn
	roles       = [aws_iam_role.CloudGuardAWPCrossAccountRole.name]
}
# END Cross account role policy

# AWP proxy lambda function
resource "aws_lambda_function" "CloudGuardAWPSnapshotsUtilsFunction" {
	function_name    = local.remote_snapshots_utils_function_name
	handler          = "snapshots_utils.lambda_handler"
	description      = "CloudGuard AWP Proxy for managing remote actions and resources"
	role             = aws_iam_role.CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.arn
	runtime          = "python3.9"
	memory_size      = 256
	timeout          = local.remote_snapshots_utils_function_time_out
	s3_bucket        = local.agentless_bucket_name
	s3_key           = "${local.remote_functions_prefix_key}/${local.remote_snapshots_utils_function_name}7.zip"

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

resource "aws_cloudwatch_log_group" "CloudGuardAWPSnapshotsUtilsLogGroup" {
	name              = "/aws/lambda/CloudGuardAWPSnapshotsUtils"
	retention_in_days = 30
	depends_on = [
		aws_lambda_function.CloudGuardAWPSnapshotsUtilsFunction
	]
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

resource "aws_iam_policy_attachment" "CloudGuardAWPLambdaExecutionRolePolicyAttachment" {
	count       = local.scan_mode == "inAccount" ? 1 : 0
	name = "CloudGuardAWPLambdaExecutionRolePolicyAttachment"
	policy_arn  = aws_iam_policy.CloudGuardAWPLambdaExecutionRolePolicy[count.index].arn
	roles       = [aws_iam_role.CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.name]
}

resource "aws_iam_policy_attachment" "CloudGuardAWPLambdaExecutionRolePolicyAttachment_SaaS" {
	count       = local.scan_mode == "saas" ? 1 : 0
	name = "CloudGuardAWPLambdaExecutionRolePolicyAttachment"
	policy_arn  = aws_iam_policy.CloudGuardAWPLambdaExecutionRolePolicy_SaaS[count.index].arn
	roles       = [aws_iam_role.CloudGuardAWPSnapshotsUtilsLambdaExecutionRole.name]
}
# END AWP proxy lambda function role policy

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

resource "aws_kms_alias" "CloudGuardAWPKeyAlias" {
	count      = local.scan_mode == "saas" ? 1 : 0
	name       = "alias/CloudGuardAWPKey"
	target_key_id = aws_kms_key.CloudGuardAWPKey[count.index].arn
	depends_on = [
		aws_kms_key.CloudGuardAWPKey
	]
}


resource "dome9_awp_aws_onboarding" "awp_aws_onboarding_test" {
	cloudguard_account_id = dome9_cloudaccount_aws.aws_onboarding_account_test.id
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

data "dome9_awp_aws_onboarding" "awp_aws_onboarding_test" {
	id = dome9_awp_aws_onboarding.awp_aws_onboarding_test.cloudguard_account_id
	depends_on = [
		dome9_awp_aws_onboarding.awp_aws_onboarding_test
	]
}