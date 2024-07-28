package usersManager

import (
	"context"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager/entities"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager/repositories"
)

type UsersManagerUsecase struct {
	grpcClient repositories.IUsersManagerClient
}

func NewUsersManagerUsecase(grpcClient repositories.IUsersManagerClient) *UsersManagerUsecase {
	return &UsersManagerUsecase{
		grpcClient: grpcClient,
	}
}

func (usecase *UsersManagerUsecase) GetUserData(ctx context.Context, userId string) *entities.UserProfile {
	userProfile := usecase.grpcClient.GetClientById(ctx, userId)
	return userProfile
}
