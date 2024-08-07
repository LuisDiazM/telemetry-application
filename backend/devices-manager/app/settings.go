package app

import (
	"github.com/phuslu/log"

	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/messaging"
	"github.com/kelseyhightower/envconfig"
)

type AppSettings struct {
	Port   string `envconfig:"PORT" default:"50052"`
	HiveMQ *messaging.HiveMQSettings
}

func GetAppSettings() *AppSettings {
	var s AppSettings
	err := envconfig.Process("backups", &s)
	if err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
	return &s
}

func GetHiveMQSettings() *messaging.HiveMQSettings {
	return GetAppSettings().HiveMQ
}
