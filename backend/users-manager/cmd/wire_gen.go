// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/LuisDiazM/backend/users-manager/app"
	"github.com/LuisDiazM/backend/users-manager/infraestructure/server"
	"github.com/LuisDiazM/backend/users-manager/infraestructure/server/controllers"
)

// Injectors from wire.go:

func CreateApp() *app.Application {
	usersController := controllers.NewUsersController()
	serverGRPCCustom := server.NewServerGRPCCustom(usersController)
	appSettings := app.GetAppSettings()
	application := app.NewApplication(serverGRPCCustom, appSettings)
	return application
}
