package app

import (
	"context"

	"github.com/LuisDiazM/backend/telemetry-gateway/infraestructure/messaging"
	"github.com/phuslu/log"
)

type Application struct {
	MessagingBroker     *messaging.MessagingBroker
	ApplicationSettings *AppSettings
}

func NewApplication(messaging *messaging.MessagingBroker, appSettings *AppSettings) *Application {
	return &Application{
		MessagingBroker:     messaging,
		ApplicationSettings: appSettings,
	}
}

func (app *Application) Start(ctx context.Context) {
	app.SaveDevicesData(ctx)

	//run forever
	log.Info().Msg("Application started!!")
	done := make(chan struct{})
	<-done
}
