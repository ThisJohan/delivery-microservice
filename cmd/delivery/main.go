package main

import (
	"log"

	"github.com/ThisJohan/delivery-microservice/internal/app/shipping"
	"github.com/ThisJohan/delivery-microservice/pkg/db"
	"github.com/ThisJohan/delivery-microservice/pkg/di"
	"github.com/ThisJohan/delivery-microservice/pkg/env"
	"github.com/ThisJohan/delivery-microservice/pkg/grpcext"
	"github.com/ThisJohan/delivery-microservice/pkg/redisext"
	"github.com/ThisJohan/delivery-microservice/pkg/workerext"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Config struct {
	Grpc     grpcext.Config
	Shipping shipping.Config
	Database db.Config
	Redis    redisext.Config
}

const (
	service = "delivery"
)

var (
	configs Config
)

func init() {
	if err := env.LoadConfig(service, &configs); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}

func main() {
	database, err := db.OpenDBConnection(configs.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := shipping.Migrate(database); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	rdb := redisext.Connect(configs.Redis)

	// Worker
	worker := shipping.NewWorker(configs.Shipping)
	workerext.StartWorker(
		worker,
		di.DIBuilder(db.Service, func() *gorm.DB {
			return database
		}),
		di.DIBuilder(redisext.Service, func() *redis.Client {
			return rdb
		}),
	)

	// GRPC Server
	s := grpc.NewServer(
		di.GrpcProvide(
			di.DIBuilder(db.Service, func() *gorm.DB {
				return database
			}),
			di.DIBuilder(redisext.Service, func() *redis.Client {
				return rdb
			}),
		),
	)

	shipping.NewService(s, configs.Shipping)

	if err := grpcext.Serve(s, configs.Grpc); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
