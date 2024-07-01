package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/LuisDiazM/backend/users-manager/cmd"
)

func main() {

	app := cmd.CreateApp()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
