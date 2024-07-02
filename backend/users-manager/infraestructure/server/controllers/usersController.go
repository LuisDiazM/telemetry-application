package controllers

import (
	"context"
	"fmt"

	pb "github.com/LuisDiazM/backend/users-manager/infraestructure/server/protos/users"
)

type UsersController struct {
	pb.UsersServiceServer
	Name string
}

func NewUsersController() *UsersController {
	return &UsersController{
		Name: "Users Controller",
	}
}

func (s *UsersController) SayHello(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	return &pb.PingReply{Message: fmt.Sprintf(`Hey %s!!!`, in.Name)}, nil
}
