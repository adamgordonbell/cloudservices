package blackerfriday

import (
	"bytes"
	"log"
	"strings"

	"github.com/adamgordonbell/cloudservices/lambda-api/parse"
	"github.com/adamgordonbell/cloudservices/lambda-api/util"
	"github.com/russross/blackfriday"
)

func ToEssayMarkdown(body string) string {
	log.Println("Processing Markdown to Esssay Text")
	r := string(blackfriday.Markdown([]byte(body), EssayRenderer(0), 0))
	return util.CleanNewLines(r)
}

type EssayMarkdown struct {
	article parse.Article
}

// EssayMarkdownRenderer creates and configures a EssayMarkdown object, which
// satisfies the Renderer interface.
//
// flags is a set of EssayMarkdown_* options ORed together (currently no such options
// are defined).
func EssayRenderer(flags int) blackfriday.Renderer {
	return &EssayMarkdown{}
}

func (options *EssayMarkdown) GetFlags() int {
	return 0
}

// render code chunks using verbatim, or listings if we have a language
func (options *EssayMarkdown) BlockCode(out *bytes.Buffer, text []byte, info string) {
}

func (options *EssayMarkdown) TitleBlock(out *bytes.Buffer, text []byte) {
}

func (options *EssayMarkdown) BlockQuote(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (options *EssayMarkdown) BlockHtml(out *bytes.Buffer, text []byte) {
}

func (options *EssayMarkdown) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	// options.article.Nodes = append(options.article.Nodes, parse.HeadingNode{content: })
	marker := out.Len()
	for i := 0; i < level; i++ {
		out.WriteString("#")
	}
	out.WriteString(" ")
	if !text() {
		out.Truncate(marker)
		return
	}
	out.WriteString("\n\n")
}

func (options *EssayMarkdown) HRule(out *bytes.Buffer) {
}

func (options *EssayMarkdown) List(out *bytes.Buffer, text func() bool, flags int) {
	text()
}

func (options *EssayMarkdown) ListItem(out *bytes.Buffer, text []byte, flags int) {
}

func (options *EssayMarkdown) Paragraph(out *bytes.Buffer, text func() bool) {
	text()
	out.WriteString("\n\n")
}

func (options *EssayMarkdown) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
}

func (options *EssayMarkdown) TableRow(out *bytes.Buffer, text []byte) {
}

func (options *EssayMarkdown) TableHeaderCell(out *bytes.Buffer, text []byte, align int) {
}

func (options *EssayMarkdown) TableCell(out *bytes.Buffer, text []byte, align int) {
}

// TODO: this
func (options *EssayMarkdown) Footnotes(out *bytes.Buffer, text func() bool) {
	text()
}

func (options *EssayMarkdown) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {

}

func (options *EssayMarkdown) AutoLink(out *bytes.Buffer, link []byte, kind int) {
}

func (options *EssayMarkdown) CodeSpan(out *bytes.Buffer, text []byte) {
	//change periods to hacker around sentence parser
	s := strings.ReplaceAll(string(text), ".", "•")
	out.Write([]byte(s))
}

func (options *EssayMarkdown) DoubleEmphasis(out *bytes.Buffer, text []byte) {
}

func (options *EssayMarkdown) Emphasis(out *bytes.Buffer, text []byte) {
}

func (options *EssayMarkdown) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
}

func (options *EssayMarkdown) LineBreak(out *bytes.Buffer) {
	out.WriteString("\n")
}

func (options *EssayMarkdown) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	out.Write(content)
}

func (options *EssayMarkdown) RawHtmlTag(out *bytes.Buffer, tag []byte) {
}

func (options *EssayMarkdown) TripleEmphasis(out *bytes.Buffer, text []byte) {
}

func (options *EssayMarkdown) StrikeThrough(out *bytes.Buffer, text []byte) {
}

func (options *EssayMarkdown) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
}

func (options *EssayMarkdown) Entity(out *bytes.Buffer, entity []byte) {
	out.Write(entity)
}

func (options *EssayMarkdown) NormalText(out *bytes.Buffer, text []byte) {
	log.Println("NormalText", string(text))
	out.Write(text)
}

// header and footer
func (options *EssayMarkdown) DocumentHeader(out *bytes.Buffer) {
}

func (options *EssayMarkdown) DocumentFooter(out *bytes.Buffer) {
}
