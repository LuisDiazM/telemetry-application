package devices

import (
	"context"

	"github.com/LuisDiazM/backend/telemetry-gateway/domain/devices/entities"
	"github.com/LuisDiazM/backend/telemetry-gateway/domain/devices/repositories"
	"github.com/LuisDiazM/backend/telemetry-gateway/infraestructure/database"
	"github.com/phuslu/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DATABASE_DEVICES    = "devices-data"
	COLLECTION_DEVICES  = "devices"
	COLLECTION_MEASURES = "measures"
)

type DevicesDbRepository struct {
	Mongo *database.MongoImplementation
}

func NewDevicesDbRepository(mongo *database.MongoImplementation) repositories.IDbRepository {
	return &DevicesDbRepository{
		Mongo: mongo,
	}
}

func (rep *DevicesDbRepository) GetDeviceData(ctx context.Context, sensorId string) *entities.MetaInfo {
	coll := rep.Mongo.GetCollection(DATABASE_DEVICES, COLLECTION_DEVICES)
	filter := bson.D{{Key: "sensor_id", Value: sensorId}}
	var sensorInfo entities.MetaInfo
	err := coll.FindOne(ctx, filter).Decode(&sensorInfo)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}
	return &sensorInfo
}

func (rep *DevicesDbRepository) SaveDeviceMeasureData(ctx context.Context, input entities.DeviceMeasures) string {
	coll := rep.Mongo.GetCollection(DATABASE_DEVICES, COLLECTION_MEASURES)
	result, err := coll.InsertOne(ctx, input)
	if err != nil {
		log.Error().Msg(err.Error())
		return ""
	}
	return result.InsertedID.(primitive.ObjectID).String()

}
