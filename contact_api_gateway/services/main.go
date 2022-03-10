package services

import (
	"contact_api_gateway/config"
	"contact_api_gateway/genproto/user_service"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceManager interface {
	UserService() user_service.UserServiceClient
}

type grpcClients struct {
	userService user_service.UserServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connUserService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		userService: user_service.NewUserServiceClient(connUserService),
	}, nil
}

func (g *grpcClients) UserService() user_service.UserServiceClient {
	return g.userService
}
