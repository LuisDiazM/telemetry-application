package entities

type DeviceDataRequest struct {
	DeviceId    string  `json:"device_id,omitempty"`
	Temperature float32 `json:"temperature,omitempty"`
	Humidity    float32 `json:"humidity,omitempty"`
}
