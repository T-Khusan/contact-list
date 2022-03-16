package services

import (
	"contact_api_gateway/config"
	"contact_api_gateway/genproto/contact_service"
	"contact_api_gateway/genproto/user_service"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceManager interface {
	ContactService() contact_service.ContactServiceClient
	UserService() user_service.UserServiceClient
}

type grpcClients struct {
	contactService contact_service.ContactServiceClient
	userService    user_service.UserServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connContactService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.ContactServiceHost, conf.ContactServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connUserService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		contactService: contact_service.NewContactServiceClient(connContactService),
		userService:    user_service.NewUserServiceClient(connUserService),
	}, nil
}

func (g *grpcClients) ContactService() contact_service.ContactServiceClient {
	return g.contactService
}

func (g *grpcClients) UserService() user_service.UserServiceClient {
	return g.userService
}
