package repositories

import "github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager/entities"

type IUsersManagerCacheRepo interface {
	GetUserByExternalId(userId string) *entities.UserProfile
	SetUser(userProfile entities.UserProfile)
}
