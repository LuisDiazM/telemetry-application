package server

import (
	"github.com/LuisDiazM/backend/users-manager/infraestructure/server/controllers"
	"github.com/LuisDiazM/backend/users-manager/infraestructure/server/protos/users"
	"google.golang.org/grpc"
)

type ServerGRPCCustom struct {
	UserController *controllers.UsersController
}

func NewServerGRPCCustom(userController *controllers.UsersController) *ServerGRPCCustom {
	return &ServerGRPCCustom{
		UserController: userController,
	}
}

func (srv *ServerGRPCCustom) RegisterControllers(s *grpc.Server) {
	users.RegisterUsersServiceServer(s, srv.UserController)
}
