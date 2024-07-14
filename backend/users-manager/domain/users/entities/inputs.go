package entities

type UserRequest struct {
	UserId string `json:"user_id,omitempty"`
}

type SaveUserRequest struct {
	ExternalId  string
	FullName    string
	Email       string
	Preferences interface{}
	Role        string
	Status      string
}
