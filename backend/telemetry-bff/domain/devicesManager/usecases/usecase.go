package usecases

import (
	"context"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/devicesManager/entities"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/devicesManager/repositories"
)

type DeviceManagerUsecase struct {
	grpcDeviceClient repositories.IgrpcClientDeviceRepo
}

func NewDeviceManagerUsecase(grpcDeviceClient repositories.IgrpcClientDeviceRepo) *DeviceManagerUsecase {
	return &DeviceManagerUsecase{grpcDeviceClient: grpcDeviceClient}
}

func (usecase *DeviceManagerUsecase) GetDeviceById(ctx context.Context, deviceId string) *entities.DeviceData {
	return usecase.grpcDeviceClient.GetDeviceById(ctx, deviceId)
}

func (usecase *DeviceManagerUsecase) GetDevicesByLocation(ctx context.Context, location entities.Location) *[]entities.DeviceData {
	return usecase.grpcDeviceClient.GetDevicesByLocation(ctx, location)
}

func (usecase *DeviceManagerUsecase) GetStatusDevicesByLocation(ctx context.Context, location entities.Location) *[]entities.Status {
	return usecase.grpcDeviceClient.GetStatusByLocation(ctx, location)
}
