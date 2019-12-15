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
	CloudAccountAzureUpdatedAccountName   = "updated_cloud_account_name"
	CloudAccountAzureOperationMode        = "Read"
	CloudAccountAzureVendor               = "azure"
)

// GCP resource/data source
const (
	CloudAccountGCPCreationResourceName    = "test_cloudaccount_gcp"
	CloudAccountGCPUpdatedAccountName      = "updated_cloud_account_name"
	CloudAccountGCPType                    = "service_account"
	CloudAccountGCPVendor                  = "google"
	CloudAccountGCPAuthURL                 = "https://accounts.google.com/o/oauth2/auth"
	CloudAccountGCPTokenURL                = "https://oauth2.googleapis.com/token"
	CloudAccountGCPAuthProviderX509CertURL = "https://www.googleapis.com/oauth2/v1/certs"
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
	AWSSecurityGroupTagValue    = "tag_val_1"
)

// Role resource/data source
const (
	RoleName        = "test_role"
	RoleDescription = "this is role test"

	// Update const
	RoleUpdateDescription = "this is update role test"
)

// OrganizationalUnit resource/data source
const (
	OrganizationalUnitName       = "test_ou"
	OrganizationalUnitNameUpdate = "test_ou_update"
	ParentID                     = "" // empty string as parent id creates ou under Dome9 main root ou
	Dome9RootOuID                = "00000000-0000-0000-0000-000000000000"
)
