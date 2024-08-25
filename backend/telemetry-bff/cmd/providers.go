package cmd

import (
	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/devicesManager/usecases"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/cache"
	cacheUsers "github.com/LuisDiazM/backend/telemetry-bff/infraestructure/cache/repositories/usersManager"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/clients/devices_manager"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/clients/users"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/web"

	"github.com/google/wire"
)

// infraestructure
var WebServerProvider = wire.NewSet(web.NewHTTPServer)
var SettingsProvider = wire.NewSet(app.GetAppSettings)
var CacheProvider = wire.NewSet(cache.NewRedisImp)
var SettingsUserService = wire.NewSet(app.GetUsersManagerHostSettings)
var CacheSettings = wire.NewSet(app.GetCacheSettings)
var SettingsDevicesManager = wire.NewSet(app.GetDevicesManagerSettings)

// repositories
var UsersManagerRepo = wire.NewSet(users.NewUsersManagerGrpcClient)
var UsersCacheRepo = wire.NewSet(cacheUsers.NewUsersManagerCacheRepo)
var DevicesManagerRepo = wire.NewSet(devices_manager.NewDevicesManagerGRPCClientRepo)

// usecases
var UsersManagerUsecase = wire.NewSet(usersManager.NewUsersManagerUsecase)
var DevicesMAnagerUsecase = wire.NewSet(usecases.NewDeviceManagerUsecase)

// Application
var ApplicationProvider = wire.NewSet(app.NewApplication)
