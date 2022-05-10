package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
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
	S3 *s3.S3
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
		result, err := app.get(event.QueryStringParameters.Url)
		if errors.Is(err, errNoKey) {
			log.Println("No cache value found")
			resp, err := process(event.QueryStringParameters.Url)
			if err != nil {
				return resp, err
			}
			log.Println("Caching value")
			err = app.put(event.QueryStringParameters.Url, resp.Body)
			return resp, err
		} else if err != nil {
			return Response{StatusCode: 500}, err
		}
		return Response{Body: result, StatusCode: 200, Headers: headersTXT}, nil
	}
}

func process(url string) (Response, error) {
	article, err := readability.FromURL(url, 30*time.Second)
	if err != nil {
		log.Printf("Error: failed to parse (422) %v: %v", url, err)
		return Response{StatusCode: 422}, err
	}
	cmd := exec.Command("lynx", "--stdin", "--dump", "--nolist", "--assume_charset=utf8")
	cmd.Stdin = strings.NewReader(article.Content)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: failed to lynx %v: %v", url, err)
		return Response{StatusCode: 500}, err
	}
	body := article.Title + "\n\n" + string(out)
	return Response{Body: body, StatusCode: 200, Headers: headersTXT}, nil
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

var errNoKey = errors.New(s3.ErrCodeNoSuchKey)

func (app App) get(url string) (string, error) {
	req := &s3.GetObjectInput{
		Bucket: aws.String("text-mode"),
		Key:    aws.String(url),
	}
	r, err := app.S3.GetObject(req)
	if err != nil {
		if s3Err, ok := err.(awserr.Error); ok && s3Err.Code() == s3.ErrCodeNoSuchKey {
			return "", errNoKey
		} else {
			return "", fmt.Errorf("failed to get result: %w", err)
		}
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
	app := App{S3: s3}

	lambda.Start(app.HandleLambdaEvent)
}
