package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	api "github.com/adamgordonbell/cloudservices/activity-log/api/v1"
)

var grpcServerEndpoint = ":8080"

func main() {
	log.Println("Starting listening on port 8081")
	port := ":8081"
	mux := runtime.NewServeMux()
	tlsCreds, err := credentials.NewClientTLSFromFile("../activity-log/certs/ca.pem", "")
	if err != nil {
		log.Fatalf("No cert found: %v", err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCreds)}
	err = api.RegisterActivityLogServiceHandlerFromEndpoint(context.Background(), mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
