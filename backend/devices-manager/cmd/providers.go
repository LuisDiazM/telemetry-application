package cmd

import (
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/app"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/messaging"
	"github.com/google/wire"
)

// infraestructure
var AppSettingsProv = wire.NewSet(app.GetAppSettings)
var HiveMQSettinsProv = wire.NewSet(app.GetHiveMQSettings)
var HiveMQBroker = wire.NewSet(messaging.NewMessagingBroker)
var ApplicationProv = wire.NewSet(app.NewApplication)
