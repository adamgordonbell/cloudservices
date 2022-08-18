package textmode

import (
	"errors"
	"fmt"
	"log"
	nurl "net/url"
	"os/exec"
	"regexp"
	"strings"

	textrank "github.com/DavidBelicza/TextRank"
	"github.com/JesusIslam/tldr"
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
	//maybe switch to https://github.com/DavidBelicza/TextRank
	text, err := ConvertHTMLToReadablePlainText(body, pageURL)
	if err != nil {
		return "", err
	}
	space := regexp.MustCompile(`\s+`)
	text = space.ReplaceAllString(text, " ")
	intoSentences := 2
	bag := tldr.New()
	result, err := bag.Summarize(text, intoSentences)
	return strings.Join(result, "\n\n"), err
}

type summary struct {
	title  string
	topic  string
	quotes []string
}

type section struct {
	title   string
	content string
}

func ConvertSectionToSummary(section section) summary {
	tr := textrank.NewTextRank()
	rule := textrank.NewDefaultRule()
	language := textrank.NewDefaultLanguage()
	algorithmDef := textrank.NewChainAlgorithm()
	tr.Populate(section.content, language, rule)
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

	println(rankedPhrases)
	return summary{
		title:  section.title,
		topic:  rankedPhrases[0].Right + " " + rankedPhrases[0].Left,
		quotes: quotes,
	}
}
