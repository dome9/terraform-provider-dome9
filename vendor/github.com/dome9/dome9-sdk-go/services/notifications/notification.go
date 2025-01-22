package notifications

import (
	"fmt"
	"net/http"
)

// Constants
const (
	RESTfulServicePathNotification = "notification"
)

// Models
type BaseNotificationViewModel struct {
	Name                 string                               `json:"name" validate:"required"`
	Description          string                               `json:"description"`
	AlertsConsole        bool                                 `json:"alertsConsole" default:"true"`
	SendOnEachOccurrence bool                                 `json:"sendOnEachOccurrence"`
	Origin               string                               `json:"origin" validate:"required"`
	IntegrationSettings  NotificationIntegrationSettingsModel `json:"integrationSettingsModel" validate:"required"`
}

type NotificationIntegrationSettingsModel struct {
	ReportsIntegrationSettings            []ReportNotificationIntegrationSettings    `json:"reportsIntegrationSettings"`
	SingleNotificationIntegrationSettings []SingleNotificationIntegrationSettings    `json:"singleNotificationIntegrationSettings"`
	ScheduledIntegrationSettings          []ScheduledNotificationIntegrationSettings `json:"scheduledIntegrationSettings"`
}

type BaseNotificationIntegrationSettings struct {
	IntegrationId string                        `json:"integrationId" validate:"required"`
	OutputType    string                        `json:"outputType"`
	Filter        *ComplianceNotificationFilter `json:"filter"`
}

type SingleNotificationIntegrationSettings struct {
	BaseNotificationIntegrationSettings
	Payload string `json:"payload"`
}

type ReportNotificationIntegrationSettings struct {
	BaseNotificationIntegrationSettings
}

type ScheduledNotificationIntegrationSettings struct {
	BaseNotificationIntegrationSettings
	CronExpression string `json:"cronExpression" validate:"required,cron"`
}

type ComplianceNotificationFilter struct {
	Severities       []string        `json:"severities"`
	RuleEntityTypes  []string        `json:"rule_entity_types"`
	EntityTags       []TagRuleEntity `json:"entity_tags"`
	EntityNames      []string        `json:"entity_names"`
	EntityIds        []string        `json:"entity_ids"`
	EntityCategories []string        `json:"entity_categories"`
}

type TagRuleEntity struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PutNotificationViewModel struct {
	BaseNotificationViewModel
	Id string `json:"id" validate:"required"`
}

type PostNotificationViewModel struct {
	BaseNotificationViewModel
}

type ResponseNotificationViewModel struct {
	BaseNotificationViewModel
	Id string `json:"id" validate:"required"`
}

// APIs

func (service *Service) Create(body PostNotificationViewModel) (*ResponseNotificationViewModel, *http.Response, error) {
	v := new(ResponseNotificationViewModel)
	resp, err := service.Client.NewRequestDoRetry("POST", RESTfulServicePathNotification, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*[]ResponseNotificationViewModel, *http.Response, error) {
	v := new([]ResponseNotificationViewModel)
	resp, err := service.Client.NewRequestDoRetry("GET", RESTfulServicePathNotification, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetById(id string) (*ResponseNotificationViewModel, *http.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id parameter must be passed")
	}

	v := new(ResponseNotificationViewModel)
	relativeURL := fmt.Sprintf("%s/%s", RESTfulServicePathNotification, id)
	resp, err := service.Client.NewRequestDoRetry("GET", relativeURL, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetByName(name string) (*ResponseNotificationViewModel, *http.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name parameter must be passed")
	}

	v := new(ResponseNotificationViewModel)
	relativeURL := fmt.Sprintf("%s?name=%s", RESTfulServicePathNotification, name)
	resp, err := service.Client.NewRequestDoRetry("GET", relativeURL, nil, nil, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(body PutNotificationViewModel) (*ResponseNotificationViewModel, *http.Response, error) {
	if body.Id == "" {
		return nil, nil, fmt.Errorf("id parameter must be passed")
	}

	v := new(ResponseNotificationViewModel)
	resp, err := service.Client.NewRequestDoRetry("PUT", RESTfulServicePathNotification, nil, body, v, nil)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Delete(id string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", RESTfulServicePathNotification, id)
	resp, err := service.Client.NewRequestDoRetry("DELETE", relativeURL, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
