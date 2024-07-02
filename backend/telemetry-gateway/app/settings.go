package app

import (
	"log"

	"github.com/LuisDiazM/backend/telemetry-gateway/infraestructure/messaging"
	"github.com/kelseyhightower/envconfig"
)

type AppSettings struct {
	Port   string `envconfig:"PORT" default:"50051"`
	HiveMQ *messaging.HiveMQSettings
}

func GetAppSettings() *AppSettings {
	var s AppSettings
	err := envconfig.Process("backups", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &s
}

func GetHiveMQSettings() *messaging.HiveMQSettings {
	return GetAppSettings().HiveMQ
}
