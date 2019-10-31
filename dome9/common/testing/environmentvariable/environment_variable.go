package environmentvariable

// global
const (
	OrganizationalUnitName = "ORGANIZATIONAL_UNIT_NAME"
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

// Continuous Compliance Policy environment variable
const (
	ContinuousCompliancePolicyEnvVarNotificationId1 = "CONTINUOUS_COMPLIANCE_NOTIFICATION_ID1"
	ContinuousCompliancePolicyEnvVarNotificationId2 = "CONTINUOUS_COMPLIANCE_NOTIFICATION_ID2"
)
