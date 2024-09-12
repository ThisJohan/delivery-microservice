package uber

import (
	"context"

	"github.com/ThisJohan/snapp-assignment/api"
	"google.golang.org/grpc"
)

func NewService(server *grpc.Server) {
	s := new(Service)
	api.RegisterUberServiceServer(server, s)
}

type Service struct {
	api.UnimplementedUberServiceServer
}

func (s *Service) Test(context.Context, *api.TodoRequest) (*api.TodoResponse, error) {
	panic("unimplemented")
}
