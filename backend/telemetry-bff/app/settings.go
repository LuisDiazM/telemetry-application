package app

import (
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/cache"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/clients/devices_manager"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/clients/telemetry_analysis"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/clients/users"

	"github.com/kelseyhightower/envconfig"
)

var appSettings *Settings

type Settings struct {
	Port                  uint64 `envconfig:"PORT" required:"true"`
	Auth0Domain           string `envconfig:"AUTH0_DOMAIN" required:"true"`
	UsersManagerHost      *users.UsersManagerServiceSettings
	DeviceManagerHost     *devices_manager.DeviceSettings
	TelemetryAnalysisHost *telemetry_analysis.TelemetryAnalysisSettings
	Cache                 *cache.RedisSettings
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

func GetCacheSettings() *cache.RedisSettings {
	return loadAppSettings().Cache
}

func GetDevicesManagerSettings() *devices_manager.DeviceSettings {
	return loadAppSettings().DeviceManagerHost
}

func GetTelemetryAnalysisSettings() *telemetry_analysis.TelemetryAnalysisSettings {
	return loadAppSettings().TelemetryAnalysisHost
}
