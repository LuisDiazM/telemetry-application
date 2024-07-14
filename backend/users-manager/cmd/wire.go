//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/LuisDiazM/backend/users-manager/app"
	"github.com/google/wire"
)

func CreateApp() *app.Application {
	wire.Build(AppProvider,
		ServerCustomProvider,
		UsersControllerProvider,
		SettingsProvider,

		usersRepoProvider,
		usersUsecaseProvider,
		MongoProvider,
		settingsMongoProvider,
	)
	return new(app.Application)
}
