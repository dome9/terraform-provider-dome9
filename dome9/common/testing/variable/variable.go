package variable

// alibaba resource/data source
const (
	CloudAccountAlibabaCreationResourceName = "test_cloudaccount_alibaba"
	CloudAccountAlibabaUpdatedAccountName   = "updated_cloud_account_name"
	CloudAccountAlibabaVendor               = "alibaba"
)

// oci resource/data source
const (
	CloudAccountOciCreationResourceName = "test_cloudaccount_oci"
	CloudAccountOciVendor               = "oci"
)

// aws resource/data source
const (
	CloudAccountAWSCreationResourceName    = "test_cloudaccount_aws"
	CloudAccountAWSVendor                  = "aws"
	CloudAccountAWSOriginalAccountName     = "original_cloud_account_name_before_change"
	CloudAccountAWSUpdatedAccountName      = "updated_cloud_account_name"
	CloudAccountAWSFetchedRegion           = "us_east_1"
	CloudAccountAWSReadOnlyGroupBehavior   = "ReadOnly"
	CloudAccountAWSFullManageGroupBehavior = "FullManage"
)

// azure resource/data source
const (
	CloudAccountAzureCreationResourceName = "test_cloudaccount_azure"
	CloudAccountAzureOperationMode        = "Read"
	CloudAccountAzureVendor               = "azure"

	// update const
	CloudAccountAzureUpdatedAccountName  = "updated_cloud_account_name"
	CloudAccountAzureUpdateOperationMode = "Manage"
)

// gcp resource/data source
const (
	CloudAccountGCPCreationResourceName = "test_cloudaccount_gcp"
	CloudAccountGCPUpdatedAccountName   = "updated_cloud_account_name"
	CloudAccountGCPVendor               = "google"
)

// kubernetes resource/data source
const (
	CloudAccountKubernetesVendor              = "kubernetes"
	CloudAccountKubernetesOriginalAccountName = "test_cloudaccount_kubernetes_resource_name"
	CloudAccountKubernetesUpdatedAccountName  = "test_cloudaccount_kubernetes_updated_resource_name"

	// features
	CloudAccountKubernetesAdmissionControlEnabled   = true
	CloudAccountKubernetesRuntimeProtectionEnabled  = true
	CloudAccountKubernetesImageAssuranceEnabled     = true
	CloudAccountKubernetesThreatIntelligenceEnabled = true

	// features updates
	CloudAccountKubernetesAdmissionControlUpdateEnabled   = false
	CloudAccountKubernetesRuntimeProtectionUpdateEnabled  = false
	CloudAccountKubernetesImageAssuranceUpdateEnabled     = false
	CloudAccountKubernetesThreatIntelligenceUpdateEnabled = false
)

// ip list resource/data source
const (
	IPListCreationResourceName      = "test_iplist"
	IPListDescriptionResource       = "acceptance-test"
	IPListUpdateDescriptionResource = "update-acceptance-test"
)

// continuous Compliance Notification resource/data source
const (
	ContinuousComplianceNotificationName               = "test_notification_tf_old"
	ContinuousComplianceNotificationDescription        = "this notification for testing_old"
	ContinuousComplianceNotificationAlertsConsole      = true
	ContinuousComplianceNotificationEnabled            = "Enabled"
	ContinuousComplianceNotificationDisabled           = "Disabled"
	ContinuousComplianceNotificationCronExpression     = "0 0 10 1/1 * ? *"
	ContinuousComplianceNotificationType               = "Detailed"
	ContinuousComplianceNotificationRecipient          = "test-old@test.com"
	ContinuousComplianceNotificationJsonWithFullEntity = "JsonWithFullEntity"

	// update const
	ContinuousComplianceNotificationUpdateName          = "test_notification_update_old"
	ContinuousComplianceNotificationUpdateDescription   = "this notification for update testing old"
	ContinuousComplianceNotificationUpdateAlertsConsole = true
)

// Notification resource/data source
const (
	NotificationName                 = "test_notification_tf"
	NotificationDescription          = "this notification for testing"
	NotificationAlertsConsole        = true
	NotificationSendOnEachOccurrence = true
	NotificationCronExpression       = "0 0 10 1/1 * ? *"
	NotificationIntegrationID        = "9c792681-e25e-4dcf-8009-7fb8cc1279b3"
	NotificationOutputType           = "Default"
	NotificationJsonWithFullEntity   = "JsonWithFullEntity"

	// update const
	NotificationUpdateName          = "test_notification_update"
	NotificationUpdateDescription   = "this notification for update testing"
	NotificationUpdateAlertsConsole = true
)

// Integration resource/data source
const (
	IntegrationName              = "test_integration_tf"
	IntegrationType              = "Webhook"
	IntegrationUrl               = "https://webhook.site/test"
	IntegrationMethodType        = "Post"
	IntegrationAuthType          = "BasicAuth"
	IntegrationUsername          = "yogev"
	IntegrationPassword          = "password"
	IntegrationIgnoreCertificate = true

	// update const
	IntegrationUpdateName              = "test_integration_update"
	IntegrationUpdateUrl               = "https://webhook.site/test"
	IntegrationUpdateMethodType        = "Post"
	IntegrationUpdateAuthType          = "BasicAuth"
	IntegrationUpdateUsername          = "usr"
	IntegrationUpdatePassword          = "password123"
	IntegrationUpdateIgnoreCertificate = false
)

// continuous Compliance Policy resource/data source
const (
	ContinuousCompliancePolicyRulesetId = -14
)

// ruleset resource/data source
const (
	RuleSetName              = "test_rule_set"
	RuleSetDescription       = "this is acceptance test"
	RuleSetDescriptionUpdate = "this is acceptance test"
)

// aws security group resource/data source
const (
	AWSSecurityGroupDescription   = "this is aws security group test"
	AWSSecurityGroupRegionID      = "us_east_1"
	WaitUntilSecurityGroupCreated = 45

	// Update
	AWSSecurityGroupTagValue = "value"
)

// azure security group resource/data source
const (
	AzureSecurityGroupDescription       = "this is azure security group test"
	AzureSecurityGroupRegion            = "australiaeast"
	AzureSecurityGroupTagValue          = "tag_val_1"
	AzureSecurityGroupIsTamperProtected = false

	// 	update const
	AzureSecurityGroupUpdateDescription       = "this is azure security group update test"
	AzureSecurityGroupUpdateIsTamperProtected = true
	AzureSecurityGroupUpdateTagValue          = "val"
)

// role resource/data source
const (
	RoleName                    = "test_role"
	RoleDescription             = "this is role test"
	RoleToPermittedAlertActions = false

	// update const
	RoleUpdateDescription             = "this is update role test"
	RoleUpdateToPermittedAlertActions = true
)

// AwsUnifiedOnbording resource/data
const (
	AwsUnifiedOnbordingOnboardType                    = "Simple"
	AwsUnifiedOnbordingFullProtection                 = "true"
	AwsUnifiedOnbordingCloudVendor                    = "aws"
	AwsUnifiedOnbordingEnableStackModify              = "true"
	AwsUnifiedOnbordingPostureManagementConfiguration = `{rulesets: "[0]"}`

	AwsUnifiedOnbordingServerlessConfiguration    = `{	enabled: true}`
	AwsUnifiedOnbordingIntelligenceConfigurations = `{
	enabled: true
	rulesets: "[0]"
	}`

	DataSourceSuffix                    = "Data"
	AwsUnifiedOnbordingIamCapabilities0 = `CAPABILITY_IAM`
	AwsUnifiedOnbordingIamCapabilities1 = `CAPABILITY_NAMED_IAM`
	AwsUnifiedOnbordingIamCapabilities2 = `CAPABILITY_AUTO_EXPAND`
)

// organizational unit resource/data source
const (
	OrganizationalUnitName       = "test_ou"
	OrganizationalUnitNameUpdate = "test_ou_update"
	ParentID                     = "" // empty string as parent id creates ou under Dome9 main root ou
)

// users resource/data source
const (
	UserFirstName    = "first_name_for_test"
	UserLastName     = "last_name_for_test"
	UserIsSsoEnabled = false
)

// iam entity resource
const (
	IAMSafeEntityProtect       = "Protect"
	IAMSafeEntityTypeUser      = "User"
	IAMSafeEntityName          = "user_for_testing_dont_remove"
	WaitUntilAttachIAMSafeDone = 300

	// 	update const
	IAMSafeEntityProtectWithElevation = "ProtectWithElevation"
)

// service account entity resource
const (
	ServiceAccountName       = "serviceAccount"
	ServiceAccountNameUpdate = "serviceAccountUpdate"
)

// Admission Control Policy resource
const (
	AdmissionControlPolicyDefaultRulesetId = -2001
	AdmissionControlPolicyDetectAction     = "Detection"
	AdmissionControlPolicyPreventAction    = "Prevention"
	AdmissionControlPolicyTargetType       = "Environment"
	AdmissionControlKubernetesAccountName  = "test_admission_control_kubernetes_resource_name"
)

const (
	ImageAssurancePolicyDefaultRulesetId = -2002
	ImageAssurancePolicyDetectAction     = "Detection"
	ImageAssurancePolicyPreventAction    = "Prevention"
	ImageAssurancePolicyTargetType       = "Environment"
	ImageAssuranceKubernetesAccountName  = "test_image_assurance_policy_kubernetes_resource_name"
)

const (
	VulnerabilityPolicyDefaultRulesetId = -2002
	VulnerabilityPolicyDetectAction     = "Detection"
	VulnerabilityPolicyPreventAction    = "Prevention"
	VulnerabilityPolicyTargetType       = "Environment"
	VulnerabilityKubernetesAccountName  = "test_vulnerability_policy_kubernetes_resource_name"
)

// assessment resource/data source
const (
	BundleID                = 1155085
	CloudAccountID          = "dd7678bf-04d9-49ab-a653-1e03ced3726c"
	CloudAccountType        = "Azure"
	RequestID               = "c7475133-33a9-4a48-8601-dc2ecf407453"
	TriggeredBy             = "Manual"
	AssessmentPassed        = "true"
	HasErrors               = "false"
	HasDataSyncStatusIssues = "false"
)

// awp aws onboarding data resource/data source
const (
	OnboardedAwsCloudGuardAccountID   = "b4b6e1dd-e405-47a1-b2e5-6cc49bfbad00"
	AwpAwsCrossAccountRoleName        = "CloudGuardAWPCrossAccountRole"
	ScanMode                          = "inAccount"
	DisabledRegions                   = `["us-east-1", "us-west-1"]`
	DisabledRegionsUpdate             = `["us-east-1", "us-west-1", "ap-northeast-1", "ap-southeast-2"]`
	ScanMachineIntervalInHours        = "6"
	ScanMachineIntervalInHoursUpdate  = "11"
	MaxConcurrentScansPerRegion       = "4"
	MaxConcurrentScansPerRegionUpdate = "8"
	CustomTags                        = `{
			tag1 = "value1"
			tag2 = "value2"
		}`
	CustomTagsUpdate = `{
			tag1 = "value1"
			tag2 = "value2"
			tag3 = "value3"
		}`
)

// aws organization onboarding resource/data source
const (
	AwsOrganizationOnboardingCreationResourceName = "TestAwsOrganizationOnboarding"
)

// aws organization onboarding management stack data source
const (
	AwsFakeAccountId = "111111111111"
)

// awp azure onboarding data resource/data source
const (
	OnboardedAzureCloudGuardAccountID   = "dd7678bf-04d9-49ab-a653-1e03ced3726c"
	AzureDisabledRegions                   = `["eastus", "eastus2"]`
	AzureDisabledRegionsUpdate             = `["eastus", "eastus2", "westus", "westus2"]`
)
