package entities

import "time"

type Location struct {
	TypeL       string    `json:"type,omitempty" bson:"type"`
	Coordinates []float64 `json:"coordinates,omitempty" bson:"coordinates"`
}

type MetaInfo struct {
	SensorId   string   `json:"sensor_id,omitempty" bson:"sensor_id"`
	TypeSensor string   `json:"type_sensor,omitempty" bson:"type_sensor"`
	Location   Location `json:"location,omitempty" bson:"location"`
	Units      []string `json:"units,omitempty" bson:"units"`
}

type SensorValues struct {
	Temperature float32 `json:"temperature,omitempty" bson:"temperature"`
	Humidity    float32 `json:"humidity,omitempty" bson:"humidity"`
}

type DeviceMeasures struct {
	Metadata  MetaInfo     `json:"metadata,omitempty" bson:"metadata"`
	Timestamp time.Time    `json:"timestamp,omitempty" bson:"timestamp"`
	Values    SensorValues `json:"values,omitempty" bson:"values"`
}
