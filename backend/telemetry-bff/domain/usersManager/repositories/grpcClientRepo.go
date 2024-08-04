package repositories

import (
	"context"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager/entities"
)

type IUsersManagerClient interface {
	GetClientById(ctx context.Context, userId string) *entities.UserProfile
	SaveUser(ctx context.Context, userData *entities.SaveUserData) (*entities.SaveUserResponse, error)
}
