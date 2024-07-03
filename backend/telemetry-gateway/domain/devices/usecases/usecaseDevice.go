package usecases

import (
	"context"

	"github.com/LuisDiazM/backend/telemetry-gateway/domain/devices/entities"
	"github.com/LuisDiazM/backend/telemetry-gateway/domain/devices/repositories"
	"github.com/LuisDiazM/backend/telemetry-gateway/domain/devices/usecases/utils"
)

type DeviceUsecase struct {
	DbRepository    repositories.IDbRepository
	CacheRepository repositories.IDeviceCacheRepo
}

func NewDeviceUsecase(dbRepo repositories.IDbRepository, cacheRepo repositories.IDeviceCacheRepo) *DeviceUsecase {
	return &DeviceUsecase{
		DbRepository:    dbRepo,
		CacheRepository: cacheRepo,
	}
}

func (usecase *DeviceUsecase) SaveDeviceRegistries(ctx context.Context, dataRequest entities.DeviceDataRequest) {

	deviceData := usecase.CacheRepository.FindDeviceBySensorId(dataRequest.DeviceId)
	if deviceData == nil {
		deviceData = usecase.DbRepository.GetDeviceData(ctx, dataRequest.DeviceId)
		usecase.CacheRepository.SetDeviceData(dataRequest.DeviceId, *deviceData)
	}

	deviceMeasures := utils.MapDeviceData(dataRequest, *deviceData)
	if deviceMeasures == nil {
		return
	}

	usecase.DbRepository.SaveDeviceMeasureData(ctx, *deviceMeasures)
}
