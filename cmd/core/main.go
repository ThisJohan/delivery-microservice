package main

import (
	"log"

	"github.com/ThisJohan/delivery-microservice/internal/app/manager"
	"github.com/ThisJohan/delivery-microservice/pkg/env"
	"github.com/ThisJohan/delivery-microservice/pkg/grpcext"
	"github.com/ThisJohan/delivery-microservice/pkg/workerext"
	"google.golang.org/grpc"
)

const (
	service = "core"
)

type Config struct {
	Grpc    grpcext.Config
	Manager manager.Config
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
	// Worker
	worker := manager.NewWorker(configs.Manager)
	workerext.StartWorker(
		worker,
	)

	// GRPC Server
	s := grpc.NewServer()

	manager.NewService(s)

	if err := grpcext.Serve(s, configs.Grpc); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
