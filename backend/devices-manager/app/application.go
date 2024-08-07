package app

import (
	"context"

	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/messaging"
	"github.com/phuslu/log"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	MessagingBroker     *messaging.MessagingBroker
	ApplicationSettings *AppSettings
}

func NewApplication(messaging *messaging.MessagingBroker,
	appSettings *AppSettings) *Application {
	return &Application{
		MessagingBroker:     messaging,
		ApplicationSettings: appSettings,
	}
}

func (app *Application) Start(ctx context.Context) {
	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		//subscribe handlers
		app.NotifyStatusWS(ctx)
		//run forever
		log.Info().Msg("Application started!!")
		done := make(chan struct{})
		<-done
		return nil
	})

	g.Wait()
}
