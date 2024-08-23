package cmd

import (
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/app"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/messaging"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/server"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/server/controllers"
	"github.com/google/wire"
)

// infraestructure
var AppSettingsProv = wire.NewSet(app.GetAppSettings)
var HiveMQSettinsProv = wire.NewSet(app.GetHiveMQSettings)
var HiveMQBroker = wire.NewSet(messaging.NewMessagingBroker)
var GrpcServer = wire.NewSet(server.NewServerGRPC)

// controllers
var DevicesControllerProv = wire.NewSet(controllers.NewDevicesController)

// application
var ApplicationProv = wire.NewSet(app.NewApplication)
