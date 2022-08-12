package textmode

import (
	"errors"
	"fmt"
	"log"
	nurl "net/url"
	"os/exec"
	"strings"

	readability "github.com/go-shiori/go-readability"
)

type Conversion func(string, string) (string, error)

var failedToParse = errors.New("failed to parse")
var lynxFailure = errors.New("failed to run lynx")
var pandocFailure = errors.New("failed to run pandoc")

func ConvertHTMLToReadablePlainText(body string, pageURL string) (string, error) {
	body, err := ConvertHTMLToReadableHTML(body, pageURL)
	if err != nil {
		return "", err
	}
	return ConvertHTMLToPlainText(body, pageURL)
}

func ConvertHTMLToReadableMarkDown(body string, pageURL string) (string, error) {
	body, err := ConvertHTMLToReadableHTML(body, pageURL)
	if err != nil {
		return "", err
	}
	return ConvertHTMLToMarkDown(body, pageURL)
}

func ConvertHTMLToReadableHTML(body string, pageURL string) (string, error) {
	log.Println("Processing HTML to Plain Text")
	parsedURL, err := nurl.ParseRequestURI(pageURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %v", err)
	}
	article, err := readability.FromReader(strings.NewReader(body), parsedURL)
	if err != nil {
		log.Printf("Error: failed to parse url: %s", pageURL)
		return "", failedToParse
	}
	return article.Content, nil
}

func ConvertHTMLToPlainText(body string, pageURL string) (string, error) {
	log.Println("Processing HTML to Markdown using lynx")
	cmd := exec.Command("lynx", "--stdin", "--dump", "--nolist", "--assume_charset=utf8")
	cmd.Stdin = strings.NewReader(body)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: failed to lynx %v: %v", pageURL, err)
		return "", lynxFailure
	}
	return string(out), nil
}

func ConvertHTMLToMarkDown(body string, pageURL string) (string, error) {
	log.Println("Processing HTML to Markdown using lynx")
	cmd := exec.Command("pandoc", "-s", "--from=html", "--to=markdown_strict-raw_html-native_divs-native_spans-fenced_divs-bracketed_spans")
	cmd.Stdin = strings.NewReader(body)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: failed to pandoc %v: %v", pageURL, err)
		return "", pandocFailure
	}
	return string(out), nil
}
