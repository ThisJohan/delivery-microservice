package uber

import (
	"context"
	"fmt"

	"github.com/ThisJohan/snapp-assignment/api"
	"github.com/ThisJohan/snapp-assignment/pkg/grpcext"
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
	fmt.Println("Here")
	if req.ShipmentID%2 == 0 {
		fmt.Println("Here2")
		_, _ = s.shippingService.StatusChange(ctx, &api.ShipmentStatusChange{
			ID:     req.ShipmentID,
			Status: api.Shipment_ASSIGNED,
		})
	}

	return nil, nil
}
