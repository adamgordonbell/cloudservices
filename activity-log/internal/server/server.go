package server

import (
	"context"
	"log"

	api "github.com/adamgordonbell/cloudservices/activity-log/api/v1"

	"google.golang.org/grpc"
)

// var _ api1.Activity = (*grpcServer)(nil)

type grpcServer struct {
	api.UnimplementedActivity_LogServer
	Activities *Activities
}

func (s *grpcServer) Retrieve(ctx context.Context, req *api.RetrieveRequest) (*api.Activity, error) {
	resp, err := s.Activities.Retrieve(int(req.Id))
	return nil, nil
}

func (s *grpcServer) Insert(ctx context.Context, activity *api.Activity) (*api.Activity, error) {
	return nil, nil
}

func (s *grpcServer) List(ctx context.Context, req *api.ListRequest) (*api.Activities, error) {
	return nil, nil
}

func NewGRPCServer() *grpc.Server {
	var acc *Activities
	var err error
	if acc, err = NewActivities(); err != nil {
		log.Fatal(err)
	}
	gsrv := grpc.NewServer()
	srv := grpcServer{
		Activities: acc,
	}
	api.RegisterActivity_LogServer(gsrv, &srv)
	return gsrv
}
