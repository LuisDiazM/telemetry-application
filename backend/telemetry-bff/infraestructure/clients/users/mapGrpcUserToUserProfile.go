package users

import (
	"github.com/phuslu/log"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager/entities"
)

func MapGRPCUserToUserProfile(userResponse *UserActiveResponse) *entities.UserProfile {
	if userResponse == nil {
		log.Warn().Msg("grpc response nill")
		return nil
	}
	userProfile := entities.UserProfile{ExternalId: userResponse.ExternalId,
		FullName:    userResponse.FullName,
		Email:       userResponse.Email,
		Preferences: userResponse.Preferences,
		CreatedAt:   userResponse.CreatedAt.AsTime(),
		UpdatedAt:   userResponse.UpdatedAt.AsTime(),
		Role:        entities.Role(userResponse.Role),
		Status:      entities.Status(userResponse.Status)}
	return &userProfile
}
