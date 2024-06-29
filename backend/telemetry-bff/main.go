package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/LuisDiazM/backend/telemetry-bff/cmd"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/web"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/web/routes"
)

func main() {
	app := cmd.CreateApp()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()
	web.EnableCors(app)
	routes.SetUpRoutes(app)
	err := app.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
