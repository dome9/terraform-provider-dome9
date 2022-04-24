resource "dome9_aws_unified_onboarding" "aws_unified_onboarding" {
}

resource "aws_cloudformation_stack" "stack"{
	name = dome9_aws_unified_onboarding.aws_unified_onboarding.stack_name
	template_url = dome9_aws_unified_onboarding.aws_unified_onboarding.template_url
	parameters = dome9_aws_unified_onboarding.aws_unified_onboarding.parameters
	capabilities = dome9_aws_unified_onboarding.aws_unified_onboarding.iam_capabilities
}