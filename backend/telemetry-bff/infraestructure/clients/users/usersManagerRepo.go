package users

import (
	"context"
	"flag"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager/entities"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/usersManager/repositories"
	"github.com/phuslu/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UsersManagerGrpcClient struct {
	conn        *grpc.ClientConn
	serviceHost string
}

func NewUsersManagerGrpcClient(settings *UsersManagerServiceSettings) repositories.IUsersManagerClient {
	serverAddr := flag.String("addr", settings.UsersManagerHost, "The server address in the format of host:port")

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Error().Msg(err.Error())
	}

	return &UsersManagerGrpcClient{conn: conn, serviceHost: *serverAddr}
}

func (repo *UsersManagerGrpcClient) GetClientById(ctx context.Context, userId string) *entities.UserProfile {
	client := NewUsersServiceClient(repo.conn)
	userRequest := &UserRequest{UserId: userId}
	response, err := client.GetActiveUserByExternalId(ctx, userRequest)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}
	userProfile := MapGRPCUserToUserProfile(response)
	return userProfile
}

func (repo *UsersManagerGrpcClient) CloseConn() {
	repo.conn.Close()
}
