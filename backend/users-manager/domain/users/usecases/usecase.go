package usecases

import (
	"context"
	"fmt"

	"github.com/LuisDiazM/backend/users-manager/domain/users/entities"
	"github.com/LuisDiazM/backend/users-manager/domain/users/repositories"
	"github.com/LuisDiazM/backend/users-manager/domain/users/usecases/utils"
	"github.com/phuslu/log"
)

type UserUsecases struct {
	UsersDBRepo repositories.IUsersDbRepository
}

func NewUserUsecases(dbRepo repositories.IUsersDbRepository) *UserUsecases {
	return &UserUsecases{
		UsersDBRepo: dbRepo,
	}
}

func (usecase *UserUsecases) GetUserByExternalId(ctx context.Context, userId string) *entities.UserProfile {
	user, err := usecase.UsersDBRepo.GetUserById(ctx, userId)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}
	return user
}

func (usecase *UserUsecases) SaveUser(ctx context.Context, userRequest entities.SaveUserRequest) (bool, error) {
	user, _ := usecase.UsersDBRepo.GetUserById(ctx, userRequest.ExternalId)

	if user != nil {
		log.Warn().Msg(fmt.Sprintf(`user %s with email %s already exists!!`, userRequest.ExternalId, userRequest.Email))
		return false, fmt.Errorf(`user %s already exists`, user.Email)
	}

	userData, err := utils.NewparseUser(&userRequest)
	if err != nil {
		log.Warn().Msg(err.Error())
		return false, err
	}
	userId, err := usecase.UsersDBRepo.SaveUser(ctx, userData)
	if userId == "" || err != nil {
		return false, err
	}
	return true, nil
}

func (usecase *UserUsecases) DeleteUser(ctx context.Context, userId string) error {
	return usecase.UsersDBRepo.DeleteUserById(ctx, userId)
}
