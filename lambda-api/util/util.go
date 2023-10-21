package util

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func RequestBody(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func CleanWhitespace(s string) string {
	space := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(space.ReplaceAllString(s, " "))
}

func CleanNewLines(s string) string {
	newlines := regexp.MustCompile(`\n\n+`)
	s2 := newlines.ReplaceAllString(s, "{{{BREAK}}}")
	s3 := strings.ReplaceAll(s2, "\n", " ")
	return strings.ReplaceAll(s3, "{{{BREAK}}}", "\n\n")
}
