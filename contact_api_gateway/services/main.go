package services

import (
	"fmt"

	"bitbucket.org/udevs/example_api_gateway/config"
	"google.golang.org/grpc"
)

type ServiceManager interface {
	// ProfessionService() position_service.ProfessionServiceClient
	
}

type grpcClients struct {
	// professionService position_service.ProfessionServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connPositionService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PositionServiceHost, conf.PositionServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		// professionService: position_service.NewProfessionServiceClient(connPositionService),
	}, nil
}

// func (g *grpcClients) ProfessionService() position_service.ProfessionServiceClient {
// 	return g.professionService
// }
