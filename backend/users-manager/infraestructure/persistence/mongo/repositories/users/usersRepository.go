package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/LuisDiazM/backend/users-manager/domain/users/entities"
	"github.com/LuisDiazM/backend/users-manager/domain/users/repositories"
	"github.com/phuslu/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/LuisDiazM/backend/users-manager/infraestructure/persistence/mongo"
)

const (
	DATABASE                = "users"
	COLLECTION_USER_PROFILE = "users_profile"
)

type UserDBRepository struct {
	mongo *mongo.MongoImplementation
}

func NewUserDBRepository(mongo *mongo.MongoImplementation) repositories.IUsersDbRepository {
	return &UserDBRepository{mongo: mongo}
}

func (repo *UserDBRepository) SaveUser(ctx context.Context, user entities.UserProfile) (string, error) {
	collection := repo.mongo.GetCollection(DATABASE, COLLECTION_USER_PROFILE)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", fmt.Errorf(`document for user %s`, user.Email)
	}
	if result != nil {
		docId, ok := result.InsertedID.(primitive.ObjectID)
		if !ok {
			return "", errors.New("problem decoded objectId in in SaveUser")
		}
		return docId.String(), nil
	}
	return "", nil
}

func (repo *UserDBRepository) GetUserById(ctx context.Context, userId string) (*entities.UserProfile, error) {
	collection := repo.mongo.GetCollection(DATABASE, COLLECTION_USER_PROFILE)
	filter := bson.D{{Key: "external_id", Value: userId}}
	singleResult := collection.FindOne(ctx, filter)

	var user entities.UserProfile

	err := singleResult.Decode(&user)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, fmt.Errorf(`error with user %s`, userId)
	}

	return &user, nil

}

func (repo *UserDBRepository) DeleteUserById(ctx context.Context, userId string) error {
	collection := repo.mongo.GetCollection(DATABASE, COLLECTION_USER_PROFILE)
	filter := bson.D{{Key: "external_id", Value: userId}}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error().Msg(err.Error())
		return fmt.Errorf(`error deleting user %s`, userId)
	}
	return nil
}
