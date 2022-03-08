package main

import (
	"context"
	"log"
	"net/http"

	"github.com/adamgordonbell/cloudservices/activity-log/internal/server"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	api "github.com/adamgordonbell/cloudservices/activity-log/api/v1"
)

func main() {

	// GRPC Server
	_, srv := server.NewGRPCServer()

	// Rest Server
	mux := runtime.NewServeMux()
	err := api.RegisterActivityLogServiceHandlerServer(context.Background(), mux, &srv)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Starting listening on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
