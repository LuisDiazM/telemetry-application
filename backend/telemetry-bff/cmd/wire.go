//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/google/wire"
)

func CreateApp() *app.Application {
	wire.Build(WebServerProvider,
		CacheProvider,
		SettingsProvider,
		SettingsUserService,
		CacheSettings,
		UsersCacheRepo,
		UsersManagerUsecase,
		UsersManagerRepo,
		DevicesManagerRepo,
		SettingsDevicesManager,
		DevicesMAnagerUsecase,
		ApplicationProvider)
	return new(app.Application)
}
