package app

import (
	"context"
	"fmt"
	"net"

	"github.com/LuisDiazM/backend/users-manager/infraestructure/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Application struct {
	ServerGRPCCustom *server.ServerGRPCCustom
	GRPCServer       *grpc.Server
	AppSettings      *AppSettings
}

func NewApplication(srvCustom *server.ServerGRPCCustom, appSettings *AppSettings) *Application {
	s := grpc.NewServer()
	return &Application{
		ServerGRPCCustom: srvCustom,
		GRPCServer:       s,
		AppSettings:      appSettings,
	}
}

func (app *Application) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", app.AppSettings.Port))
	if err != nil {
		return err

	}
	app.ServerGRPCCustom.RegisterControllers(app.GRPCServer)

	reflection.Register(app.GRPCServer)
	if err = app.GRPCServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
