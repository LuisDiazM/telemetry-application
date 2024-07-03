package cmd

import (
	"github.com/LuisDiazM/backend/telemetry-gateway/app"
	"github.com/LuisDiazM/backend/telemetry-gateway/domain/devices/usecases"
	"github.com/LuisDiazM/backend/telemetry-gateway/infraestructure/cache"
	cacheDevices "github.com/LuisDiazM/backend/telemetry-gateway/infraestructure/cache/repositories/devices"

	"github.com/LuisDiazM/backend/telemetry-gateway/infraestructure/database"
	"github.com/LuisDiazM/backend/telemetry-gateway/infraestructure/database/repositories/devices"
	"github.com/LuisDiazM/backend/telemetry-gateway/infraestructure/messaging"
	"github.com/google/wire"
)

// infraestructure
var MongoProvider = wire.NewSet(database.NewMongoImplementation)
var HiveProvider = wire.NewSet(messaging.NewMessagingBroker)
var CacheProvider = wire.NewSet(cache.NewRedisImp)

// repositories
var DeviceDBProvider = wire.NewSet(devices.NewDevicesDbRepository)
var DeviceCacheProvider = wire.NewSet(cacheDevices.NewCacheDeviceRepository)

// usecases
var DeviceUsecaseProvider = wire.NewSet(usecases.NewDeviceUsecase)

// application
var SettingsProvider = wire.NewSet(app.GetAppSettings)
var ApplicationProvider = wire.NewSet(app.NewApplication)
var HiveSettingsProvider = wire.NewSet(app.GetHiveMQSettings)
var MongoSettingsProvider = wire.NewSet(app.GetMongoSettings)
var CacheSettingsProvider = wire.NewSet(app.GetCacheSettings)
