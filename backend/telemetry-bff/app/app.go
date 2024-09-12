package app

import (
	"context"
	"fmt"
	"log"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/devicesManager/usecases"
	measures "github.com/LuisDiazM/backend/telemetry-bff/domain/telemetryAnalysis/usecases"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	WebServer               *gin.Engine
	Settings                *Settings
	UsersManagerUsecase     *usersManager.UsersManagerUsecase
	DevicesManagerUsecase   *usecases.DeviceManagerUsecase
	TelemetryAnalysisUsecae *measures.MeasuresUsecase
}

func NewApplication(webServer *gin.Engine,
	settings *Settings,
	usersManager *usersManager.UsersManagerUsecase,
	devicesManager *usecases.DeviceManagerUsecase,
	measures *measures.MeasuresUsecase) *Application {
	return &Application{
		WebServer:               webServer,
		Settings:                settings,
		UsersManagerUsecase:     usersManager,
		DevicesManagerUsecase:   devicesManager,
		TelemetryAnalysisUsecae: measures,
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
