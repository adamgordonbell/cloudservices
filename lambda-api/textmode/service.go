package textmode

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	readability "github.com/go-shiori/go-readability"
)

type App struct {
	S3 *s3.S3
}

func NewApp() App {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		log.Fatalf("failed to create AWS session, %v", err)
	}
	s3 := s3.New(sess)
	return App{S3: s3}
}

func (app App) TextModeHandler(w http.ResponseWriter, r *http.Request) {
	result, err := app.ProcessWithCache("https://www.google.com")
	if err != nil {
		log.Printf("Error: failed to process: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, result)
}

var failedToParse = errors.New("failed to parse")
var lynxFailure = errors.New("failed to run lynx")
var cacheFailure = errors.New("failed to access cache")

func (app App) ProcessWithCache(url string) (string, error) {
	result, err := app.get(url)
	if errors.Is(err, errNoKey) {
		log.Println("No cache value found")
		resp, err := process(url)
		if err != nil {
			return resp, err
		}
		log.Println("Caching value")
		err = app.put(url, resp)
		return resp, err
	} else if err != nil {
		log.Printf("Error: failed to get cache: %v", err)
		return "", cacheFailure
	} else {
		log.Println("Cache hit")
		return result, nil
	}
}

func process(url string) (string, error) {
	article, err := readability.FromURL(url, 30*time.Second)
	if err != nil {
		log.Printf("Error: failed to parse (422) %v: %v", url, err)
		return "", failedToParse
	}
	cmd := exec.Command("lynx", "--stdin", "--dump", "--nolist", "--assume_charset=utf8")
	cmd.Stdin = strings.NewReader(article.Content)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: failed to lynx %v: %v", url, err)
		return "", lynxFailure
	}
	return string(out), nil
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
