package entities

import "time"

type Status struct {
	LastPing time.Time  `json:"last_ping,omitempty" bson:"last_ping"`
	Device   DeviceData `json:"device,omitempty" bson:"device"`
}
