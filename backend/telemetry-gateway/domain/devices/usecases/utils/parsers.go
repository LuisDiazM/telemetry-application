package utils

import (
	"time"

	"github.com/LuisDiazM/backend/telemetry-gateway/domain/devices/entities"
)

const (
	minHumidity    = 0
	minTemperature = -60.0
	maxTemperature = 60.0
)

func MapDeviceData(requestData entities.DeviceDataRequest, deviceMetaData entities.MetaInfo) *entities.DeviceMeasures {
	if !isDataValid(requestData) {
		return nil
	}
	timestamp := time.Now().UTC()
	return &entities.DeviceMeasures{
		Metadata:  deviceMetaData,
		Timestamp: timestamp,
		Values: entities.SensorValues{
			Temperature: requestData.Temperature,
			Humidity:    requestData.Humidity,
		},
	}
}

func isDataValid(requestData entities.DeviceDataRequest) bool {
	if requestData.Humidity < minHumidity || requestData.Humidity > 100 || requestData.Temperature < minTemperature || requestData.Temperature > maxTemperature {
		return false
	}
	return true
}
