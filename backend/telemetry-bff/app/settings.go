package app

import "github.com/kelseyhightower/envconfig"

var appSettings *Settings

type Settings struct {
	Port        uint64 `envconfig:"PORT" required:"true"`
	Auth0Domain string `envconfig:"AUTH0_DOMAIN" required:"true"`
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
