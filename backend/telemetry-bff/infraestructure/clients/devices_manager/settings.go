package devices_manager

type DeviceSettings struct {
	DeviceManagerHost string `envconfig:"DEVICES_MANAGER_HOST" required:"true"`
}
