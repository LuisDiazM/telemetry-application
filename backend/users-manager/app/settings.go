package app

import (
	"log"

	"github.com/LuisDiazM/backend/users-manager/infraestructure/persistence/mongo"
	"github.com/kelseyhightower/envconfig"
)

type AppSettings struct {
	Mongo *mongo.MongoSettings
}

func GetAppSettings() *AppSettings {
	var s AppSettings
	err := envconfig.Process("backups", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &s
}

func GetMongoSettings() *mongo.MongoSettings {
	return GetAppSettings().Mongo
}
