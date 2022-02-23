package main

import (
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/adamgordonbell/cloudservices/activity-log/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(tamird): point to merged gRPC code rather than a PR.
		// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

var grpcServerEndpoint = "localhost:9090"

func main() {
	log.Println("Starting listening on port 8080")
	port := ":8080"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", port)

	// mux := runtime.NewServeMux()
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// err := api.RegisterActivity_LogHandlerFromEndpoint(context.Background(), mux, grpcServerEndpoint, opts)
	// if err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	srv := server.NewGRPCServer()
	// Register reflection service on gRPC server.
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// http.ListenAndServe(port, mux)
}
