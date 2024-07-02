package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/LuisDiazM/backend/telemetry-gateway/app"
	"github.com/LuisDiazM/backend/telemetry-gateway/infraestructure/messaging"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()
	appSettings := app.GetAppSettings()
	hivemq := messaging.NewMessagingBroker(appSettings.HiveMQ)

	app := app.NewApplication(hivemq, appSettings)

	app.Start(ctx)

}
