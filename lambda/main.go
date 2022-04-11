package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	QueryStringParameters QueryStringParameters `json:"queryStringParameters"`
}
type QueryStringParameters struct {
	Url string `json:"url"`
}

type Response struct {
	StatusCode int               `json:"statusCode"`
	Body       string            `json:"body"`
	Headers    map[string]string `json:"headers"`
}

var headersTXT = map[string]string{"Content-Type": "application/json"}

func HandleLambdaEvent(event Event) (Response, error) {
	if event.QueryStringParameters.Url == "" {
		bytes, err := ioutil.ReadFile("index.txt")
		if err != nil {
			fmt.Print(err)
			return Response{}, err
		}
		return Response{Body: string(bytes), StatusCode: 200, Headers: headersTXT}, nil
	} else {
		resp, err := http.Get(event.QueryStringParameters.Url)
		if err != nil {
			log.Printf("Error: %v", err)
			return Response{StatusCode: 500}, err
		}
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error: %v", err)
			return Response{StatusCode: 500}, err
		}
		return Response{Body: string(bytes), StatusCode: 200, Headers: headersTXT}, nil
	}
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
