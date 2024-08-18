package assessment

import (
	"fmt"
	"net/http"
)

const (
	assessmentResourcePath    = "assessment/bundleV2"
	assessmentHistoryBasePath = "AssessmentHistoryV2"
)

type RunBundleRequest struct {
	BundleID               int    `json:"id"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	Dome9CloudAccountID    string `json:"dome9CloudAccountId"`
	ExternalCloudAccountID string `json:"externalCloudAccountId"`
	CloudAccountID         string `json:"cloudAccountId"`
	CloudAccountType       string `json:"cloudAccountType"`
	RequestID              string `json:"requestId"`
	ShouldMinimizeResult   bool   `json:"shouldMinimizeResult"`
}

type RunBundleResponse struct {
	Request                 Request                  `json:"request"`
	Tests                   []Test                   `json:"tests"`
	TestEntities            map[string][]interface{} `json:"testEntities"`
	Exclusions              []Exclusion              `json:"exclusions"`
	Remediations            []Remediation            `json:"remediations"`
	DataSyncStatus          []DataSyncStatus         `json:"dataSyncStatus"`
	CreatedTime             string                   `json:"createdTime"`
	ID                      int                      `json:"id"`
	AssessmentId            string                   `json:"assessmentId"`
	TriggeredBy             string                   `json:"triggeredBy"`
	AssessmentPassed        bool                     `json:"assessmentPassed"`
	HasErrors               bool                     `json:"hasErrors"`
	Stats                   Stats                    `json:"stats"`
	HasDataSyncStatusIssues bool                     `json:"hasDataSyncStatusIssues"`
	ComparisonCustomId      string                   `json:"comparisonCustomId"`
	AdditionalFields        map[string][]interface{} `json:"additionalFields"`
}

type DeleteRequest struct {
	HistoryId int `json:"historyId"`
}

type Request struct {
	IsTemplate             bool   `json:"isTemplate"`
	BundleID               int    `json:"id"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	Dome9CloudAccountID    string `json:"dome9CloudAccountId"`
	ExternalCloudAccountID string `json:"externalCloudAccountId"`
	CloudAccountID         string `json:"cloudAccountId"`
	CloudAccountType       string `json:"cloudAccountType"`
	RequestID              string `json:"requestId"`
	ShouldMinimizeResult   bool   `json:"shouldMinimizeResult"`
}

type Test struct {
	Error             string         `json:"error"`
	TestedCount       int            `json:"testedCount"`
	RelevantCount     int            `json:"relevantCount"`
	NonComplyingCount int            `json:"nonComplyingCount"`
	ExclusionStats    ExclusionStats `json:"exclusionStats"`
	EntityResults     []EntityResult `json:"entityResults"`
	Rule              Rule           `json:"rule"`
	TestPassed        bool           `json:"testPassed"`
}

type ExclusionStats struct {
	TestedCount       int `json:"testedCount"`
	RelevantCount     int `json:"relevantCount"`
	NonComplyingCount int `json:"nonComplyingCount"`
}

type EntityResult struct {
	ValidationStatus string                          `json:"validationStatus"`
	IsRelevant       bool                            `json:"isRelevant"`
	IsValid          bool                            `json:"isValid"`
	IsExcluded       bool                            `json:"isExcluded"`
	ExclusionID      string                          `json:"exclusionId"`
	RemediationID    string                          `json:"remediationId"`
	Error            string                          `json:"error"`
	TestObj          RuleEngineFailedEntityReference `json:"testObj"`
}

type RuleEngineFailedEntityReference struct {
	Id                         string `json:"id"`
	Dome9Id                    string `json:"dome9Id"`
	EntityType                 string `json:"entityType"`
	EntityIndex                int    `json:"entityIndex"`
	CustomEntityComparisonHash string `json:"customEntityComparisonHash"`
}

type Rule struct {
	Name          string   `json:"name"`
	Severity      string   `json:"severity"`
	Logic         string   `json:"logic"`
	Description   string   `json:"description"`
	Remediation   string   `json:"remediation"`
	Cloudbots     string   `json:"cloudbots"`
	ComplianceTag string   `json:"complianceTag"`
	Domain        string   `json:"domain"`
	Priority      string   `json:"priority"`
	ControlTitle  string   `json:"controlTitle"`
	RuleID        string   `json:"ruleId"`
	Category      string   `json:"category"`
	Labels        []string `json:"labels"`
	LogicHash     string   `json:"logicHash"`
	IsDefault     bool     `json:"isDefault"`
}

type Exclusion struct {
	Platform              string                       `json:"platform"`
	ID                    int                          `json:"id"`
	Rules                 []ExclusionOrRemediationRule `json:"rules"`
	LogicExpressions      []string                     `json:"logicExpressions"`
	RulesetId             int                          `json:"rulesetId"`
	CloudAccountIds       []string                     `json:"cloudAccountIds"`
	Comment               string                       `json:"comment"`
	OrganizationalUnitIds []string                     `json:"organizationalUnitIds"`
	DateRange             Date                         `json:"dateRange"`
}

type Remediation struct {
	Platform              string                       `json:"platform"`
	ID                    int                          `json:"id"`
	Rules                 []ExclusionOrRemediationRule `json:"rules"`
	LogicExpressions      []string                     `json:"logicExpressions"`
	RulesetId             int                          `json:"rulesetId"`
	CloudAccountIds       []string                     `json:"cloudAccountIds"`
	Comment               string                       `json:"comment"`
	CloudBots             []string                     `json:"cloudBots"`
	OrganizationalUnitIds []string                     `json:"organizationalUnitIds"`
	DateRange             Date                         `json:"dateRange"`
}

type DataSyncStatus struct {
	EntityType                   string                         `json:"entityType"`
	RecentlySuccessfulSync       bool                           `json:"recentlySuccessfulSync"`
	GeneralFetchPermissionIssues bool                           `json:"generalFetchPermissionIssues"`
	EntitiesWithPermissionIssues []EntitiesWithPermissionIssues `json:"entitiesWithPermissionIssues"`
}

type Stats struct {
	Passed                  int           `json:"passed"`
	PassedRulesBySeverity   RulesSeverity `json:"passedRulesBySeverity"`
	Failed                  int           `json:"failed"`
	FailedRulesBySeverity   RulesSeverity `json:"failedRulesBySeverity"`
	Error                   int           `json:"error"`
	FailedTests             int           `json:"failedTests"`
	LogicallyTested         int           `json:"logicallyTested"`
	FailedEntities          int           `json:"failedEntities"`
	ExcludedTests           int           `json:"excludedTests"`
	ExcludedFailedTests     int           `json:"excludedFailedTests"`
	ExcludedRules           int           `json:"excludedRules"`
	ExcludedRulesBySeverity RulesSeverity `json:"excludedRulesBySeverity"`
}

type EntitiesWithPermissionIssues struct {
	ExternalID            string `json:"externalId"`
	Name                  string `json:"name"`
	CloudVendorIdentifier string `json:"cloudVendorIdentifier"`
}

type ExclusionOrRemediationRule struct {
	LogicHash string `json:"logicHash"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
}

type Date struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type RulesSeverity struct {
	Informational int `json:"informational"`
	Low           int `json:"low"`
	Medium        int `json:"medium"`
	High          int `json:"high"`
	Critical      int `json:"critical"`
}

func (service *Service) Run(body *RunBundleRequest) (*RunBundleResponse, *http.Response, error) {
	v := new(RunBundleResponse)
	resp, err := service.Client.NewRequestDo("POST", assessmentResourcePath, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id int) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%v", assessmentHistoryBasePath, id)

	deleteRequest := DeleteRequest{
		HistoryId: id,
	}

	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, deleteRequest, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (service *Service) Get(id string) (*RunBundleResponse, *http.Response, error) {
	v := new(RunBundleResponse)
	relativeURL := fmt.Sprintf("%s/%s", assessmentHistoryBasePath, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
