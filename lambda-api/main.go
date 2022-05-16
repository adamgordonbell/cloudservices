package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	_ "github.com/motemen/go-loghttp/global"
)

func TextModeHomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Text Mode page request")
	bytes, err := ioutil.ReadFile("textmode.txt")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home page request")
	log.Println(r.RequestURI)
	bytes, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("Starting up")
	r := mux.NewRouter()
	r.HandleFunc("/text-mode", TextModeHomeHandler)
	r.HandleFunc("/lambda-api/text-mode", TextModeHomeHandler)
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)

	r.Use(loggingMiddleware)

	adapter := gorillamux.NewV2(r)

	lambda.Start(adapter.ProxyWithContext)
}
