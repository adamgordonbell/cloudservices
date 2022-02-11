package server

import (
	"context"
	"log"
	"time"

	api "github.com/adamgordonbell/cloudservices/activity-log/api/v1"

	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// var _ api1.Activity = (*grpcServer)(nil)

type grpcServer struct {
	api.UnimplementedActivity_LogServer
	Activities *Activities
}

func (s *grpcServer) Retrieve(ctx context.Context, req *api.RetrieveRequest) (*api.Activity, error) {
	resp, err := s.Activities.Retrieve(int(req.Id))
	if err == ErrIDNotFound {
		return nil, status.Error(codes.NotFound, "id was not found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *grpcServer) Insert(ctx context.Context, activity *api.Activity) (*api.InsertResponse, error) {
	id, err := s.Activities.Insert(activity)
	if err != nil {
		log.Printf("Error:%s", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	res := api.InsertResponse{Id: int32(id)}
	return &res, nil
}

func (s *grpcServer) List(ctx context.Context, req *api.ListRequest) (*api.Activities, error) {
	activities, err := s.Activities.List(int(req.Offset))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
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
