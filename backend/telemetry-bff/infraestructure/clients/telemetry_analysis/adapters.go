package telemetry_analysis

import "github.com/LuisDiazM/backend/telemetry-bff/domain/telemetryAnalysis/entities"

func ParseMeasuresToEntities(response *MeasuresByDeviceResponse) *entities.MeasureData {
	if response == nil {
		return nil
	}
	metadata := entities.SensorMetadata{Units: response.SensorMetadata.Units,
		SensorID:   response.SensorMetadata.SensorId,
		TypeSensor: response.SensorMetadata.TypeSensor,
		Location: entities.Location{Coordinates: response.SensorMetadata.Location.Coordinates,
			Type: response.SensorMetadata.Location.Type}}

	var values []entities.Value
	for _, value := range response.Values {
		val := entities.Value{Temperature: value.Temperature, Humidity: value.Humidity, Timestamp: value.Timestamp.AsTime()}
		values = append(values, val)
	}
	return &entities.MeasureData{Values: values, SensorMetadata: metadata}
}
