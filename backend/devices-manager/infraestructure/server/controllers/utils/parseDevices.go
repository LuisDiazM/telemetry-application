package utils

import (
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/domain/devices/entities"
	pb "github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/server/protos/devices"
)

func ParseDevicesToGrpc(device *entities.DeviceData) (*pb.DeviceData, error) {
	if device == nil {
		return nil, errors.New("device nil")
	}
	return &pb.DeviceData{SensorId: device.SensorId,
		TypeSensor: device.TypeSensor,
		Location:   &pb.Location{Coordiantes: []float32{float32(device.Location.Coordinates[0]), float32(device.Location.Coordinates[1])}},
		Units:      device.Units}, nil
}

func ParseDevicesListToGRPCResponse(devices []entities.DeviceData) *pb.DevicesDataResponse {
	var devicesGrpc pb.DevicesDataResponse
	for _, device := range devices {
		deviceGRPC, err := ParseDevicesToGrpc(&device)
		if err == nil {
			devicesGrpc.Devices = append(devicesGrpc.Devices, deviceGRPC)
		}
	}
	return &devicesGrpc
}

func ParseDeviceGRPCtoDevice(device *pb.DeviceData) *entities.DeviceData {
	if device == nil {
		return nil
	} else {
		return &entities.DeviceData{SensorId: device.SensorId, Location: entities.Location{Type: entities.Point, Coordinates: []float64{float64(device.Location.Coordiantes[0]), float64(device.Location.Coordiantes[1])}}, TypeSensor: device.TypeSensor, Units: device.Units}
	}
}

func ParseStatusToGrpc(status []entities.Status) *pb.StatusDataResponse {
	var response pb.StatusDataResponse
	for _, st := range status {
		device, err := ParseDevicesToGrpc(&st.Device)
		if err == nil {
			response.Status = append(response.Status, &pb.StatusData{LastPing: timestamppb.New(st.LastPing), Device: device})
		}
	}
	return &response
}
