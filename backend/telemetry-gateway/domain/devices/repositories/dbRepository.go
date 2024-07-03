package repositories

import (
	"context"

	"github.com/LuisDiazM/backend/telemetry-gateway/domain/devices/entities"
)

type IDbRepository interface {
	GetDeviceData(ctx context.Context, sensorId string) *entities.MetaInfo
	SaveDeviceMeasureData(ctx context.Context, input entities.DeviceMeasures) string
}
