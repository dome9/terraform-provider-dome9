package cloudaccounts

const (
	RESTfulPathAWS   = "cloudaccounts"
	RESTfulPathAzure = "AzureCloudAccount"
	RESTfulPathGCP   = "GoogleCloudAccount"
)

// AWS service paths
const (
	RESTfulServicePathAWSName               = "name"
	RESTfulServicePathAWSRegionConfig       = "region-conf"
	RESTfulServicePathAWSOrganizationalUnit = "organizationalUnit"
	RESTfulServicePathAWSCredentials        = "credentials"
	RESTfulServicePathAWSCloudAccounts      = "cloudaccounts"
	RESTfulServicePathAWSIAMSafe            = "iam-safe"
	RESTfulPathRestrictedIamEntities        = "restrictedIamEntities"
	RESTfulPathUser                         = "user"
	RESTfulPathIAM                          = "iam"
	DeleteForce                             = "DeleteForce"
)

// Azure service paths
const (
	RESTfulServicePathAzureName               = "AccountName"
	RESTfulServicePathAzureOperationMode      = "OperationMode"
	RESTfulServicePathAzureOrganizationalUnit = "organizationalUnit"
	RESTfulServicePathAzureCredentials        = "Credentials"
)

// GCP service paths
const (
	RESTfulServicePathGCPName               = "AccountName"
	RESTfulServicePathGCPCredentialsGSuite  = "Credentials/Gsuite"
	RESTfulServicePathGCPOrganizationalUnit = "organizationalUnit"
	RESTfulServicePathGCPCredentials        = "Credentials"
)

type QueryParameters struct {
	ID string
}
