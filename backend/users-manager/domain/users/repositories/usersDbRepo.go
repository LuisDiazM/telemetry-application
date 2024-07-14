package repositories

import (
	"context"

	"github.com/LuisDiazM/backend/users-manager/domain/users/entities"
)

type IUsersDbRepository interface {
	SaveUser(ctx context.Context, user entities.UserProfile) (string, error)
	GetUserById(ctx context.Context, userId string) (*entities.UserProfile, error)
	DeleteUserById(ctx context.Context, userId string) error
}
