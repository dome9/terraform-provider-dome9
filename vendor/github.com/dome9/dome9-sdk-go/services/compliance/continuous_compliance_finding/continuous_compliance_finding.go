package continuous_compliance_finding

import (
	"fmt"
	"net/http"
)

const (
	continuousComplianceFindingPath = "Finding"
	searchFindingPath               = "search"
)

type Sorting struct {
	FieldName string `json:"fieldName,omitempty"`
	Direction int    `json:"direction,omitempty"`
}

type DateRange struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type FieldFilter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Filter struct {
	FreeTextPhrase   string         `json:"freeTextPhrase,omitempty"`
	Fields           *[]FieldFilter `json:"fields,omitempty"`
	OnlyCIEM         bool           `json:"onlyCIEM,omitempty"`
	IncludedFeatures []string       `json:"includedFeatures,omitempty"`
	CreationTime     *DateRange     `json:"creationTime,omitempty"`
}

type ContinuousComplianceFindingRequest struct {
	PageSize     int        `json:"pageSize,omitempty"`
	Sorting      *Sorting   `json:"sorting,omitempty"`
	MultiSorting *[]Sorting `json:"multiSorting,omitempty"`
	Filter       *Filter    `json:"filter,omitempty"`
	SearchAfter  *[]string  `json:"searchAfter,omitempty"`
	DataSource   string     `json:"dataSource,omitempty"`
}

type FindingComment struct {
	Text      string `json:"text,omitempty"`
	Timestamp string `json:"timestamp"`
	UserName  string `json:"userName"`
}

type TagRule struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Magellan struct {
	AlertWindowStartTime string `json:"alertWindowStartTime"`
	AlertWindowEndTime   string `json:"alertWindowEndTime"`
}

type WebhookResponse struct {
	RequestTime     string                 `json:"requestTime"`
	ResponseContent map[string]interface{} `json:"responseContent"`
}

type AdditionalField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Finding struct {
	Id                          string                     `json:"id"`
	FindingKey                  string                     `json:"findingKey"`
	CreatedTime                 string                     `json:"createdTime"`
	UpdatedTime                 string                     `json:"updatedTime"`
	CloudAccountType            string                     `json:"cloudAccountType"`
	Comments                    []FindingComment           `json:"comments,omitempty"`
	CloudAccountId              string                     `json:"cloudAccountId"`
	CloudAccountExternalId      string                     `json:"cloudAccountExternalId"`
	OrganizationalUnitId        string                     `json:"organizationalUnitId,omitempty"`
	OrganizationalUnitPath      string                     `json:"organizationalUnitPath,omitempty"`
	BundleId                    float64                    `json:"bundleId"`
	AlertType                   string                     `json:"alertType"`
	RuleId                      string                     `json:"ruleId,omitempty"`
	RuleName                    string                     `json:"ruleName"`
	RuleLogic                   string                     `json:"ruleLogic"`
	EntityDome9Id               string                     `json:"entityDome9Id,omitempty"`
	EntityExternalId            string                     `json:"entityExternalId"`
	EntityType                  string                     `json:"entityType"`
	EntityTypeByEnvironmentType string                     `json:"entityTypeByEnvironmentType"`
	EntityName                  string                     `json:"entityName"`
	EntityNetwork               string                     `json:"entityNetwork,omitempty"`
	EntityTags                  []TagRule                  `json:"entityTags,omitempty"`
	Severity                    string                     `json:"severity"`
	Description                 string                     `json:"description"`
	Remediation                 string                     `json:"remediation"`
	Tag                         string                     `json:"tag"`
	Region                      string                     `json:"region"`
	BundleName                  string                     `json:"bundleName"`
	Acknowledged                bool                       `json:"acknowledged"`
	Origin                      string                     `json:"origin"`
	LastSeenTime                string                     `json:"lastSeenTime"`
	OwnerUserName               string                     `json:"ownerUserName,omitempty"`
	Magellan                    Magellan                   `json:"magellan,omitempty"`
	IsExcluded                  bool                       `json:"isExcluded"`
	WebhookResponses            map[string]WebhookResponse `json:"webhookResponses,omitempty"`
	RemediationActions          []string                   `json:"remediationActions,omitempty"`
	AdditionalFields            []AdditionalField          `json:"additionalFields"`
	Occurrences                 []string                   `json:"occurrences"`
	ScanId                      string                     `json:"scanId"`
	Status                      string                     `json:"status"`
	Category                    string                     `json:"category"`
	Action                      string                     `json:"action"`
	Labels                      []string                   `json:"labels"`
}

type FieldAggregation struct {
	Value string  `json:"value"`
	Count float64 `json:"count"`
}

type ContinuousComplianceFindingResponse struct {
	SearchRequest      ContinuousComplianceFindingRequest `json:"searchRequest"`
	Findings           []Finding                          `json:"findings"`
	TotalFindingsCount float64                            `json:"totalFindingsCount"`
	Aggregations       map[string][]FieldAggregation      `json:"aggregations"`
	SearchAfter        []string                           `json:"searchAfter"`
}

func (service *Service) Search(body *ContinuousComplianceFindingRequest) (*ContinuousComplianceFindingResponse, *http.Response, error) {
	v := new(ContinuousComplianceFindingResponse)
	relativeURL := fmt.Sprintf("%s/%s", continuousComplianceFindingPath, searchFindingPath)
	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
