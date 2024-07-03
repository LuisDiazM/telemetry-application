package repositories

import "github.com/LuisDiazM/backend/telemetry-gateway/domain/devices/entities"

type IDeviceCacheRepo interface {
	FindDeviceBySensorId(sensorId string) *entities.MetaInfo
	SetDeviceData(sensorId string, deviceData entities.MetaInfo)
}
