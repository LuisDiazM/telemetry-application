package controllers

import (
	"context"

	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/domain/devices/entities"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/domain/devices/usecases"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/server/controllers/utils"
	pb "github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/server/protos/devices"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DevicesController struct {
	pb.DevicesServiceServer
	Usecase *usecases.DevicesUsecase
}

func NewDevicesController(deviceUsecase *usecases.DevicesUsecase) *DevicesController {
	return &DevicesController{Usecase: deviceUsecase}
}

func (controller *DevicesController) GetDeviceById(ctx context.Context, in *pb.DeviceByIdRequest) (*pb.DeviceData, error) {
	deviceData := controller.Usecase.GetDeviceById(ctx, in.SensorId)
	if deviceData == nil {
		return nil, status.Error(codes.NotFound, "")
	}
	response, err := utils.ParseDevicesToGrpc(deviceData)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (controller *DevicesController) GetDevicesByLocation(ctx context.Context, in *pb.LocationRequests) (*pb.DevicesDataResponse, error) {
	location := entities.Location{Type: entities.Point, Coordinates: []float64{float64(in.Latitude), float64(in.Longitude)}}
	devices := controller.Usecase.GetDevicesFromLocation(ctx, location)
	if devices == nil {
		return nil, status.Error(codes.NotFound, "")
	}
	response := utils.ParseDevicesListToGRPCResponse(*devices)
	return response, nil
}

func (controller *DevicesController) CreateDevice(ctx context.Context, in *pb.DeviceData) (*empty.Empty, error) {
	device := utils.ParseDeviceGRPCtoDevice(in)
	if device == nil {
		return nil, status.Error(codes.InvalidArgument, "")
	}
	response := controller.Usecase.CreateDevice(ctx, *device)
	if response != nil {
		return &empty.Empty{}, nil
	}
	return nil, status.Error(codes.Aborted, "can not create device")
}

func (controller *DevicesController) DeleteDeviceById(ctx context.Context, in *pb.DeviceByIdRequest) (*empty.Empty, error) {
	controller.Usecase.DeleteDeviceById(ctx, in.SensorId)
	return &empty.Empty{}, nil
}

func (controller *DevicesController) GetStatusByLocation(ctx context.Context, in *pb.LocationRequests) (*pb.StatusDataResponse, error) {
	location := entities.Location{Type: entities.Point, Coordinates: []float64{float64(in.Latitude), float64(in.Longitude)}}
	devicesStatus := controller.Usecase.GetStatusByLocation(ctx, location)
	if devicesStatus == nil {
		return nil, status.Error(codes.NotFound, "no data")
	}
	response := utils.ParseStatusToGrpc(*devicesStatus)
	return response, nil
}
