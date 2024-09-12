package logistics

import (
	"context"

	"github.com/ThisJohan/snapp-assignment/api"
	"github.com/ThisJohan/snapp-assignment/pkg/grpcext"
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
	Todo(ctx context.Context, in *api.TodoRequest, opts ...grpc.CallOption) (*api.TodoResponse, error)
}

type Service struct {
	api.UnimplementedLogisticsServiceServer
	providers map[string]LogisticProvider
}

func (s *Service) Todo(ctx context.Context, req *api.TodoRequest) (*api.TodoResponse, error) {
	provider := s.providers["uber"]
	return provider.Todo(ctx, req)
}
