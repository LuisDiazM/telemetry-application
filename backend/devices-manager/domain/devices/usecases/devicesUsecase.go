package usecases

import (
	"context"

	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/domain/devices/entities"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/domain/devices/repositories"
)

type DevicesUsecase struct {
	devicesRepoDB repositories.IDevicesDBRepo
}

func NewDevicesUsecase(devicesRepoDb repositories.IDevicesDBRepo) *DevicesUsecase {
	return &DevicesUsecase{devicesRepoDB: devicesRepoDb}
}

func (usecase *DevicesUsecase) GetDevicesFromLocation(ctx context.Context, location entities.Location) *[]entities.DeviceData {
	return usecase.devicesRepoDB.GetDevicesByLocation(ctx, location)
}

func (usecase *DevicesUsecase) GetDeviceById(ctx context.Context, deviceId string) *entities.DeviceData {
	return usecase.devicesRepoDB.GetDeviceById(ctx, deviceId)
}

func (usecase *DevicesUsecase) CreateDevice(ctx context.Context, device entities.DeviceData) *entities.DeviceData {
	isCreated := usecase.devicesRepoDB.CreateDevice(ctx, device)
	if isCreated {
		return &device
	} else {
		return nil
	}
}

func (usecase *DevicesUsecase) DeleteDeviceById(ctx context.Context, deviceId string) bool {
	return usecase.devicesRepoDB.DeleteDeviceById(ctx, deviceId)
}

func (usecase *DevicesUsecase) UpdateDeviceStatus(ctx context.Context, sensorId string) bool {
	return usecase.devicesRepoDB.UpdateDeviceStatus(ctx, sensorId)
}

func (usecase *DevicesUsecase) GetStatusByLocation(ctx context.Context, location entities.Location) *[]entities.Status {
	return usecase.devicesRepoDB.GetStatusByLocation(ctx, location)
}
