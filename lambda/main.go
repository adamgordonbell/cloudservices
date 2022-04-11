package main

import (
	"fmt"
	"io/ioutil"

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
		str := string(bytes) // convert content to a 'string'
		return Response{Body: str, StatusCode: 200, Headers: headersTXT}, nil
	} else {
		return Response{Body: fmt.Sprintf("got url: %s", event.QueryStringParameters.Url), StatusCode: 200}, nil
	}
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
