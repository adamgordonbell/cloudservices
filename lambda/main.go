package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	readability "github.com/go-shiori/go-readability"
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
		article, err := readability.FromURL(event.QueryStringParameters.Url, 30*time.Second)
		if err != nil {
			log.Printf("Error: failed to parse (422) %v: %v", event.QueryStringParameters.Url, err)
			return Response{StatusCode: 422}, err
		}
		cmd := exec.Command("lynx", "--stdin", "--dump", "--nolist", "--assume_charset=utf8")
		cmd.Stdin = strings.NewReader(article.Content)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Error: failed to lynx %v: %v", event.QueryStringParameters.Url, err)
			return Response{StatusCode: 500}, err
		}
		body := article.Title + "\n\n" + string(out) + "\n\n" + "Text-Mode By Earthly.dev"
		return Response{Body: body, StatusCode: 200, Headers: headersTXT}, nil
	}
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
