package variable

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

// k8s resource/data source
const (
	CloudAccountK8SVendor              = "kubernetes"
	CloudAccountK8SOriginalAccountName = "test_cloudaccount_k8s_resource_name"
	CloudAccountK8SUpdatedAccountName  = "test_cloudaccount_k8s_updated_resource_name"
)

// ip list resource/data source
const (
	IPListCreationResourceName      = "test_iplist"
	IPListDescriptionResource       = "acceptance-test"
	IPListUpdateDescriptionResource = "update-acceptance-test"
)

// continuous Compliance Notification resource/data source
const (
	ContinuousComplianceNotificationName               = "test_notification"
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
