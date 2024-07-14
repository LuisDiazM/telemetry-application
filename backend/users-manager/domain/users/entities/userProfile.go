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
	ExternalId  string      `json:"external_id,omitempty" bson:"external_id"`
	FullName    string      `json:"full_name,omitempty" bson:"full_name"`
	Email       string      `json:"email,omitempty" bson:"email"`
	Preferences interface{} `json:"preferences,omitempty" bson:"preferences"`
	CreatedAt   time.Time   `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty" bson:"updated_at"`
	Role        Role        `json:"role,omitempty" bson:"role"`
	Status      Status      `json:"status,omitempty" bson:"status"`
}
