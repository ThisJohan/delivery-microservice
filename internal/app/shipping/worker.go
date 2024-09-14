package shipping

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ThisJohan/delivery-microservice/api"
	"github.com/ThisJohan/delivery-microservice/pkg/grpcext"
	"github.com/ThisJohan/delivery-microservice/pkg/redisext"
	"github.com/go-co-op/gocron/v2"
	"github.com/redis/go-redis/v9"
)

const (
	shipmentsStack = "queued_shipments"
)

func NewWorker(config Config) *Worker {
	return &Worker{
		logistics: api.NewLogisticsServiceClient(grpcext.NewConnection(config.TPLService)),
	}
}

type Worker struct {
	ctx       context.Context
	logistics api.LogisticsServiceClient
}

func (w *Worker) RegisterWorker(ctx context.Context) {
	w.ctx = ctx

	go Queue(w)

	Cron(w)
}

func Cron(w *Worker) {
	s, _ := gocron.NewScheduler()

	s.NewJob(
		gocron.DurationJob(
			time.Second*15,
		),
		gocron.NewTask(w.CheckPendingShipments),
	)
	s.NewJob(
		gocron.DurationJob(
			time.Second*10,
		),
		gocron.NewTask(w.CheckOnProcessShipments),
	)
	s.Start()

	select {}
}

func (w *Worker) CheckOnProcessShipments() error {
	rdb := redisext.InjectRedis(w.ctx)
	now := time.Now()
	expiredItems, err := rdb.ZRangeByScore(w.ctx, shipmentsStack, &redis.ZRangeBy{
		Min: "-inf", Max: fmt.Sprintf("%v", now.Unix()),
	}).Result()
	if err != nil {
		return err
	}

	fmt.Printf("Found %d expired shipments\n", len(expiredItems))

	for _, item := range expiredItems {
		shipment, _ := ShipmentsRepo.Get(w.ctx, item)
		if shipment.Status != api.Shipment_QUEUED.String() {
			continue
		}
		if now.Hour() >= 23 {
			shipment.Status = api.Shipment_CANCELED.String()
			ShipmentsRepo.Update(w.ctx, shipment)
		} else {
			enqueueShipment(w.ctx, shipment)
		}
	}

	// delete expired shipments
	return rdb.ZRemRangeByScore(w.ctx, shipmentsStack, "-inf", fmt.Sprintf("%v", now.Unix())).Err()
}

func (w *Worker) CheckPendingShipments() error {
	shipments, err := ShipmentsRepo.GetPendingShipments(w.ctx, time.Hour)
	if err != nil {
		return err
	}
	fmt.Printf("Found %d pending shipments\n", len(shipments))
	for _, shipment := range shipments {
		if err := enqueueShipment(w.ctx, &shipment); err != nil {
			fmt.Printf("Failed to enqueue shipment: %v\n", err)
		}

	}
	return nil
}

func Queue(w *Worker) {
	rdb := redisext.InjectRedis(w.ctx)
	pubsub := rdb.Subscribe(w.ctx, "tpl_queue")
	ch := pubsub.Channel()

	for msg := range ch {
		var shipment Shipment
		json.Unmarshal([]byte(msg.Payload), &shipment)
		var zero Shipment
		if shipment == zero {
			continue
		}

		if _, err := w.logistics.Deliver(w.ctx, &api.DeliverRequest{
			ShipmentID: uint32(shipment.ID),
		}); err != nil {
			// Retry until success
			// TODO: this could continue forever, better have some max retry count
			enqueueShipment(w.ctx, &shipment)
		}

		expiration := time.Now().Add(time.Second * 15).Unix()
		rdb.ZAdd(w.ctx, shipmentsStack, redis.Z{
			Score:  float64(expiration),
			Member: shipment.ID,
		})
	}
}
