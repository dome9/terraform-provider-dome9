package environmentvariable

// global
const (
	OrganizationalUnitName  	  = "ORGANIZATIONAL_UNIT_NAME"
	UpdatedOrganizationalUnitId   = "UPDATED_ORGANIZATIONAL_UNIT_ID"
	UpdatedOrganizationalUnitName = "UPDATED_ORGANIZATIONAL_UNIT_NAME"
	UpdatedOrganizationalUnitPath = "UPDATED_ORGANIZATIONAL_UNIT_PATH"
)

// AWS environment variable
const (
	CloudAccountAWSEnvVarArn        = "ARN"
	CloudAccountUpdatedAWSEnvVarArn = "ARN_UPDATE"
	CloudAccountAWSEnvVarSecret     = "SECRET"
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
