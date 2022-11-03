package k8s

import (
	"time"
)

type CloudAccountRequest struct {
	Name                 string `json:"name"`
	OrganizationalUnitID string `json:"organizationalUnitId,omitempty"`
}

type CloudAccountResponse struct {
	ID                        string    `json:"id"` //The k8s cluster ID
	Name                      string    `json:"name"`
	CreationDate              time.Time `json:"creationDate"`
	Vendor                    string    `json:"vendor"`
	OrganizationalUnitID      string    `json:"organizationalUnitId,omitempty"`
	OrganizationalUnitPath    string    `json:"organizationalUnitPath,omitempty"`
	OrganizationalUnitName    string    `json:"organizationalUnitName,omitempty"`
	ClusterVersion            string    `json:"clusterVersion"`
	RuntimeProtectionEnabled  bool      `json:"runtimeProtection"`
	AdmissionControlEnabled   bool      `json:"admissionControl"`
	ImageAssuranceEnabled     bool      `json:"vulnerabilityAssessment"`
	ThreatIntelligenceEnabled bool      `json:"magellan"`
}

type CloudAccountUpdateNameRequest struct {
	Name string `json:"name"`
}

type CloudAccountUpdateOrganizationalIDRequest struct {
	OrganizationalUnitId string `json:"organizationalUnitId,omitempty"`
}

type RuntimeProtectionEnableRequest struct {
	CloudAccountId string `json:"k8sAccountId"`
	Enabled        bool   `json:"enabled"`
}

type AdmissionControlEnableRequest struct {
	CloudAccountId string `json:"k8sAccountId"`
	Enabled        bool   `json:"enabled"`
}

type ImageAssuranceEnableRequest struct {
	CloudAccountId string `json:"cloudAccountId"`
	Enabled        bool   `json:"enabled"`
}

type ThreatIntelligenceEnableRequest struct {
	CloudAccountId string `json:"k8sAccountId"`
	Enabled        bool   `json:"enabled"`
}