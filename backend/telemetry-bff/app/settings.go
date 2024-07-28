package app

import (
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/clients/users"
	"github.com/kelseyhightower/envconfig"
)

var appSettings *Settings

type Settings struct {
	Port             uint64                             `envconfig:"PORT" required:"true"`
	Auth0Domain      string                             `envconfig:"AUTH0_DOMAIN" required:"true"`
	UsersManagerHost *users.UsersManagerServiceSettings `envconfig:"USERS_MANAGER_HOST" required:"true"`
}

func loadAppSettings() *Settings {
	if appSettings == nil {
		settings := Settings{}
		if err := envconfig.Process("", &settings); err != nil {
			panic(err)
		}

		appSettings = &settings
	}

	return appSettings
}

func GetAppSettings() *Settings {
	return loadAppSettings()
}

func GetUsersManagerHostSettings() *users.UsersManagerServiceSettings {
	return loadAppSettings().UsersManagerHost
}
