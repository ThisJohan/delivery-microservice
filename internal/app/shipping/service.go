package shipping

import (
	"context"
	"time"

	"github.com/ThisJohan/snapp-assignment/api"
	"github.com/ThisJohan/snapp-assignment/pkg/grpcext"
	"google.golang.org/grpc"
)

func NewService(server *grpc.Server, config Config) {
	s := &Service{
		logistics: api.NewLogisticsServiceClient(grpcext.NewConnection(config.TPLService)),
	}
	api.RegisterShippingServiceServer(server, s)
}

type Service struct {
	api.UnimplementedShippingServiceServer
	logistics api.LogisticsServiceClient
}

func (s *Service) Create(ctx context.Context, req *api.CreateShipmentRequest) (*api.Shipment, error) {
	shipment := &Shipment{
		OrderID:     req.OrderID,
		UserAddress: req.UserAddress,
		Origin:      req.Origin,
		Destination: req.Destination,
		TimeSlot:    time.Unix(req.TimeSlot, 0),
		Status:      api.Shipment_PENDING.String(),
	}

	if err := ShipmentsRepo.Create(ctx, shipment); err != nil {
		return nil, err
	}

	if time.Until(shipment.TimeSlot) <= time.Hour {
		// Should notify tpl immediately
		queueShipment(ctx, shipment)
	}

	return &api.Shipment{
		ID:          uint32(shipment.ID),
		OrderID:     shipment.OrderID,
		UserAddress: shipment.UserAddress,
		Origin:      shipment.Origin,
		Destination: shipment.Destination,
		TimeSlot:    shipment.TimeSlot.Unix(),
		Status:      api.ShipmentStatus(api.ShipmentStatus_value[shipment.Status]),
	}, nil
}
