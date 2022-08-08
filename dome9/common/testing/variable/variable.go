package variable

// alibaba resource/data source
const (
	CloudAccountAlibabaCreationResourceName = "test_cloudaccount_alibaba"
	CloudAccountAlibabaUpdatedAccountName   = "updated_cloud_account_name"
	CloudAccountAlibabaVendor               = "alibaba"
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
	CloudAccountKubernetesAdmissionControlEnabled  = true
	CloudAccountKubernetesRuntimeProtectionEnabled = true
	CloudAccountKubernetesImageAssuranceEnabled    = true

	// features updates
	CloudAccountKubernetesAdmissionControlUpdateEnabled  = false
	CloudAccountKubernetesRuntimeProtectionUpdateEnabled = false
	CloudAccountKubernetesImageAssuranceUpdateEnabled    = false
)

// ip list resource/data source
const (
	IPListCreationResourceName      = "test_iplist"
	IPListDescriptionResource       = "acceptance-test"
	IPListUpdateDescriptionResource = "update-acceptance-test"
)

// continuous Compliance Notification resource/data source
const (
	ContinuousComplianceNotificationName               = "test_notification_tf"
	ContinuousComplianceNotificationDescription        = "this notification for testing"
	ContinuousComplianceNotificationAlertsConsole      = true
	ContinuousComplianceNotificationEnabled            = "Enabled"
	ContinuousComplianceNotificationDisabled           = "Disabled"
	ContinuousComplianceNotificationCronExpression     = "0 0 10 1/1 * ? *"
	ContinuousComplianceNotificationType               = "Detailed"
	ContinuousComplianceNotificationRecipient          = "test@test.com"
	ContinuousComplianceNotificationJsonWithFullEntity = "JsonWithFullEntity"

	// update const
	ContinuousComplianceNotificationUpdateName          = "test_notification_update"
	ContinuousComplianceNotificationUpdateDescription   = "this notification for update testing"
	ContinuousComplianceNotificationUpdateAlertsConsole = false
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
	AwsUnifiedOnbordingTemplateUrl      = `https://cloudguard-unified-onboarding-common.s3.amazonaws.com/4.5.0/templates/role_based/onboarding.yml`
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

// assessment resource/data source
const (
	BundleID                = 469163
	Dome9CloudAccountID     = "ac08e6a9-64a2-47d3-a3fd-135aa3c062b5"
	CloudAccountID          = "591910677983"
	CloudAccountType        = "Aws"
	RequestID               = "ac08e6a9-64a2-47d3-a3fd-135aa3c062cc"
	ShouldMinimizeResult    = false
	TriggeredBy             = "Manual"
	AssessmentPassed        = "false"
	HasErrors               = "false"
	HasDataSyncStatusIssues = "true"
)
