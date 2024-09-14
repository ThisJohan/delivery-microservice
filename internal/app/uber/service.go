package uber

import (
	"context"

	"github.com/ThisJohan/delivery-microservice/api"
	"github.com/ThisJohan/delivery-microservice/pkg/grpcext"
	"google.golang.org/grpc"
)

func NewService(server *grpc.Server, config Config) {
	s := &Service{
		shippingService: api.NewShippingServiceClient(grpcext.NewConnection(config.DeliveryService)),
	}
	api.RegisterUberServiceServer(server, s)
}

type Service struct {
	api.UnimplementedUberServiceServer
	shippingService api.ShippingServiceClient
}

// Deliver implements api.UberServiceServer.
func (s *Service) Deliver(ctx context.Context, req *api.DeliverRequest) (*api.DeliverResponse, error) {
	model := &api.ShipmentStatusChange{
		ID:     req.ShipmentID,
		Status: api.Shipment_ASSIGNED,
	}
	if req.ShipmentID%2 == 0 {
		if req.ShipmentID%3 == 0 {
			model.Status = api.Shipment_DELIVERED
		}
		s.shippingService.StatusChange(ctx, model)
	}

	return nil, nil
}
