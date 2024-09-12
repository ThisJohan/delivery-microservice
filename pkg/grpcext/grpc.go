package grpcext

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	GrpcPort string `env:"GRPC_PORT" envDefault:"6565"`
}

func Serve(s *grpc.Server, c Config) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", c.GrpcPort))
	if err != nil {
		return fmt.Errorf("cannot create listener: %s", err)
	}
	fmt.Println("Grpc Listening on port " + c.GrpcPort)
	return s.Serve(lis)
}

func NewInternalConnection(configs Config) grpc.ClientConnInterface {
	conn, err := grpc.NewClient(
		fmt.Sprintf("localhost:%v", configs.GrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	return conn
}

func NewConnection(addr string) grpc.ClientConnInterface {
	conn, err := grpc.NewClient(addr)
	if err != nil {
		panic(err)
	}
	return conn
}
