package entities

type SaveUserData struct {
	FullName    string
	Email       string
	Preferences *Preferences
	Role        string
	Status      string
	Id          string
}

type SaveUserResponse struct {
	Key   string
	Value string
}
