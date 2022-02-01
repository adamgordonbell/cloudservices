package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

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
		return nil, errors.New("not found")
	}
	if err != nil {
		return nil, errors.New("Internal Error")
	}
	return &resp, nil
}

func (s *grpcServer) Insert(ctx context.Context, activity *api.Activity) (*api.Activity, error) {
	id, err := s.Activities.Insert(*activity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := api.IDDocument{ID: id}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
