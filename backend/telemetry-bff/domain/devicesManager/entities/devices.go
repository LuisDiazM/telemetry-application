package entities

type LocationType string

const (
	Point LocationType = "Point"
)

type Location struct {
	Type        LocationType `json:"type,omitempty"`
	Coordinates []float64    `json:"coordinates,omitempty"`
}
type DeviceData struct {
	SensorId   string   `json:"sensor_id,omitempty" bson:"sensor_id"`
	Location   Location `json:"location,omitempty" bson:"location"`
	TypeSensor string   `json:"type_sensor,omitempty" bson:"type_sensor"`
	Units      []string `json:"units,omitempty" bson:"units"`
}
