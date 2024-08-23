package server

import (
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/server/controllers"
	devicesPb "github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/server/protos/devices"

	"google.golang.org/grpc"
)

type ServerGRPC struct {
	DevicesController *controllers.DevicesController
}

func NewServerGRPC(devicesController *controllers.DevicesController) *ServerGRPC {
	return &ServerGRPC{DevicesController: devicesController}
}

func (srv *ServerGRPC) RegisterControllers(s *grpc.Server) {
	devicesPb.RegisterDevicesServiceServer(s, srv.DevicesController)
}
