package uber

import (
	"context"
	"fmt"

	"github.com/ThisJohan/snapp-assignment/api"
	"github.com/ThisJohan/snapp-assignment/pkg/di"
	"google.golang.org/grpc"
)

func NewService(server *grpc.Server) {
	s := new(Service)
	api.RegisterUberServiceServer(server, s)
}

type Service struct {
	api.UnimplementedUberServiceServer
}

func (s *Service) Todo(ctx context.Context, req *api.TodoRequest) (*api.TodoResponse, error) {
	x := di.Inject[string](ctx, "Hi")
	fmt.Println(x)
	return &api.TodoResponse{Data: "Hello There"}, nil
}
