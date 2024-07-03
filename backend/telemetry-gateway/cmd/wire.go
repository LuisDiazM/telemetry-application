//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/LuisDiazM/backend/telemetry-gateway/app"
	"github.com/google/wire"
)

func CreateApp() *app.Application {
	wire.Build(HiveProvider,
		MongoProvider,
		SettingsProvider,
		ApplicationProvider,
		DeviceUsecaseProvider,
		DeviceDBProvider,
		HiveSettingsProvider,
		MongoSettingsProvider,

		CacheProvider,
		DeviceCacheProvider,
		CacheSettingsProvider,
	)
	return new(app.Application)
}
