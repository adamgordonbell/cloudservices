package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

type App struct {
	AwsSession *session.Session
	S3         *s3.S3
}

var headersTXT = map[string]string{"Content-Type": "application/json"}

func (app App) HandleLambdaEvent(event Event) (Response, error) {

	if event.QueryStringParameters.Url == "" {
		log.Println("Home page request")
		bytes, err := ioutil.ReadFile("index.txt")
		if err != nil {
			fmt.Print(err)
			return Response{}, err
		}
		return Response{Body: string(bytes), StatusCode: 200, Headers: headersTXT}, nil
	} else {
		log.Printf("request: %v", event.QueryStringParameters.Url)
		cache, err := app.get(event.QueryStringParameters.Url)
		if err != nil {
			log.Printf("cache miss %v", err)
			log.Printf("No cache found: %v", event.QueryStringParameters.Url)
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
			err = app.put(event.QueryStringParameters.Url, body)
			if err != nil {
				log.Printf("Error: failed to put %v: %v", event.QueryStringParameters.Url, err)
			}
			return Response{Body: body, StatusCode: 200, Headers: headersTXT}, nil
		} else {
			log.Printf("cache found: %v", event.QueryStringParameters.Url)
			return Response{Body: cache, StatusCode: 200, Headers: headersTXT}, nil
		}
	}
}

func (app App) put(url string, result string) error {
	input := &s3.PutObjectInput{
		Body:   strings.NewReader(result),
		Bucket: aws.String("text-mode"),
		Key:    aws.String(url),
	}
	r, err := app.S3.PutObject(input)
	if err != nil {
		return fmt.Errorf("failed to store result: %w", err)
	}
	log.Printf("Stored result: %v", r)
	return nil
}

func (app App) get(url string) (string, error) {
	req := &s3.GetObjectInput{
		Bucket: aws.String("text-mode"),
		Key:    aws.String(url),
	}
	r, err := app.S3.GetObject(req)
	if err != nil {
		return "", fmt.Errorf("failed to get result: %w", err)
	}
	defer r.Body.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(r.Body)
	if err != nil {
		return "", fmt.Errorf("failed to get S3 result: %w", err)
	}
	return buf.String(), nil
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		log.Fatalf("failed to create AWS session, %v", err)
	}
	s3 := s3.New(sess)
	app := App{AwsSession: sess, S3: s3}

	lambda.Start(app.HandleLambdaEvent)
}
