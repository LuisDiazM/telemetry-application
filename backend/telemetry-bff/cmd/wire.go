//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/google/wire"
)

func CreateApp() *app.Application {
	wire.Build(WebServerProvider,
		SettingsProvider,
		SettingsUserService,
		UsersManagerUsecase,
		UsersManagerRepo,
		ApplicationProvider)
	return new(app.Application)
}
