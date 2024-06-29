package app

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	WebServer *gin.Engine
	Settings  *Settings
}

func NewApplication(webServer *gin.Engine, settings *Settings) *Application {
	return &Application{
		WebServer: webServer,
		Settings:  settings,
	}
}

func (app *Application) Start(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		if err := app.WebServer.Run(fmt.Sprintf(`:%d`, app.Settings.Port)); err != nil {
			return err
		}
		log.Println("web server started!!")
		return nil
	})
	return g.Wait()
}
