package manager

import (
	"context"
	"fmt"
	"time"

	"github.com/ThisJohan/delivery-microservice/api"
	"github.com/ThisJohan/delivery-microservice/pkg/grpcext"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

type Worker struct {
	ctx             context.Context
	shippingService api.ShippingServiceClient
}

func NewWorker(config Config) *Worker {
	return &Worker{
		shippingService: api.NewShippingServiceClient(grpcext.NewConnection(config.DeliveryService)),
	}
}

func (w *Worker) RegisterWorker(ctx context.Context) {
	w.ctx = ctx

	Cron(w)
}

func Cron(w *Worker) {
	s, _ := gocron.NewScheduler()

	s.NewJob(
		gocron.DurationJob(
			time.Second*15,
		),
		gocron.NewTask(w.Seed),
	)
	s.Start()

	select {}
}

func (w *Worker) Seed() error {
	shipment, err := w.shippingService.Create(context.Background(), &api.CreateShipmentRequest{
		OrderID:     uuid.New().String(),
		UserAddress: "test address",
		Origin:      "test origin",
		Destination: "test destination",
		TimeSlot:    time.Now().Add(time.Minute * time.Duration(randInt(30, 90))).Unix(),
	})
	fmt.Printf("Created shipment: %v\n", shipment)
	return err
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}
