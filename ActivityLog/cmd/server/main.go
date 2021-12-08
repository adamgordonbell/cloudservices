package main

import (
	"github.com/adamgordonbell/cloudservices/activitylog/internal/server"
)

func main() {
	println("Starting on http://localhost:8080")
	srv := server.NewHTTPServer(":8080")
	srv.ListenAndServe()
}
