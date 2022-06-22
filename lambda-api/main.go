package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/adamgordonbell/cloudservices/lambda-api/textmode"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	_ "github.com/motemen/go-loghttp/global"
)

type State struct {
	adapter *gorillamux.GorillaMuxAdapterV2
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

// func logRawRequest
func (s *State) logRawRequestAndProxy(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	jRequest, _ := json.Marshal(event)
	log.Printf("Raw Input:\n %s\n", string(jRequest))
	resp, err := s.adapter.Proxy(event)
	jResp, _ := json.Marshal(resp)
	log.Printf("Raw Output:\n %s\n", string(jResp))
	return resp, err
}

func main() {
	app := textmode.NewApp()
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Not found", r.RequestURI)
		http.Error(w, fmt.Sprintf("Not found: %s", r.RequestURI), http.StatusNotFound)
	})

	s := r.PathPrefix("/default").Subrouter()
	s.HandleFunc("/text-mode", app.Handler)
	s.HandleFunc("/", HomeHandler)
	r.Use(loggingMiddleware)

	if runtime_api, _ := os.LookupEnv("AWS_LAMBDA_RUNTIME_API"); runtime_api != "" {
		log.Println("Starting up in Lambda Runtime")
		adapter := gorillamux.NewV2(r)
		state := &State{adapter: adapter}
		lambda.Start(state.logRawRequestAndProxy)
	} else {
		log.Println("Starting up on own")
		srv := &http.Server{
			Addr:    ":8080",
			Handler: r,
		}
		_ = srv.ListenAndServe()
	}
}
