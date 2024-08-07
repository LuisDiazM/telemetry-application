package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/cmd"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()

	application := cmd.CreateApp()
	application.Start(ctx)

}
