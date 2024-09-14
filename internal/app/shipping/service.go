package shipping

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ThisJohan/delivery-microservice/api"
	"github.com/ThisJohan/delivery-microservice/pkg/grpcext"
	"google.golang.org/grpc"
)

func NewService(server *grpc.Server, config Config) {
	s := &Service{
		logistics: api.NewLogisticsServiceClient(grpcext.NewConnection(config.TPLService)),
		manager:   api.NewManagerServiceClient(grpcext.NewConnection(config.ManagerService)),
	}
	api.RegisterShippingServiceServer(server, s)
}

type Service struct {
	api.UnimplementedShippingServiceServer
	logistics api.LogisticsServiceClient
	manager   api.ManagerServiceClient
}

func (s *Service) StatusChange(ctx context.Context, req *api.ShipmentStatusChange) (*api.Void, error) {
	id := strconv.FormatUint(uint64(req.ID), 10)
	// TODO: good idea to delete shipment from redis processes
	shipment, err := ShipmentsRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	shipment.Status = req.Status.String()
	if err := ShipmentsRepo.Update(ctx, shipment); err != nil {
		return nil, err
	}
	_, err = s.manager.UpdateShipment(ctx, &api.UpdateShipmentRequest{
		Shipment: convertShipment(shipment),
	})
	if err != nil {
		// we need a retry queue for this error, but redis is pain in the ass, so we can't do that
		fmt.Printf("Failed to notify manager the shipment: %v\n", err)
	}
	return &api.Void{}, nil
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
		enqueueShipment(ctx, shipment)
	}

	return convertShipment(shipment), nil
}

func convertShipment(shipment *Shipment) *api.Shipment {
	return &api.Shipment{
		ID:          uint32(shipment.ID),
		OrderID:     shipment.OrderID,
		UserAddress: shipment.UserAddress,
		Origin:      shipment.Origin,
		Destination: shipment.Destination,
		TimeSlot:    shipment.TimeSlot.Unix(),
		Status:      api.ShipmentStatus(api.ShipmentStatus_value[shipment.Status]),
	}
}
