package shipping

import (
	"context"
	"fmt"
	"time"

	"github.com/ThisJohan/snapp-assignment/pkg/db"
	"gorm.io/gorm"
)

type (
	shipmentRepository struct{}
)

var (
	ShipmentsRepo = shipmentRepository{}
)

// Todo: write migration instead of auto migrate
// CREATE TYPE shipment_status AS ENUM ('PENDING', 'ASSIGNED', 'DELIVERED', 'CANCELED');

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

func (s *shipmentRepository) GetPendingShipments(ctx context.Context, timeMargin time.Duration) ([]Shipment, error) {
	db := db.InjectDB(ctx)
	var shipments []Shipment
	err := db.Where(
		"status = ? AND time_slot > ? AND time_slot <= ?",
		"PENDING",
		time.Now(),
		time.Now().Add(timeMargin),
	).Find(&shipments).Error

	return shipments, err
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Shipment{})
}
