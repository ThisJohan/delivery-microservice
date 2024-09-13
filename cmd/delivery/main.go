package main

import (
	"fmt"
	"log"

	"github.com/ThisJohan/snapp-assignment/internal/app/shipping"
	"github.com/ThisJohan/snapp-assignment/pkg/db"
	"github.com/ThisJohan/snapp-assignment/pkg/di"
	"github.com/ThisJohan/snapp-assignment/pkg/env"
	"github.com/ThisJohan/snapp-assignment/pkg/grpcext"
	"google.golang.org/grpc"
)

type Config struct {
	Grpc     grpcext.Config
	Shipping shipping.Config
	Database db.Config
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
	fmt.Println(configs)

	database, err := db.OpenDBConnection(configs.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := shipping.Migrate(database); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	s := grpc.NewServer(
		di.GrpcProvide("db", database),
	)

	shipping.NewService(s, configs.Shipping)

	if err := grpcext.Serve(s, configs.Grpc); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
