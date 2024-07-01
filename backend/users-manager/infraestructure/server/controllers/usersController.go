package controllers

import (
	"context"
	"log"

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

func (s *UsersController) SayHello(context.Context, *pb.PingRequest) (*pb.PingReply, error) {
	log.Println(s.Name)
	return &pb.PingReply{Message: "hey there!!"}, nil
}
