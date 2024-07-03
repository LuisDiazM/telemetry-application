package app

import (
	"log"

	"github.com/LuisDiazM/backend/telemetry-gateway/infraestructure/cache"
	"github.com/LuisDiazM/backend/telemetry-gateway/infraestructure/database"
	"github.com/LuisDiazM/backend/telemetry-gateway/infraestructure/messaging"
	"github.com/kelseyhightower/envconfig"
)

type AppSettings struct {
	Port    string `envconfig:"PORT" default:"50051"`
	HiveMQ  *messaging.HiveMQSettings
	MongoDb *database.MongoSettings
	Cache   *cache.RedisSettings
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

func GetMongoSettings() *database.MongoSettings {
	return GetAppSettings().MongoDb
}

func GetCacheSettings() *cache.RedisSettings {
	return GetAppSettings().Cache
}
