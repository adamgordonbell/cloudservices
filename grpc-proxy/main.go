package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	api "github.com/adamgordonbell/cloudservices/activity-log/api/v1"
)

var grpcServerEndpoint = "localhost:8080"

func main() {
	log.Println("Starting listening on port 8081")
	port := ":8081"

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := api.RegisterActivity_LogHandlerFromEndpoint(context.Background(), mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	http.ListenAndServe(port, mux)
}
