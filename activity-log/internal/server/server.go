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

type GrpcServer struct {
	api.UnimplementedActivityLogServiceServer
	Activities *Activities
}

func (s *GrpcServer) Retrieve(ctx context.Context, req *api.RetrieveRequest) (*api.RetrieveResponse, error) {
	activity, err := s.Activities.Retrieve(int(req.Id))
	if err == ErrIDNotFound {
		return nil, status.Error(codes.NotFound, "id was not found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api.RetrieveResponse{Activity: activity}, nil
}

func (s *GrpcServer) Insert(ctx context.Context, req *api.InsertRequest) (*api.InsertResponse, error) {
	id, err := s.Activities.Insert(req.Activity)
	if err != nil {
		log.Printf("Error:%s", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	res := api.InsertResponse{Id: int32(id)}
	return &res, nil
}

func (s *GrpcServer) List(ctx context.Context, req *api.ListRequest) (*api.ListResponse, error) {
	activities, err := s.Activities.List(int(req.Offset))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api.ListResponse{Activities: activities}, nil
}

func NewGRPCServer() (*grpc.Server, GrpcServer) {
	var acc *Activities
	var err error
	if acc, err = NewActivities(); err != nil {
		log.Fatal(err)
	}
	gsrv := grpc.NewServer()
	srv := GrpcServer{
		Activities: acc,
	}
	api.RegisterActivityLogServiceServer(gsrv, &srv)
	return gsrv, srv
}

type Activity struct {
	Time        time.Time
	Description string
	ID          int
}
