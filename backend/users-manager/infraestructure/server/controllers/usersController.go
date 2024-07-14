package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/LuisDiazM/backend/users-manager/domain/users/usecases"
	"github.com/LuisDiazM/backend/users-manager/infraestructure/server/controllers/utils"
	pb "github.com/LuisDiazM/backend/users-manager/infraestructure/server/protos/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UsersController struct {
	pb.UsersServiceServer
	Name        string
	UserUsecase *usecases.UserUsecases
}

func NewUsersController(usecaseUser *usecases.UserUsecases) *UsersController {
	return &UsersController{
		Name:        "Users Controller",
		UserUsecase: usecaseUser,
	}
}

func (s *UsersController) SayHello(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	return &pb.PingReply{Message: fmt.Sprintf(`Hey %s!!!`, in.Name)}, nil
}

func (s *UsersController) GetActiveUserByExternalId(ctx context.Context, in *pb.UserRequest) (*pb.UserActiveResponse, error) {
	user := s.UserUsecase.GetUserByExternalId(ctx, in.UserId)
	if user == nil {

		return nil, status.Error(codes.NotFound, "")
	}
	userData, err := utils.ParseUserResponse(user)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (s *UsersController) SaveUser(ctx context.Context, in *pb.SaveUserRequest) (*pb.MetadataResponse, error) {
	userRequest, err := utils.ParseSaveUser(in)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}
	isSavedUser, err := s.UserUsecase.SaveUser(ctx, userRequest)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}
	return &pb.MetadataResponse{Key: "status", Description: strconv.FormatBool(isSavedUser)}, nil
}

func (s *UsersController) DeleteUser(ctx context.Context, in *pb.UserRequest) (*pb.MetadataResponse, error) {
	err := s.UserUsecase.DeleteUser(ctx, in.UserId)
	if err != nil {

		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.MetadataResponse{Key: "status", Description: fmt.Sprintf(`user %s deleted`, in.UserId)}, nil
}
