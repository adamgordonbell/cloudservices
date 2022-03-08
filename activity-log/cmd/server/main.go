package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/adamgordonbell/cloudservices/activity-log/internal/server"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	api "github.com/adamgordonbell/cloudservices/activity-log/api/v1"
)

func grpcHandlerFunc(grpcServer grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(
			r.Header.Get("Content-Type"), "application/grpc") {
			log.Println("GRPC")
			grpcServer.ServeHTTP(w, r)
		} else {
			log.Println("REST")
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func main() {

	// GRPC Server
	grpcServer, srv := server.NewGRPCServer()
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// Rest Server
	mux := runtime.NewServeMux()
	err := api.RegisterActivityLogServiceHandlerServer(context.Background(), mux, &srv)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Starting listening on port 8080")
	err = http.ListenAndServeTLS(":8080", "./certs/server.pem", "./certs/server-key.pem", grpcHandlerFunc(*grpcServer, mux))
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
