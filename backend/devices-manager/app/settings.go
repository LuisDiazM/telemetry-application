package app

import (
	"github.com/phuslu/log"

	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/database/mongodb"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/messaging"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/pusher"
	"github.com/kelseyhightower/envconfig"
)

type AppSettings struct {
	Port    string `envconfig:"PORT" default:"50052"`
	HiveMQ  *messaging.HiveMQSettings
	MongoDB *mongodb.MongoSettings
	Pusher  *pusher.PusherSettings
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

func GetMongoSettings() *mongodb.MongoSettings {
	return GetAppSettings().MongoDB
}

func GetPusherSettings() *pusher.PusherSettings {
	return GetAppSettings().Pusher
}
