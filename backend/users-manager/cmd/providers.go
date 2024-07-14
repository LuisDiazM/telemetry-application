package cmd

import (
	"github.com/LuisDiazM/backend/users-manager/app"
	"github.com/LuisDiazM/backend/users-manager/domain/users/usecases"
	"github.com/LuisDiazM/backend/users-manager/infraestructure/persistence/mongo"
	"github.com/LuisDiazM/backend/users-manager/infraestructure/persistence/mongo/repositories/users"
	"github.com/LuisDiazM/backend/users-manager/infraestructure/server"
	"github.com/LuisDiazM/backend/users-manager/infraestructure/server/controllers"
	"github.com/google/wire"
)

// Infraestructure
var ServerCustomProvider = wire.NewSet(server.NewServerGRPCCustom)
var SettingsProvider = wire.NewSet(app.GetAppSettings)
var settingsMongoProvider = wire.NewSet(app.GetMongoSettings)
var MongoProvider = wire.NewSet(mongo.NewMongoImplementation)

// repositories
var usersRepoProvider = wire.NewSet(users.NewUserDBRepository)

// usecases
var usersUsecaseProvider = wire.NewSet(usecases.NewUserUsecases)

// controllers
var UsersControllerProvider = wire.NewSet(controllers.NewUsersController)

// Application
var AppProvider = wire.NewSet(app.NewApplication)
