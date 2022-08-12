package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type App struct {
	S3 *s3.S3
}

// var cacheFailure = errors.New("failed to access cache")

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

func getCacheKey(path string, url string) string {
	return path + "|" + url
}

func (app App) put(key string, result string) error {
	input := &s3.PutObjectInput{
		Body:   strings.NewReader(result),
		Bucket: aws.String("text-mode"),
		Key:    aws.String(key),
	}
	r, err := app.S3.PutObject(input)
	if err != nil {
		return fmt.Errorf("failed to store result: %w", err)
	}
	log.Printf("Stored result: %v", r)
	return nil
}

var errNoKey = errors.New(s3.ErrCodeNoSuchKey)

func (app App) get(key string) (string, error) {
	req := &s3.GetObjectInput{
		Bucket: aws.String("text-mode"),
		Key:    aws.String(key),
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
