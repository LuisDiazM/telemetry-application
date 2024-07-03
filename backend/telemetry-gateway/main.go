package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/LuisDiazM/backend/telemetry-gateway/cmd"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()
	app := cmd.CreateApp()

	app.Start(ctx)

}
