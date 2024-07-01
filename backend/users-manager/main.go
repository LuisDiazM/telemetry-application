package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/LuisDiazM/backend/users-manager/infraestructure/server"
	"github.com/LuisDiazM/backend/users-manager/infraestructure/server/controllers"
	pb "github.com/LuisDiazM/backend/users-manager/infraestructure/server/protos/users"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type Server struct {
	pb.UsersServiceServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	germa := controllers.NewUsersController()

	srvCustom := server.NewServerGRPCCustom(germa)
	srvCustom.RegisterControllers(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("faile to serve %v\n", err)
	}
}
