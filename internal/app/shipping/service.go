package shipping

import (
	"github.com/ThisJohan/snapp-assignment/api"
	"github.com/ThisJohan/snapp-assignment/pkg/grpcext"
	"google.golang.org/grpc"
)

func NewService(server *grpc.Server, config Config) {
	s := &Service{
		logistics: api.NewLogisticsServiceClient(grpcext.NewConnection(config.TPLService)),
	}
	// res, err := s.logistics.Todo(context.Background(), &api.TodoRequest{})
	// fmt.Println(res, err)
	api.RegisterShippingServiceServer(server, s)
}

type Service struct {
	api.UnimplementedShippingServiceServer
	logistics api.LogisticsServiceClient
}
