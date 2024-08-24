package devices

import (
	"context"
	"log"
	"time"

	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/domain/devices/entities"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/domain/devices/repositories"
	"github.com/LuisDiazM/telemetry-application/backend/devices-manager/infraestructure/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type DevicesDbRepository struct {
	mongo *mongodb.MongoImplementation
}

func NewDevicesDbRepository(mongo *mongodb.MongoImplementation) repositories.IDevicesDBRepo {
	return &DevicesDbRepository{
		mongo: mongo,
	}
}

func (repo *DevicesDbRepository) GetDeviceById(ctx context.Context, deviceId string) *entities.DeviceData {
	coll := repo.mongo.GetCollection(DEVICES_DATABASE, DEVICES_COLLECTION)
	filter := bson.D{{Key: "sensor_id", Value: deviceId}}
	cursor := coll.FindOne(ctx, filter)
	var result entities.DeviceData
	if cursor == nil {
		return nil
	}
	err := cursor.Decode(&result)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &result
}

func getGeoNearAggregation(location entities.Location) bson.A {
	return bson.A{
		bson.D{
			{Key: "$geoNear",
				Value: bson.D{
					{Key: "near",
						Value: bson.D{
							{Key: "type", Value: location.Type},
							{Key: "coordinates",
								Value: bson.A{
									location.Coordinates[0],
									location.Coordinates[1],
								},
							},
						},
					},
					{Key: "distanceField", Value: "distance"},
					{Key: "maxDistance", Value: 20000},
					{Key: "query", Value: bson.D{}},
					{Key: "spherical", Value: false},
				},
			},
		},
	}
}
func (repo *DevicesDbRepository) GetDevicesByLocation(ctx context.Context, location entities.Location) *[]entities.DeviceData {
	coll := repo.mongo.GetCollection(DEVICES_DATABASE, DEVICES_COLLECTION)
	aggregation := getGeoNearAggregation(location)
	var results []entities.DeviceData
	cursor, err := coll.Aggregate(ctx, aggregation)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Println(err.Error())
		return nil
	}

	return &results
}

func (repo *DevicesDbRepository) CreateDevice(ctx context.Context, device entities.DeviceData) bool {
	coll := repo.mongo.GetCollection(DEVICES_DATABASE, DEVICES_COLLECTION)
	_, err := coll.InsertOne(ctx, device)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func (repo *DevicesDbRepository) DeleteDeviceById(ctx context.Context, deviceId string) bool {
	coll := repo.mongo.GetCollection(DEVICES_DATABASE, DEVICES_COLLECTION)
	filter := bson.D{{Key: "sensor_id", Value: deviceId}}
	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return false
	}
	if result.DeletedCount == 0 {
		return false
	}
	return true
}

func (repo *DevicesDbRepository) UpdateDeviceStatus(ctx context.Context, sensorId string) bool {
	coll := repo.mongo.GetCollection(DEVICES_DATABASE, STATUS_COLLECTION)
	filter := bson.D{{Key: "device.sensor_id", Value: sensorId}}
	updateDoc := bson.D{{Key: "$set", Value: bson.D{{Key: "last_ping", Value: time.Now().UTC()}}}}

	result, err := coll.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	matchCount := result.MatchedCount
	if matchCount == 0 {
		collDevices := repo.mongo.GetCollection(DEVICES_DATABASE, DEVICES_COLLECTION)
		result := collDevices.FindOne(ctx, bson.D{{Key: "sensor_id", Value: sensorId}})
		if result != nil {
			var deviceData entities.DeviceData
			err := result.Decode(&deviceData)
			if err != nil {
				return false
			}
			var statusDocument entities.Status = entities.Status{LastPing: time.Now().UTC(), Device: deviceData}
			_, err = coll.InsertOne(ctx, statusDocument)
			if err != nil {
				return false
			}
		}
	}

	return true
}

func (repo *DevicesDbRepository) GetStatusByLocation(ctx context.Context, location entities.Location) *[]entities.Status {
	coll := repo.mongo.GetCollection(DEVICES_DATABASE, STATUS_COLLECTION)
	agg := getGeoNearAggregation(location)
	var results []entities.Status

	cursor, err := coll.Aggregate(ctx, agg)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Println(err.Error())
		return nil
	}

	return &results
}
