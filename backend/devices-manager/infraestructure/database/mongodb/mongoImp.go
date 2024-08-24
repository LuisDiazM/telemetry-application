package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoImplementation struct {
	Client *mongo.Client
}

func NewMongoImplementation(settings *MongoSettings) *MongoImplementation {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	opt := options.Client()
	opt.SetMaxPoolSize(settings.MaxPoolSize)
	opt.ApplyURI(settings.Url)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		panic(err)
	}
	return &MongoImplementation{Client: client}
}

func (imp *MongoImplementation) GetCollection(databaseName string, collectionName string) *mongo.Collection {
	return imp.Client.Database(databaseName).Collection(collectionName)
}
