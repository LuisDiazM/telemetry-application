package entities

import "time"

type Preferences struct {
}

type Role string
type Status string

const (
	ANALYST     Role = "ANALYST"
	ADMIN       Role = "ADMIN"
	SUPER_ADMIN Role = "SUPER_ADMIN"

	TRIAL      Status = "TRIAL"
	PRODUCTION Status = "PRODUCTION"
)

type UserProfile struct {
	ExternalId  string      `json:"external_id,omitempty"`
	FullName    string      `json:"full_name,omitempty"`
	Email       string      `json:"email,omitempty"`
	Preferences interface{} `json:"preferences,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty"`
	Role        Role        `json:"role,omitempty"`
	Status      Status      `json:"status,omitempty"`
}
