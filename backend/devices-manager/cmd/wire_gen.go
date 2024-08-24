// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/app"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/domain/devices/usecases"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/database/mongodb"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/database/mongodb/repositories/devices"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/messaging"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/server"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/server/controllers"
)

// Injectors from wire.go:

func CreateApp() *app.Application {
	hiveMQSettings := app.GetHiveMQSettings()
	messagingBroker := messaging.NewMessagingBroker(hiveMQSettings)
	appSettings := app.GetAppSettings()
	mongoSettings := app.GetMongoSettings()
	mongoImplementation := mongodb.NewMongoImplementation(mongoSettings)
	iDevicesDBRepo := devices.NewDevicesDbRepository(mongoImplementation)
	devicesUsecase := usecases.NewDevicesUsecase(iDevicesDBRepo)
	devicesController := controllers.NewDevicesController(devicesUsecase)
	serverGRPC := server.NewServerGRPC(devicesController)
	application := app.NewApplication(messagingBroker, appSettings, serverGRPC, devicesUsecase)
	return application
}
