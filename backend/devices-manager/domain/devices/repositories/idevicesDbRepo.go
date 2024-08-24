package repositories

import (
	"context"

	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/domain/devices/entities"
)

type IDevicesDBRepo interface {
	GetDeviceById(ctx context.Context, deviceId string) *entities.DeviceData
	GetDevicesByLocation(ctx context.Context, location entities.Location) *[]entities.DeviceData
	CreateDevice(ctx context.Context, device entities.DeviceData) bool
	DeleteDeviceById(ctx context.Context, deviceId string) bool
	UpdateDeviceStatus(ctx context.Context, sensorId string) bool
	GetStatusByLocation(ctx context.Context, location entities.Location) *[]entities.Status
}
