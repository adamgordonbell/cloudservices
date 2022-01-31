package server

import (
	"log"

	api1 "github.com/adamgordonbell/cloudservices/activity-log/api/v1"
	"google.golang.org/grpc"
)

type grpcServer struct {
	api1.UnimplementedActivity_LogServer
	Activities *Activities
}

// var _ api1.Activity = (*grpcServer)(nil)

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
	api1.RegisterActivity_LogServer(gsrv, srv)
	return gsrv
}
