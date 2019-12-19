package variable

// AWS resource/data source
const (
	CloudAccountAWSCreationResourceName    = "test_cloudaccount_aws"
	CloudAccountAWSVendor                  = "aws"
	CloudAccountAWSOriginalAccountName     = "original_cloud_account_name_before_change"
	CloudAccountAWSUpdatedAccountName      = "updated_cloud_account_name"
	CloudAccountAWSFetchedRegion           = "us_east_1"
	CloudAccountAWSReadOnlyGroupBehavior   = "ReadOnly"
	CloudAccountAWSFullManageGroupBehavior = "FullManage"
)

// Azure resource/data source
const (
	CloudAccountAzureCreationResourceName = "test_cloudaccount_azure"
	CloudAccountAzureOperationMode        = "Read"
	CloudAccountAzureVendor               = "azure"

	// Update const
	CloudAccountAzureUpdatedAccountName  = "updated_cloud_account_name"
	CloudAccountAzureUpdateOperationMode = "Manage"
)

// GCP resource/data source
const (
	CloudAccountGCPCreationResourceName = "test_cloudaccount_gcp"
	CloudAccountGCPUpdatedAccountName   = "updated_cloud_account_name"
	CloudAccountGCPVendor               = "google"
)

// IpList resource/data source
const (
	IPListCreationResourceName      = "test_iplist"
	IPListDescriptionResource       = "acceptance-test"
	IPListUpdateDescriptionResource = "update-acceptance-test"
)

// Continuous Compliance Notification resource/data source
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

	// Update const
	ContinuousComplianceNotificationUpdateName          = "test_notification_update"
	ContinuousComplianceNotificationUpdateDescription   = "this notification for update testing"
	ContinuousComplianceNotificationUpdateAlertsConsole = false
)

// Ruleset resource/data source
const (
	RuleSetName              = "test_rule_set"
	RuleSetDescription       = "this is acceptance test"
	RuleSetDescriptionUpdate = "this is acceptance test"
)

// AWS security group resource/data source
const (
	AWSSecurityGroupDescription = "this is aws security group test"
	AWSSecurityGroupRegionID    = "us_east_1"
)

// Azure security group resource/data source
const (
	AzureSecurityGroupDescription       = "this is azure security group test"
	AzureSecurityGroupRegion            = "australiaeast"
	AzureSecurityGroupTagValue          = "tag_val_1"
	AzureSecurityGroupIsTamperProtected = false
)

// Role resource/data source
const (
	RoleName        = "test_role"
	RoleDescription = "this is role test"

	// Update const
	RoleUpdateDescription = "this is update role test"
)

// Organizational unit resource/data source
const (
	OrganizationalUnitName       = "test_ou"
	OrganizationalUnitNameUpdate = "test_ou_update"
	ParentID                     = "" // empty string as parent id creates ou under Dome9 main root ou
)

// Users resource/data source
const (
	UserFirstName    = "first_name_for_test"
	UserLastName     = "last_name_for_test"
	UserIsSsoEnabled = false
)
