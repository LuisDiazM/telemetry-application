package users

type UsersManagerServiceSettings struct {
	UsersManagerHost string `envconfig:"USERS_MANAGER_HOST" required:"true"`
}
