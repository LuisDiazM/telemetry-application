package usersManager

import (
	"time"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager/entities"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager/repositories"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/cache"
)

type UsersManagerCacheRepo struct {
	Redis *cache.RedisImp
}

const (
	PREFIX = "TELEMETRY_BFF_USERS"
)

func NewUsersManagerCacheRepo(cache *cache.RedisImp) repositories.IUsersManagerCacheRepo {
	return &UsersManagerCacheRepo{Redis: cache}
}

func (repo *UsersManagerCacheRepo) GetUserByExternalId(userId string) *entities.UserProfile {
	var userProfile entities.UserProfile
	err := repo.Redis.Get(PREFIX, userId, &userProfile)
	if err != nil {
		return nil
	}
	return &userProfile
}

func (repo *UsersManagerCacheRepo) SetUser(userProfile entities.UserProfile) {
	repo.Redis.Set(PREFIX, userProfile.ExternalId, userProfile, 24*time.Hour)
}
