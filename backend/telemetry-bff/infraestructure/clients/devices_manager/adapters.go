package devices_manager

import "github.com/LuisDiazM/backend/telemetry-bff/domain/devicesManager/entities"

func ParseDeviceGRPCtoDevice(device *DeviceData) *entities.DeviceData {
	if device == nil {
		return nil
	} else {
		return &entities.DeviceData{SensorId: device.SensorId, Location: entities.Location{Type: entities.Point, Coordinates: []float64{float64(device.Location.Coordiantes[0]), float64(device.Location.Coordiantes[1])}}, TypeSensor: device.TypeSensor, Units: device.Units}
	}
}

func ParseLocationToGRPC(location *entities.Location) *LocationRequests {
	if location == nil {
		return nil
	}
	return &LocationRequests{Latitude: float32(location.Coordinates[0]), Longitude: float32(location.Coordinates[1])}
}

func ParseDevicesGRPCtoDevices(devices *DevicesDataResponse) *[]entities.DeviceData {
	if devices == nil {
		return nil
	}
	var devicesEnt []entities.DeviceData
	for _, dev := range devices.Devices {
		device := entities.DeviceData{SensorId: dev.SensorId, TypeSensor: dev.TypeSensor, Units: dev.Units, Location: entities.Location{Type: entities.Point, Coordinates: []float64{float64(dev.Location.Coordiantes[0]), float64(dev.Location.Coordiantes[1])}}}
		devicesEnt = append(devicesEnt, device)
	}
	return &devicesEnt
}

func ParseStatusDataResponseToStatus(status *StatusDataResponse) *[]entities.Status {
	if status == nil {
		return nil
	}
	var statusData []entities.Status
	for _, st := range status.Status {
		device := entities.DeviceData{SensorId: st.Device.SensorId, TypeSensor: st.Device.TypeSensor, Units: st.Device.Units, Location: entities.Location{Type: entities.Point, Coordinates: []float64{float64(st.Device.Location.Coordiantes[0]), float64(st.Device.Location.Coordiantes[1])}}}
		statusEntity := entities.Status{LastPing: st.LastPing.AsTime(), Device: device}
		statusData = append(statusData, statusEntity)
	}
	return &statusData
}
