package shipping

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ThisJohan/snapp-assignment/api"
	"github.com/ThisJohan/snapp-assignment/pkg/grpcext"
	"github.com/ThisJohan/snapp-assignment/pkg/redisext"
	"github.com/go-co-op/gocron/v2"
	"github.com/redis/go-redis/v9"
)

const (
	shipmentsStack = "pending_shipments"
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
			10*time.Minute,
		),
		gocron.NewTask(w.CheckPendingShipments),
	)
	s.NewJob(
		gocron.DurationJob(
			time.Minute,
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
	// do some process on expired items
	for _, item := range expiredItems {
		shipment, _ := ShipmentsRepo.Get(w.ctx, item)
		if now.Hour() >= 23 {
			shipment.Status = api.Shipment_CANCELED.String()
			ShipmentsRepo.Update(w.ctx, shipment)
		} else {
			queueShipment(w.ctx, shipment)
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
	fmt.Println(len(shipments))
	for _, shipment := range shipments {
		queueShipment(w.ctx, &shipment)
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

		if _, err := w.logistics.Todo(w.ctx, &api.TodoRequest{}); err != nil {
			// Retry until success
			queueShipment(w.ctx, &shipment)
		}

		expiration := time.Now().Add(time.Minute * 15).Unix()
		rdb.ZAdd(w.ctx, shipmentsStack, redis.Z{
			Score:  float64(expiration),
			Member: shipment.ID,
		})
	}
}
