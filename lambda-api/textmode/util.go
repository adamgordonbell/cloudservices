package textmode

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

func cleanWhitespace(s string) string {
	space := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(space.ReplaceAllString(s, " "))
}
