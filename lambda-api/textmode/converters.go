package textmode

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	nurl "net/url"
	"os/exec"
	"strings"

	textrank "github.com/DavidBelicza/TextRank"
	"github.com/adamgordonbell/cloudservices/lambda-api/parse"
	readability "github.com/go-shiori/go-readability"
	"github.com/gomarkdown/markdown"
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
	cmd := exec.Command("lynx", "--stdin", "--dump", "--nolist", "--assume_charset=utf8", "--display_charset=utf-8")
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

func ConvertHTMLToTLDR(body string, pageURL string) (string, error) {
	log.Println("Processing HTML to TLDR")
	body, err := ConvertHTMLToReadablePlainText(body, pageURL)
	if err != nil {
		return "", err
	}
	tmpl := template.Must(template.ParseFiles("./textmode/tldr.md"))
	section := parse.Article{Title: "", Url: pageURL, Content: body}
	summary := ConvertSectionToSummary(section)
	w := bytes.NewBufferString("")
	err = tmpl.Execute(w, summary)
	if err != nil {
		return "", err
	}
	html := markdown.ToHTML(w.Bytes(), nil, nil)
	return string(html), err
}

type summary struct {
	Title   string
	Author  string
	URL     string
	Topic   string
	Quotes  []string
	Phrases []string
}

func ConvertSectionToSummary(article parse.Article) summary {
	tr := textrank.NewTextRank()
	rule := textrank.NewDefaultRule()
	language := textrank.NewDefaultLanguage()
	algorithmDef := textrank.NewChainAlgorithm()
	tr.Populate(article.Content, language, rule)
	tr.Ranking(algorithmDef)
	rankedPhrases := textrank.FindPhrases(tr)

	quotes := []string{}
	relationQuote := textrank.FindSentencesByRelationWeight(tr, 3)
	for _, value := range relationQuote {
		sentenses := textrank.FindSentencesFrom(tr, value.ID, 3)
		q := ""
		for _, s := range sentenses {
			q += s.Value + " "
		}
		quotes = append(quotes, cleanWhitespace(q))
	}

	phrases := []string{}
	for _, valeu := range rankedPhrases[:3] {
		phrases = append(phrases, cleanWhitespace(valeu.Right+" "+valeu.Left))
	}

	return summary{
		Title:   article.Title,
		Author:  "Unknown",
		URL:     article.Url,
		Topic:   rankedPhrases[0].Right + " " + rankedPhrases[0].Left,
		Quotes:  quotes,
		Phrases: phrases,
	}
}
