package usersManager

import (
	"context"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager/entities"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager/repositories"
)

type UsersManagerUsecase struct {
	grpcClient repositories.IUsersManagerClient
	cacheRepo  repositories.IUsersManagerCacheRepo
}

func NewUsersManagerUsecase(grpcClient repositories.IUsersManagerClient, cacheRepo repositories.IUsersManagerCacheRepo) *UsersManagerUsecase {
	return &UsersManagerUsecase{
		grpcClient: grpcClient,
		cacheRepo:  cacheRepo,
	}
}

func (usecase *UsersManagerUsecase) GetUserData(ctx context.Context, userId string) *entities.UserProfile {
	userProfile := usecase.cacheRepo.GetUserByExternalId(userId)
	if userProfile == nil {
		userProfile = usecase.grpcClient.GetClientById(ctx, userId)
		if userProfile != nil {
			usecase.cacheRepo.SetUser(*userProfile)
		}
	}
	return userProfile
}
