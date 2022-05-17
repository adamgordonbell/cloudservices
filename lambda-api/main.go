package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/adamgordonbell/cloudservices/lambda-api/textmode"
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
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	app := textmode.NewApp()
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Not found", r.RequestURI)
		http.Error(w, fmt.Sprintf("Not found: %s", r.RequestURI), http.StatusNotFound)
	})

	s := r.PathPrefix("/default/lambda-api").Subrouter()
	s.HandleFunc("/text-mode", TextModeHomeHandler)
	s.HandleFunc("/text-mode/{url}", app.TextModeHandler)
	s.HandleFunc("/", HomeHandler)
	r.Use(loggingMiddleware)

	if _, inLambda := os.LookupEnv("AWS_LAMBDA_RUNTIME_API"); inLambda {
		log.Println("Starting up in Lambda Runtime")
		adapter := gorillamux.NewV2(r)
		lambda.Start(adapter.ProxyWithContext)
	} else {

		log.Println("Starting up on own")
		srv := &http.Server{
			Addr:    ":8080",
			Handler: r,
		}
		_ = srv.ListenAndServe()
	}
}
