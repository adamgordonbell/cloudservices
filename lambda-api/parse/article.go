package parse

type Article struct {
	Title   string
	Url     string
	Content string
	Nodes   []ArticleNode
}

func NewArticle(title string, url string, content string) *Article {

	return &Article{
		Title:   title,
		Url:     url,
		Content: content,
	}
}

type ArticleNode interface {
	getContent() string
}

type HeadingNode struct {
	content string
}

func (h *HeadingNode) getContent() string {
	return h.content
}

type ParagraphNode struct {
	content string
}

func (p *ParagraphNode) getContent() string {
	return p.content
}

func (article *Article) getParagraphs() []ArticleNode {
	paragraphs := []ArticleNode{}
	for _, node := range article.Nodes {
		switch node.(type) {
		case *ParagraphNode:
			paragraphs = append(paragraphs, node)
		}
	}
	return paragraphs
}

func (article *Article) getHeadings() []ArticleNode {
	headings := []ArticleNode{}
	for _, node := range article.Nodes {
		switch node.(type) {
		case *HeadingNode:
			headings = append(headings, node)
		}
	}
	return headings
}
