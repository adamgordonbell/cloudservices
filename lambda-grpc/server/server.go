package server

import (
	api "github.com/adamgordonbell/cloudservices/lambda-grpc/api/v1"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	api.UnimplementedNumberServiceServer
}

func NewGRPCServer() *grpc.Server {
	gsrv := grpc.NewServer()
	srv := GrpcServer{}
	api.RegisterNumberServiceServer(gsrv, &srv)
	return gsrv
}
