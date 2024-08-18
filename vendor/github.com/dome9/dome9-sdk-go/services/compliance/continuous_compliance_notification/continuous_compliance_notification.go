package continuous_compliance_notification

import (
	"fmt"
	"net/http"
)

const (
	continuousComplianceResourcePath = "Compliance/ContinuousComplianceNotification"
)

type ContinuousComplianceNotificationRequest struct {
	Name                                string                               `json:"name"`
	Description                         string                               `json:"description,omitempty"`
	AlertsConsole                       bool                                 `json:"alertsConsole,omitempty"`
	ScheduledReport                     *ScheduledReport                     `json:"scheduledReport"`
	ChangeDetection                     ChangeDetection                      `json:"changeDetection"`
	GCPSecurityCommandCenterIntegration *GCPSecurityCommandCenterIntegration `json:"gcpSecurityCommandCenterIntegration"`
}

type ContinuousComplianceNotificationResponse struct {
	ID                                  string                               `json:"id"`
	Name                                string                               `json:"name"`
	Description                         string                               `json:"description"`
	AlertsConsole                       bool                                 `json:"alertsConsole"`
	ScheduledReport                     *ScheduledReport                     `json:"scheduledReport"`
	ChangeDetection                     ChangeDetection                      `json:"changeDetection"`
	GCPSecurityCommandCenterIntegration *GCPSecurityCommandCenterIntegration `json:"gcpSecurityCommandCenterIntegration"`
}

type ScheduledReport struct {
	EmailSendingState string        `json:"emailSendingState,omitempty"`
	ScheduleData      *ScheduleData `json:"ScheduleData,omitempty"`
}

type ScheduleData struct {
	CronExpression string   `json:"cronExpression"`
	Type           string   `json:"type"`
	Recipients     []string `json:"recipients"`
}

type ChangeDetection struct {
	EmailSendingState              string                     `json:"emailSendingState,omitempty"`
	EmailPerFindingSendingState    string                     `json:"emailPerFindingSendingState,omitempty"`
	SNSSendingState                string                     `json:"snsSendingState,omitempty"`
	ExternalTicketCreatingState    string                     `json:"externalTicketCreatingState,omitempty"`
	AWSSecurityHubIntegrationState string                     `json:"awsSecurityHubIntegrationState,omitempty"`
	WebhookIntegrationState        string                     `json:"webhookIntegrationState,omitempty"`
	SlackIntegrationState          string                     `json:"slackIntegrationState,omitempty"`
	TeamsIntegrationState          string                     `json:"teamsIntegrationState,omitempty"`
	EmailData                      *EmailData                 `json:"emailData,omitempty"`
	EmailPerFindingData            *EmailPerFindingData       `json:"emailPerFindingData,omitempty"`
	SNSData                        *SNSData                   `json:"snsData,omitempty"`
	TicketingSystemData            *TicketingSystemData       `json:"ticketingSystemData,omitempty"`
	AWSSecurityHubIntegration      *AWSSecurityHubIntegration `json:"awsSecurityHubIntegration,omitempty"`
	WebhookData                    *WebhookData               `json:"webhookData,omitempty"`
	SlackData                      *SlackData                 `json:"slackData,omitempty"`
	TeamsData                      *TeamsData                 `json:"teamsData,omitempty"`
}

type EmailData struct {
	Recipients []string `json:"recipients"`
}

type EmailPerFindingData struct {
	Recipients               []string `json:"recipients"`
	NotificationOutputFormat string   `json:"notificationOutputFormat"`
}

type SNSData struct {
	SNSTopicArn     string `json:"snsTopicArn"`
	SNSOutputFormat string `json:"snsOutputFormat"`
}

type TicketingSystemData struct {
	SystemType         string `json:"systemType"`
	ShouldCloseTickets bool   `json:"shouldCloseTickets"`
	Domain             string `json:"domain,omitempty"`
	User               string `json:"user,omitempty"`
	Pass               string `json:"pass"`
	ProjectKey         string `json:"projectKey,omitempty"`
	IssueType          string `json:"issueType,omitempty"`
}

type AWSSecurityHubIntegration struct {
	ExternalAccountID string `json:"externalAccountId"`
	Region            string `json:"region"`
}

type WebhookData struct {
	URL               string                 `json:"url"`
	HTTPMethod        string                 `json:"httpMethod"`
	AuthMethod        string                 `json:"authMethod"`
	Username          string                 `json:"username,omitempty"`
	Password          string                 `json:"password,omitempty"`
	FormatType        string                 `json:"formatType"`
	PayloadFormat     map[string]interface{} `json:"payloadFormat"`
	IgnoreCertificate bool                   `json:"ignoreCertificate"`
	AdvancedUrl       bool                   `json:"advancedUrl"`
}

type SlackData struct {
	URL string `json:"url"`
}

type TeamsData struct {
	URL string `json:"url"`
}

type GCPSecurityCommandCenterIntegration struct {
	State     string `json:"state"`
	ProjectID string `json:"projectId,omitempty"`
	SourceID  string `json:"sourceId,omitempty"`
}

func (service *Service) Get(id string) (*ContinuousComplianceNotificationResponse, *http.Response, error) {
	v := new(ContinuousComplianceNotificationResponse)
	relativeURL := fmt.Sprintf("%s/%s", continuousComplianceResourcePath, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]ContinuousComplianceNotificationResponse, *http.Response, error) {
	v := new([]ContinuousComplianceNotificationResponse)
	resp, err := service.Client.NewRequestDo("GET", continuousComplianceResourcePath, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(body *ContinuousComplianceNotificationRequest) (*ContinuousComplianceNotificationResponse, *http.Response, error) {
	v := new(ContinuousComplianceNotificationResponse)
	resp, err := service.Client.NewRequestDo("POST", continuousComplianceResourcePath, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(id string, body *ContinuousComplianceNotificationRequest) (*ContinuousComplianceNotificationResponse, *http.Response, error) {
	v := new(ContinuousComplianceNotificationResponse)
	relativeURL := fmt.Sprintf("%s/%s", continuousComplianceResourcePath, id)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, body, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", continuousComplianceResourcePath, id)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
