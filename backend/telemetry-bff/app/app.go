package app

import (
	"context"
	"fmt"
	"log"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/devicesManager/usecases"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	WebServer             *gin.Engine
	Settings              *Settings
	UsersManagerUsecase   *usersManager.UsersManagerUsecase
	DevicesManagerUsecase *usecases.DeviceManagerUsecase
}

func NewApplication(webServer *gin.Engine,
	settings *Settings,
	usersManager *usersManager.UsersManagerUsecase,
	devicesManager *usecases.DeviceManagerUsecase) *Application {
	return &Application{
		WebServer:             webServer,
		Settings:              settings,
		UsersManagerUsecase:   usersManager,
		DevicesManagerUsecase: devicesManager,
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
