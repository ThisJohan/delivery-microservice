package shipping

import (
	"context"

	"github.com/ThisJohan/snapp-assignment/pkg/db"
	"gorm.io/gorm"
)

type (
	shipmentRepository struct{}
)

var (
	Shipments = shipmentRepository{}
)

type Shipment struct {
	gorm.Model
	Name string
}

func (s *shipmentRepository) Create(ctx context.Context, data *Shipment) error {
	db := db.InjectDB(ctx)
	return db.Create(data).Error
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Shipment{})
}
