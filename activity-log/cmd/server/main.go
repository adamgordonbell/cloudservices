package main

import (
	"context"
	"log"
	"net/http"

	"github.com/adamgordonbell/cloudservices/activity-log/internal/server"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	api "github.com/adamgordonbell/cloudservices/activity-log/api/v1"
)

func grpcHandlerFunc(grpcServer grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(tamird): point to merged gRPC code rather than a PR.
		// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		log.Printf("Got: %s", r.Header.Get("Content-Type"))
		// if r.ProtoMajor == 2 && strings.HasPrefix(
		// 	r.Header.Get("Content-Type"), "application/grpc") {
		// log.Println("GRPC")
		// grpcServer.ServeHTTP(w, r)
		// } else {
		// 	log.Println("other")
		otherHandler.ServeHTTP(w, r)
		// }
	})
}

func main() {

	// GRPC Server
	grpcServer, srv := server.NewGRPCServer()
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// Rest Server
	mux := runtime.NewServeMux()
	err := api.RegisterActivity_LogHandlerServer(context.Background(), mux, &srv)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Starting listening on port 8080")
	err = http.ListenAndServe(":8080", grpcHandlerFunc(*grpcServer, mux))
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
