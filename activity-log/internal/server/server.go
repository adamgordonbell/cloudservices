package server

import (
	"context"
	"fmt"
	"log"
	"time"

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
	if err == ErrIDNotFound {
		return nil, fmt.Errorf("%d not found", req.Id)
	}
	if err != nil {
		return nil, fmt.Errorf("Internal Error: %w", err)
	}
	return &resp, nil
}

func (s *grpcServer) Insert(ctx context.Context, activity *api.Activity) (*api.InsertResponse, error) {
	id, err := s.Activities.Insert(*activity)
	if err != nil {
		log.Printf("Error:%s", err.Error())
		return nil, fmt.Errorf("Internal Error: %w", err)
	}
	res := api.InsertResponse{Id: int32(id)}
	return &res, nil
}

func (s *grpcServer) List(ctx context.Context, req *api.ListRequest) (*api.Activities, error) {
	activities, err := s.Activities.List(int(req.Offset))
	if err != nil {
		return nil, fmt.Errorf("Internal Error: %w", err)
	}
	return &api.Activities{Activities: activities}, nil
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

type Activity struct {
	Time        time.Time
	Description string
	ID          int
}

func convert(a api.Activity) Activity {
	return Activity{
		ID:          int(a.Id),
		Description: a.Description,
		Time:        a.Time.AsTime(),
	}
}
