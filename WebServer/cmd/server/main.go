package main

import (
	"github.com/adamgordonbell/cloudservices/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	srv.ListenAndServe()
}
