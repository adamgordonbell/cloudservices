package textmode

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	readability "github.com/go-shiori/go-readability"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if url := r.FormValue("url"); url != "" {
		TextModeHandler(w, r)
	} else {
		TextModeHomeHandler(w, r)
	}
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

func TextModeHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	mode := r.FormValue("mode")
	result, err := Process(url, mode)
	if err != nil {
		log.Printf("Error: failed to process: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, result)
}

var failedToParse = errors.New("failed to parse")
var lynxFailure = errors.New("failed to run lynx")
var pandocFailure = errors.New("failed to run pandoc")

func Process(url string, mode string) (string, error) {
	var resp string
	var err error
	//ToDo: These should be replaced with handlers that call converters
	switch mode {
	case "markdown":
		resp, err = processToMarkdown(url)
	case "html":
		resp, err = processToReadable(url)
	default:
		resp, err = processToPlainText(url)
	}
	return resp, err
}

func processToReadable(url string) (string, error) {
	log.Println("Processing to readability")
	article, err := readability.FromURL(url, 30*time.Second)
	if err != nil {
		log.Printf("Error: failed to parse (422) %v: %v", url, err)
		return "", failedToParse
	}
	return article.Content, nil
}

func processToPlainText(url string) (string, error) {
	log.Println("Processing to Plain Text")
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

func processToMarkdown(url string) (string, error) {
	log.Println("Processing to markdown")
	article, err := readability.FromURL(url, 30*time.Second)
	if err != nil {
		log.Printf("Error: failed to parse (422) %v: %v", url, err)
		return "", failedToParse
	}
	cmd := exec.Command("pandoc", "-s", "--from=html", "--to=markdown_strict-raw_html-native_divs-native_spans-fenced_divs-bracketed_spans")
	cmd.Stdin = strings.NewReader(article.Content)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error: failed to lynx %v: %v", url, err)
		return "", lynxFailure
	}
	return string(out), nil
}
