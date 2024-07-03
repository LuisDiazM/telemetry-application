package app

import (
	"context"

	"github.com/LuisDiazM/backend/telemetry-gateway/domain/devices/usecases"
	"github.com/LuisDiazM/backend/telemetry-gateway/infraestructure/messaging"
	"github.com/phuslu/log"
)

type Application struct {
	MessagingBroker     *messaging.MessagingBroker
	ApplicationSettings *AppSettings
	DeviceUsecase       *usecases.DeviceUsecase
}

func NewApplication(messaging *messaging.MessagingBroker,
	appSettings *AppSettings, deviceUsecase *usecases.DeviceUsecase) *Application {
	return &Application{
		MessagingBroker:     messaging,
		ApplicationSettings: appSettings,
		DeviceUsecase:       deviceUsecase,
	}
}

func (app *Application) Start(ctx context.Context) {
	//subscribe handlers
	app.SaveDevicesData(ctx)

	//run forever
	log.Info().Msg("Application started!!")
	done := make(chan struct{})
	<-done
}
