//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/app"
	"github.com/google/wire"
)

func CreateApp() *app.Application {
	wire.Build(AppSettingsProv,
		HiveMQSettinsProv,
		HiveMQBroker,
		GrpcServer,
		DevicesControllerProv,
		DevicesRepoProv,
		DevicesUsecaseProv,
		MongoProv,
		SettingsMongoProv,
		PusherSettings,
		PusherProv,
		ApplicationProv)

	return new(app.Application)
}
