package services

import (
	"contact_api_gateway/config"
	"contact_api_gateway/genproto/contact_service"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceManager interface {
	ContactService() contact_service.ContactServiceClient
}

type grpcClients struct {
	contactService contact_service.ContactServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connContactService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.ContactServiceHost, conf.ContactServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		contactService: contact_service.NewContactServiceClient(connContactService),
	}, nil
}

func (g *grpcClients) ContactService() contact_service.ContactServiceClient {
	return g.contactService
}
