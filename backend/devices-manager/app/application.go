package app

import (
	"context"
	"fmt"
	"net"

	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/domain/devices/usecases"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/messaging"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/pusher"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/server"
	"github.com/phuslu/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Application struct {
	MessagingBroker     *messaging.MessagingBroker
	ApplicationSettings *AppSettings
	GRPCServer          *grpc.Server
	Pusher              *pusher.PusherImp
	SrvCustom           *server.ServerGRPC
	DevicesUsecase      *usecases.DevicesUsecase
}

func NewApplication(messaging *messaging.MessagingBroker,
	appSettings *AppSettings, pusher *pusher.PusherImp, srvCustom *server.ServerGRPC,
	devicesUsecase *usecases.DevicesUsecase) *Application {
	s := grpc.NewServer()
	return &Application{
		MessagingBroker:     messaging,
		ApplicationSettings: appSettings,
		GRPCServer:          s,
		SrvCustom:           srvCustom,
		DevicesUsecase:      devicesUsecase,
		Pusher:              pusher,
	}
}

func (app *Application) Start(ctx context.Context) {
	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		//subscribe handlers
		app.NotifyStatusWS(ctx)
		//run forever
		log.Info().Msg("MQTT Listener Run!!!")
		done := make(chan struct{})
		<-done
		return nil
	})

	g.Go(func() error {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", app.ApplicationSettings.Port))
		if err != nil {
			return err

		}
		app.SrvCustom.RegisterControllers(app.GRPCServer)
		reflection.Register(app.GRPCServer)
		log.Info().Msg(fmt.Sprintf("grpc server start at port %s!!", app.ApplicationSettings.Port))
		if err = app.GRPCServer.Serve(lis); err != nil {
			return err
		}
		return nil
	})

	g.Wait()
}
