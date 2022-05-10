package main

import (
	"log"
	"net/http"

	"github.com/adamgordonbell/cloudservices/lambda-grpc/server"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("Starting Up")

	grpcServer := server.NewGRPCServer()

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	lambda.Start(handlerfunc.NewV2(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		grpcServer.ServeHTTP(w, req)
	}).ProxyWithContext)

	log.Println("Started")
}
