package repositories

import (
	"context"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/devicesManager/entities"
)

type IgrpcClientDeviceRepo interface {
	GetDeviceById(ctx context.Context, deviceId string) *entities.DeviceData
	GetDevicesByLocation(ctx context.Context, locationReq entities.Location) *[]entities.DeviceData
	GetStatusByLocation(ctx context.Context, locationReq entities.Location) *[]entities.Status
}
