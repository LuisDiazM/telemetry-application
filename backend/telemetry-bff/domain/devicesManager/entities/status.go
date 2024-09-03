package entities

import "time"

type Status struct {
	LastPing time.Time  `json:"last_ping,omitempty"`
	Device   DeviceData `json:"device,omitempty"`
}
