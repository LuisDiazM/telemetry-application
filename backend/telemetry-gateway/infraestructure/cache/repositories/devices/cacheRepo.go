package devices

import (
	"time"

	"github.com/LuisDiazM/backend/telemetry-gateway/domain/devices/entities"
	"github.com/LuisDiazM/backend/telemetry-gateway/domain/devices/repositories"
	"github.com/LuisDiazM/backend/telemetry-gateway/infraestructure/cache"
)

type CacheRepositoryDevices struct {
	Redis *cache.RedisImp
}

const (
	DEVICES_PREFIX = "devices"
)

func NewCacheDeviceRepository(cache *cache.RedisImp) repositories.IDeviceCacheRepo {
	return &CacheRepositoryDevices{Redis: cache}
}

func (repo *CacheRepositoryDevices) FindDeviceBySensorId(sensorId string) *entities.MetaInfo {
	var sensorData entities.MetaInfo
	err := repo.Redis.Get(DEVICES_PREFIX, sensorId, &sensorData)
	if err != nil {
		return nil
	}
	return &sensorData
}

func (repo *CacheRepositoryDevices) SetDeviceData(sensorId string, deviceData entities.MetaInfo) {
	repo.Redis.Set(DEVICES_PREFIX, sensorId, deviceData, 24*time.Hour)
}
