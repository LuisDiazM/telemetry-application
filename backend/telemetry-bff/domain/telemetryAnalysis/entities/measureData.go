package entities

import "time"

type MeasureData struct {
	Values         []Value        `json:"values"`
	SensorMetadata SensorMetadata `json:"sensor_metadata"`
}

type SensorMetadata struct {
	Units      []string `json:"units"`
	SensorID   string   `json:"sensor_id"`
	TypeSensor string   `json:"type_sensor"`
	Location   Location `json:"location"`
}

type Location struct {
	Coordinates []float64 `json:"coordinates"`
	Type        string    `json:"type"`
}

type Value struct {
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	Timestamp   time.Time `json:"timestamp"`
}
