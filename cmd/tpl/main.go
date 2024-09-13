package main

import (
	"log"

	"github.com/ThisJohan/snapp-assignment/internal/app/logistics"
	"github.com/ThisJohan/snapp-assignment/internal/app/uber"
	"github.com/ThisJohan/snapp-assignment/pkg/env"
	"github.com/ThisJohan/snapp-assignment/pkg/grpcext"
	"google.golang.org/grpc"
)

type Config struct {
	Grpc      grpcext.Config
	Uber      uber.Config
	Logistics logistics.Config
}

const (
	service = "tpl"
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
	s := grpc.NewServer()

	uber.NewService(s, configs.Uber)
	logistics.NewService(s, configs.Logistics)

	if err := grpcext.Serve(s, configs.Grpc); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
