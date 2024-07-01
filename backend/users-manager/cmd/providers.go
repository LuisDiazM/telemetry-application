package cmd

import (
	"github.com/LuisDiazM/backend/users-manager/app"
	"github.com/LuisDiazM/backend/users-manager/infraestructure/server"
	"github.com/LuisDiazM/backend/users-manager/infraestructure/server/controllers"
	"github.com/google/wire"
)

// Infraestructure
var ServerCustomProvider = wire.NewSet(server.NewServerGRPCCustom)
var SettingsProvider = wire.NewSet(app.GetAppSettings)

// repositories

//usecases

// controllers
var UsersControllerProvider = wire.NewSet(controllers.NewUsersController)

// Application
var AppProvider = wire.NewSet(app.NewApplication)
