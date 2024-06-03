package cloudaccounts

// Single onboarding
const (
	RESTfulPathAWS     = "cloudaccounts"
	RESTfulPathAzure   = "AzureCloudAccount"
	RESTfulPathGCP     = "GoogleCloudAccount"
	RESTfulPathK8S     = "kubernetes/account"
	RESTfulPathAlibaba = "AlibabaCloudAccount"
	RESTfulPathOci     = "oci-cloud-account"
)

// Organization onboarding
const (
	RESTfulServicePathAwsOrgMgmtOnboarding = "aws-organization-management-onboarding"
	RESTfulServicePathAwsOrgMgmt           = "aws-organization-management"
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

// K8S service paths
const (
	RESTfulServicePathK8SName               = "AccountName"
	RESTfulServicePathK8SOrganizationalUnit = "organizationalUnit"
	RESTfulPathK8sEnable                    = "enable"
	RESTfulPathK8sDisable                   = "disable"
	//runtime-protection
	RESTfulPathK8SRuntimeProtection = "runtimeProtection"
	//admission-control
	RESTfulPathK8SAdmissionControl = "admissionControl"
	//image-assurance
	RESTfulPathK8SImageAssurance = "imageAssurance"
	//threat-intelligence
	RESTfulPathK8SThreatIntelligence = "threatIntelligence"
)

// Alibaba service paths
const (
	RESTfulServicePathAlibabaName               = "AccountName"
	RESTfulServicePathAlibabaOrganizationalUnit = "organizationalUnit"
	RESTfulServicePathAlibabaCredentials        = "Credentials"
)

// Oci service paths
const (
	RESTfulServicePathOciTempData           = "save-temp-data"
	RESTfulServicePathOciOrganizationalUnit = "organizational-Unit"
)

// AWS organization onboarding service paths
const (
	RESTfulServicePathAwsOrgMgmtOnboardingMgmtStack          = "management-stack"
	RESTfulServicePathAwsOrgMgmtOnboardingMemberAccountStack = "member-account-configuration"
	RESTfulServicePathAwsOrgMgmtStacksetArn                  = "stackset-arn"
	RESTfulServicePathAwsOrgMgmtConfiguration                = "configuration"
)

type QueryParameters struct {
	ID string
}
