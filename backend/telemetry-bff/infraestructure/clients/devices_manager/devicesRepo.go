package devices_manager

import (
	"context"
	"flag"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/devicesManager/entities"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/devicesManager/repositories"
	"github.com/phuslu/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DevicesManagerGRPCClientRepo struct {
	conn        *grpc.ClientConn
	serviceHost string
}

func NewDevicesManagerGRPCClientRepo(settings *DeviceSettings) repositories.IgrpcClientDeviceRepo {
	serverAddr := flag.String("addrDevicesManager", settings.DeviceManagerHost, "The server address in the format of host:port")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	return &DevicesManagerGRPCClientRepo{
		serviceHost: *serverAddr,
		conn:        conn,
	}
}

func (repo *DevicesManagerGRPCClientRepo) GetDeviceById(ctx context.Context, deviceId string) *entities.DeviceData {
	client := NewDevicesServiceClient(repo.conn)
	request := DeviceByIdRequest{SensorId: deviceId}
	response, err := client.GetDeviceById(ctx, &request)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}
	return ParseDeviceGRPCtoDevice(response)
}

func (repo *DevicesManagerGRPCClientRepo) GetDevicesByLocation(ctx context.Context, locationReq entities.Location) *[]entities.DeviceData {
	client := NewDevicesServiceClient(repo.conn)
	location := ParseLocationToGRPC(&locationReq)
	response, err := client.GetDevicesByLocation(ctx, location)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}
	return ParseDevicesGRPCtoDevices(response)
}

func (repo *DevicesManagerGRPCClientRepo) GetStatusByLocation(ctx context.Context, locationReq entities.Location) *[]entities.Status {
	client := NewDevicesServiceClient(repo.conn)
	location := ParseLocationToGRPC(&locationReq)
	response, err := client.GetStatusByLocation(ctx, location)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}
	return ParseStatusDataResponseToStatus(response)
}
