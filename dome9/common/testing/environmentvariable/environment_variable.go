package environmentvariable

// global
const (
	OrganizationalUnitName = "ORGANIZATIONAL_UNIT_NAME"
)

// Alibaba environment variable
const (
	CloudAccountAlibabaEnvVarAccessKey    = "ALIBABA_ACCESS_KEY"
	CloudAccountAlibabaEnvVarAccessSecret = "ALIBABA_ACCESS_SECRET"
)

// Oci environment variable
const (
	CloudAccountOciEnvVarTenancyId  = "OCI_TENANCY_ID"
	CloudAccountOciEnvVarHomeRegion = "OCI_HOME_REGION"
	CloudAccountOciEnvVarUserOcid   = "OCI_USER_OCID"
)

// AWS environment variable
const (
	CloudAccountAWSEnvVarArn               = "ARN"
	CloudAccountUpdatedAWSEnvVarArn        = "ARN_UPDATE"
	CloudAccountAWSEnvVarSecret            = "SECRET"
	AwpAwsCrossAccountRoleExternalIdEnvVar = "AWP_EXTERNAL_ID"
)

// Azure environment variable
const (
	CloudAccountAzureEnvVarClientId       = "CLIENT_ID"
	CloudAccountAzureEnvVarSubscriptionId = "SUBSCRIPTION_ID"
	CloudAccountAzureEnvVarClientPassword = "CLIENT_PASSWORD"
	CloudAccountAzureEnvVarTenantId       = "TENANT_ID"
)

// GCP environment variable
const (
	CloudAccountGCPEnvVarProjectId         = "PROJECT_ID"
	CloudAccountGCPEnvVarPrivateKeyId      = "PRIVATE_KEY_ID"
	CloudAccountGCPEnvVarPrivateKey        = "PRIVATE_KEY"
	CloudAccountGCPEnvVarClientEmail       = "CLIENT_EMAIL"
	CloudAccountGCPEnvVarClientId          = "CLIENT_ID"
	CloudAccountGCPEnvVarClientX509CertUrl = "CLIENT_X509_CERT_URL"
)

// Azure security group
const (
	AzureSecurityGroupResourceGroup = "AZURE_RESOURCE_GROUP"
)

// Attach IAM safe
const (
	AttachIAMSafeEnvVarGroupArn  = "ATTACH_IAM_SAFE_GROUP_ARN"
	AttachIAMSafeEnvVarPolicyArn = "ATTACH_IAM_SAFE_POLICY_ARN"
)

// Aws organization onboarding environment variable
const (
	AwsOrganizationOnboardingEnvVarRoleArn     = "AWS_ORG_ROLE_ARN"
	AwsOrganizationOnboardingEnvVarSecret      = "AWS_ORG_SECRET"
	AwsOrganizationOnboardingEnvVarStackSetArn = "AWS_ORG_STACK_SET_ARN"
)

// Azure organization onboarding environment variable
const (
	CloudAccountAzureEnvVarManagementGroupId = "MGMT_GROUP_ID"
)
