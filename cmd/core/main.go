package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ThisJohan/snapp-assignment/api"
	"github.com/ThisJohan/snapp-assignment/pkg/env"
	"github.com/ThisJohan/snapp-assignment/pkg/grpcext"
)

const (
	service = "core"
)

type Config struct {
	Grpc            grpcext.Config
	DeliveryService string `env:"DELIVERY_SERVICE" envDefault:"delivery:6565"`
}

var (
	configs Config
)

func init() {
	if err := env.LoadConfig(service, &configs); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}

func main() {
	shipping := api.NewShippingServiceClient(grpcext.NewConnection(configs.DeliveryService))

	x, err := shipping.Create(context.Background(), &api.CreateShipmentRequest{
		OrderID:     "test_id",
		UserAddress: "test address",
		Origin:      "test origin",
		Destination: "test destination",
		TimeSlot:    time.Now().Add(time.Hour * 3).Unix(),
	})
	fmt.Println(x, err)
}
