package shipping

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ThisJohan/delivery-microservice/api"
	"github.com/ThisJohan/delivery-microservice/pkg/db"
	"github.com/ThisJohan/delivery-microservice/pkg/redisext"
	"gorm.io/gorm"
)

type (
	shipmentRepository struct{}
)

var (
	ShipmentsRepo = shipmentRepository{}
)

// TODO: write migration instead of auto migrate
// CREATE TYPE shipment_status AS ENUM ('PENDING', 'ASSIGNED', 'DELIVERED', 'CANCELED', 'QUEUED');
// TODO: it is a good idea to have a retry count column to track how much time it took for a 3pl to accept delivery

type Shipment struct {
	gorm.Model
	OrderID     string
	UserAddress string
	Origin      string
	Destination string
	TimeSlot    time.Time
	Status      string `gorm:"type:shipment_status"`
}

func (s *Shipment) BeforeSave(*gorm.DB) error {
	now := time.Now()
	futureLimit := now.AddDate(0, 0, 4)

	if s.TimeSlot.Before(now) || s.TimeSlot.After(futureLimit) {
		return fmt.Errorf("time slot out of range")
	}
	hour := s.TimeSlot.Hour()
	if hour < 9 || hour >= 23 {
		return fmt.Errorf("time slot out of range")
	}
	return nil
}

func (s *shipmentRepository) Create(ctx context.Context, data *Shipment) error {
	db := db.InjectDB(ctx)
	return db.Create(data).Error
}

func (s *shipmentRepository) Get(ctx context.Context, id string) (*Shipment, error) {
	db := db.InjectDB(ctx)
	var shipment Shipment
	err := db.First(&shipment, id).Error
	return &shipment, err
}

func (s *shipmentRepository) Update(ctx context.Context, data *Shipment) error {
	db := db.InjectDB(ctx)
	return db.Save(data).Error
}

func (s *shipmentRepository) GetPendingShipments(ctx context.Context, timeMargin time.Duration) ([]Shipment, error) {
	db := db.InjectDB(ctx)
	var shipments []Shipment
	now := time.Now()
	err := db.Where(
		"status = ? AND time_slot > ? AND time_slot <= ?",
		api.Shipment_PENDING.String(),
		now,
		now.Add(timeMargin),
	).Find(&shipments).Error

	return shipments, err
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Shipment{})
}

func enqueueShipment(ctx context.Context, s *Shipment) error {
	return db.InjectDB(ctx).Transaction(func(tx *gorm.DB) error {
		s.Status = api.Shipment_QUEUED.String()
		if err := tx.Save(s).Error; err != nil {
			return err
		}
		rdb := redisext.InjectRedis(ctx)
		payload, _ := json.Marshal(s)
		if err := rdb.Publish(ctx, "tpl_queue", payload).Err(); err != nil {
			return err
		}
		return nil
	})
}
