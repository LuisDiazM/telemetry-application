package utils

import (
	"errors"

	"github.com/LuisDiazM/backend/users-manager/domain/users/entities"
	pb "github.com/LuisDiazM/backend/users-manager/infraestructure/server/protos/users"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func newSaveUserRequest(userPb *pb.SaveUserRequest) entities.SaveUserRequest {
	return entities.SaveUserRequest{
		ExternalId:  userPb.ExternalId,
		FullName:    userPb.FullName,
		Email:       userPb.Email,
		Preferences: userPb.Preferences,
		Role:        userPb.Role,
		Status:      userPb.Status,
	}
}

func ParseSaveUser(userPb *pb.SaveUserRequest) (entities.SaveUserRequest, error) {
	if userPb == nil {
		return entities.SaveUserRequest{}, errors.New("userPb is nil")
	}
	return newSaveUserRequest(userPb), nil
}

func ParseUserResponse(user *entities.UserProfile) (*pb.UserActiveResponse, error) {
	if user == nil {
		return nil, errors.New("user empty")
	}
	return &pb.UserActiveResponse{ExternalId: user.ExternalId,
		FullName:    user.FullName,
		Email:       user.Email,
		Preferences: &pb.UserPreferences{},
		CreatedAt:   timestamppb.New(user.CreatedAt),
		UpdatedAt:   timestamppb.New(user.UpdatedAt),
		Role:        string(user.Role),
		Status:      string(user.Status)}, nil
}
