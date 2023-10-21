package blackerfriday

import (
	"testing"

	"github.com/russross/blackfriday"
)

func runMarkdownBlock(input string) string {

	renderer := EssayRenderer(0)

	return string(blackfriday.Markdown([]byte(input), renderer, 0))
}

func doTestsBlock(t *testing.T, tests []string) {
	doTestsBlockWithRunner(t, tests, runMarkdownBlock)
}

func doTestsBlockWithRunner(t *testing.T, tests []string, runner func(string) string) {
	// catch and report panics
	var candidate string
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("\npanic while processing [%#v]: %s\n", candidate, err)
		}
	}()

	for i := 0; i+1 < len(tests); i += 2 {
		input := tests[i]
		candidate = input
		expected := tests[i+1]
		actual := runner(candidate)
		if actual != expected {
			t.Errorf("\nInput   [%#v]\nExpected[%#v]\nActual  [%#v]",
				candidate, expected, actual)
		}
	}
}

func TestPrefixHeaderNoExtensions(t *testing.T) {
	var tests = []string{
		"# Header 1\n",
		"# Header 1\n",

		"## Header 2\n",
		"## Header 2\n",

		"# Heading\nText Text Text\n",
		"# Heading\nText Text Text\n",

		// "# Heading\n* list\n",
		// "# Heading\n",

		// `* list\n\n* list\nxyz adfasdf asdfs\n`,
		// "",

		// "bla bla bla\nbla blabla\n",
		// "bla bla bla\nbla blabla\n",

		// "bla bla bla\n\n* l\n",
		// "bla bla bla\n",
	}
	doTestsBlock(t, tests)
}
