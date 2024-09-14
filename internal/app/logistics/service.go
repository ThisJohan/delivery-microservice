package logistics

import (
	"context"

	"github.com/ThisJohan/delivery-microservice/api"
	"github.com/ThisJohan/delivery-microservice/pkg/grpcext"
	"google.golang.org/grpc"
)

func NewService(server *grpc.Server, config Config) {
	internalConn := grpcext.NewInternalConnection(config.Grpc)
	s := &Service{
		providers: map[string]LogisticProvider{
			"uber": api.NewUberServiceClient(internalConn),
		},
	}
	api.RegisterLogisticsServiceServer(server, s)
}

type LogisticProvider interface {
	Deliver(ctx context.Context, in *api.DeliverRequest, opts ...grpc.CallOption) (*api.DeliverResponse, error)
}

type Service struct {
	api.UnimplementedLogisticsServiceServer
	providers map[string]LogisticProvider
}

func (s *Service) Deliver(ctx context.Context, req *api.DeliverRequest) (*api.DeliverResponse, error) {
	provider := s.providers["uber"]
	return provider.Deliver(ctx, req)
}
