package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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

func (s *State) logRawRequestAndProxy(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	jRequest, _ := json.Marshal(event)
	log.Printf("Raw Input:\n %s\n", string(jRequest))
	resp, err := s.adapter.Proxy(event)
	jResp, _ := json.Marshal(resp)
	log.Printf("Raw Output:\n %s\n", string(jResp))
	return resp, err
}

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Not found", r.RequestURI)
		http.Error(w, fmt.Sprintf("Not found: %s", r.RequestURI), http.StatusNotFound)
	})

	s := r.PathPrefix("/default").Subrouter()
	s.HandleFunc("/text-mode", textmode.Handler)
	s.HandleFunc("/markdown-mode", handlerCreator(textmode.ConvertHTMLToMarkDown, "text/plain; charset=utf-8"))
	s.HandleFunc("/markdown-mode2", handlerCreator(textmode.ConvertHTMLToReadableMarkDown, "text/plain; charset=utf-8"))
	s.HandleFunc("/", HomeHandler)
	// app := NewApp()
	// r.Use(app.cachingMiddleWare)

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

func handlerCreator(opp textmode.Conversion, contentType string) func(http.ResponseWriter,
	*http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if url := r.FormValue("url"); url != "" {
			body, err := RequestBody(url)
			if err != nil {
				log.Printf("Error getting content: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			result, err := opp(body, url)
			if err != nil {
				log.Printf("Error converting content: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", contentType)
			w.WriteHeader(http.StatusOK)
			_, _ = io.WriteString(w, result)
		} else {
			textmode.TextModeHomeHandler(w, r)
		}
	}
}

func RequestBody(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
