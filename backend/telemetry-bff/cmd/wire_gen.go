// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/devicesManager/usecases"
	usecases2 "github.com/LuisDiazM/backend/telemetry-bff/domain/telemetryAnalysis/usecases"
	usersManager2 "github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/cache"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/cache/repositories/usersManager"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/clients/devices_manager"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/clients/telemetry_analysis"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/clients/users"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/web"
)

// Injectors from wire.go:

func CreateApp() *app.Application {
	engine := web.NewHTTPServer()
	settings := app.GetAppSettings()
	usersManagerServiceSettings := app.GetUsersManagerHostSettings()
	iUsersManagerClient := users.NewUsersManagerGrpcClient(usersManagerServiceSettings)
	redisSettings := app.GetCacheSettings()
	redisImp := cache.NewRedisImp(redisSettings)
	iUsersManagerCacheRepo := usersManager.NewUsersManagerCacheRepo(redisImp)
	usersManagerUsecase := usersManager2.NewUsersManagerUsecase(iUsersManagerClient, iUsersManagerCacheRepo)
	deviceSettings := app.GetDevicesManagerSettings()
	igrpcClientDeviceRepo := devices_manager.NewDevicesManagerGRPCClientRepo(deviceSettings)
	deviceManagerUsecase := usecases.NewDeviceManagerUsecase(igrpcClientDeviceRepo)
	telemetryAnalysisSettings := app.GetTelemetryAnalysisSettings()
	iTelemetryAnalysisGRPC := telemetry_analysis.NewTelemetryAnalysisGRPCClient(telemetryAnalysisSettings)
	measuresUsecase := usecases2.NewMeasuresUsecase(iTelemetryAnalysisGRPC)
	application := app.NewApplication(engine, settings, usersManagerUsecase, deviceManagerUsecase, measuresUsecase)
	return application
}
