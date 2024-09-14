package manager

import (
	"context"
	"fmt"

	"github.com/ThisJohan/delivery-microservice/api"
	"google.golang.org/grpc"
)

func NewService(server *grpc.Server) {
	s := &Service{}
	api.RegisterManagerServiceServer(server, s)
}

type Service struct {
	api.UnimplementedManagerServiceServer
}

func (s *Service) UpdateShipment(ctx context.Context, req *api.UpdateShipmentRequest) (*api.Void, error) {
	fmt.Printf("UpdateShipment: %v\n", req)
	return &api.Void{}, nil
}
