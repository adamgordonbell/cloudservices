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

	app := NewApp()
	s := r.PathPrefix("/default").Subrouter()
	s.HandleFunc("/text-mode", app.handlerCreator(textmode.ConvertHTMLToReadablePlainText, "text/plain; charset=utf-8"))
	s.HandleFunc("/text-mode-raw", app.handlerCreator(textmode.ConvertHTMLToPlainText, "text/plain; charset=utf-8"))

	s.HandleFunc("/markdown-mode", app.handlerCreator(textmode.ConvertHTMLToReadableMarkDown, "text/plain; charset=utf-8"))
	s.HandleFunc("/markdown-mode-raw", app.handlerCreator(textmode.ConvertHTMLToMarkDown, "text/plain; charset=utf-8"))

	s.HandleFunc("/html-mode", app.handlerCreator(textmode.ConvertHTMLToReadableHTML, "text/html; charset=utf-8"))
	s.HandleFunc("/tldr-mode", app.handlerCreator(textmode.ConvertHTMLToTLDR, "text/plain; charset=utf-8"))
	s.HandleFunc("/", HomeHandler)

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

func (app App) handlerCreator(opp textmode.Conversion, contentType string) func(http.ResponseWriter,
	*http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if url := r.FormValue("url"); url != "" {
			result, err := app.RunAndCache(opp, url, r.RequestURI)
			if err != nil {
				log.Printf("Error getting content: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", contentType)
			w.WriteHeader(http.StatusOK)
			_, _ = io.WriteString(w, result)
		} else {
			TextModeHomeHandler(w, r)
		}
	}
}

func (app App) RunAndCache(opp textmode.Conversion, url string, path string) (string, error) {
	key := getCacheKey(url, path)
	result, err := app.get(key)
	if true { //errors.Is(err, errNoKey) {
		log.Println("No cache value found")
		body, err := textmode.RequestBody(url)
		if err != nil {
			return "", fmt.Errorf("Error getting content: %w", err)
		}
		result, err := opp(body, url)
		if err != nil {
			return "", fmt.Errorf("Error converting content: %w", err)
		}
		_ = app.put(key, result)
		return result, nil
	} else if err != nil {
		return "", fmt.Errorf("Error: failed to get cache: %w, calling method", err)
	} else {
		log.Println("Cache hit")
		return result, nil
	}
}
