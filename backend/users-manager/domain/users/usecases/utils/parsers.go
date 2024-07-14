package utils

import (
	"errors"
	"time"

	"github.com/LuisDiazM/backend/users-manager/domain/users/entities"
)

func NewparseUser(user *entities.SaveUserRequest) (entities.UserProfile, error) {
	if user == nil {
		return entities.UserProfile{}, errors.New("user already exists")
	}
	return entities.UserProfile{ExternalId: user.ExternalId,
		FullName:    user.FullName,
		Email:       user.Email,
		Preferences: user.Preferences,
		Role:        entities.Role(user.Role),
		Status:      entities.Status(user.Status),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}, nil
}
