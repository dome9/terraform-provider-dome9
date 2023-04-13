package cloudaccounts

const (
	RESTfulPathAWS     = "cloudaccounts"
	RESTfulPathAzure   = "AzureCloudAccount"
	RESTfulPathGCP     = "GoogleCloudAccount"
	RESTfulPathK8S     = "KubernetesAccount"
	RESTfulPathAlibaba = "AlibabaCloudAccount"
	RESTfulPathOci     = "oci-cloud-account"
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
	//runtime-protection
	RESTfulPathK8SRuntimeProtection = "runtime-protection"
	//admission-control
	RESTfulPathK8SAdmissionControl = "admission-control"
	//image-assurance
	RESTfulPathK8SImageAssurance = "vulnerabilityAssessment"
	//threat-intelligence
	RESTfulPathK8SThreatIntelligence = "magellan-kubernetes-flowlogs"
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

type QueryParameters struct {
	ID string
}
