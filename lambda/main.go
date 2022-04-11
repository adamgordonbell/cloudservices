package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	QueryStringParameters QueryStringParameters `json:"queryStringParameters"`
}
type QueryStringParameters struct {
	url string
}

type Response struct {
	StatusCode int               `json:"statusCode"`
	Body       string            `json:"body"`
	Headers    map[string]string `json:"headers"`
}

func HandleLambdaEvent(event Event) (Response, error) {
	return Response{Body: fmt.Sprintf("got url: %s", event.QueryStringParameters.url), StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
