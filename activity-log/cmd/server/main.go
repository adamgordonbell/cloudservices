package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/adamgordonbell/cloudservices/activity-log/internal/server"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(tamird): point to merged gRPC code rather than a PR.
		// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		log.Printf("Got: %s", r)
		// grpcServer.ServeHTTP(w, r)
		if false {
			fmt.Fprintf(w, "get\n")
		} else if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			log.Println("grpcServer")
			grpcServer.ServeHTTP(w, r)
		} else {
			log.Println("REST")
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func main() {

	// GRPC Server
	grpcServer := server.NewGRPCServer()
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// Rest Server
	gwmux := runtime.NewServeMux()

	log.Println("Starting listening on port 8080")
	err := http.ListenAndServe(":8080", grpcHandlerFunc(grpcServer, gwmux))
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
