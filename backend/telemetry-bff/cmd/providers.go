package cmd

import (
	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/clients/users"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/web"
	"github.com/google/wire"
)

// infraestructure
var WebServerProvider = wire.NewSet(web.NewHTTPServer)
var SettingsProvider = wire.NewSet(app.GetAppSettings)
var SettingsUserService = wire.NewSet(app.GetUsersManagerHostSettings)

// repositories
var UsersManagerRepo = wire.NewSet(users.NewUsersManagerGrpcClient)

// usecases
var UsersManagerUsecase = wire.NewSet(usersManager.NewUsersManagerUsecase)

// Application
var ApplicationProvider = wire.NewSet(app.NewApplication)
