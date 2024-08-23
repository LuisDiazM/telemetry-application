package controllers

import (
	"context"

	pb "github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/server/protos/devices"
)

type DevicesController struct {
	pb.DevicesServiceServer
}

func NewDevicesController() *DevicesController {
	return &DevicesController{}
}

func (*DevicesController) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "Hellow"}, nil
}
